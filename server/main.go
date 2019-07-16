package main

import (
	"fmt"
	"github.com/soyum2222/slog"
	"gonat/server/config"
	"gonat/server/conn"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

func main() {

	config.Load()

	fmt.Println("config load success")
	if config.Debug {
		go http.ListenAndServe(":8808", nil)
	}

	var err error
	slog.Logger, err = slog.DefaultNew(func() slog.SLogConfig {
		cfg := slog.TestSLogConfig()
		cfg.Debug = config.Debug
		return cfg
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("log create success")

	conn.Start(strconv.Itoa(config.Port))
}