package leveler

import (
	"log"
	"bytes"
	"context"
	"reflect"
	data "leveler/data"
	util "leveler/util"
	grpc "google.golang.org/grpc"
	endpoints "leveler/endpoints"
	jsonpb "github.com/golang/protobuf/jsonpb"
	proto "github.com/golang/protobuf/proto"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
)

var DatabaseLabelsMap = map[string]int{
	"action": 0,
	"requirement": 1,
	"role": 2,
	"host": 3,
}

type EndpointServer struct {
	Database data.Database
}

func (s *EndpointServer) RegisterEndpoints (grpcServer *grpc.Server) {
	endpoints.RegisterActionEndpointServer(grpcServer, s)
	endpoints.RegisterRequirementEndpointServer(grpcServer, s)
	endpoints.RegisterRoleEndpointServer(grpcServer, s)
	endpoints.RegisterHostEndpointServer(grpcServer, s)
}

// ACTION ENDPOINTS

func (s *EndpointServer) genericCreate(t string, obj proto.Message, dest interface{}) error {
	log.Printf("Creating %s: %v", t, obj)

	var id string

	m, err := util.ConvertProtoToJsonMap(obj)
	if err != nil {
		return err
	}

	id, err = s.Database.Create(t, m)
	if err != nil {
		return err
	}

	v := reflect.ValueOf(dest).Elem().FieldByName("Id")
	ptr := v.Addr().Interface().(*string)
	*ptr = id

	return nil
}

func (s *EndpointServer) genericGet(t string, id string, dest interface{}) error {
	log.Printf("Retrieving %s: %s", t, id)

	var jsonString []byte

	result, err := s.Database.Get(t, id)
	if err != nil {
		return err
	}

	jsonString, err = util.ConvertMapToJson(result)
	if err != nil {
		log.Printf("Error converting map to JSON: %v", err)
		return err
	}

	err = jsonpb.Unmarshal(bytes.NewReader(jsonString), dest.(proto.Message))
	if err != nil {
		return err
	}

	return nil
}

func (s *EndpointServer) genericList(t string, query string, dest interface{}) error {
	log.Printf("Retrieiving %s list", t)

	var jsonString []byte

	result, err := s.Database.List(t, query)
	if err != nil {
		return err
	}

	jsonString, err = util.ConvertMapToJson(result)
	if err != nil {
		return err
	}

	err = jsonpb.Unmarshal(bytes.NewReader(jsonString), dest.(proto.Message))
	if err != nil {
		return err
	}

	return nil
}

func (s *EndpointServer) genericUpdate(t string, id string, obj proto.Message, dest interface{}) error {
	log.Printf("Updating %s: %v", obj)

	var jsonString []byte

	m, err := util.ConvertProtoToJsonMap(obj)
	if err != nil {
		return err
	}

	result, err := s.Database.Update(t, id, m)
	if err != nil {
		return err
	}

	jsonString, err = util.ConvertMapToJson(result)
	if err != nil {
		return err
	}

	err = jsonpb.Unmarshal(bytes.NewReader(jsonString), dest.(proto.Message))
	if err != nil {
		return err
	}

	return nil
}

func (s *EndpointServer) genericDelete(t string, id string) error {
	log.Printf("Deleting %s: %s", t, id)

	err := s.Database.Delete(t, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *EndpointServer) CreateAction(ctx context.Context, action *endpoints.Action) (*endpoints.ActionId, error) {
	var actionId = &endpoints.ActionId{}
	err := s.genericCreate("action", action, actionId)
	if err != nil {
		return actionId, err
	}

	return actionId, nil
}

func (s *EndpointServer) GetAction(ctx context.Context, actionId *endpoints.ActionId) (*endpoints.Action, error) {
	var action = &endpoints.Action{}
	err := s.genericGet("action", actionId.Id, action)
	if err != nil {
		return action, err
	}

	return action, nil
}

func (s *EndpointServer) ListActions(ctx context.Context, query *endpoints.Query) (*endpoints.ActionList, error) {
	var actionList = &endpoints.ActionList{}
	err := s.genericList("action", query.Query, actionList)
	if err != nil {
		return actionList, err
	}

	return actionList, nil
}

func (s *EndpointServer) UpdateAction(ctx context.Context, action *endpoints.Action) (*endpoints.Action, error) {
	var updatedAction = &endpoints.Action{}
	err := s.genericUpdate("action", action.Id, action, updatedAction)
	if err != nil {
		return updatedAction, err
	}

	return updatedAction, nil
}

func (s *EndpointServer) DeleteAction(ctx context.Context, actionId *endpoints.ActionId) (*google_protobuf.Empty, error) {
	var emptyMessage = &google_protobuf.Empty{}
	err := s.genericDelete("action", actionId.Id)
	if err != nil {
		return emptyMessage, err
	}

	return emptyMessage, nil
}

// REQUIREMENT ENDPOINTS

func (s *EndpointServer) CreateRequirement(ctx context.Context, requirement *endpoints.Requirement) (*endpoints.RequirementId, error) {
	var requirementId = &endpoints.RequirementId{}

	err := s.genericCreate("requirement", requirement, requirementId)
	if err != nil {
		return requirementId, err
	}

	return requirementId, nil
}

func (s *EndpointServer) GetRequirement(ctx context.Context, requirementId *endpoints.RequirementId) (*endpoints.Requirement, error) {
	var requirement = &endpoints.Requirement{}

	err := s.genericGet("requirement", requirementId.Id, requirement)
	if err != nil {
		return requirement, err
	}

	return requirement, nil
}

func (s *EndpointServer) ListRequirements(ctx context.Context, query *endpoints.Query) (*endpoints.RequirementList, error) {
	var requirementList = &endpoints.RequirementList{}

	err := s.genericList("requirement", query.Query, requirementList)
	if err != nil {
		return requirementList, err
	}

	return requirementList, nil
}

func (s *EndpointServer) UpdateRequirement(ctx context.Context, requirement *endpoints.Requirement) (*endpoints.Requirement, error) {
	var updatedRequirement = &endpoints.Requirement{}
	err := s.genericUpdate("requirement", requirement.Id, requirement, updatedRequirement)
	if err != nil {
		return updatedRequirement, err
	}

	return updatedRequirement, nil
}

func (s *EndpointServer) DeleteRequirement(ctx context.Context, requirementId *endpoints.RequirementId) (*google_protobuf.Empty, error) {
	var emptyMessage = &google_protobuf.Empty{}
	err := s.genericDelete("requirement", requirementId.Id)
	if err != nil {
		return emptyMessage, err
	}

	return emptyMessage, nil
}

// ROLE ENDPOINTS 

func (s *EndpointServer) CreateRole(ctx context.Context, role *endpoints.Role) (*endpoints.RoleId, error) {
	var roleId = &endpoints.RoleId{}

	err := s.genericCreate("role", role, roleId)
	if err != nil {
		return roleId, err
	}

	return roleId, nil
}

func (s *EndpointServer) GetRole(ctx context.Context, roleId *endpoints.RoleId) (*endpoints.Role, error) {
	var role = &endpoints.Role{}

	err := s.genericGet("role", roleId.Id, role)
	if err != nil {
		return role, err
	}

	return role, nil
}

func (s *EndpointServer) ListRoles(ctx context.Context, query *endpoints.Query) (*endpoints.RoleList, error) {
	var roleList = &endpoints.RoleList{}

	err := s.genericList("role", query.Query, roleList)
	if err != nil {
		return roleList, err
	}

	return roleList, nil
}

func (s *EndpointServer) UpdateRole(ctx context.Context, role *endpoints.Role) (*endpoints.Role, error) {
	var updatedRole = &endpoints.Role{}
	err := s.genericUpdate("role", role.Id, role, updatedRole)
	if err != nil {
		return updatedRole, err
	}

	return updatedRole, nil
}

func (s *EndpointServer) DeleteRole(ctx context.Context, roleId *endpoints.RoleId) (*google_protobuf.Empty, error) {
	var emptyMessage = &google_protobuf.Empty{}
	err := s.genericDelete("role", roleId.Id)
	if err != nil {
		return emptyMessage, err
	}

	return emptyMessage, nil
}

// HOST ENDPOINTS

func (s *EndpointServer) CreateHost(ctx context.Context, host *endpoints.Host) (*endpoints.HostId, error) {
	var hostId = &endpoints.HostId{}

	err := s.genericCreate("host", host, hostId)
	if err != nil {
		return hostId, err
	}

	return hostId, nil
}

func (s *EndpointServer) GetHost(ctx context.Context, hostId *endpoints.HostId) (*endpoints.Host, error) {
	var host = &endpoints.Host{}

	err := s.genericGet("host", hostId.Id, host)
	if err != nil {
		return host, err
	}

	return host, nil
}

func (s *EndpointServer) ListHosts(ctx context.Context, query *endpoints.Query) (*endpoints.HostList, error) {
	var hostList = &endpoints.HostList{}

	err := s.genericList("host", query.Query, hostList)
	if err != nil {
		return hostList, err
	}

	return hostList, nil
}

func (s *EndpointServer) UpdateHost(ctx context.Context, host *endpoints.Host) (*endpoints.Host, error) {
	var updatedHost = &endpoints.Host{}
	err := s.genericUpdate("host", host.Id, host, updatedHost)
	if err != nil {
		return updatedHost, err
	}

	return updatedHost, nil
}

func (s *EndpointServer) DeleteHost(ctx context.Context, hostId *endpoints.HostId) (*google_protobuf.Empty, error) {
	var emptyMessage = &google_protobuf.Empty{}
	err := s.genericDelete("host", hostId.Id)
	if err != nil {
		return emptyMessage, err
	}

	return emptyMessage, nil
}