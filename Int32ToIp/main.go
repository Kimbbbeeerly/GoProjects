package main

import (
	"fmt"
	"strconv"
)

func Int32ToIp(n32 uint32) (ip string) {
	// Converting type from UINT32 to INT64
	n64 := int64(n32)

	// Declaring 4 octets of IPv4 Address
	o0 := strconv.FormatInt((n64>>24)&0xff, 10)
	o1 := strconv.FormatInt((n64>>16)&0xff, 10)
	o2 := strconv.FormatInt((n64>>8)&0xff, 10)
	o3 := strconv.FormatInt((n64 & 0xff), 10)

	// Combine octets to return
	ip = o0 + "." + o1 + "." + o2 + "." + o3

	return
}

func main() {
	//For example, transforming 32 bit number "2149583361" to IPv4 Address
	result := Int32ToIp(2149583361)
	fmt.Print(result)
}
