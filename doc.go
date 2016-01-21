/*
Package goherokuredis inspects environment values for well known Redis-As-A-Service
connection parameters being used at Heroku. Currently - these environment variables are supported (in this order):

- REDISTOGO_URL - see for details (https://devcenter.heroku.com/articles/redistogo#using-with-node-js)

- OPENREDIS_URL - see for details (https://devcenter.heroku.com/articles/openredis#using-redis-from-node-js)

- REDISCLOUD_URL - see for details (https://devcenter.heroku.com/articles/heroku-redis#connecting-in-node-js)

- REDISGREEN_URL - see for details (https://devcenter.heroku.com/articles/redisgreen#using-redis-with-node-js)

- REDIS_URL - see for details (https://devcenter.heroku.com/articles/heroku-redis#connecting-in-node-js)

if all of this environment values are empty, the default connection string is being used - `redis://:@localhost:6379` - that is equal to
connecting to 6373 port of localhost without password - stack settings for running Redis server on localost.
Also it is possible to override default setting by providing argument to `Init()` function.


*/
package goherokuredis
