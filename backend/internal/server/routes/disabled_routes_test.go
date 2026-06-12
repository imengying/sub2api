//go:build unit

package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Wei-Shaw/sub2api/internal/handler"
	servermiddleware "github.com/Wei-Shaw/sub2api/internal/server/middleware"
	"github.com/Wei-Shaw/sub2api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestDisabledOAuthRoutesAreNotRegistered(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.Use(gin.Recovery())
	v1 := r.Group("/api/v1")
	RegisterAuthRoutes(v1, &handler.Handlers{}, passJWT, nil, nil)

	for _, tc := range []struct {
		method string
		path   string
	}{
		{http.MethodGet, "/api/v1/auth/oauth/linuxdo/start"},
		{http.MethodGet, "/api/v1/auth/oauth/linuxdo/bind/start"},
		{http.MethodGet, "/api/v1/auth/oauth/linuxdo/callback"},
		{http.MethodGet, "/api/v1/auth/oauth/github/start"},
		{http.MethodGet, "/api/v1/auth/oauth/github/callback"},
		{http.MethodGet, "/api/v1/auth/oauth/google/start"},
		{http.MethodGet, "/api/v1/auth/oauth/google/callback"},
		{http.MethodGet, "/api/v1/auth/oauth/wechat/start"},
		{http.MethodGet, "/api/v1/auth/oauth/wechat/bind/start"},
		{http.MethodGet, "/api/v1/auth/oauth/wechat/callback"},
		{http.MethodGet, "/api/v1/auth/oauth/oidc/start"},
		{http.MethodGet, "/api/v1/auth/oauth/oidc/bind/start"},
		{http.MethodGet, "/api/v1/auth/oauth/oidc/callback"},
		{http.MethodGet, "/api/v1/auth/oauth/dingtalk/start"},
		{http.MethodGet, "/api/v1/auth/oauth/dingtalk/bind/start"},
		{http.MethodGet, "/api/v1/auth/oauth/dingtalk/callback"},
		{http.MethodPost, "/api/v1/auth/oauth/bind-token"},
	} {
		w := httptest.NewRecorder()
		r.ServeHTTP(w, httptest.NewRequest(tc.method, tc.path, nil))
		require.Equal(t, http.StatusNotFound, w.Code, tc.method+" "+tc.path)
	}
}

func TestDisabledAnnouncementRoutesAreNotRegistered(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.Use(gin.Recovery())
	v1 := r.Group("/api/v1")
	RegisterUserRoutes(v1, &handler.Handlers{}, passJWT, nil)
	RegisterAdminRoutes(v1, &handler.Handlers{Admin: &handler.AdminHandlers{}}, passAdmin, nil)

	for _, tc := range []struct {
		method string
		path   string
	}{
		{http.MethodGet, "/api/v1/announcements"},
		{http.MethodPost, "/api/v1/announcements/1/read"},
		{http.MethodGet, "/api/v1/admin/announcements"},
		{http.MethodPost, "/api/v1/admin/announcements"},
		{http.MethodGet, "/api/v1/admin/announcements/1"},
		{http.MethodPut, "/api/v1/admin/announcements/1"},
		{http.MethodDelete, "/api/v1/admin/announcements/1"},
		{http.MethodGet, "/api/v1/admin/announcements/1/read-status"},
	} {
		w := httptest.NewRecorder()
		r.ServeHTTP(w, httptest.NewRequest(tc.method, tc.path, nil))
		require.Equal(t, http.StatusNotFound, w.Code, tc.method+" "+tc.path)
	}
}

func TestDisabledGeminiAntigravityAdminRoutesAreNotRegistered(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.Use(gin.Recovery())
	v1 := r.Group("/api/v1")
	RegisterAdminRoutes(v1, &handler.Handlers{Admin: &handler.AdminHandlers{}}, passAdmin, nil)

	for _, tc := range []struct {
		method string
		path   string
	}{
		{http.MethodPost, "/api/v1/admin/gemini/oauth/auth-url"},
		{http.MethodPost, "/api/v1/admin/gemini/oauth/exchange-code"},
		{http.MethodGet, "/api/v1/admin/gemini/oauth/capabilities"},
		{http.MethodPost, "/api/v1/admin/antigravity/oauth/auth-url"},
		{http.MethodPost, "/api/v1/admin/antigravity/oauth/exchange-code"},
		{http.MethodPost, "/api/v1/admin/antigravity/oauth/refresh-token"},
		{http.MethodGet, "/api/v1/admin/accounts/antigravity/default-model-mapping"},
	} {
		w := httptest.NewRecorder()
		r.ServeHTTP(w, httptest.NewRequest(tc.method, tc.path, nil))
		require.Equal(t, http.StatusNotFound, w.Code, tc.method+" "+tc.path)
	}
}

func passJWT(c *gin.Context) {
	c.Set(string(servermiddleware.ContextKeyUser), servermiddleware.AuthSubject{UserID: 1})
	c.Set(string(servermiddleware.ContextKeyUserRole), service.RoleUser)
	c.Next()
}

func passAdmin(c *gin.Context) {
	c.Set(string(servermiddleware.ContextKeyUser), servermiddleware.AuthSubject{UserID: 1})
	c.Set(string(servermiddleware.ContextKeyUserRole), service.RoleAdmin)
	c.Next()
}
