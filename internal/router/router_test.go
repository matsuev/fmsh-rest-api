package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_handleGetHello(t *testing.T) {
	testCases := []struct {
		name   string
		code   int
		result string
	}{
		{
			name:   "valid",
			code:   http.StatusOK,
			result: "Hello!",
		},
	}

	r := setupRouter("default")

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/hello", nil)

			r.ServeHTTP(rec, req)

			assert.Equal(t, tc.code, rec.Code)
			assert.Equal(t, tc.result, rec.Body.String())
		})
	}
}

// func Benchmark_handleGetHello(b *testing.B) {
// 	r := setupRouter("gin")

// 	rec := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)

// 	for i := 0; i < b.N; i++ {
// 		r.ServeHTTP(rec, req)
// 	}
// }
