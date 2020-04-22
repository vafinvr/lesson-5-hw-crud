package main

import (
	"flag"
	"lesson-5-hw-crud/infrastructure"
	"lesson-5-hw-crud/infrastructure/logger"
	"lesson-5-hw-crud/infrastructure/pgdb"
	"lesson-5-hw-crud/infrastructure/webServer"
	"lesson-5-hw-crud/interfaces/config"
	"lesson-5-hw-crud/interfaces/migrations"
	"lesson-5-hw-crud/interfaces/usersRepo"
	"lesson-5-hw-crud/interfaces/webService"
	"lesson-5-hw-crud/usecases/usersInteractor"
)

func main() {
	migrateMode := flag.Bool("migrate", false, "run as migrate mode")

	log := logger.New()

	cfg, err := config.NewConfig(new(infrastructure.Env))
	if err != nil {
		log.Fatal(err.Error())
	}

	db := pgdb.New(
		cfg.GetDbHost(),
		cfg.GetDbBase(),
		cfg.GetDbUser(),
		cfg.GetDbPassword(),
		cfg.GetDbNetwork(),
	)

	flag.Parse()
	if migrateMode != nil && *migrateMode == true {
		log.Info("Service run as migrate mode")
		mig := migrations.New(db)
		if err := mig.Up(); err != nil {
			log.Fatal(err.Error())
		}
		log.Info("Migrations successful applied")
		return
	}

	ur := usersRepo.New(db)
	ui := usersInteractor.New(ur)

	srv := webServer.New(cfg.GetListenAddr(), log)
	ws := webService.New(srv, ui, log)

	ws.Start()
}
