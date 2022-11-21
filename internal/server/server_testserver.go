package server

import "context"

// TestServer ...§
type TestServer struct{}

// NewTsetServer function
func NewTestServer() *TestServer {
	return &TestServer{}
}

// Listen function
func (s *TestServer) Listen() error {
	return nil
}

// Shutdown function
func (s *TestServer) Shutdown(ctx context.Context) error {
	_ = ctx
	return nil
}
