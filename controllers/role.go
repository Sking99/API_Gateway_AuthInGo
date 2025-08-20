package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/services"
	"AuthInGo/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type RoleController struct {
	RoleService services.RoleService
}

func NewRoleController(_roleService services.RoleService) *RoleController {
	return &RoleController{
		RoleService: _roleService,
	}
}

func (rc *RoleController) AssignRoleToUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	roleId := chi.URLParam(r, "roleId")
	if userId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "User ID is required", nil)
		return
	}

	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", nil)
		return
	}

	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	roleIdInt, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}

	err = rc.RoleService.AssignRoleToUser(userIdInt, roleIdInt)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error assigning role to user", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role assigned to user successfully", nil)
}

func (rc *RoleController) GetRoleById(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", nil)
		return
	}

	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}

	role, err := rc.RoleService.GetRoleById(id)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error fetching role by ID", err)
		return
	}

	if role == nil {
		utils.WriteJsonErrorResponse(w, http.StatusNotFound, "Role not found", nil)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role fetched successfully", role)
}

func (rc *RoleController) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := rc.RoleService.GetAllRoles()
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error fetching roles", err)
		return
	}

	if len(roles) == 0 {
		utils.WriteJsonSuccessResponse(w, http.StatusOK, "No roles found", nil)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Roles fetched successfully", roles)
}

func (rc *RoleController) CreateRole(w http.ResponseWriter, r *http.Request) {
	payload := r.Context().Value("payload").(dto.CreateRoleRequestDTO)

	role, err := rc.RoleService.CreateRole(payload.Name, payload.Description)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error creating role", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusCreated, "Role created successfully", role)
}

func (rc *RoleController) UpdateRole(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", nil)
		return
	}

	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}

	payload := r.Context().Value("payload").(dto.UpdateRoleRequestDTO)

	role, err := rc.RoleService.UpdateRole(id, payload.Name, payload.Description)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error updating role", err)
		return
	}

	if role == nil {
		utils.WriteJsonErrorResponse(w, http.StatusNotFound, "Role not found", nil)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role updated successfully", role)
}

func (rc *RoleController) DeleteRole(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", nil)
		return
	}

	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}

	err = rc.RoleService.DeleteRole(id)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error deleting role", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role deleted successfully", nil)
}

func (rc *RoleController) GetRolePermissions(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", nil)
		return
	}

	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}

	permissions, err := rc.RoleService.GetRolePermissions(id)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error fetching role permissions", err)
		return
	}

	if len(permissions) == 0 {
		utils.WriteJsonSuccessResponse(w, http.StatusOK, "No permissions found for this role", nil)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role permissions fetched successfully", permissions)
}

func (rc *RoleController) AssignPermissionToRole(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", nil)
		return
	}

	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}

	payload := r.Context().Value("payload").(dto.AssignPermissionRequestDTO)

	err = rc.RoleService.AddPermissionToRole(id, payload.PermissionId)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error assigning permission to role", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Permission assigned to role successfully", nil)
}

func (rc *RoleController) RemovePermissionFromRole(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", nil)
		return
	}

	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}

	payload := r.Context().Value("payload").(dto.RemovePermissionRequestDTO)

	err = rc.RoleService.RemovePermissionFromRole(id, payload.PermissionId)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error removing permission from role", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Permission removed from role successfully", nil)
}

func (rc *RoleController) GetAllRolePermissions(w http.ResponseWriter, r *http.Request) {
	rolePermissions, err := rc.RoleService.GetAllRolePermissions()
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error fetching all role permissions", err)
		return
	}

	if len(rolePermissions) == 0 {
		utils.WriteJsonSuccessResponse(w, http.StatusOK, "No role permissions found", nil)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "All role permissions fetched successfully", rolePermissions)
}
