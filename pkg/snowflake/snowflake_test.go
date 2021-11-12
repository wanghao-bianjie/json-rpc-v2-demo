package snowflake

import (
	"fmt"
	"testing"
)

func TestSnowflakeUUID(t *testing.T) {
	sf, err := NewSnowflakeUUID(1, 1)
	if err != nil {
		t.Fatal(err)
	}
	uuid := sf.GenerateUUID()
	fmt.Printf("uuid: %d\n", uuid)
	fmt.Printf("uuid: %b\n", uuid)
	dataCenterId, workerId := sf.GetDeviceID(uuid)
	fmt.Printf("data center id: %d,worker id: %d\n", dataCenterId, workerId)
	fmt.Printf("timestamp: %d\n", sf.GetTimestamp(uuid))
	fmt.Printf("generate timestamp: %d\n", sf.GetGenTimestamp(uuid))
	fmt.Printf("generate time: %s\n", sf.GetGenTime(uuid))
	fmt.Printf("timestamp status: %f\n", sf.GetTimestampStatus())
}
