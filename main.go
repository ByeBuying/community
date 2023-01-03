package main

import (
	"context"
	"flag"
	"gin-mongo-api/conf"
	"golang.org/x/sync/errgroup"
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

	// model

}
