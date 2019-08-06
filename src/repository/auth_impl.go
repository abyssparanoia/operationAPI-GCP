package repository

import (
	"context"
	"os/exec"

	"github.com/abyssparanoia/operationAPI-GCP/src/lib/log"
)

type auth struct {
	credentialPath string
}

func (r auth) Activate(ctx context.Context) error {

	err := exec.Command("gcloud", "auth", "activate-service-account", "--key-file", r.credentialPath).Run()
	if err != nil {
		log.Errorf(ctx, "exec.Command :%s", err)
		return err
	}

	return nil
}

// NewAuth ... get auth repository
func NewAuth(credentialPath string) Auth {
	return &auth{credentialPath}
}
