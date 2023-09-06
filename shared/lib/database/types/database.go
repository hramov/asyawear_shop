package types

import (
	"time"
)

type ConnectOptions struct {
	Host            string
	Port            int
	User            string
	Password        string
	Database        string
	SslMode         string
	MaxOpenCons     int
	MaxIdleCons     int
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration
}
