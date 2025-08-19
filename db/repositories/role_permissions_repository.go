package db

import (
	"AuthInGo/models"
	"database/sql"
)

type RolePermissionsRepository interface {
	GetRolePermissionById(id int64) (*models.RolePermissions, error)
	GetRolePermissionByRoleId(roleId int64) ([]*models.RolePermissions, error)
	AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermissions, error)
	RemovePermissionFromRole(roleId int64, permissionId int64) error
	GetAllRolePermissions() ([]*models.RolePermissions, error)
}

type RolePermissionsRepositoryImpl struct {
	db *sql.DB
}

func NewRolePermissionsRepository(_db *sql.DB) RolePermissionsRepository {
	return &RolePermissionsRepositoryImpl{
		db: _db,
	}
}

func (r *RolePermissionsRepositoryImpl) GetRolePermissionById(id int64) (*models.RolePermissions, error) {
	query := "SELECT id, role_id, permission_id, created_at, updated_at FROM role_permissions WHERE id = ?"
	row := r.db.QueryRow(query, id)

	rolePermission := &models.RolePermissions{}
	if err := row.Scan(&rolePermission.Id, &rolePermission.RoleId, &rolePermission.PermissionId, &rolePermission.Created_at, &rolePermission.Updated_at); err != nil {
		return nil, err
	}
	return rolePermission, nil
}

func (r *RolePermissionsRepositoryImpl) GetRolePermissionByRoleId(roleId int64) ([]*models.RolePermissions, error) {
	query := "SELECT id, role_id, permission_id, created_at, updated_at FROM role_permissions WHERE role_id = ?"
	rows, err := r.db.Query(query, roleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rolePermissions []*models.RolePermissions
	for rows.Next() {
		rolePermission := &models.RolePermissions{}
		if err := rows.Scan(&rolePermission.Id, &rolePermission.RoleId, &rolePermission.PermissionId, &rolePermission.Created_at, &rolePermission.Updated_at); err != nil {
			return nil, err
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}
	return rolePermissions, nil
}

func (r *RolePermissionsRepositoryImpl) AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermissions, error) {
	query := "INSERT INTO role_permissions (role_id, permission_id, created_at, updated_at) VALUES (?, ?, NOW(), NOW())"
	result, err := r.db.Exec(query, roleId, permissionId)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return r.GetRolePermissionById(id)
}

func (r *RolePermissionsRepositoryImpl) RemovePermissionFromRole(roleId int64, permissionId int64) error {
	query := "DELETE FROM role_permissions WHERE role_id = ? AND permission_id = ?"
	result, err := r.db.Exec(query, roleId, permissionId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *RolePermissionsRepositoryImpl) GetAllRolePermissions() ([]*models.RolePermissions, error) {
	query := "SELECT id, role_id, permission_id, created_at, updated_at FROM role_permissions"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rolePermissions []*models.RolePermissions
	for rows.Next() {
		rolePermission := &models.RolePermissions{}
		if err := rows.Scan(&rolePermission.Id, &rolePermission.RoleId, &rolePermission.PermissionId, &rolePermission.Created_at, &rolePermission.Updated_at); err != nil {
			return nil, err
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}
	return rolePermissions, nil
}
