package repositories

import (
	"context"
	"watcharis/migrate-profile-loadtest/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	GetProfile(ctx context.Context, id int, limit int) ([]models.Profile, error)
	InsertProfileAnnouncement(ctx context.Context, profiles []models.Profile) error
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{
		db: db,
	}
}

func (r *profileRepository) GetProfile(ctx context.Context, id int, limit int) ([]models.Profile, error) {
	var profiles []models.Profile
	if err := r.db.WithContext(ctx).Table(`_del_profiles_ p`).Where("p.id > ?", id).Order("p.id asc").Limit(limit).Find(&profiles).Error; err != nil {
		return nil, err
	}
	return profiles, nil
}

func (r *profileRepository) InsertProfileAnnouncement(ctx context.Context, profiles []models.Profile) error {
	if err := r.db.WithContext(ctx).Table(`profile_announcements`).Create(&profiles).Error; err != nil {
		return err
	}
	return nil
}
