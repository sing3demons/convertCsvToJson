package main

// import (
// 	"encoding/csv"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"os"
// )

// type Blog struct {
// 	ID     string `json:"id"`
// 	Title  string `json:"title"`
// 	Author string `json:"author"`
// }

// func main() {
// 	file, err := os.Open("test.csv")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer file.Close()

// 	// Create a new reader.
// 	reader := csv.NewReader(file)
// 	reader.FieldsPerRecord = -1
// 	data, err := reader.ReadAll()
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	blogList := []Blog{}

// 	for index, each := range data {
// 		var blog Blog
// 		if index > 0 {
// 			blog.ID = each[0]
// 			blog.Title = each[1]
// 			blog.Author = each[2]
// 			blogList = append(blogList, blog)
// 		}
// 	}

// 	fmt.Println(blogList)

// 	b, err := json.Marshal(blogList)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	http.HandleFunc("/blog", func(w http.ResponseWriter, _ *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write([]byte(b))
// 	})

// 	http.ListenAndServe(":8080", nil)
// }
