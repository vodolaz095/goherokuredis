package goherokuredis

import (
	"os"
	"testing"
)

func TestLocalDB(t *testing.T) {
	redisClient, err := Init()
	if err != nil {
		t.Error(err)
	}
	info := redisClient.Info()
	if info.Err() != nil {
		t.Errorf("Error calling info %s", err.Error())
	}
	if redisClient.String() != "Redis<localhost:6379 db:0>" {
		t.Errorf("Wrong redis database client!")
	}
	if redisClient.Close() != nil {
		t.Errorf("Error closing redis client %s", err.Error())
	}
}

func TestRedisToGo(t *testing.T) {
	err := os.Setenv("REDISTOGO_URL", "redis://redistogo:@localhost:6379/")
	if err != nil {
		t.Errorf("Error setting environment for RedisToGo - %s", err.Error())
	}
	redisClient, err := Init()
	if err != nil {
		t.Error(err)
	}

	info := redisClient.Info()
	if info.Err() != nil {
		t.Errorf("Error calling info %s", err.Error())
	}
	if redisClient.String() != "Redis<localhost:6379 db:0>" {
		t.Errorf("Wrong redis database client!")
	}
	if redisClient.Close() != nil {
		t.Errorf("Error closing redis client %s", err.Error())
	}

	err = os.Unsetenv("REDISTOGO_URL")
	if err != nil {
		t.Errorf("Error unsetting environment for RedisToGo - %s", err.Error())
	}
}

func TestRedisToGoNoAuth(t *testing.T) {
	err := os.Setenv("REDISTOGO_URL", "redis://localhost:6379/")
	if err != nil {
		t.Errorf("Error setting environment for RedisToGo - %s", err.Error())
	}
	redisClient, err := Init()
	if err != nil {
		t.Error(err)
	}

	info := redisClient.Info()
	if info.Err() != nil {
		t.Errorf("Error calling info %s", err.Error())
	}
	if redisClient.String() != "Redis<localhost:6379 db:0>" {
		t.Errorf("Wrong redis database client!")
	}
	if redisClient.Close() != nil {
		t.Errorf("Error closing redis client %s", err.Error())
	}

	err = os.Unsetenv("REDISTOGO_URL")
	if err != nil {
		t.Errorf("Error unsetting environment for RedisToGo - %s", err.Error())
	}
}

func TestOpenRedis(t *testing.T) {
	err := os.Setenv("OPENREDIS_URL", "redis://openredis:@localhost:6379/")
	if err != nil {
		t.Errorf("Error setting environment for OpenRedis - %s", err.Error())
	}

	redisClient, err := Init()
	if err != nil {
		t.Error(err)
	}

	info := redisClient.Info()
	if info.Err() != nil {
		t.Errorf("Error calling info %s", err.Error())
	}
	if redisClient.String() != "Redis<localhost:6379 db:0>" {
		t.Errorf("Wrong redis database client!")
	}
	if redisClient.Close() != nil {
		t.Errorf("Error closing redis client %s", err.Error())
	}

	err = os.Unsetenv("OPENREDIS_URL")
	if err != nil {
		t.Errorf("Error setting environment for OpenRedis - %s", err.Error())
	}
}

func TestRedisCloud(t *testing.T) {
	err := os.Setenv("REDISCLOUD_URL", "redis://rediscloud:@localhost:6379/")
	if err != nil {
		t.Errorf("Error setting environment for RedisCloud - %s", err.Error())
	}
	redisClient, err := Init()
	if err != nil {
		t.Error(err)
	}

	info := redisClient.Info()
	if info.Err() != nil {
		t.Errorf("Error calling info %s", err.Error())
	}
	if redisClient.String() != "Redis<localhost:6379 db:0>" {
		t.Errorf("Wrong redis database client!")
	}
	if redisClient.Close() != nil {
		t.Errorf("Error closing redis client %s", err.Error())
	}

	err = os.Unsetenv("REDISCLOUD_URL")
	if err != nil {
		t.Errorf("Error unsetting environment for RedisCloud - %s", err.Error())
	}
}

func TestRedisGreen(t *testing.T) {
	err := os.Setenv("REDISGREEN_URL", "redis://redisgreen:@localhost:6379/")
	if err != nil {
		t.Errorf("Error setting environment for RedisGreen - %s", err.Error())
	}
	redisClient, err := Init()
	if err != nil {
		t.Error(err)
	}

	info := redisClient.Info()
	if info.Err() != nil {
		t.Errorf("Error calling info %s", err.Error())
	}
	if redisClient.String() != "Redis<localhost:6379 db:0>" {
		t.Errorf("Wrong redis database client!")
	}
	if redisClient.Close() != nil {
		t.Errorf("Error closing redis client %s", err.Error())
	}

	err = os.Unsetenv("REDISGREEN_URL")
	if err != nil {
		t.Errorf("Error unsetting environment for RedisToGo - %s", err.Error())
	}
}

func TestHerokuRedis(t *testing.T) {
	err := os.Setenv("REDIS_URL", "redis://herokuredis:@localhost:6379/")
	if err != nil {
		t.Errorf("Error setting environment for Heroku Redis - %s", err.Error())
	}
	redisClient, err := Init()
	if err != nil {
		t.Error(err)
	}

	info := redisClient.Info()
	if info.Err() != nil {
		t.Errorf("Error calling info %s", err.Error())
	}
	if redisClient.String() != "Redis<localhost:6379 db:0>" {
		t.Errorf("Wrong redis database client!")
	}
	if redisClient.Close() != nil {
		t.Errorf("Error closing redis client %s", err.Error())
	}

	err = os.Unsetenv("REDIS_URL")
	if err != nil {
		t.Errorf("Error unsetting environment for RedisToGo - %s", err.Error())
	}
}

func TestBadSchemaHttp(t *testing.T) {
	err := os.Setenv("REDIS_URL", "http://example.org/")
	if err != nil {
		t.Errorf("Error setting environment for Heroku Redis - %s", err.Error())
	}
	_, err = Init()
	if err.Error() != "herokuredis : wrong database connection string schema - http" {
		t.Errorf("Bad error response - %s", err.Error())
	}
	err = os.Unsetenv("REDIS_URL")
	if err != nil {
		t.Errorf("Error unsetting environment for RedisToGo - %s", err.Error())
	}

}

func TestBadSchemaHttps(t *testing.T) {
	err := os.Setenv("REDIS_URL", "https://example.org/")
	if err != nil {
		t.Errorf("Error setting environment for Heroku Redis - %s", err.Error())
	}
	_, err = Init()

	if err.Error() != "herokuredis : wrong database connection string schema - https" {
		t.Errorf("Bad error response - %s", err.Error())
	}
	err = os.Unsetenv("REDIS_URL")
	if err != nil {
		t.Errorf("Error unsetting environment for RedisToGo - %s", err.Error())
	}
}

func TestDuplicateConnectionStrings(t *testing.T) {
	err := os.Setenv("REDIS_URL", "redis://herokuredis:@localhost:6379/")
	if err != nil {
		t.Errorf("Error setting environment for Heroku Redis - %s", err.Error())
	}
	err = os.Setenv("REDISGREEN_URL", "redis://redisgreen:@localhost:6379/")
	if err != nil {
		t.Errorf("Error setting environment for RedisGreen - %s", err.Error())
	}

	_, err = Init()
	if err.Error() != "goherokuredis : Duplicate redis connection extracted [REDISGREEN_URL REDIS_URL]" {
		t.Error(err)
	}

	err = os.Unsetenv("REDIS_URL")
	if err != nil {
		t.Errorf("Error unsetting environment for RedisToGo - %s", err.Error())
	}
	err = os.Unsetenv("REDISGREEN_URL")
	if err != nil {
		t.Errorf("Error unsetting environment for RedisToGo - %s", err.Error())
	}
}

func TestOverrideConnectionString(t *testing.T) {
	err := os.Setenv("REDIS_URL", "redis://herokuredis:@example.org:6379/")
	if err != nil {
		t.Errorf("Error setting environment for Heroku Redis - %s", err.Error())
	}
	redisClient, err := Init("redis://herokuredis:@localhost:6379/")
	if err != nil {
		t.Error(err)
	}

	info := redisClient.Info()
	if info.Err() != nil {
		t.Errorf("Error calling info %s", err.Error())
	}
	if redisClient.String() != "Redis<localhost:6379 db:0>" {
		t.Errorf("Wrong redis database client!")
	}
	if redisClient.Close() != nil {
		t.Errorf("Error closing redis client %s", err.Error())
	}

	err = os.Unsetenv("REDIS_URL")
	if err != nil {
		t.Errorf("Error unsetting environment for RedisToGo - %s", err.Error())
	}
}

func TestBadOverrideConnectionString(t *testing.T) {
	err := os.Setenv("REDIS_URL", "redis://herokuredis:@example.org:6379/")
	if err != nil {
		t.Errorf("Error setting environment for Heroku Redis - %s", err.Error())
	}
	_, err = Init("redis://herokuredis:@localhost:6379/", "redis://herokuredis:@otherhost.com:6379/")
	if err.Error() != "goherokuredis : Multiple connection override strings" {
		t.Error(err)
	}

	err = os.Unsetenv("REDIS_URL")
	if err != nil {
		t.Errorf("Error unsetting environment for RedisToGo - %s", err.Error())
	}

}
