// https://www.digitalocean.com/community/tutorials/how-to-use-json-in-go
// https://blog.logrocket.com/using-json-go-guide/
// https://www.sohamkamani.com/golang/json/
// https://www.sohamkamani.com/golang/json-encode-decode/
// https://www.sohamkamani.com/golang/json-marshalling/

// https://blog.logrocket.com/using-json-go-guide/ !!!

package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {

	playground()
	another()
	unmarsh()
	unmarsh2()
	unmarshaller()

}

func playground() {
	data := map[string]interface{}{
		"intValue":    1234,
		"boolValue":   true,
		"stringValue": "hello!",
		"DateValue":   time.Date(2022, 3, 2, 9, 10, 0, 0, time.UTC),
		"ObjectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	fmt.Printf("json data: %s\n", jsonData)

	//var data map[string]interface{}
	error := json.Unmarshal([]byte(jsonData), &data)
	if error != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}

	fmt.Printf("json map: %v\n", data)
	fmt.Printf("json struct: %#v\n", data)
	fmt.Printf("dateValue: %#v\n", data["DateValue"])     // Access "DateValue" using map indexing
	fmt.Printf("objectValue: %#v\n", data["ObjectValue"]) // Access "ObjectValue" using map indexing

	fileCount := map[string]int{
		"cpp": 10,
		"js":  8,
		"go":  10,
	}
	bytes, _ := json.Marshal(fileCount)
	fmt.Println(string(bytes))

}

func another() {

	type Seller struct {
		Id      int
		Name    string
		Country string
	}

	type Product struct {
		Id     int
		Name   string
		Seller Seller
		Price  int
	}

	products := []Product{
		Product{
			Id:     50,
			Name:   "Writing Book",
			Seller: Seller{1, "ABC Company", "US"},
			Price:  100,
		},
		Product{
			Id:     51,
			Name:   "Kettle",
			Seller: Seller{20, "John Store", "DE"},
			Price:  500,
		},
	}
	bytes, _ := json.Marshal(products)
	fmt.Println(string(bytes))
}

func unmarsh() {

	type Window struct {
		Width  int    `json:"width"`
		Height int    `json:"height"`
		Title  string `json:"title"`
	}

	jsonInput := `{
		"width": 500,
		"height": 200,
		"title": "Hello Go!"
	}`

	var window Window
	err := json.Unmarshal([]byte(jsonInput), &window)

	if err != nil {
		fmt.Println("JSON decode error!")
		return
	}

	fmt.Println(window) // {500 200 Hello Go!}

}

func unmarsh2() {
	jsonInput := `{
        "apples": 10,
        "mangos": 20,
        "grapes": 20
    }`
	var fruitBasket map[string]int
	err := json.Unmarshal([]byte(jsonInput), &fruitBasket)

	if err != nil {
		fmt.Println("JSON decode error!")
		return
	}

	fmt.Println(fruitBasket) // map[apples:10 grapes:20 mangos:20]
}

// Unmarshalling more complex data
type Product struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Seller struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		CountryCode string `json:"country_code"`
	} `json:"seller"`
	Price int `json:"price"`
}

func unmarshaller() {
	jsonInput := `[
    {
        "id":50,
        "name":"Writing Book",
        "seller":{
            "id":1,
            "name":"ABC Company",
            "country_code":"US"
        },
        "price":100
    },
    {
        "id":51,
        "name":"Kettle",
        "seller":{
            "id":20,
            "name":"John Store",
            "country_code":"DE"
        },
        "price":500
    }]
    `
	var products []Product
	err := json.Unmarshal([]byte(jsonInput), &products)

	if err != nil {
		fmt.Println("JSON decode error!")
		return
	}

	fmt.Println(products)
	// [{50 Writing Book {1 ABC Company US} 100} {51 Kettle {20 John Store DE} 500}]
}

// Writing JSON files to the filesystem
// In previous examples, we printed the encoded JSON content to the console via the Println function. Now we can save those JSON strings as files with the ioutil.WriteFile function, as shown below.

/*

package main
import (
    "io/ioutil"
    "encoding/json"
)
type Window struct {
    Width int `json:"width"`
    Height int `json:"height"`
    X int `json:"x"`
    Y int `json:"y"`
}
type Config struct {
    Timeout float32 `json:"timeout"`
    PluginsPath string `json:"pluginsPath"`
    Window Window `json:"window"`
}
func main() {
    config := Config {
        Timeout: 40.420,
        PluginsPath: "~/plugins/etc",
        Window: Window {500, 200, 20, 20},
    }
    bytes, _ := json.MarshalIndent(config, "", "  ")
    ioutil.WriteFile("config.json", bytes, 0644)
}

*/

//The above code writes config.json by encoding the config object as a JSON object. Here we used two spaces for indentation and converted struct fields to camel case JSON keys by using struct tags.
