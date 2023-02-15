package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sever-codebase/application/configs"
	"sever-codebase/application/models"
	"syscall"
	"time"
)

func main() {
	//init env and config
	env, err := configs.LoadEnv()
	cfg, err := configs.LoadConfig("conf/", env)
	if err != nil {
		return
	}
	fmt.Println("Config: ", cfg)

	//force shutdown
	forceQuit(cfg)

	//config grpc gateway server
	StartGrpcGatewayServer(cfg)

	// start grpc server
	StartGrpcServer(cfg)
}

func forceQuit(cfg *models.Config) {
	// handle signal
	_, ctxCancel := context.WithCancel(context.Background())
	go func() {
		osSignal := make(chan os.Signal, 1)
		signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)
		<-osSignal
		ctxCancel()
		// Wait for maximum 15s
		go func() {
			var durationSec time.Duration = 15
			if cfg.Server.Env == "D" {
				durationSec = 1
			}
			timer := time.NewTimer(durationSec * time.Second)
			<-timer.C
			//ll.Fatal("Force shutdown due to timeout!")
		}()
	}()
}
