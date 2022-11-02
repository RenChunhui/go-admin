package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/renchunhui/go-admin/internal/migration"
	"github.com/renchunhui/go-admin/internal/router"
	"github.com/renchunhui/go-admin/pkg/config"
	"github.com/renchunhui/go-admin/pkg/db"
)

func init() {
	err := config.NewConfig()

	if err != nil {
		return
	}

	err = db.Open()

	if err != nil {
		return
	}

	migration.New()
}

func main() {
	r := router.NewRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.Http.Port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
