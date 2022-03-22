package main

import (
	"SwipeRight_Profile_Service/config"
	Router "SwipeRight_Profile_Service/router/http"
	"flag"
	"fmt"
	"os"
)


func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	Router.Init()
}
