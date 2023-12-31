// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	base64 "encoding/base64"
	entproto "entgo.io/contrib/entproto"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	fmt "fmt"
	ent "github.com/ecshreve/dndgen/ent"
	abilityscore "github.com/ecshreve/dndgen/ent/abilityscore"
	skill "github.com/ecshreve/dndgen/ent/skill"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	strconv "strconv"
)

// SkillService implements SkillServiceServer
type SkillService struct {
	client *ent.Client
	UnimplementedSkillServiceServer
}

// NewSkillService returns a new SkillService
func NewSkillService(client *ent.Client) *SkillService {
	return &SkillService{
		client: client,
	}
}

// toProtoSkill transforms the ent type to the pb type
func toProtoSkill(e *ent.Skill) (*Skill, error) {
	v := &Skill{}
	desc := e.Desc
	v.Desc = desc
	id := int64(e.ID)
	v.Id = id
	indx := e.Indx
	v.Indx = indx
	name := e.Name
	v.Name = name
	if edg := e.Edges.AbilityScore; edg != nil {
		id := int64(edg.ID)
		v.AbilityScore = &AbilityScore{
			Id: id,
		}
	}
	return v, nil
}

// toProtoSkillList transforms a list of ent type to a list of pb type
func toProtoSkillList(e []*ent.Skill) ([]*Skill, error) {
	var pbList []*Skill
	for _, entEntity := range e {
		pbEntity, err := toProtoSkill(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements SkillServiceServer.Create
func (svc *SkillService) Create(ctx context.Context, req *CreateSkillRequest) (*Skill, error) {
	skill := req.GetSkill()
	m, err := svc.createBuilder(skill)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoSkill(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Get implements SkillServiceServer.Get
func (svc *SkillService) Get(ctx context.Context, req *GetSkillRequest) (*Skill, error) {
	var (
		err error
		get *ent.Skill
	)
	id := int(req.GetId())
	switch req.GetView() {
	case GetSkillRequest_VIEW_UNSPECIFIED, GetSkillRequest_BASIC:
		get, err = svc.client.Skill.Get(ctx, id)
	case GetSkillRequest_WITH_EDGE_IDS:
		get, err = svc.client.Skill.Query().
			Where(skill.ID(id)).
			WithAbilityScore(func(query *ent.AbilityScoreQuery) {
				query.Select(abilityscore.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoSkill(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements SkillServiceServer.Update
func (svc *SkillService) Update(ctx context.Context, req *UpdateSkillRequest) (*Skill, error) {
	skill := req.GetSkill()
	skillID := int(skill.GetId())
	m := svc.client.Skill.UpdateOneID(skillID)
	skillDesc := skill.GetDesc()
	m.SetDesc(skillDesc)
	skillIndx := skill.GetIndx()
	m.SetIndx(skillIndx)
	skillName := skill.GetName()
	m.SetName(skillName)
	if skill.GetAbilityScore() != nil {
		skillAbilityScore := int(skill.GetAbilityScore().GetId())
		m.SetAbilityScoreID(skillAbilityScore)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoSkill(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Delete implements SkillServiceServer.Delete
func (svc *SkillService) Delete(ctx context.Context, req *DeleteSkillRequest) (*emptypb.Empty, error) {
	var err error
	id := int(req.GetId())
	err = svc.client.Skill.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements SkillServiceServer.List
func (svc *SkillService) List(ctx context.Context, req *ListSkillRequest) (*ListSkillResponse, error) {
	var (
		err      error
		entList  []*ent.Skill
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.Skill.Query().
		Order(ent.Desc(skill.FieldID)).
		Limit(pageSize + 1)
	if req.GetPageToken() != "" {
		bytes, err := base64.StdEncoding.DecodeString(req.PageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		token, err := strconv.ParseInt(string(bytes), 10, 32)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		pageToken := int(token)
		listQuery = listQuery.
			Where(skill.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListSkillRequest_VIEW_UNSPECIFIED, ListSkillRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListSkillRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithAbilityScore(func(query *ent.AbilityScoreQuery) {
				query.Select(abilityscore.FieldID)
			}).
			All(ctx)
	}
	switch {
	case err == nil:
		var nextPageToken string
		if len(entList) == pageSize+1 {
			nextPageToken = base64.StdEncoding.EncodeToString(
				[]byte(fmt.Sprintf("%v", entList[len(entList)-1].ID)))
			entList = entList[:len(entList)-1]
		}
		protoList, err := toProtoSkillList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &ListSkillResponse{
			SkillList:     protoList,
			NextPageToken: nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements SkillServiceServer.BatchCreate
func (svc *SkillService) BatchCreate(ctx context.Context, req *BatchCreateSkillsRequest) (*BatchCreateSkillsResponse, error) {
	requests := req.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.SkillCreate, len(requests))
	for i, req := range requests {
		skill := req.GetSkill()
		var err error
		bulk[i], err = svc.createBuilder(skill)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.Skill.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoSkillList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &BatchCreateSkillsResponse{
			Skills: protoList,
		}, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *SkillService) createBuilder(skill *Skill) (*ent.SkillCreate, error) {
	m := svc.client.Skill.Create()
	skillDesc := skill.GetDesc()
	m.SetDesc(skillDesc)
	skillIndx := skill.GetIndx()
	m.SetIndx(skillIndx)
	skillName := skill.GetName()
	m.SetName(skillName)
	if skill.GetAbilityScore() != nil {
		skillAbilityScore := int(skill.GetAbilityScore().GetId())
		m.SetAbilityScoreID(skillAbilityScore)
	}
	return m, nil
}
