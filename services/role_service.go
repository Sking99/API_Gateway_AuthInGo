package services

import (
	repo "AuthInGo/db/repositories"
	"AuthInGo/models"
)

type RoleService interface {
	GetRoleById(id int64) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRoles() ([]*models.Role, error)
	CreateRole(name string, description string) (*models.Role, error)
	UpdateRole(id int64, name string, description string) (*models.Role, error)
	DeleteRole(id int64) error
	GetRolePermissions(roleId int64) ([]*models.Permission, error)
	AddPermissionToRole(roleId int64, permissionId int64) error
	RemovePermissionFromRole(roleId int64, permissionId int64) error
	GetAllRolePermissions() ([]*models.RolePermissions, error)
	AssignRoleToUser(userId int64, roleId int64) error
}

type RoleServiceImpl struct {
	roleRepository           repo.RoleRepository
	permissionRepository     repo.PermissionRepository
	rolePermissionRepository repo.RolePermissionsRepository
	userRoleRepository       repo.UserRoleRepository
}

func NewRoleService(roleRepo repo.RoleRepository, rolePermRepo repo.RolePermissionsRepository, userRoleRepo repo.UserRoleRepository) RoleService {
	return &RoleServiceImpl{
		roleRepository:           roleRepo,
		rolePermissionRepository: rolePermRepo,
		userRoleRepository:       userRoleRepo,
	}
}

func (s *RoleServiceImpl) GetRoleById(id int64) (*models.Role, error) {
	return s.roleRepository.GetRoleById(id)
}

func (s *RoleServiceImpl) GetRoleByName(name string) (*models.Role, error) {
	return s.roleRepository.GetRoleByName(name)
}

func (s *RoleServiceImpl) GetAllRoles() ([]*models.Role, error) {
	return s.roleRepository.GetAllRoles()
}

func (s *RoleServiceImpl) CreateRole(name string, description string) (*models.Role, error) {
	return s.roleRepository.CreateRole(name, description)
}

func (s *RoleServiceImpl) UpdateRole(id int64, name string, description string) (*models.Role, error) {
	return s.roleRepository.UpdateRole(id, name, description)
}

func (s *RoleServiceImpl) DeleteRole(id int64) error {
	return s.roleRepository.DeleteRole(id)
}

func (s *RoleServiceImpl) GetRolePermissions(roleId int64) ([]*models.Permission, error) {
	rolePermissions, err := s.rolePermissionRepository.GetRolePermissionByRoleId(roleId)
	if err != nil {
		return nil, err
	}

	var permissions []*models.Permission
	for _, rp := range rolePermissions {
		permission, err := s.permissionRepository.GetPermissionById(rp.PermissionId)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

func (s *RoleServiceImpl) AddPermissionToRole(roleId int64, permissionId int64) error {
	_, err := s.rolePermissionRepository.AddPermissionToRole(roleId, permissionId)
	return err
}

func (s *RoleServiceImpl) RemovePermissionFromRole(roleId int64, permissionId int64) error {
	return s.rolePermissionRepository.RemovePermissionFromRole(roleId, permissionId)
}

func (s *RoleServiceImpl) GetAllRolePermissions() ([]*models.RolePermissions, error) {
	return s.rolePermissionRepository.GetAllRolePermissions()
}

func (s *RoleServiceImpl) AssignRoleToUser(userId int64, roleId int64) error {
	return s.userRoleRepository.AssignRoleToUser(userId, roleId)
}
