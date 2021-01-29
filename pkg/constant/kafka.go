package constant

import "os"

const KAFKA_HOSTS = "127.0.0.1:9092"

func init() {
	if os.Getenv("KAFKA_HOSTS") != "" {
		POSTGRESQL_DB_NAME = os.Getenv("KAFKA_HOSTS")
	}
}
