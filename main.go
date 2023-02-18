package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Blog struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {

	blogs, err := ReadCsv("test.csv")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		if err := toJson(w, blogs);err!=nil{
			log.Printf("Error sending JSON response: %v", err)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
func toJson(w http.ResponseWriter, data any) error {
	return json.NewEncoder(w).Encode(data)
}

func ReadCsv(filename string) ([]Blog, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	blogs := make([]Blog, 0)

	for index, data := range records {
		if index == 0 {
			continue
		}

		id, err := strconv.Atoi(data[0])
		if err != nil {
			return nil, fmt.Errorf("error opening file: %v", err)
		}

		blog := Blog{
			ID:     id,
			Title:  data[1],
			Author: data[2],
		}

		blogs = append(blogs, blog)
	}

	return blogs, nil
}
