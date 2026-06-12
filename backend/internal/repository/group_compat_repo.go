package repository

import (
	"context"
	"database/sql"

	dbent "github.com/Wei-Shaw/sub2api/ent"
	"github.com/Wei-Shaw/sub2api/internal/pkg/pagination"
	"github.com/Wei-Shaw/sub2api/internal/service"
)

type groupRepositoryCompat struct{}

func NewGroupRepository(_ *dbent.Client, _ *sql.DB) service.GroupRepository {
	return groupRepositoryCompat{}
}

func (groupRepositoryCompat) ListWithFilters(ctx context.Context, params pagination.PaginationParams, platform, status, search string, isExclusive *bool) ([]service.Group, *pagination.PaginationResult, error) {
	return nil, &pagination.PaginationResult{Page: params.Page, PageSize: params.PageSize}, nil
}

func (groupRepositoryCompat) ListActive(ctx context.Context) ([]service.Group, error) {
	return nil, nil
}

func (groupRepositoryCompat) ListActiveByPlatform(ctx context.Context, platform string) ([]service.Group, error) {
	return nil, nil
}

func (groupRepositoryCompat) GetByID(ctx context.Context, id int64) (*service.Group, error) {
	return nil, service.ErrGroupNotFound
}

func (groupRepositoryCompat) GetByIDLite(ctx context.Context, id int64) (*service.Group, error) {
	return nil, service.ErrGroupNotFound
}

func (groupRepositoryCompat) Create(ctx context.Context, group *service.Group) error {
	return service.ErrGroupNotFound
}

func (groupRepositoryCompat) Update(ctx context.Context, group *service.Group) error {
	return service.ErrGroupNotFound
}

func (groupRepositoryCompat) DeleteCascade(ctx context.Context, id int64) ([]int64, error) {
	return nil, service.ErrGroupNotFound
}

func (groupRepositoryCompat) DeleteAccountGroupsByGroupID(ctx context.Context, groupID int64) (int64, error) {
	return 0, nil
}

func (groupRepositoryCompat) BindAccountsToGroup(ctx context.Context, groupID int64, accountIDs []int64) error {
	return nil
}

func (groupRepositoryCompat) GetAccountIDsByGroupIDs(ctx context.Context, groupIDs []int64) ([]int64, error) {
	return nil, nil
}

func (groupRepositoryCompat) UpdateSortOrders(ctx context.Context, updates []service.GroupSortOrderUpdate) error {
	return nil
}

type userGroupRateRepositoryCompat struct{}

func NewUserGroupRateRepository(_ *sql.DB) service.UserGroupRateRepository {
	return userGroupRateRepositoryCompat{}
}

func (userGroupRateRepositoryCompat) GetByUserID(ctx context.Context, userID int64) (map[int64]float64, error) {
	return nil, nil
}

func (userGroupRateRepositoryCompat) GetByGroupID(ctx context.Context, groupID int64) ([]service.UserGroupRateEntry, error) {
	return nil, nil
}

func (userGroupRateRepositoryCompat) GetRPMOverrideByUserAndGroup(ctx context.Context, userID, groupID int64) (*int, error) {
	return nil, nil
}

func (userGroupRateRepositoryCompat) SyncUserGroupRates(ctx context.Context, userID int64, rates map[int64]*float64) error {
	return nil
}

func (userGroupRateRepositoryCompat) DeleteByGroupID(ctx context.Context, groupID int64) error {
	return nil
}

func (userGroupRateRepositoryCompat) SyncGroupRateMultipliers(ctx context.Context, groupID int64, entries []service.GroupRateMultiplierInput) error {
	return nil
}

func (userGroupRateRepositoryCompat) ClearGroupRPMOverrides(ctx context.Context, groupID int64) error {
	return nil
}

func (userGroupRateRepositoryCompat) SyncGroupRPMOverrides(ctx context.Context, groupID int64, entries []service.GroupRPMOverrideInput) error {
	return nil
}
