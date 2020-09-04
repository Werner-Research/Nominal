package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	key, active := os.LookupEnv("NOMINAL_KEY")
	if active {
		content, err := json.Marshal(map[string]string{
			"value1": "Executed successfully!",
		})
		if err != nil {
			os.Exit(2)
		}
		resp, err := http.Post("https://maker.ifttt.com/trigger/terminal_trigger/with/key/"+key, "application/json", bytes.NewBuffer(content))
		if err != nil {
			fmt.Printf("Something went wrong: %v \n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("ok?")
		fmt.Printf("%v", string(body))
	} else {
		fmt.Println("You forgot to set your 'NOMINAL_KEY' var")
		os.Exit(1)
	}
}
