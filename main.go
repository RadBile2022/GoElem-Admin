package main

import (
	"elementary-admin/config"
	"elementary-admin/src"
	"sync"
	"time"
)

var cfg config.Config

func init() {
	cfg = config.InitConfig()
}
func main() {
	srv := src.InitServer(cfg)

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Local = time.UTC
		srv.Run()
	}()

	wg.Wait()
}
