package boot

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	"bythen-takehome/internal/config"

	"github.com/fsnotify/fsnotify"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	httpServer "bythen-takehome/internal/delivery/http"

	authData "bythen-takehome/internal/data/auth"
	authHandler "bythen-takehome/internal/delivery/http/auth"
	authService "bythen-takehome/internal/service/auth"

	blogData "bythen-takehome/internal/data/blog"
	blogHandler "bythen-takehome/internal/delivery/http/blog"
	blogService "bythen-takehome/internal/service/blog"
)

func HTTP() error {
	err := config.Init()
	if err != nil {
		log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
	}
	cfg := config.Get()

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = cfg.Database.Master
	}

	db, err := openConnectionPool("mysql", dsn)
	if err != nil {
		log.Fatalf("[DB] Failed to open mysql connection pool: %v", err)
	}

	authD := authData.New(db)
	authS := authService.New(authD)
	authH := authHandler.New(authS)

	blogD := blogData.New(db)
	blogS := blogService.New(blogD)
	blogH := blogHandler.New(blogS)

	config.PrepareWatchPath()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		err := config.Init()
		if err != nil {
			log.Printf("[VIPER] Error get config file, %v", err)
		}
		cfg = config.Get()

		dbNew, err := openConnectionPool("mysql", cfg.Database.Master)
		if err != nil {
			log.Printf("[VIPER] Error open db connection, %v", err)
		} else {
			*db = *dbNew
			blogD.InitStmt()
		}
	})

	s := httpServer.Server{
		Auth: authH,
		Blog: blogH,
	}

	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}

func openConnectionPool(driver string, connString string) (db *sqlx.DB, err error) {
	const maxRetries = 10
	const retryDelay = 3 * time.Second

	for i := 1; i <= maxRetries; i++ {
		db, err = sqlx.Open(driver, connString)
		if err == nil {
			err = db.Ping()
			if err == nil {
				log.Println("Database connected")
				return db, nil
			}
		}

		log.Printf(
			"Database not ready (attempt %d/%d): %v",
			i, maxRetries, err,
		)

		time.Sleep(retryDelay)
	}

	return nil, fmt.Errorf("Failed to connect to database after %d attempts: %w", maxRetries, err)
}
