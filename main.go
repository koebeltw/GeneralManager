package main

import (
	"bufio"
	"log"
	"os"
)

func main(){
	for {
		r := bufio.NewReader(os.Stdin)
		b, _, _ := r.ReadLine()
		log.Println(string(b))
		switch {
		case string(b) == "send":
			//Client.Send(1, 1, []byte{1})
		}
	}
}
