package main

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/abyssparanoia/operationAPI-GCP/src/handler"
	"github.com/abyssparanoia/operationAPI-GCP/src/lib/accesscontrol"
)

// Routing ... ルーティング設定
func Routing(r *chi.Mux, d *Dependency) {
	// ブラウザのCORS対応
	r.Use(accesscontrol.Handle)

	// ログをリクエスト単位でまとめるため、情報をContextに保持する
	r.Use(d.Log.Handle)

	// 障害検知でサーバーの生存確認のため、pingリクエストを用意する
	r.Get("/ping", handler.Ping)

	// 例: サブルーティング
	r.Route("/v1", func(r chi.Router) {
		r.Get("/backup/firestore", d.BackupHandler.Firestore)
	})

	http.Handle("/", r)
}
