package repository

import "context"

// Firestore ... firestore interface
type Firestore interface {
	Backup(ctx context.Context, date string) error
}
