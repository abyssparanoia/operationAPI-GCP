package repository

import "context"

// Firestore ... firestore interface
type Firestore interface {
	Backup(ctx context.Context, bucketPath string) error
}
