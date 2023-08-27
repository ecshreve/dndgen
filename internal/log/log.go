package log

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	colorRed    = 31
	colorYellow = 33
	colorPurple = 35
	colorBlue   = 36
	colorGray   = 37
)

// CustomFormatter is our implementation of a log formatter.
type CustomFormatter struct {
	log.TextFormatter
}

// Format handles applying custom formatting to a log entry.
func (f *CustomFormatter) Format(entry *log.Entry) ([]byte, error) {
	levelColor := getColorByLevel(entry.Level)
	return []byte(fmt.Sprintf("\x1b[%dm%s\x1b[0m - [%s] - \x1b[100m%s\x1b[0m - %s\n", levelColor, strings.ToUpper(entry.Level.String())[:4], entry.Time.Format(f.TimestampFormat), strings.TrimPrefix(entry.Caller.Function, "github.com/ecshreve/dndgen"), entry.Message)), nil
}

// getColorByLevel handles level to color mapping for the formatter.
func getColorByLevel(level log.Level) int {
	switch level {
	case log.TraceLevel:
		return colorBlue
	case log.WarnLevel:
		return colorYellow
	case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
		return colorRed
	case log.DebugLevel, log.InfoLevel:
		return colorPurple
	default:
		return colorGray
	}
}

// init initializes the global logger with the custom formatter.
func init() {
	log.SetFormatter(&CustomFormatter{log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
	},
	})
	log.SetLevel(log.TraceLevel)
	log.SetReportCaller(true)
}
