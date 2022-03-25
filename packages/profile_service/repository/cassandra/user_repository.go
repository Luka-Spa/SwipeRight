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

func GetUser() map[string]interface{} {
	var query = "SELECT value FROM profile.user_profile_table WHERE pk=0;"
	user := repository.ReadCassandraQuery(query)[0]
	return user;
}
