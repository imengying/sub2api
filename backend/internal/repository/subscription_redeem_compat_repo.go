package repository

import (
	"context"
	"time"

	dbent "github.com/Wei-Shaw/sub2api/ent"
	"github.com/Wei-Shaw/sub2api/internal/pkg/pagination"
	"github.com/Wei-Shaw/sub2api/internal/service"
)

type announcementRepositoryCompat struct{}
type announcementReadRepositoryCompat struct{}

func NewAnnouncementRepository(_ *dbent.Client) service.AnnouncementRepository {
	return announcementRepositoryCompat{}
}

func NewAnnouncementReadRepository(_ *dbent.Client) service.AnnouncementReadRepository {
	return announcementReadRepositoryCompat{}
}

type userSubscriptionRepositoryCompat struct{}

func NewUserSubscriptionRepository(_ *dbent.Client) service.UserSubscriptionRepository {
	return userSubscriptionRepositoryCompat{}
}

func (userSubscriptionRepositoryCompat) Create(ctx context.Context, sub *service.UserSubscription) error {
	return nil
}

func (userSubscriptionRepositoryCompat) GetByID(ctx context.Context, id int64) (*service.UserSubscription, error) {
	return nil, service.ErrSubscriptionNotFound
}

func (userSubscriptionRepositoryCompat) GetByUserIDAndGroupID(ctx context.Context, userID, groupID int64) (*service.UserSubscription, error) {
	return nil, service.ErrSubscriptionNotFound
}

func (userSubscriptionRepositoryCompat) GetActiveByUserIDAndGroupID(ctx context.Context, userID, groupID int64) (*service.UserSubscription, error) {
	return nil, service.ErrSubscriptionNotFound
}

func (userSubscriptionRepositoryCompat) Update(ctx context.Context, sub *service.UserSubscription) error {
	return nil
}

func (userSubscriptionRepositoryCompat) Delete(ctx context.Context, id int64) error {
	return nil
}

func (userSubscriptionRepositoryCompat) ListByUserID(ctx context.Context, userID int64) ([]service.UserSubscription, error) {
	return nil, nil
}

func (userSubscriptionRepositoryCompat) ListActiveByUserID(ctx context.Context, userID int64) ([]service.UserSubscription, error) {
	return nil, nil
}

func (userSubscriptionRepositoryCompat) ListByGroupID(ctx context.Context, groupID int64, params pagination.PaginationParams) ([]service.UserSubscription, *pagination.PaginationResult, error) {
	return nil, &pagination.PaginationResult{Page: params.Page, PageSize: params.PageSize}, nil
}

func (userSubscriptionRepositoryCompat) List(ctx context.Context, params pagination.PaginationParams, userID, groupID *int64, status, platform, sortBy, sortOrder string) ([]service.UserSubscription, *pagination.PaginationResult, error) {
	return nil, &pagination.PaginationResult{Page: params.Page, PageSize: params.PageSize}, nil
}

func (userSubscriptionRepositoryCompat) ExistsByUserIDAndGroupID(ctx context.Context, userID, groupID int64) (bool, error) {
	return false, nil
}

func (userSubscriptionRepositoryCompat) ExtendExpiry(ctx context.Context, subscriptionID int64, newExpiresAt time.Time) error {
	return service.ErrSubscriptionNotFound
}

func (userSubscriptionRepositoryCompat) UpdateStatus(ctx context.Context, subscriptionID int64, status string) error {
	return service.ErrSubscriptionNotFound
}

func (userSubscriptionRepositoryCompat) UpdateNotes(ctx context.Context, subscriptionID int64, notes string) error {
	return service.ErrSubscriptionNotFound
}

func (userSubscriptionRepositoryCompat) ActivateWindows(ctx context.Context, id int64, start time.Time) error {
	return service.ErrSubscriptionNotFound
}

func (userSubscriptionRepositoryCompat) ResetDailyUsage(ctx context.Context, id int64, newWindowStart time.Time) error {
	return service.ErrSubscriptionNotFound
}

func (userSubscriptionRepositoryCompat) ResetWeeklyUsage(ctx context.Context, id int64, newWindowStart time.Time) error {
	return service.ErrSubscriptionNotFound
}

func (userSubscriptionRepositoryCompat) ResetMonthlyUsage(ctx context.Context, id int64, newWindowStart time.Time) error {
	return service.ErrSubscriptionNotFound
}

func (userSubscriptionRepositoryCompat) IncrementUsage(ctx context.Context, id int64, costUSD float64) error {
	return service.ErrSubscriptionNotFound
}

func (userSubscriptionRepositoryCompat) BatchUpdateExpiredStatus(ctx context.Context) (int64, error) {
	return 0, nil
}

type redeemCodeRepositoryCompat struct{}

func NewRedeemCodeRepository(_ *dbent.Client) service.RedeemCodeRepository {
	return redeemCodeRepositoryCompat{}
}

func (redeemCodeRepositoryCompat) Create(ctx context.Context, code *service.RedeemCode) error {
	return nil
}

func (redeemCodeRepositoryCompat) CreateBatch(ctx context.Context, codes []service.RedeemCode) error {
	return nil
}

func (redeemCodeRepositoryCompat) GetByID(ctx context.Context, id int64) (*service.RedeemCode, error) {
	return nil, service.ErrRedeemCodeNotFound
}

func (redeemCodeRepositoryCompat) GetByCode(ctx context.Context, code string) (*service.RedeemCode, error) {
	return nil, service.ErrRedeemCodeNotFound
}

func (redeemCodeRepositoryCompat) Update(ctx context.Context, code *service.RedeemCode) error {
	return service.ErrRedeemCodeNotFound
}

func (redeemCodeRepositoryCompat) BatchUpdate(ctx context.Context, ids []int64, fields service.RedeemCodeBatchUpdateFields) (int64, error) {
	return 0, nil
}

func (redeemCodeRepositoryCompat) Delete(ctx context.Context, id int64) error {
	return service.ErrRedeemCodeNotFound
}

func (redeemCodeRepositoryCompat) Use(ctx context.Context, id, userID int64) error {
	return service.ErrRedeemCodeNotFound
}

func (redeemCodeRepositoryCompat) List(ctx context.Context, params pagination.PaginationParams) ([]service.RedeemCode, *pagination.PaginationResult, error) {
	return nil, &pagination.PaginationResult{Page: params.Page, PageSize: params.PageSize}, nil
}

func (redeemCodeRepositoryCompat) ListWithFilters(ctx context.Context, params pagination.PaginationParams, codeType, status, search string) ([]service.RedeemCode, *pagination.PaginationResult, error) {
	return nil, &pagination.PaginationResult{Page: params.Page, PageSize: params.PageSize}, nil
}

func (redeemCodeRepositoryCompat) ListByUser(ctx context.Context, userID int64, limit int) ([]service.RedeemCode, error) {
	return nil, nil
}

func (redeemCodeRepositoryCompat) ListByUserPaginated(ctx context.Context, userID int64, params pagination.PaginationParams, codeType string) ([]service.RedeemCode, *pagination.PaginationResult, error) {
	return nil, &pagination.PaginationResult{Page: params.Page, PageSize: params.PageSize}, nil
}

func (redeemCodeRepositoryCompat) SumPositiveBalanceByUser(ctx context.Context, userID int64) (float64, error) {
	return 0, nil
}
