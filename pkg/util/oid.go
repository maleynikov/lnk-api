package util

import (
	"fmt"
	"time"

	"github.com/howeyc/crc16"
)

func OID(data string) string {
	checksum := crc16.Checksum([]byte(
		fmt.Sprintf("%s%d", data, time.Now().UnixNano()),
	), crc16.IBMTable)

	// fmt.Println("Checksum: ", checksum)
	// fmt.Println(strconv.FormatInt(int64(checksum), 16))

	return fmt.Sprintf("%x", checksum)
}
