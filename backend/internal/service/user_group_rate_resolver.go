package service

import (
	"context"
	"fmt"
	"time"

	gocache "github.com/patrickmn/go-cache"
	"golang.org/x/sync/singleflight"
)

type userGroupRateResolver struct {
	repo     UserGroupRateRepository
	cache    *gocache.Cache
	ttl      time.Duration
	sf       *singleflight.Group
	logScope string
}

func newUserGroupRateResolver(repo UserGroupRateRepository, cache *gocache.Cache, ttl time.Duration, sf *singleflight.Group, logScope string) *userGroupRateResolver {
	return &userGroupRateResolver{
		repo:     repo,
		cache:    cache,
		ttl:      ttl,
		sf:       sf,
		logScope: logScope,
	}
}

func (r *userGroupRateResolver) Resolve(ctx context.Context, userID, groupID int64, groupDefaultMultiplier float64) float64 {
	if r == nil || userID <= 0 || groupID <= 0 {
		return groupDefaultMultiplier
	}

	key := userGroupRateCacheKey(userID, groupID)
	if r.cache != nil {
		if cached, ok := r.cache.Get(key); ok {
			if multiplier, ok := cached.(float64); ok {
				userGroupRateCacheHitTotal.Add(1)
				return multiplier
			}
		}
		userGroupRateCacheMissTotal.Add(1)
	}

	if r.repo == nil {
		userGroupRateCacheFallbackTotal.Add(1)
		return groupDefaultMultiplier
	}

	load := func() (float64, error) {
		userGroupRateCacheLoadTotal.Add(1)
		multiplier, err := r.repo.GetByUserAndGroup(ctx, userID, groupID)
		if err != nil {
			return groupDefaultMultiplier, err
		}
		resolved := groupDefaultMultiplier
		if multiplier != nil && *multiplier > 0 {
			resolved = *multiplier
		}
		if r.cache != nil {
			expiration := r.ttl
			if expiration <= 0 {
				expiration = gocache.DefaultExpiration
			}
			r.cache.Set(key, resolved, expiration)
		}
		return resolved, nil
	}

	if r.sf == nil {
		resolved, err := load()
		if err != nil {
			userGroupRateCacheFallbackTotal.Add(1)
			return groupDefaultMultiplier
		}
		return resolved
	}

	value, err, shared := r.sf.Do(key, func() (any, error) {
		return load()
	})
	if shared {
		userGroupRateCacheSFSharedTotal.Add(1)
	}
	if err != nil {
		userGroupRateCacheFallbackTotal.Add(1)
		return groupDefaultMultiplier
	}
	resolved, ok := value.(float64)
	if !ok {
		userGroupRateCacheFallbackTotal.Add(1)
		return groupDefaultMultiplier
	}
	return resolved
}

func userGroupRateCacheKey(userID, groupID int64) string {
	return fmt.Sprintf("%d:%d", userID, groupID)
}
