package api

import (
	"net/http"

	"github.com/abyssparanoia/operationAPI-GCP/src/lib/renderer"
	"github.com/abyssparanoia/operationAPI-GCP/src/service"
)

// BackupHandler ... backup handler
type BackupHandler struct {
	buSvc service.Backup
}

// Firestore ... backup firestore data
func (h *BackupHandler) Firestore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := h.buSvc.Firestore(ctx)
	if err != nil {
		renderer.HandleError(ctx, w, "h.buSvc.Firestore", err)
		return
	}

	renderer.Success(ctx, w)
}

// NewBackupHandler ... get backup handler
func NewBackupHandler(buSvc service.Backup) *BackupHandler {
	return &BackupHandler{buSvc}
}
