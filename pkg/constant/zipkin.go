package constant

import "os"

var TRACE_REPORTER_URL = "http://127.0.0.1:9411/api/v2/spans"

func init() {
	if os.Getenv("TRACE_REPORTER_URL") != "" {
		TRACE_REPORTER_URL = os.Getenv("TRACE_REPORTER_URL")
	}
}
