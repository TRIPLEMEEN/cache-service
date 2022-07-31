package main

import (
	"context"
	"errors"
)

type bytes struct {
	New string
}

type API int

var database []bytes

// Set Method response
func Set(ctx context.Context, key string, value []bytes) error {
	return errors.New("not implemented yet. Al-ameen will implement me")
}

/// Get Method response
func Get(ctx context.Context, key string) ([]bytes, error) {
	return []bytes("Al-ameen will implement me"), nil
}
