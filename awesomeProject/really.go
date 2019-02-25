package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Data struct {
	User []UsersArray `json:"data"`
}

type UsersArray struct {
	Name string `json:"name"`
	Phone int `json:"phone"`
	Email string `json:"email"`
}

type Sensor struct {
	Variables []Variable `json:"variables"`
}

type Variable struct {
	Name string `json:"name"`
	Measurements []Measurement `json:"measurements"`
}

type Measurement struct {
	Value float32 `json:"value"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	response, err := http.Get("http://localhost:3000/users")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	responseData, _ := ioutil.ReadAll(response.Body)

	var responseObject Data
	json.Unmarshal(responseData, &responseObject)

	output := make(map[string][]UsersArray)

	for _, value := range responseObject.User {
		output[value.Name] = append(output[value.Name], value)
	}

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
	json.NewEncoder(w).Encode(output)
}

func main() {

	//fmt.Println(responseObject.User.Name)

	//for i := 0; i < len(responseObject.User);i++ {
	//	fmt.Println(responseObject.User[i])
	//}

	r := mux.NewRouter()
	r.HandleFunc("/get_users_filtered", getUsers).Methods("GET")
	log.Fatal(http.ListenAndServe(":3001", r))

}