package boot

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"bythen-takehome/internal/config"

	"github.com/fsnotify/fsnotify"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	blogServer "bythen-takehome/internal/delivery/http"

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

	db, err := openConnectionPool("mysql", cfg.Database.Master)
	if err != nil {
		log.Fatalf("[DB] Failed to open mysql connection pool: %v", err)
	}

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

	s := blogServer.Server{
		Blog: blogH,
	}

	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}

func openConnectionPool(driver string, connString string) (db *sqlx.DB, err error) {
	db, err = sqlx.Open(driver, connString)
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, err
}
