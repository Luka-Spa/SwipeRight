package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/Luka-Spa/SwipeRight/packages/recommendation/config"
	httpRouter "github.com/Luka-Spa/SwipeRight/packages/recommendation/router/http"
	"github.com/Luka-Spa/SwipeRight/packages/recommendation/service/consumer"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		print("Usage: server -e {mode} \n")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	httpRouter.Init()
	consumer.NewKafkaConsumer().ConsumeUserProfile()
	waitForSyscall()
}

func waitForSyscall() {
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
}
