package main

import "fmt"

const APPID      = "your app id"
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
your private key
-----END RSA PRIVATE KEY-----
`)

func main() {
	fmt.Println("Hello")
}
