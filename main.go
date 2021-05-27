package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	p1 := person{
		First: "Alex",
	}

	p2 := person{
		First: "Jenny",
	}

	xp := []person{p1, p2}
	bs, err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(bs))

	xp2 := []person{}

	if err = json.Unmarshal(bs, &xp2); err != nil {
		log.Panic(err)
	}
	fmt.Println(xp2)

	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)
	http.ListenAndServe(":8080", nil)
}

func encode(w http.ResponseWriter, r *http.Request) {

}

func decode(w http.ResponseWriter, r *http.Request) {

}
