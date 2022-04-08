package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Luka-Spa/SwipeRight/packages/profile/config"
	"github.com/Luka-Spa/SwipeRight/packages/profile/repository"
	httpRouter "github.com/Luka-Spa/SwipeRight/packages/profile/router/http"
	log "github.com/sirupsen/logrus"
)


func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	log.Infof("Server starting in a %s environment", *environment)
	config.Init(*environment)
	repository.Init()
	httpRouter.Init()

	defer repository.DB.Close()
}
