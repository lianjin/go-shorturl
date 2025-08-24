package main

import (
	"gsurl/httpsvr"
	"gsurl/log"
	"gsurl/service"
	"gsurl/storage"
)

func main() {
	log.Init()
	log.Logger.Infof("Starting gsurl service...")
	service.InitIdGenerator()
	storage.Init()
	httpsvr.Init()
}
