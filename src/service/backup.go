package service

import "context"

// Backup ... backup interface
type Backup interface {
	Firestore(ctx context.Context) error
}
