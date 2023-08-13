package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
	"github.com/nyelwa-senguji/ticketing_system_backend/endpoint"
	"github.com/nyelwa-senguji/ticketing_system_backend/service"
	httptransport "github.com/nyelwa-senguji/ticketing_system_backend/transport"
	"github.com/nyelwa-senguji/ticketing_system_backend/utils"
)

func main() {
	dbUsername := utils.LoadEnviromentalVariables("DB_USER")
	dbPassword := utils.LoadEnviromentalVariables("DB_PASSWORD")
	dbPort := utils.LoadEnviromentalVariables("DB_PORT")
	dbName := utils.LoadEnviromentalVariables("DB_NAME")

	dbsource := dbUsername + ":" + dbPassword + "@tcp(localhost:" + dbPort + ")/" + dbName + "?parseTime=true"

	var httpAddr = flag.String("http", ":9000", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "ticketing",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var conn *sql.DB
	{
		var err error

		conn, err = sql.Open("mysql", dbsource)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	flag.Parse()
	ctx := context.Background()
	var srv service.Service
	{
		repository := db.NewRepository(conn)
		srv = service.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := endpoint.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := httptransport.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
