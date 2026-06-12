package service

import (
	"context"
	"time"

	"github.com/Wei-Shaw/sub2api/internal/domain"
	infraerrors "github.com/Wei-Shaw/sub2api/internal/pkg/errors"
	"github.com/Wei-Shaw/sub2api/internal/pkg/pagination"
)

var ErrGroupNotFound = infraerrors.NotFound("GROUP_NOT_FOUND", "group not found")

type OpenAIMessagesDispatchModelConfig = domain.OpenAIMessagesDispatchModelConfig
type GroupModelsListConfig = domain.GroupModelsListConfig

type Group struct {
	ID                              int64
	Name                            string
	Description                     string
	Platform                        string
	RateMultiplier                  float64
	IsExclusive                     bool
	Status                          string
	Hydrated                        bool
	SubscriptionType                string
	DailyLimitUSD                   *float64
	WeeklyLimitUSD                  *float64
	MonthlyLimitUSD                 *float64
	AllowImageGeneration            bool
	ImageRateIndependent            bool
	ImageRateMultiplier             float64
	ImagePrice1K                    *float64
	ImagePrice2K                    *float64
	ImagePrice4K                    *float64
	DefaultValidityDays             int
	ClaudeCodeOnly                  bool
	FallbackGroupID                 *int64
	FallbackGroupIDOnInvalidRequest *int64
	ModelRouting                    map[string][]int64
	ModelRoutingEnabled             bool
	MCPXMLInject                    bool
	SupportedModelScopes            []string
	AccountGroups                   []AccountGroup
	AccountCount                    int64
	ActiveAccountCount              int64
	RateLimitedAccountCount         int64
	SortOrder                       int
	AllowMessagesDispatch           bool
	RequireOAuthOnly                bool
	RequirePrivacySet               bool
	DefaultMappedModel              string
	MessagesDispatchModelConfig     OpenAIMessagesDispatchModelConfig
	ModelsListConfig                GroupModelsListConfig
	RPMLimit                        int
	CreatedAt                       time.Time
	UpdatedAt                       time.Time
}

type AccountGroup struct {
	AccountID int64
	GroupID   int64
	Priority  int
	CreatedAt time.Time
	Account   *Account
	Group     *Group
}

type GroupSortOrderUpdate struct {
	ID        int64 `json:"id"`
	SortOrder int   `json:"sort_order"`
}

type UserGroupRateEntry struct {
	UserID         int64
	Username       string
	Email          string
	RateMultiplier float64
	RPMOverride    *int
}

type GroupRateMultiplierInput struct {
	UserID         int64   `json:"user_id"`
	RateMultiplier float64 `json:"rate_multiplier"`
}

type GroupRPMOverrideInput struct {
	UserID      int64 `json:"user_id"`
	RPMOverride *int  `json:"rpm_override"`
}

type GroupRepository interface {
	ListWithFilters(ctx context.Context, params pagination.PaginationParams, platform, status, search string, isExclusive *bool) ([]Group, *pagination.PaginationResult, error)
	ListActive(ctx context.Context) ([]Group, error)
	ListActiveByPlatform(ctx context.Context, platform string) ([]Group, error)
	GetByID(ctx context.Context, id int64) (*Group, error)
	GetByIDLite(ctx context.Context, id int64) (*Group, error)
	Create(ctx context.Context, group *Group) error
	Update(ctx context.Context, group *Group) error
	DeleteCascade(ctx context.Context, id int64) ([]int64, error)
	DeleteAccountGroupsByGroupID(ctx context.Context, groupID int64) (int64, error)
	BindAccountsToGroup(ctx context.Context, groupID int64, accountIDs []int64) error
	GetAccountIDsByGroupIDs(ctx context.Context, groupIDs []int64) ([]int64, error)
	UpdateSortOrders(ctx context.Context, updates []GroupSortOrderUpdate) error
}

type UserGroupRateRepository interface {
	GetByUserID(ctx context.Context, userID int64) (map[int64]float64, error)
	GetByGroupID(ctx context.Context, groupID int64) ([]UserGroupRateEntry, error)
	GetRPMOverrideByUserAndGroup(ctx context.Context, userID, groupID int64) (*int, error)
	SyncUserGroupRates(ctx context.Context, userID int64, rates map[int64]*float64) error
	DeleteByGroupID(ctx context.Context, groupID int64) error
	SyncGroupRateMultipliers(ctx context.Context, groupID int64, entries []GroupRateMultiplierInput) error
	ClearGroupRPMOverrides(ctx context.Context, groupID int64) error
	SyncGroupRPMOverrides(ctx context.Context, groupID int64, entries []GroupRPMOverrideInput) error
}

func (g *Group) IsSubscriptionType() bool {
	return g != nil && g.SubscriptionType == SubscriptionTypeSubscription
}

func (g *Group) HasDailyLimit() bool {
	return g != nil && g.DailyLimitUSD != nil
}

func (g *Group) HasWeeklyLimit() bool {
	return g != nil && g.WeeklyLimitUSD != nil
}

func (g *Group) HasMonthlyLimit() bool {
	return g != nil && g.MonthlyLimitUSD != nil
}

func (g *Group) GetRoutingAccountIDs(model string) []int64 {
	if g == nil || !g.ModelRoutingEnabled || len(g.ModelRouting) == 0 {
		return nil
	}
	ids := g.ModelRouting[model]
	if len(ids) == 0 {
		return nil
	}
	out := make([]int64, len(ids))
	copy(out, ids)
	return out
}

func (g *Group) CustomModelsListEnabled() bool {
	return g != nil && g.ModelsListConfig.Enabled && len(g.ModelsListConfig.Models) > 0
}

func IsGroupContextValid(group *Group) bool {
	return group != nil && group.Hydrated
}
