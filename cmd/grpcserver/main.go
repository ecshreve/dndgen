package main

import (
	"context"
	"fmt"
	"net"

	"github.com/ecshreve/dndgen/ent"
	entpb "github.com/ecshreve/dndgen/ent/proto/entpb"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

const (
	component = "grpc-example"
	grpcAddr  = ":8080"
	httpAddr  = ":8081"
)

// interceptorLogger adapts go-kit logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func interceptorLogger(l log.FieldLogger) logging.Logger {
	return logging.LoggerFunc(func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
		f := make(map[string]any, len(fields)/2)
		i := logging.Fields(fields).Iterator()
		if i.Next() {
			k, v := i.At()
			f[k] = v
		}
		l := l.WithFields(f)

		switch lvl {
		case logging.LevelDebug:
			l.Debug(msg)
		case logging.LevelInfo:
			l.Info(msg)
		case logging.LevelWarn:
			l.Warn(msg)
		case logging.LevelError:
			l.Error(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

func main() {
	// Setup logging.
	logger := log.New()

	logTraceID := func(ctx context.Context) logging.Fields {
		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			return logging.Fields{"traceID", span.TraceID().String()}
		}
		return nil
	}
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.PayloadReceived, logging.PayloadSent, logging.FinishCall),
		logging.WithFieldsFromContext(logTraceID),
	}

	// Initialize an ent client.
	client, err := ent.Open("sqlite3", "file:dev.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Run the migration tool (creating tables, etc).
	// if err := client.Schema.Create(context.Background(), schema.WithGlobalUniqueID(true)); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }

	// Initialize the generated service.
	svc := entpb.NewAbilityScoreService(client)

	// Create a new gRPC server (you can wire multiple services to a single server).
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(interceptorLogger(logger), opts...),
		),

		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(interceptorLogger(logger), opts...),
		),
	)
	// Register the service with the server.
	entpb.RegisterAbilityScoreServiceServer(server, svc)

	// Open port 5000 for listening to traffic.
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}

	// Listen for traffic indefinitely.
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
