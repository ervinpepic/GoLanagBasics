package main

import(
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

func main() {
	url := "http://localhost:8181"
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var p Payload
	err = json.Unmarshal(body, &p)

	if err != nil {
		panic(err)
	}

	fmt.Println(p.Envelope.CarsByDoor, "\n", p.Envelope.TrucksByTon)
}

type Payload struct {
	Envelope Data
}

type Data struct {
	CarsByDoor Cars
	TrucksByTon Trucks
}

type Cars map[string]int 
type Trucks map[string]int

