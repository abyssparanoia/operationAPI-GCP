package main

import (
	"github.com/abyssparanoia/operationAPI-GCP/src/handler/api"
	"github.com/abyssparanoia/operationAPI-GCP/src/lib/deploy"
	"github.com/abyssparanoia/operationAPI-GCP/src/lib/log"
	"github.com/abyssparanoia/operationAPI-GCP/src/repository"
	"github.com/abyssparanoia/operationAPI-GCP/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	Log           *log.Middleware
	BackupHandler *api.BackupHandler
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject(e *Environment) {

	var lCli log.Writer
	if deploy.IsLocal() {
		lCli = log.NewWriterStdout()
	} else {
		lCli = log.NewWriterStackdriver(e.ProjectID)
	}

	// Repository
	authRepo := repository.NewAuth(e.CredentialsPath)
	fsRepo := repository.NewFirestore(e.ProjectID, e.BackupBucketName)

	buSvc := service.NewBackup(authRepo, fsRepo)

	// Middleware
	d.Log = log.NewMiddleware(lCli, e.MinLogSeverity)

	// Handler
	d.BackupHandler = api.NewBackupHandler(buSvc)
}
