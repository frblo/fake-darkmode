package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

var darkmode bool // disgusting global variable

func setBool(args []string) {
	defBool := false

	if len(args) < 2 {
		log.Printf("No argument, assuming default value %v", defBool)
		darkmode = defBool
	}
	dm, err := strconv.ParseBool(os.Args[1])
	if err != nil {
		log.Printf("Arg %v not parsed correctly, assuming default value %v", os.Args[1], defBool)
		darkmode = defBool
	} else {
		log.Printf("%v parsed correctly", args[1])
		darkmode = dm
	}
}

func serveDarkmode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(darkmode)
}

func main() {
	setBool(os.Args)

	http.HandleFunc("/", serveDarkmode)

	http.ListenAndServe(":8090", nil)
}
