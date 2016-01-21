package goherokuredis

import (
	"fmt"
	"net/url"
	"os"

	"gopkg.in/redis.v3"
)

//Init returns ready to use redis client
//with connection parameters extracted from environment values for known Redis-as-a-Service providers
//being used at Heroku. Currently - these environment variables are supported:
// 'REDISTOGO_URL',//https://devcenter.heroku.com/articles/redistogo#using-with-node-js
// 'OPENREDIS_URL',//https://devcenter.heroku.com/articles/openredis#using-redis-from-node-js
// 'REDISCLOUD_URL', //https://devcenter.heroku.com/articles/heroku-redis#connecting-in-node-js
// 'REDISGREEN_URL', //https://devcenter.heroku.com/articles/redisgreen#using-redis-with-node-js
// 'REDIS_URL' //https://devcenter.heroku.com/articles/heroku-redis#connecting-in-node-js
func Init(overrideRedisConnectionString ...string) (*redis.Client, error) {
	redisConnectionUrl := "redis://:@localhost:6379/"
	extractedConnectionString := ""
	knownProviders := []string{
		"REDISTOGO_URL",
		"OPENREDIS_URL",
		"REDISCLOUD_URL",
		"REDISGREEN_URL",
		"REDIS_URL",
	}
	duplicateConnectionStrings := []string{}
	for _, v := range knownProviders {
		extractedConnectionString = os.Getenv(v)

		if extractedConnectionString != "" {
			duplicateConnectionStrings = append(duplicateConnectionStrings, v)
		}

		if len(duplicateConnectionStrings) > 1 {
			return nil, fmt.Errorf("goherokuredis : Duplicate redis connection extracted %s!", duplicateConnectionStrings)
		}
	}

	if extractedConnectionString == "" {
		extractedConnectionString = redisConnectionUrl
	}

	if len(overrideRedisConnectionString) > 1 {
		return nil, fmt.Errorf("goherokuredis : Multiple connection override strings!")
	}

	if len(overrideRedisConnectionString) == 1 {
		extractedConnectionString = overrideRedisConnectionString[0]
	}

	connectionParameters, err := url.Parse(extractedConnectionString)
	if err != nil {
		return nil, err
	}

	if connectionParameters.Scheme != "redis" {
		return nil, fmt.Errorf("herokuredis : wrong database connection string schema - %s", connectionParameters.Scheme)
	}

	passwd := ""
	userAndPasswd := connectionParameters.User
	if userAndPasswd != nil {
		passwd, _ = userAndPasswd.Password()
	} else {
		passwd = ""
	}
	client := redis.NewClient(&redis.Options{
		Addr:     connectionParameters.Host,
		Password: passwd,
		DB:       0, // use default DB
	})
	err = client.Ping().Err()
	if err != nil {
		return nil, err
	}
	return client, nil
}
