package cassandra

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var DB *gocql.Session

func Connect(config *viper.Viper) *gocql.Session {
	cluster := gocql.NewCluster(config.GetStringSlice("cassandra.hosts")...)
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10
	cluster.DisableInitialHostLookup = true
	cluster.Keyspace = config.GetString("cassandra.keyspace")
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: config.GetString("cassandra.username"), Password: config.GetString("cassandra.password")} //replace the username and password fields with their real settings.
	var err error
	DB, err = cluster.CreateSession()
	if err != nil {
		log.Errorln(err)
		log.Fatal("Unable to connect to Cassandra")
	} else {
		log.Infof("Connected to Cassandra")
	}
	return DB
}

func CassandraWrite(query string, values ...interface{}) error {
	if err := DB.Query(query).Bind(values...).Exec(); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func CassandraReadSingle[T any](query string, result T, values ...interface{}) T {
	err := DB.Query(query).Bind(values...).Scan(&result)
	if err != nil {
		log.Error(err)
	}
	return result
}

func CassandraRead[T any](query string, result T, values ...interface{}) []T {
	var results = []T{}
	m := map[string]interface{}{}
	iter := DB.Query(query).Iter()
	for iter.MapScan(m) {
		var result T
		if err := mapstructure.Decode(m, &result); err != nil {
			log.Error(err)
		}
		results = append(results, result)
		m = map[string]interface{}{}
	}
	return results
}
