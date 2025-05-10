package utils

import (
	"app/modules/club/domain/config"
	"context"
)

// GetContext is a helper function to get a context with a timeout.
func GetContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, config.DefaultTimeout)
}
