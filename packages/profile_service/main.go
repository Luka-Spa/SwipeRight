package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Luka-Spa/SwipeRight/packages/profile_service/config"
	httpRouter "github.com/Luka-Spa/SwipeRight/packages/profile_service/router/http"
)


func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	httpRouter.Init()

}
