package models

type Role struct {
	Id          int64
	Name        string
	Description string
	Created_at  string
	Updated_at  string
}

type Permission struct {
	Id          int64
	Name        string
	Description string
	Resource    string
	Action      string
	Created_at  string
	Updated_at  string
}

type RolePermissions struct {
	Id           int64
	RoleId       int64
	PermissionId int64
	Created_at   string
	Updated_at   string
}
