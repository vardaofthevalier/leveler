
package leveler

import (
	"log"
	"context"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"

)

func (s *EndpointServer) CreateAction(ctx context.Context, action *Action) (*ActionId, error) {
	var actionId = &ActionId{}
	err := s.genericCreate("action", action, actionId)
	if err != nil {
		log.Print(err)
		return actionId, err
	}

	return actionId, nil
}

func (s *EndpointServer) GetAction(ctx context.Context, actionId *ActionId) (*Action, error) {
	var action = &Action{}
	err := s.genericGet("action", actionId.Id, action)
	if err != nil {
		log.Print(err)
		return action, err
	}

	return action, nil
}

func (s *EndpointServer) ListActions(ctx context.Context, query *Query) (*ActionList, error) {
	var actionList = &ActionList{}
	err := s.genericList("action", query.Query, actionList)
	if err != nil {
		log.Print(err)
		return actionList, err
	}

	return actionList, nil
}

func (s *EndpointServer) UpdateAction(ctx context.Context, action *Action) (*Action, error) {
	var updatedAction = &Action{}
	err := s.genericUpdate("action", action.Id, action, updatedAction)
	if err != nil {
		log.Print(err)
		return updatedAction, err
	}

	return updatedAction, nil
}

func (s *EndpointServer) DeleteAction(ctx context.Context, actionId *ActionId) (*google_protobuf.Empty, error) {
	var emptyMessage = &google_protobuf.Empty{}
	err := s.genericDelete("action", actionId.Id)
	if err != nil {
		log.Print(err)
		return emptyMessage, err
	}

	return emptyMessage, nil
}

// REQUIREMENT ENDPOINTS

func (s *EndpointServer) CreateRequirement(ctx context.Context, requirement *Requirement) (*RequirementId, error) {
	var requirementId = &RequirementId{}

	err := s.genericCreate("requirement", requirement, requirementId)
	if err != nil {
		log.Print(err)
		return requirementId, err
	}

	return requirementId, nil
}

func (s *EndpointServer) GetRequirement(ctx context.Context, requirementId *RequirementId) (*Requirement, error) {
	var requirement = &Requirement{}

	err := s.genericGet("requirement", requirementId.Id, requirement)
	if err != nil {
		log.Print(err)
		return requirement, err
	}

	return requirement, nil
}

func (s *EndpointServer) ListRequirements(ctx context.Context, query *Query) (*RequirementList, error) {
	var requirementList = &RequirementList{}

	err := s.genericList("requirement", query.Query, requirementList)
	if err != nil {
		log.Print(err)
		return requirementList, err
	}

	return requirementList, nil
}

func (s *EndpointServer) UpdateRequirement(ctx context.Context, requirement *Requirement) (*Requirement, error) {
	var updatedRequirement = &Requirement{}
	err := s.genericUpdate("requirement", requirement.Id, requirement, updatedRequirement)
	if err != nil {
		log.Print(err)
		return updatedRequirement, err
	}

	return updatedRequirement, nil
}

func (s *EndpointServer) DeleteRequirement(ctx context.Context, requirementId *RequirementId) (*google_protobuf.Empty, error) {
	var emptyMessage = &google_protobuf.Empty{}
	err := s.genericDelete("requirement", requirementId.Id)
	if err != nil {
		log.Print(err)
		return emptyMessage, err
	}

	return emptyMessage, nil
}

// ROLE ENDPOINTS 

func (s *EndpointServer) CreateRole(ctx context.Context, role *Role) (*RoleId, error) {
	var roleId = &RoleId{}

	err := s.genericCreate("role", role, roleId)
	if err != nil {
		log.Print(err)
		return roleId, err
	}

	return roleId, nil
}

func (s *EndpointServer) GetRole(ctx context.Context, roleId *RoleId) (*Role, error) {
	var role = &Role{}

	err := s.genericGet("role", roleId.Id, role)
	if err != nil {
		log.Print(err)
		return role, err
	}

	return role, nil
}

func (s *EndpointServer) ListRoles(ctx context.Context, query *Query) (*RoleList, error) {
	var roleList = &RoleList{}

	err := s.genericList("role", query.Query, roleList)
	if err != nil {
		log.Print(err)
		return roleList, err
	}

	return roleList, nil
}

func (s *EndpointServer) UpdateRole(ctx context.Context, role *Role) (*Role, error) {
	var updatedRole = &Role{}
	err := s.genericUpdate("role", role.Id, role, updatedRole)
	if err != nil {
		log.Print(err)
		return updatedRole, err
	}

	return updatedRole, nil
}

func (s *EndpointServer) DeleteRole(ctx context.Context, roleId *RoleId) (*google_protobuf.Empty, error) {
	var emptyMessage = &google_protobuf.Empty{}
	err := s.genericDelete("role", roleId.Id)
	if err != nil {
		log.Print(err)
		return emptyMessage, err
	}

	return emptyMessage, nil
}

// HOST ENDPOINTS

func (s *EndpointServer) CreateHost(ctx context.Context, host *Host) (*HostId, error) {
	var hostId = &HostId{}

	err := s.genericCreate("host", host, hostId)
	if err != nil {
		log.Print(err)
		return hostId, err
	}

	return hostId, nil
}

func (s *EndpointServer) GetHost(ctx context.Context, hostId *HostId) (*Host, error) {
	var host = &Host{}

	err := s.genericGet("host", hostId.Id, host)
	if err != nil {
		log.Print(err)
		return host, err
	}

	return host, nil
}

func (s *EndpointServer) ListHosts(ctx context.Context, query *Query) (*HostList, error) {
	var hostList = &HostList{}

	err := s.genericList("host", query.Query, hostList)
	if err != nil {
		log.Print(err)
		return hostList, err
	}

	return hostList, nil
}

func (s *EndpointServer) UpdateHost(ctx context.Context, host *Host) (*Host, error) {
	var updatedHost = &Host{}
	err := s.genericUpdate("host", host.Id, host, updatedHost)
	if err != nil {
		log.Print(err)
		return updatedHost, err
	}

	return updatedHost, nil
}

func (s *EndpointServer) DeleteHost(ctx context.Context, hostId *HostId) (*google_protobuf.Empty, error) {
	var emptyMessage = &google_protobuf.Empty{}
	err := s.genericDelete("host", hostId.Id)
	if err != nil {
		log.Print(err)
		return emptyMessage, err
	}

	return emptyMessage, nil
}