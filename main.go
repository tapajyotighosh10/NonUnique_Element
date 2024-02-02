package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/nonunique", NonUniqueHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func NonUniqueHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Array1 []int `json:"array1"`
		Array2 []int `json:"array2"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	map1 := make(map[int]int)
	map2 := make(map[int]int)
// new
	for _, num := range requestData.Array1 {
		map1[num]++
	}

	for _, num := range requestData.Array2 {
		map2[num]++
	}
	var nonUniqueArray []int

	for num, count := range map1 {
		if count == 1 && map2[num] == 0 {
			nonUniqueArray = append(nonUniqueArray, num)
		}
	}

	for num, count := range map2 {
		if count == 1 && map1[num] == 0 {
			nonUniqueArray = append(nonUniqueArray, num)
		}
	}

	responseData := struct {
		NonUniqueArray []int `json:"non_unique_array"`
	}{
		NonUniqueArray: nonUniqueArray,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}
