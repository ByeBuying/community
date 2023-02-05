package main

import (
	"community/conf"
	ctl "community/controller"
	"community/model"
	"community/router"
	"context"
	"flag"
	"fmt"
	"go-common/klay/elog"
	"golang.org/x/sync/errgroup"
	"net/http"
	"path"
	"time"
)

var configFlag = flag.String("config", "./conf/config.toml", "configuration")
var httpFlag = flag.Int("http", 0, "router http port")
var (
	ctx context.Context
	g   errgroup.Group
)

func main() {
	flag.Parse()

	// config file
	config := conf.NewConfig(*configFlag)

	// http
	if *httpFlag != 0 {
		config.Port.Http = *httpFlag
	}

	elog.InitLog(elog.Config{
		UseTerminal:       config.Log.Terminal.Use,
		UseFile:           config.Log.File.Use,
		VerbosityTerminal: config.Log.Terminal.Verbosity,
		VerbosityFile:     config.Log.File.Verbosity,
		FilePath:          path.Join(config.Datadir.Root, config.Datadir.Log, config.Log.File.FileName),
	})

	// model
	if repositories, err := model.NewRepositories(config); err != nil {
		panic(err)
	} else if controller, err := ctl.New(config, config.Port.Http, repositories); err != nil {
		panic(fmt.Errorf("controller.New > %v", err))
	} else if rt, err := router.NewRouter(config, controller); err != nil {
		panic(fmt.Errorf("router.NewRouter > %v", err))
	} else {
		mapi := &http.Server{
			Addr:           config.Port.Server,
			Handler:        rt.Idx(),
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		// log info
		g.Go(func() error {
			return mapi.ListenAndServe()
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}

}
