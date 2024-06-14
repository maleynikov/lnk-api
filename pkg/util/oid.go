package util

import (
	"fmt"

	"github.com/howeyc/crc16"
)

func OID(data string) string {
	checksum := crc16.Checksum([]byte(data), crc16.IBMTable)

	// fmt.Println("Checksum: ", checksum)
	// fmt.Println(strconv.FormatInt(int64(checksum), 16))

	return fmt.Sprintf("%x", checksum)
}
