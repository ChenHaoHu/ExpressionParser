package main

import (
	"log"
	"strings"
)

func main() {

	str := "num{$age}"
	strs := strings.Split(str, "{$")
	log.Println(strs)
	if len(strs) == 2 {
		log.Println(strs[1][len(strs[1])-1])
		log.Println(byte('}'))
		if strs[1][len(strs[1])-1] != byte('}') {
			log.Println("false")
			return
		}
		log.Println("true")
	}
}
