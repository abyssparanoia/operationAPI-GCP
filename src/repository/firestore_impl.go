package repository

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/abyssparanoia/operationAPI-GCP/src/lib/log"
)

type firestore struct {
	projectID  string
	bucketName string
}

func (r *firestore) Backup(ctx context.Context, date string) error {

	path := fmt.Sprintf("gs://%s/firestore/%s", r.bucketName, date)

	err := exec.Command("gcloud", "alpha", "firestore", "export", path, "--project", r.projectID).Run()
	if err != nil {
		log.Errorf(ctx, "exec.Command :%s", err)
		return err
	}

	return nil
}

// NewFirestore ... get firestore repository
func NewFirestore(projectID string, bucketName string) Firestore {
	return &firestore{projectID, bucketName}
}
