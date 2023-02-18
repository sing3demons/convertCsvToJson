package main

// import (
// 	"encoding/csv"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// )

// type Blog struct {
// 	ID     string `json:"id"`
// 	Title  string `json:"title"`
// 	Author string `json:"author"`
// }

// func main() {
// 	blogList, err := readCSV("test.csv")
// 	if err != nil {
// 		log.Fatalf("Error reading CSV file: %v", err)
// 	}

// 	http.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")

// 		if r.Method == http.MethodGet {
// 			if err := sendJSON(w, blogList); err != nil {
// 				log.Printf("Error sending JSON response: %v", err)
// 			}
// 		} else {
// 			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		}
// 	})

// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func readCSV(filename string) ([]Blog, error) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return nil, fmt.Errorf("error opening file: %v", err)
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	reader.FieldsPerRecord = -1
// 	data, err := reader.ReadAll()
// 	if err != nil {
// 		return nil, fmt.Errorf("error reading CSV data: %v", err)
// 	}

// 	fmt.Println(data[0])

// 	blogList := []Blog{}
// 	for i, row := range data {
// 		if i == 0 {
// 			// Skip header row
// 			continue
// 		}

// 		blog := Blog{
// 			ID:     row[0],
// 			Title:  row[1],
// 			Author: row[2],
// 		}
// 		blogList = append(blogList, blog)
// 	}

// 	return blogList, nil
// }

// func sendJSON(w http.ResponseWriter, data interface{}) error {
// 	encoder := json.NewEncoder(w)
// 	if err := encoder.Encode(data); err != nil {
// 		return fmt.Errorf("error encoding JSON: %v", err)
// 	}
// 	return nil
// }