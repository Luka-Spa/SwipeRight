package repository

import (
	"fmt"
	"log"

	"github.com/Luka-Spa/SwipeRight/packages/profile_service/config"

	"github.com/gocql/gocql"
	"github.com/spf13/viper"
)

var DB *gocql.Session;

func Init() {
	config := config.GetConfig()
	cassandraConnection(config)
}

func cassandraConnection(config *viper.Viper) {
	// connect to the cluster
	cluster := gocql.NewCluster(config.GetStringSlice("cassandra.hosts")...) //replace PublicIP with the IP addresses used by your cluster.
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	//cluster.ConnectTimeout = time.Second * 10
	cluster.Keyspace = config.GetString("cassandra.keyspace")
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: config.GetString("cassandra.username"), Password: config.GetString("cassandra.password")} //replace the username and password fields with their real settings.
	var err error
	DB, err = cluster.CreateSession()

	if err != nil {
	log.Println(err)
	return
	}
	fmt.Println("Connected to Cassandra")
}

func ExecCassandraQuery(query string, values ...interface{}) {
	if err:= DB.Query(query).Bind(values...).Exec(); err != nil {
		log.Fatal(err)
	}
}

func ReadSingleCassandraQuery(query string, values ...interface{}) interface{} {
	var result string
	if err:= DB.Query(query).Bind(values...).Scan(&result); err != nil {
		fmt.Println(err)
	}
	return result;
}
