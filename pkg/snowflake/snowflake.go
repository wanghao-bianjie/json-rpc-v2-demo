package snowflake

import "github.com/GUAIK-ORG/go-snowflake/snowflake"

type SnowflakeUUID struct {
	sf *snowflake.Snowflake
}

func NewSnowflakeUUID(dataCenterId, workerId int64) (*SnowflakeUUID, error) {
	sf, err := snowflake.NewSnowflake(dataCenterId, workerId)
	if err != nil {
		return nil, err
	}
	return &SnowflakeUUID{
		sf: sf,
	}, nil
}

func (s SnowflakeUUID) GenerateUUID() int64 {
	return s.sf.NextVal()
}

func (s SnowflakeUUID) GetDeviceID(uuid int64) (int64, int64) {
	return snowflake.GetDeviceID(uuid)
}

func (s SnowflakeUUID) GetGenTimestamp(uuid int64) int64 {
	return snowflake.GetGenTimestamp(uuid)
}

func (s SnowflakeUUID) GetGenTime(uuid int64) string {
	return snowflake.GetGenTime(uuid)
}

func (s SnowflakeUUID) GetTimestamp(uuid int64) int64 {
	return snowflake.GetTimestamp(uuid)
}

func (s SnowflakeUUID) GetTimestampStatus() float64 {
	return snowflake.GetTimestampStatus()
}
