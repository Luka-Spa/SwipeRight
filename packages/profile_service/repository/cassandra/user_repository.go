package cassandra

import (
	"fmt"

	"github.com/Luka-Spa/SwipeRight/packages/profile_service/repository"
)

func GetUserFirstnameById() string {
	//var user model.UserProfile
	var query = "SELECT first_name FROM USER_PROFILE WHERE ID=0;"
	firstname, ok := repository.ReadSingleCassandraQuery(query).(string)
	if ok == false {
		fmt.Println("Type conversion failed")
	}
	fmt.Println(firstname)
	return firstname
}
