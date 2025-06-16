package handlers

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/server/internal/models"
	"github.com/t-saturn/central-user-manager/server/internal/services"
	userpb "github.com/t-saturn/central-user-manager/server/pb/proto/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserHandler maneja las requests gRPC para usuarios
type UserHandler struct {
	userpb.UnimplementedUserServiceServer
	service services.UserService
}

// NewUserHandler crea una nueva instancia del handler
func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// CreateUser crea un nuevo usuario
func (h *UserHandler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	// Validar request
	if req.Email == "" {
		return nil, errors.New("email is required")
	}
	if req.PasswordHash == "" {
		return nil, errors.New("password hash is required")
	}

	// Convertir request
	serviceReq := services.CreateUserRequest{
		Email:        req.Email,
		PasswordHash: req.PasswordHash,
	}

	// Campos opcionales
	if req.FirstName != "" {
		serviceReq.FirstName = &req.FirstName
	}
	if req.LastName != "" {
		serviceReq.LastName = &req.LastName
	}
	if req.Phone != "" {
		serviceReq.Phone = &req.Phone
	}
	if req.StructuralPositionId != "" {
		if id, err := uuid.Parse(req.StructuralPositionId); err == nil {
			serviceReq.StructuralPositionID = &id
		}
	}
	if req.OrganicUnitId != "" {
		if id, err := uuid.Parse(req.OrganicUnitId); err == nil {
			serviceReq.OrganicUnitID = &id
		}
	}

	// Llamar al service
	user, err := h.service.CreateUser(serviceReq)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Convertir respuesta
	return &userpb.CreateUserResponse{
		User:    h.modelToProto(user),
		Message: "User created successfully",
	}, nil
}

// GetUser obtiene un usuario por ID
func (h *UserHandler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	if req.Id == "" {
		return nil, errors.New("user ID is required")
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &userpb.GetUserResponse{
		User: h.modelToProto(user),
	}, nil
}

// GetUserByEmail obtiene un usuario por email
func (h *UserHandler) GetUserByEmail(ctx context.Context, req *userpb.GetUserByEmailRequest) (*userpb.GetUserResponse, error) {
	if req.Email == "" {
		return nil, errors.New("email is required")
	}

	user, err := h.service.GetUserByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return &userpb.GetUserResponse{
		User: h.modelToProto(user),
	}, nil
}

// UpdateUser actualiza un usuario
func (h *UserHandler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	if req.Id == "" {
		return nil, errors.New("user ID is required")
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	// Convertir request
	serviceReq := services.UpdateUserRequest{
		ID: id,
	}

	// Campos opcionales
	if req.Email != "" {
		serviceReq.Email = &req.Email
	}
	if req.FirstName != "" {
		serviceReq.FirstName = &req.FirstName
	}
	if req.LastName != "" {
		serviceReq.LastName = &req.LastName
	}
	if req.Phone != "" {
		serviceReq.Phone = &req.Phone
	}
	if req.EmailVerified {
		serviceReq.EmailVerified = &req.EmailVerified
	}
	if req.PhoneVerified {
		serviceReq.PhoneVerified = &req.PhoneVerified
	}
	if req.TwoFactorEnabled {
		serviceReq.TwoFactorEnabled = &req.TwoFactorEnabled
	}
	if req.Status != userpb.UserStatus_USER_STATUS_UNSPECIFIED {
		status := h.protoToModelStatus(req.Status)
		serviceReq.Status = &status
	}
	if req.StructuralPositionId != "" {
		positionID, parseErr := uuid.Parse(req.StructuralPositionId)
		if parseErr == nil {
			serviceReq.StructuralPositionID = &positionID
		}
	}
	if req.OrganicUnitId != "" {
		if unitID, parseErr := uuid.Parse(req.OrganicUnitId); parseErr == nil {
			serviceReq.OrganicUnitID = &unitID
		}
	}

	// Llamar al service
	user, err := h.service.UpdateUser(serviceReq)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return &userpb.UpdateUserResponse{
		User:    h.modelToProto(user),
		Message: "User updated successfully",
	}, nil
}

// DeleteUser elimina un usuario (borrado lógico)
func (h *UserHandler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	if req.Id == "" {
		return nil, errors.New("user ID is required")
	}
	if req.DeletedBy == "" {
		return nil, errors.New("deleted_by is required")
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	deletedBy, err := uuid.Parse(req.DeletedBy)
	if err != nil {
		return nil, errors.New("invalid deleted_by ID format")
	}

	err = h.service.DeleteUser(id, deletedBy)
	if err != nil {
		return nil, fmt.Errorf("failed to delete user: %w", err)
	}

	return &userpb.DeleteUserResponse{
		Message: "User deleted successfully",
	}, nil
}

// ListUsers lista usuarios con filtros y paginación
func (h *UserHandler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	serviceReq := services.ListUsersRequest{
		Page:           int(req.Page),
		PageSize:       int(req.PageSize),
		Search:         req.Search,
		IncludeDeleted: req.IncludeDeleted,
	}

	if req.Status != userpb.UserStatus_USER_STATUS_UNSPECIFIED {
		serviceReq.Status = h.protoToModelStatus(req.Status)
	}

	result, err := h.service.ListUsers(serviceReq)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	// Convertir usuarios
	users := make([]*userpb.User, len(result.Users))
	for i, user := range result.Users {
		users[i] = h.modelToProto(&user)
	}

	return &userpb.ListUsersResponse{
		Users:    users,
		Total:    int32(result.Total),
		Page:     int32(result.Page),
		PageSize: int32(result.PageSize),
	}, nil
}

// UpdatePassword actualiza la contraseña de un usuario
func (h *UserHandler) UpdatePassword(ctx context.Context, req *userpb.UpdatePasswordRequest) (*userpb.UpdatePasswordResponse, error) {
	if req.Id == "" {
		return nil, errors.New("user ID is required")
	}
	if req.NewPasswordHash == "" {
		return nil, errors.New("new password hash is required")
	}
	if req.ChangedBy == "" {
		return nil, errors.New("changed_by is required")
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	changedBy, err := uuid.Parse(req.ChangedBy)
	if err != nil {
		return nil, errors.New("invalid changed_by ID format")
	}

	err = h.service.UpdatePassword(id, req.NewPasswordHash, changedBy)
	if err != nil {
		return nil, fmt.Errorf("failed to update password: %w", err)
	}

	return &userpb.UpdatePasswordResponse{
		Message: "Password updated successfully",
	}, nil
}

// VerifyEmail marca el email como verificado
func (h *UserHandler) VerifyEmail(ctx context.Context, req *userpb.VerifyEmailRequest) (*userpb.VerifyEmailResponse, error) {
	if req.Id == "" {
		return nil, errors.New("user ID is required")
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	err = h.service.VerifyEmail(id)
	if err != nil {
		return nil, fmt.Errorf("failed to verify email: %w", err)
	}

	return &userpb.VerifyEmailResponse{
		Message: "Email verified successfully",
	}, nil
}

// VerifyPhone marca el teléfono como verificado
func (h *UserHandler) VerifyPhone(ctx context.Context, req *userpb.VerifyPhoneRequest) (*userpb.VerifyPhoneResponse, error) {
	if req.Id == "" {
		return nil, errors.New("user ID is required")
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	err = h.service.VerifyPhone(id)
	if err != nil {
		return nil, fmt.Errorf("failed to verify phone: %w", err)
	}

	return &userpb.VerifyPhoneResponse{
		Message: "Phone verified successfully",
	}, nil
}

// EnableTwoFactor habilita la autenticación de dos factores
func (h *UserHandler) EnableTwoFactor(ctx context.Context, req *userpb.EnableTwoFactorRequest) (*userpb.EnableTwoFactorResponse, error) {
	if req.Id == "" {
		return nil, errors.New("user ID is required")
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	err = h.service.EnableTwoFactor(id)
	if err != nil {
		return nil, fmt.Errorf("failed to enable two factor: %w", err)
	}

	return &userpb.EnableTwoFactorResponse{
		Message: "Two factor authentication enabled successfully",
	}, nil
}

// DisableTwoFactor deshabilita la autenticación de dos factores
func (h *UserHandler) DisableTwoFactor(ctx context.Context, req *userpb.DisableTwoFactorRequest) (*userpb.DisableTwoFactorResponse, error) {
	if req.Id == "" {
		return nil, errors.New("user ID is required")
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	err = h.service.DisableTwoFactor(id)
	if err != nil {
		return nil, fmt.Errorf("failed to disable two factor: %w", err)
	}

	return &userpb.DisableTwoFactorResponse{
		Message: "Two factor authentication disabled successfully",
	}, nil
}

// UpdateLastLogin actualiza el timestamp del último login
func (h *UserHandler) UpdateLastLogin(ctx context.Context, req *userpb.UpdateLastLoginRequest) (*userpb.UpdateLastLoginResponse, error) {
	if req.Id == "" {
		return nil, errors.New("user ID is required")
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	err = h.service.UpdateLastLogin(id)
	if err != nil {
		return nil, fmt.Errorf("failed to update last login: %w", err)
	}

	return &userpb.UpdateLastLoginResponse{
		Message: "Last login updated successfully",
	}, nil
}

// modelToProto convierte un modelo User a protobuf User
func (h *UserHandler) modelToProto(user *models.User) *userpb.User {
	protoUser := &userpb.User{
		Id:               user.ID.String(),
		Email:            user.Email,
		PasswordHash:     user.PasswordHash,
		EmailVerified:    user.EmailVerified,
		PhoneVerified:    user.PhoneVerified,
		TwoFactorEnabled: user.TwoFactorEnabled,
		Status:           h.modelToProtoStatus(user.Status),
		CreatedAt:        timestamppb.New(user.CreatedAt),
		UpdatedAt:        timestamppb.New(user.UpdatedAt),
		IsDeleted:        user.IsDeleted,
	}

	// Campos opcionales
	if user.FirstName != nil {
		protoUser.FirstName = *user.FirstName
	}
	if user.LastName != nil {
		protoUser.LastName = *user.LastName
	}
	if user.Phone != nil {
		protoUser.Phone = *user.Phone
	}
	if user.StructuralPositionID != nil {
		protoUser.StructuralPositionId = user.StructuralPositionID.String()
	}
	if user.OrganicUnitID != nil {
		protoUser.OrganicUnitId = user.OrganicUnitID.String()
	}
	if user.LastLoginAt != nil {
		protoUser.LastLoginAt = timestamppb.New(*user.LastLoginAt)
	}
	if user.DeletedAt != nil {
		protoUser.DeletedAt = timestamppb.New(*user.DeletedAt)
	}
	if user.DeletedBy != nil {
		protoUser.DeletedBy = user.DeletedBy.String()
	}

	return protoUser
}

// modelToProtoStatus convierte UserStatus del modelo a protobuf
func (h *UserHandler) modelToProtoStatus(status models.UserStatus) userpb.UserStatus {
	switch status {
	case models.UserStatusActive:
		return userpb.UserStatus_USER_STATUS_ACTIVE
	case models.UserStatusSuspended:
		return userpb.UserStatus_USER_STATUS_SUSPENDED
	case models.UserStatusDeleted:
		return userpb.UserStatus_USER_STATUS_DELETED
	default:
		return userpb.UserStatus_USER_STATUS_UNSPECIFIED
	}
}

// protoToModelStatus convierte UserStatus de protobuf al modelo
func (h *UserHandler) protoToModelStatus(status userpb.UserStatus) models.UserStatus {
	switch status {
	case userpb.UserStatus_USER_STATUS_ACTIVE:
		return models.UserStatusActive
	case userpb.UserStatus_USER_STATUS_SUSPENDED:
		return models.UserStatusSuspended
	case userpb.UserStatus_USER_STATUS_DELETED:
		return models.UserStatusDeleted
	default:
		return models.UserStatusActive
	}
}
