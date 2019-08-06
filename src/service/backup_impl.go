package service

import (
	"context"

	"github.com/abyssparanoia/operationAPI-GCP/src/lib/log"
	"github.com/abyssparanoia/operationAPI-GCP/src/lib/util"
	"github.com/abyssparanoia/operationAPI-GCP/src/repository"
)

type backup struct {
	authRepo repository.Auth
	fsRepo   repository.Firestore
}

func (s *backup) Firestore(ctx context.Context) error {
	err := s.authRepo.Activate(ctx)
	if err != nil {
		log.Errorf(ctx, "s.authRepo.Activate :%s", err)
		return err
	}

	dateStr := util.TimeTodayString()

	err = s.fsRepo.Backup(ctx, dateStr)
	if err != nil {
		log.Errorf(ctx, "s.fsRepo.Backup :%s", err)
		return err
	}

	return nil
}

// NewBackup ... get backup service
func NewBackup(authRepo repository.Auth, fsRepo repository.Firestore) Backup {
	return &backup{authRepo, fsRepo}
}
