package main

import (
	"flag"
	"os"

	"github.com/Luka-Spa/SwipeRight/packages/recommendation/config"
	"github.com/Luka-Spa/SwipeRight/packages/recommendation/repository"
	httpRouter "github.com/Luka-Spa/SwipeRight/packages/recommendation/router/http"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		print("Usage: server -e {mode} \n")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	repository.Init()
	httpRouter.Init()
}
