package Server

import (
	"log"
	"net/http"
)

func CreateSurvey(){
	
}

func Serve(bind string) {
	http.Handle("/survey/create",CreateSurvey())
	log.Printf("Sever listening on %v \n", bind)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
