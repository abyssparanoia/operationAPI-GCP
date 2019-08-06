package repository

import "context"

// Auth ... auth interface
type Auth interface {
	Activate(ctx context.Context) error
}
