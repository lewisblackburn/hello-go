package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Ping string
	Time int64
}

func main() {
	url := os.Args[1]

	api(url)
}

func api(url string) {
	res, err := http.Get(url)

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var data Response
	json.Unmarshal(body, &data)

	fmt.Printf("returned %v on %v", data.Ping, time.Unix(data.Time, 0))
}
