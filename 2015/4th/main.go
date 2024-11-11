package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func main() {
	keyBase := "iwrupvqb"
	now := time.Now()

	defer func(start time.Time) {
		fmt.Printf("Execution time: %v\n", time.Since(start))
	}(now)

	// Brute force.
	for i := 1; ; i++ {
		md5 := GetMD5Hash(fmt.Sprint(keyBase, i))

		if strings.HasPrefix(md5, "000000") {
			fmt.Println("Number you are looking for is: ", i)
			return
		}

		fmt.Println(md5)
	}
}

// 5 zeroes => 10.8232294s
// 6 zeroes =>
