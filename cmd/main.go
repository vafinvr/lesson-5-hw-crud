package main

import (
	"lesson-5-hw-crud/infrastructure"
	"lesson-5-hw-crud/infrastructure/logger"
	"lesson-5-hw-crud/infrastructure/pgdb"
	"lesson-5-hw-crud/infrastructure/webServer"
	"lesson-5-hw-crud/interfaces/config"
	"lesson-5-hw-crud/interfaces/usersRepo"
	"lesson-5-hw-crud/interfaces/webService"
	"lesson-5-hw-crud/usecases/usersInteractor"
)

func main() {
	cfg, err := config.NewConfig(new(infrastructure.Env))
	if err != nil {
		panic(err)
	}

	log := logger.New()

	db := pgdb.New(
		cfg.GetDbHost(),
		cfg.GetDbBase(),
		cfg.GetDbUser(),
		cfg.GetDbPassword(),
		cfg.GetDbNetwork(),
	)

	ur := usersRepo.New(db)
	ui := usersInteractor.New(ur)

	srv := webServer.New(cfg.GetListenAddr(), log)
	ws := webService.New(srv, ui, log)

	ws.Start()
}
