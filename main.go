package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Vladimir-Kuchinskiy/golang-arch/passwords"
)

type person struct {
	First string
}

var key []byte

func main() {
	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}

	//Password hash
	pass := "123456789"

	hashedPass, err := passwords.HashPassword(pass)
	if err != nil {
		log.Panic(err)
	}

	err = passwords.ComparePassword(pass, hashedPass)
	if err != nil {
		log.Fatalln("Not logged in", err)
	}

	log.Printf("Password: %s, hashedPassword: %s", pass, string(hashedPass))

	// Marshal unmarshal
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

	// Encode, Decode
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)
	http.HandleFunc("/encode-array", encodeArray)
	http.HandleFunc("/decode-array", decodeArray)
	http.ListenAndServe(":8080", nil)
}

func encode(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Jenny",
	}

	if err := json.NewEncoder(w).Encode(p1); err != nil {
		log.Println(err)
	}
}

func decode(w http.ResponseWriter, r *http.Request) {
	p1 := person{}

	if err := json.NewDecoder(r.Body).Decode(&p1); err != nil {
		log.Println(err)
	}
	log.Println("Person: ", p1)
}

func encodeArray(w http.ResponseWriter, r *http.Request) {
	people := []person{
		{
			First: "Jenny",
		},
		{
			First: "Mike",
		},
	}

	if err := json.NewEncoder(w).Encode(people); err != nil {
		log.Println(err)
	}
}

func decodeArray(w http.ResponseWriter, r *http.Request) {
	people := []person{}

	if err := json.NewDecoder(r.Body).Decode(&people); err != nil {
		log.Println(err)
	}
	log.Println("Person: ", people)
}
