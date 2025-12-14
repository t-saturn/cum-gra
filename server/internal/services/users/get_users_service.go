package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"

	"github.com/google/uuid"
)

func GetUsers(page, pageSize int, isDeleted bool, status, organicUnitID, positionID *string) (*dto.UsersListResponse, error) {
	db := config.DB

	query := db.Model(&models.User{}).Where("is_deleted = ?", isDeleted)

	if status != nil && *status != "" {
		query = query.Where("status = ?", *status)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	var users []models.User
	if err := query.
		Order("created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&users).Error; err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return &dto.UsersListResponse{
			Data:     []dto.UserListItemDTO{},
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		}, nil
	}

	// Obtener user IDs
	userIDs := make([]uuid.UUID, 0, len(users))
	for _, u := range users {
		userIDs = append(userIDs, u.ID)
	}

	// Obtener user details con filtros opcionales
	detailQuery := db.Preload("StructuralPosition").
		Preload("OrganicUnit").
		Preload("Ubigeo").
		Where("user_id IN ?", userIDs)

	if organicUnitID != nil && *organicUnitID != "" {
		detailQuery = detailQuery.Where("organic_unit_id = ?", *organicUnitID)
	}

	if positionID != nil && *positionID != "" {
		detailQuery = detailQuery.Where("structural_position_id = ?", *positionID)
	}

	var userDetails []models.UserDetail
	if err := detailQuery.Find(&userDetails).Error; err != nil {
		return nil, err
	}

	detailMap := make(map[uuid.UUID]*models.UserDetail)
	for i := range userDetails {
		detailMap[userDetails[i].UserID] = &userDetails[i]
	}

	out := make([]dto.UserListItemDTO, 0, len(users))
	for _, user := range users {
		detail := detailMap[user.ID]
		out = append(out, mapper.ToUserListItemDTO(user, detail))
	}

	return &dto.UsersListResponse{
		Data:     out,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}