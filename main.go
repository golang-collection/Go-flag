package main

import (
	"flag"
	"log"
)

/**
* @Author: super
* @Date: 2020-09-14 19:41
* @Description:
**/

func main() {
	var name string
	flag.StringVar(&name, "name", "GoGoGo", "help")
	flag.StringVar(&name, "n", "GoGoGo", "help")

	flag.Parse()
	log.Println("name: ", name)
}