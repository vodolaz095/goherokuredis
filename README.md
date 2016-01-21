GoHerokuRedis
=================================

[![GoDoc](https://godoc.org/github.com/vodolaz095/goherokuredis?status.svg)](https://godoc.org/github.com/vodolaz095/goherokuredis)
[![Build Status](https://travis-ci.org/vodolaz095/goherokuredis.svg?branch=master)](https://travis-ci.org/vodolaz095/goherokuredis)

Simple helper for Golang web applications for automaticaly configuring [Redis server](http://redis.io) clients by
extracting connection string from environment parameters on [Heroku](http://heroku.com/) hosting.

How it works
=================================
This module inspects environment values for well known Redis-As-A-Service
connection parameters being used at Heroku. Currently - these environment variables are supported (in this order):

- REDISTOGO_URL - see for details (https://devcenter.heroku.com/articles/redistogo#using-with-node-js)
- OPENREDIS_URL - see for details (https://devcenter.heroku.com/articles/openredis#using-redis-from-node-js)
- REDISCLOUD_URL - see for details (https://devcenter.heroku.com/articles/heroku-redis#connecting-in-node-js)
- REDISGREEN_URL - see for details (https://devcenter.heroku.com/articles/redisgreen#using-redis-with-node-js)
- REDIS_URL - see for details (https://devcenter.heroku.com/articles/heroku-redis#connecting-in-node-js)

if all of this environment values are empty, the default connection string is being used - `redis://:@localhost:6379` - that is equal to 
connecting to 6373 port of localhost without password - stack settings for running Redis server on localost

Basic usage
=================================


```go 

		package main
		
		import (
			"fmt"
		
			"github.com/vodolaz095/goherokuredis"
		)
		
		func main() {
			redisClient, err := goherokuredis.Init()
		
			if err != nil {
				panic(err)
			}
		
			err = redisClient.Ping().Err()
			if err != nil {
				panic(err)
			}
			fmt.Printf("Redis on %s is online!\n", redisClient.String())
		}


```


Usage as part of webserver
=============================

```go

	package main
	
	import (
		"fmt"
		"net/http"
		"os"
	
		"github.com/vodolaz095/goherokuredis"
	)
	
	func main() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "3000"
		}
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				err := recover()
				if err != nil {
					w.Header().Set("Content-Type", "text/plain")
					w.WriteHeader(http.StatusServiceUnavailable)
					fmt.Fprintf(w, "Redis on has this error: %s!", err)
				}
			}()
	
			redisClient, err := goherokuredis.Init()
			if err != nil {
				panic(err)
			}
			defer func() {
				redisClient.Close()
			}()
			err = redisClient.Ping().Err()
			if err != nil {
				panic(err)
			}
	
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Redis on %s is online!", redisClient.String())
		})
		err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil)
		if err != nil {
			panic(err)
		}
	}

```



License
=================================
The MIT License (MIT)

Copyright (c) 2015 Ostroumov Anatolij ostroumov095(at)gmail(dot)com et al.

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
