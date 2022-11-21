package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// setupRouter function
func setupRouter(mode string) http.Handler {
	cfg := &Config{
		Mode:      mode,
		LogEnable: false,
	}

	return New(cfg)
}

// Test_handleGetHello function
func Test_handleGetHello(t *testing.T) {
	testCases := []struct {
		name   string
		code   int
		mode   string
		result string
	}{
		{
			name:   "default valid",
			mode:   "default",
			code:   http.StatusOK,
			result: "Hello!",
		},
		{
			name:   "gin valid",
			mode:   "gin",
			code:   http.StatusOK,
			result: "Hello!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := setupRouter(tc.mode)

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/hello", nil)

			r.ServeHTTP(rec, req)

			assert.Equal(t, tc.code, rec.Code)
			assert.Equal(t, tc.result, rec.Body.String())
		})
	}
}

// Test_handlePostHello function
func Test_handlePostHello(t *testing.T) {
	testCases := []struct {
		name   string
		mode   string
		code   int
		result string
	}{
		{
			name:   "default valid",
			mode:   "default",
			code:   http.StatusOK,
			result: "Hello! post request",
		},
		{
			name:   "gin valid",
			mode:   "gin",
			code:   http.StatusOK,
			result: "Hello! post request",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := setupRouter(tc.mode)

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/hello", nil)

			r.ServeHTTP(rec, req)

			assert.Equal(t, tc.code, rec.Code)
			assert.Equal(t, tc.result, rec.Body.String())
		})
	}
}

// Benchmark_handleGetHello function
func Benchmark_handleGetHello(b *testing.B) {
	r := setupRouter(MODE_GIN)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)

	for i := 0; i < b.N; i++ {
		r.ServeHTTP(rec, req)
	}
}
