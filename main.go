package main

import (
	"gsurl/httpsvr"
	"gsurl/log"
	"gsurl/storage"
)

func main() {
	log.Init()
	log.Logger.Infof("Starting gsurl service...")
	storage.Init()
	httpsvr.Init()
}
