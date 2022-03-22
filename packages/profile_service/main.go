package main

import (
<<<<<<< HEAD
	"SwipeRight_Profile_Service/config"
	Router "SwipeRight_Profile_Service/router/http"
	"flag"
	"fmt"
	"os"
=======
	"flag"
	"fmt"
	"os"

	"github.com/Luka-Spa/SwipeRight/packages/profile_service/config"
	httpRouter "github.com/Luka-Spa/SwipeRight/packages/profile_service/router/http"
>>>>>>> 18c1201 (new commit)
)


func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
<<<<<<< HEAD
	Router.Init()
=======
	httpRouter.Init()

>>>>>>> 18c1201 (new commit)
}
