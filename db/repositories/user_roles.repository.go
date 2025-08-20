package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
	"strings"
)

type UserRoleRepository interface {
	GetUserRoles(userId int64) ([]*models.Role, error)
	AssignRoleToUser(userId int64, roleId int64) error
	RemoveRoleFromUser(userId int64, roleId int64) error
	GetUserPermissions(userId int64) ([]*models.Permission, error)
	HasPermission(userId int64, permissionName string) (bool, error)
	HasRole(userId int64, roleName string) (bool, error)
	HasAllRoles(userId int64, roleNames []string) (bool, error)
	HasAnyRole(userId int64, roleNames []string) (bool, error)
}

type UserRoleRepositoryImpl struct {
	db *sql.DB
}

func NewUserRoleRepository(_db *sql.DB) UserRoleRepository {
	return &UserRoleRepositoryImpl{
		db: _db,
	}
}

func (r *UserRoleRepositoryImpl) GetUserRoles(userId int64) ([]*models.Role, error) {
	query := `SELECT r.id, r.name, r.description, r.created_at, r.updated_at FROM user_roles ur 
				JOIN roles r ON ur.role_id = r.id WHERE ur.user_id = ?`
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*models.Role
	for rows.Next() {
		role := &models.Role{}
		if err := rows.Scan(&role.Id, &role.Name, &role.Description, &role.Created_at, &role.Updated_at); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (r *UserRoleRepositoryImpl) AssignRoleToUser(userId int64, roleId int64) error {
	query := "INSERT INTO user_roles (user_id, role_id, created_at, updated_at) VALUES (?, ?, NOW(), NOW())"
	_, err := r.db.Exec(query, userId, roleId)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRoleRepositoryImpl) RemoveRoleFromUser(userId int64, roleId int64) error {
	query := "DELETE FROM user_roles WHERE user_id = ? AND role_id = ?"
	result, err := r.db.Exec(query, userId, roleId)
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

func (r *UserRoleRepositoryImpl) GetUserPermissions(userId int64) ([]*models.Permission, error) {
	query := `SELECT p.id, p.name, p.description, p.resource, p.action, p.created_at, p.updated_at 
				FROM user_roles ur 
				JOIN role_permissions rp ON ur.role_id = rp.role_id 
				JOIN permissions p ON rp.permission_id = p.id 
				WHERE ur.user_id = ?`
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*models.Permission
	for rows.Next() {
		permission := &models.Permission{}
		if err := rows.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.Created_at, &permission.Updated_at); err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

func (r *UserRoleRepositoryImpl) HasPermission(userId int64, permissionName string) (bool, error) {
	query := `SELECT COUNT(*) FROM user_roles ur 
				JOIN role_permissions rp ON ur.role_id = rp.role_id 
				JOIN permissions p ON rp.permission_id = p.id 
				WHERE ur.user_id = ? AND p.name = ?`
	var count int
	err := r.db.QueryRow(query, userId, permissionName).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserRoleRepositoryImpl) HasRole(userId int64, roleName string) (bool, error) {
	query := `SELECT COUNT(*) FROM user_roles ur 
				JOIN roles r ON ur.role_id = r.id 
				WHERE ur.user_id = ? AND r.name = ?`
	var count int
	err := r.db.QueryRow(query, userId, roleName).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserRoleRepositoryImpl) HasAllRoles(userId int64, roleNames []string) (bool, error) {
	if len(roleNames) == 0 {
		return true, nil // If no roles are specified, return true
	}

	query := `SELECT COUNT(*) = ? 
				FROM user_roles ur
				INNER JOIN roles r ON ur.role_id = r.id
				WHERE ur.user_id = ? AND r.name IN (?)
				GROUP BY ur.user_id`

	roleNamesStr := strings.Join(roleNames, ",")
	row := r.db.QueryRow(query, len(roleNames), userId, roleNamesStr)
	var hasAllRoles bool
	if err := row.Scan(&hasAllRoles); err != nil {
		return false, err
	}

	return hasAllRoles, nil
}

func (r *UserRoleRepositoryImpl) HasAnyRole(userId int64, roleNames []string) (bool, error) {
	if len(roleNames) == 0 {
		return true, nil // If no roles are specified, return false
	}

	placeholders := strings.Repeat("?,", len(roleNames))
	placeholders = placeholders[:len(placeholders)-1]
	query := fmt.Sprintf("SELECT COUNT(*) > 0 FROM user_roles ur INNER JOIN roles r ON ur.role_id = r.id WHERE ur.user_id = ? AND r.name IN (%s)", placeholders)

	// Create args slice with userId first, then all roleNames
	args := make([]interface{}, 0, 1+len(roleNames))
	args = append(args, userId)
	for _, roleName := range roleNames {
		args = append(args, roleName)
	}

	row := r.db.QueryRow(query, args...)

	var hasAnyRole bool
	if err := row.Scan(&hasAnyRole); err != nil {
		if err == sql.ErrNoRows {
			return false, nil // No roles found for the user
		}
		return false, err // Return any other error
	}

	return hasAnyRole, nil
}
