package main

import (
	"surveyapp/Server"
	"log"
	"sync"

	"github.com/alexflint/go-arg"
)

var args struct {
	Bind string `arg:"env:GOCRUD_BIND_HTTP"`
}

func main() {
	arg.MustParse(&args)

	if args.Bind == ""{
		args.Bind = ":8082"
	}
	
	var wg sync.WaitGroup

	wg.Add(1)
	go func(){
		log.Printf("Starting Sever on...\n")
		Server.Serve(args.Bind)
		wg.Done()
	}()

	wg.Wait()

}


