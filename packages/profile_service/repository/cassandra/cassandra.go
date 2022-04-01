package cassandra

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
	"github.com/spf13/viper"
)

var DB *gocql.Session

func Connect(config *viper.Viper) *gocql.Session {
	cluster := gocql.NewCluster(config.GetStringSlice("cassandra.hosts")...)
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.Keyspace = config.GetString("cassandra.keyspace")
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: config.GetString("cassandra.username"), Password: config.GetString("cassandra.password")} //replace the username and password fields with their real settings.
	var err error
	DB, err = cluster.CreateSession()
	if err != nil {
		log.Println(err)
	} else {
	fmt.Println("Connected to Cassandra")
	}
	return DB
}

func CassandraWrite(query string, values ...interface{}) {
	if err := DB.Query(query).Bind(values...).Exec(); err != nil {
		fmt.Println(err)
	}
}

func CassandraReadSingle[T any](query string, result T, values ...interface{}) T {
	err := DB.Query(query).Bind(values...).Scan(&result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}
