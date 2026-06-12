package service

import (
	"context"
	"time"

	"github.com/Wei-Shaw/sub2api/internal/config"
	infraerrors "github.com/Wei-Shaw/sub2api/internal/pkg/errors"
	"github.com/Wei-Shaw/sub2api/internal/pkg/pagination"
)

// MaxExpiresAt is kept for migration and API compatibility.
var MaxExpiresAt = time.Date(2099, 12, 31, 23, 59, 59, 0, time.UTC)

const MaxValidityDays = 36500

var (
	ErrSubscriptionNotFound       = infraerrors.NotFound("SUBSCRIPTION_NOT_FOUND", "subscription not found")
	ErrSubscriptionExpired        = infraerrors.Forbidden("SUBSCRIPTION_EXPIRED", "subscription has expired")
	ErrSubscriptionSuspended      = infraerrors.Forbidden("SUBSCRIPTION_SUSPENDED", "subscription is suspended")
	ErrSubscriptionAlreadyExists  = infraerrors.Conflict("SUBSCRIPTION_ALREADY_EXISTS", "subscription already exists")
	ErrSubscriptionAssignConflict = infraerrors.Conflict("SUBSCRIPTION_ASSIGN_CONFLICT", "subscription assignment is disabled")
	ErrGroupNotSubscriptionType   = infraerrors.BadRequest("GROUP_NOT_SUBSCRIPTION_TYPE", "subscription groups are disabled")
	ErrInvalidInput               = infraerrors.BadRequest("INVALID_INPUT", "invalid subscription input")
	ErrDailyLimitExceeded         = infraerrors.TooManyRequests("DAILY_LIMIT_EXCEEDED", "daily usage limit exceeded")
	ErrWeeklyLimitExceeded        = infraerrors.TooManyRequests("WEEKLY_LIMIT_EXCEEDED", "weekly usage limit exceeded")
	ErrMonthlyLimitExceeded       = infraerrors.TooManyRequests("MONTHLY_LIMIT_EXCEEDED", "monthly usage limit exceeded")
	ErrSubscriptionNilInput       = infraerrors.BadRequest("SUBSCRIPTION_NIL_INPUT", "subscription input cannot be nil")
	ErrAdjustWouldExpire          = infraerrors.BadRequest("ADJUST_WOULD_EXPIRE", "adjustment would result in expired subscription")
)

// SubscriptionService is retained as a no-op compatibility shell while paid plans are removed.
type SubscriptionService struct {
	userSubRepo         UserSubscriptionRepository
	billingCacheService *BillingCacheService
	maintenanceQueue   *SubscriptionMaintenanceQueue
}

func NewSubscriptionService(args ...any) *SubscriptionService {
	var userSubRepo UserSubscriptionRepository
	var billingCacheService *BillingCacheService
	var cfg *config.Config
	for _, arg := range args {
		switch v := arg.(type) {
		case UserSubscriptionRepository:
			userSubRepo = v
		case *BillingCacheService:
			billingCacheService = v
		case *config.Config:
			cfg = v
		}
	}
	svc := &SubscriptionService{
		userSubRepo:         userSubRepo,
		billingCacheService: billingCacheService,
	}
	svc.initMaintenanceQueue(cfg)
	return svc
}

func (s *SubscriptionService) initMaintenanceQueue(cfg *config.Config) {
	if cfg == nil {
		return
	}
	mc := cfg.SubscriptionMaintenance
	if mc.WorkerCount <= 0 || mc.QueueSize <= 0 {
		return
	}
	s.maintenanceQueue = NewSubscriptionMaintenanceQueue(mc.WorkerCount, mc.QueueSize)
}

func (s *SubscriptionService) Stop() {
	if s != nil && s.maintenanceQueue != nil {
		s.maintenanceQueue.Stop()
	}
}

func (s *SubscriptionService) InvalidateSubCache(userID, groupID int64) {
	if s != nil && s.billingCacheService != nil {
		_ = s.billingCacheService.InvalidateSubscription(context.Background(), userID, groupID)
	}
}

type AssignSubscriptionInput struct {
	UserID       int64
	GroupID      int64
	ValidityDays int
	AssignedBy   int64
	Notes        string
}

func (s *SubscriptionService) AssignSubscription(ctx context.Context, input *AssignSubscriptionInput) (*UserSubscription, error) {
	return nil, ErrSubscriptionNotFound
}

func (s *SubscriptionService) AssignOrExtendSubscription(ctx context.Context, input *AssignSubscriptionInput) (*UserSubscription, bool, error) {
	return nil, false, nil
}

type BulkAssignSubscriptionInput struct {
	UserIDs      []int64
	GroupID      int64
	ValidityDays int
	AssignedBy   int64
	Notes        string
}

type BulkAssignResult struct {
	SuccessCount  int
	CreatedCount  int
	ReusedCount   int
	FailedCount   int
	Subscriptions []UserSubscription
	Errors        []string
	Statuses      map[int64]string
}

func (s *SubscriptionService) BulkAssignSubscription(ctx context.Context, input *BulkAssignSubscriptionInput) (*BulkAssignResult, error) {
	return &BulkAssignResult{Statuses: map[int64]string{}}, nil
}

func (s *SubscriptionService) RevokeSubscription(ctx context.Context, subscriptionID int64) error {
	return ErrSubscriptionNotFound
}

func (s *SubscriptionService) ExtendSubscription(ctx context.Context, subscriptionID int64, days int) (*UserSubscription, error) {
	return nil, ErrSubscriptionNotFound
}

func (s *SubscriptionService) GetByID(ctx context.Context, id int64) (*UserSubscription, error) {
	return nil, ErrSubscriptionNotFound
}

func (s *SubscriptionService) GetActiveSubscription(ctx context.Context, userID, groupID int64) (*UserSubscription, error) {
	return nil, ErrSubscriptionNotFound
}

func (s *SubscriptionService) ListUserSubscriptions(ctx context.Context, userID int64) ([]UserSubscription, error) {
	return nil, nil
}

func (s *SubscriptionService) ListActiveUserSubscriptions(ctx context.Context, userID int64) ([]UserSubscription, error) {
	return nil, nil
}

func (s *SubscriptionService) ListGroupSubscriptions(ctx context.Context, groupID int64, page, pageSize int) ([]UserSubscription, *pagination.PaginationResult, error) {
	return nil, &pagination.PaginationResult{Page: page, PageSize: pageSize}, nil
}

func (s *SubscriptionService) List(ctx context.Context, page, pageSize int, userID, groupID *int64, status, platform, sortBy, sortOrder string) ([]UserSubscription, *pagination.PaginationResult, error) {
	return nil, &pagination.PaginationResult{Page: page, PageSize: pageSize}, nil
}

func (s *SubscriptionService) CheckAndActivateWindow(ctx context.Context, sub *UserSubscription) error {
	return nil
}

func (s *SubscriptionService) AdminResetQuota(ctx context.Context, subscriptionID int64, resetDaily, resetWeekly, resetMonthly bool) (*UserSubscription, error) {
	return nil, ErrSubscriptionNotFound
}

func (s *SubscriptionService) CheckAndResetWindows(ctx context.Context, sub *UserSubscription) error {
	return nil
}

func (s *SubscriptionService) CheckUsageLimits(ctx context.Context, sub *UserSubscription, additionalCost float64) error {
	return nil
}

func (s *SubscriptionService) ValidateAndCheckLimits(sub *UserSubscription) (bool, error) {
	return false, nil
}

func (s *SubscriptionService) DoWindowMaintenance(sub *UserSubscription) {}

func (s *SubscriptionService) RecordUsage(ctx context.Context, subscriptionID int64, costUSD float64) error {
	return nil
}

type SubscriptionProgress struct {
	ID            int64                `json:"id"`
	GroupName     string               `json:"group_name"`
	ExpiresAt     time.Time            `json:"expires_at"`
	ExpiresInDays int                  `json:"expires_in_days"`
	Daily         *UsageWindowProgress `json:"daily,omitempty"`
	Weekly        *UsageWindowProgress `json:"weekly,omitempty"`
	Monthly       *UsageWindowProgress `json:"monthly,omitempty"`
}

type UsageWindowProgress struct {
	LimitUSD        float64   `json:"limit_usd"`
	UsedUSD         float64   `json:"used_usd"`
	RemainingUSD    float64   `json:"remaining_usd"`
	Percentage      float64   `json:"percentage"`
	WindowStart     time.Time `json:"window_start"`
	ResetsAt        time.Time `json:"resets_at"`
	ResetsInSeconds int64     `json:"resets_in_seconds"`
}

func (s *SubscriptionService) GetSubscriptionProgress(ctx context.Context, subscriptionID int64) (*SubscriptionProgress, error) {
	return nil, ErrSubscriptionNotFound
}

func (s *SubscriptionService) GetUserSubscriptionsWithProgress(ctx context.Context, userID int64) ([]SubscriptionProgress, error) {
	return nil, nil
}

func (s *SubscriptionService) ValidateSubscription(ctx context.Context, sub *UserSubscription) error {
	return ErrSubscriptionInvalid
}
