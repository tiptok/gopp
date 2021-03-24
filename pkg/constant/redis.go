package constant

import "os"

var REDIS_HOST = "127.0.0.1"
var REDIS_PORT = "6379"
var REDIS_AUTH = ""

func init() {
	if os.Getenv("REDIS_HOST") != "" {
		REDIS_HOST = os.Getenv("REDIS_HOST")
		REDIS_AUTH = os.Getenv("REDIS_AUTH")
	}
	if os.Getenv("REDIS_PORT") != "" {
		REDIS_PORT = os.Getenv("REDIS_PORT")
	}
	if _, ok := os.LookupEnv("REDIS_AUTH"); ok {
		REDIS_AUTH = os.Getenv("REDIS_AUTH")
	}
}
