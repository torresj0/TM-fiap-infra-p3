package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	password := "S3cr3tP@ssw0rd#2024!"
	hash := md5.Sum([]byte(password))
	fmt.Println("auth service started")
	fmt.Printf("hash: %x\n", hash)
}
