package main

import (
	"fmt"
	"imgtrace/hash"
)

func test() {
	hashValue, err := hash.Generate("/Users/sabyasacheethakur/Downloads/ChatGPT Image Apr 24, 2025, 02_54_59 PM.png") // replace with real file
	if err != nil {
		fmt.Println("Hash error:", err)
		return
	}
	fmt.Println("Hash:", hashValue)
}
