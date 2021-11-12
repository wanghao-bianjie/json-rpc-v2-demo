package global

import (
	"myjsonrpcv2/internal/app/conf"
	"myjsonrpcv2/pkg/snowflake"
)

var (
	Config *conf.Config
	Sf     *snowflake.SnowflakeUUID
)
