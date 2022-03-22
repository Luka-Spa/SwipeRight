package repository

import (
	"SwipeRight_Profile_Service/config"
	"log"
	"time"

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
	cluster := gocql.NewCluster(config.GetString("cassandra.hosts")) //replace PublicIP with the IP addresses used by your cluster.
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: config.GetString("cassandra.username"), Password: config.GetString("cassandra.password")} //replace the username and password fields with their real settings.
	DB, err := cluster.CreateSession()
	if err != nil {
	log.Println(err)
	return
	}
	defer DB.Close()
}
