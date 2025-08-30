package main

import (
	"gsurl/config"
	"gsurl/httpsvr"
	"gsurl/log"
	"gsurl/service"
	"gsurl/storage"
)

func main() {
	config.Init()
	log.Init()
	log.Logger.Infof("Starting gsurl service...")
	service.InitIdGenerator()
	service.InitCache()
	storage.Init()
	httpsvr.Init()
}
