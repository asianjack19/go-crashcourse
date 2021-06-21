package main

import (
	"flag"
	"fmt"
)

func main() {

	//parameter 1 namannya(variable), kedua default value, ketiga message
	var host *string = flag.String("host", "localhost", "Put your host")
	var user *string = flag.String("user", "root", "Put your db user")
	var password *string = flag.String("password", "root", "Put your db password")

	flag.Parse()

	fmt.Println(*host, *user, *password)
}
