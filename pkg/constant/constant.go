package constant

import "time"

const ServiceName = "gopp"

const persistent = "postgresql"
const persistentLib = "go-gp"

const TopicUserLogin = "user_login"

const MaxConn = 2
const MaxSize = 5 * 1024
const TimeOutDuration = time.Second * 5
