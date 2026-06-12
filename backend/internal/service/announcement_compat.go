package service

import (
	"context"
	"time"
)

type AnnouncementTargeting map[string]any

type Announcement struct {
	ID         int64
	Title      string
	Content    string
	Status     string
	NotifyMode string
	Targeting  AnnouncementTargeting
	StartsAt   *time.Time
	EndsAt     *time.Time
	CreatedBy  *int64
	UpdatedBy  *int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type UserAnnouncement struct {
	Announcement Announcement
	ReadAt       *time.Time
}

type AnnouncementRepository interface{}
type AnnouncementReadRepository interface{}

type AnnouncementService struct{}

func NewAnnouncementService(_ AnnouncementRepository, _ AnnouncementReadRepository, _ UserRepository, _ UserSubscriptionRepository) *AnnouncementService {
	return &AnnouncementService{}
}

func (s *AnnouncementService) ListForUser(ctx context.Context, userID int64, unreadOnly bool) ([]UserAnnouncement, error) {
	return nil, nil
}

func (s *AnnouncementService) MarkRead(ctx context.Context, userID, announcementID int64) error {
	return nil
}
