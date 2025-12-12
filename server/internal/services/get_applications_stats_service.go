package services

// import (
// 	"server/internal/dto"
// 	"server/internal/models"

// 	"gorm.io/gorm"
// )

// func GetApplicationsStats(db *gorm.DB) (*dto.ApplicationsStatsResponse, error) {
// 	var totalApps int64
// 	if err := db.Model(&models.Application{}).
// 		Count(&totalApps).Error; err != nil {
// 		return nil, err
// 	}

// 	var activeApps int64
// 	if err := db.Model(&models.Application{}).
// 		Where("is_deleted = ? AND status = ?", false, "active").
// 		Count(&activeApps).Error; err != nil {
// 		return nil, err
// 	}

// 	var totalUsers int64
// 	sub := db.Table("user_application_roles uar").
// 		Joins("JOIN users u ON u.id = uar.user_id").
// 		Where("u.is_deleted = FALSE AND uar.is_deleted = FALSE AND uar.revoked_at IS NULL").
// 		Select("DISTINCT u.id")

// 	if err := db.Table("(?) as distinct_users", sub).
// 		Count(&totalUsers).Error; err != nil {
// 		return nil, err
// 	}

// 	return &dto.ApplicationsStatsResponse{
// 		TotalApplications:  totalApps,
// 		ActiveApplications: activeApps,
// 		TotalUsers:         totalUsers,
// 	}, nil
// }
