package services

import "context"

type MigrateProfileAnnouncementService interface {
	ProcessMigrateProfileAnnouncement(ctx context.Context) error
}
