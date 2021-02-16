package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	key := "ckczppom"
	fmt.Printf("%d", FindFirstCoin(key))
}

func FindFirstCoin(key string) int {
	for i := 1; ; i++ {
		keyintstring := key + strconv.Itoa(i)
		if SixLeadingZeros(GetMD5Hash(keyintstring)) {
			return i
		}
	}
}

func GetMD5Hash(s string) string {
	toHex := md5.Sum([]byte(s))
	return hex.EncodeToString(toHex[:])
}

func SixLeadingZeros(s string) bool {
	return s[:6] == "000000"
}
