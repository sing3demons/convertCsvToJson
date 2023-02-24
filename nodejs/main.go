package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Book struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

func readCsv(filename string) ([]Book, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var books []Book
	for _, record := range records[1:] {
		books = append(books, Book{
			ID:     record[0],
			Author: record[1],
			Title:  record[2],
		})
	}

	return books, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		books, err := readCsv("test.csv")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		// var responseData []Book
		// for _, book := range books {
		// 	response := Book{
		// 		ID:     book.ID,
		// 		Author: book.Author,
		// 		Title:  book.Title,
		// 	}
		// 	responseData = append(responseData, response)
		// }

		json.NewEncoder(w).Encode(books)
	})

	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
