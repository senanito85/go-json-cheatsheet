package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Dog struct to hold attributes of a dog
type Dog struct {
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Owner        string `json:"owner"`
	Color        string `json:"color"`
	LocationCity string `json:"locationCity"`
	ID           int    `json:"id"`
}

func readJSONFile(filename string) ([]Dog, error) {
	var dogs []Dog

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &dogs); err != nil {
		return nil, err
	}

	return dogs, nil
}

func main() {
	// Specify the path to your JSON file
	filename := "dogs.json"

	dogs, err := readJSONFile(filename)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	//fmt.Println(dogs)

	//Print each dog's details
	for _, dog := range dogs {
		fmt.Printf("Name: %s\n", dog.Name)
		fmt.Printf("Age: %d\n", dog.Age)
		fmt.Printf("Owner: %s\n", dog.Owner)
		fmt.Printf("Color: %s\n", dog.Color)
		fmt.Printf("Location: %s\n", dog.LocationCity)
		fmt.Printf("ID: %d\n", dog.ID)
		fmt.Println("------------------------")
	}
}
