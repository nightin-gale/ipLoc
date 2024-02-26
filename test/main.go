package main 

import (
	"fmt"
	"log"
	"os"

	"github.com/nightin-gale/ipLoc/ipLoc"
)

func main (){
    cmdArgs := os.Args[1:]
	if len(cmdArgs) < 1 {
		log.Fatal("Pls Enter IP addr")
	}
	loc, err := ipLoc.IpLoc(cmdArgs[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(loc)
}
