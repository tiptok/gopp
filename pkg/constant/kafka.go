package constant

import "os"

const KAFKA_HOSTS = "106.52.15.41:9092"

func init() {
	if os.Getenv("KAFKA_HOSTS") != "" {
		POSTGRESQL_DB_NAME = os.Getenv("KAFKA_HOSTS")
	}
}
