package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type User struct {
	id string  `json:"id"`
	name string `json:"name"`
	phone int `json:"phone"`
	email string `json:"email"`
}

func main() {
	input := []map[string][]int{
		{
			"cristian": []int{1, 2},
			"jorge": []int{1, 3},
			"felipe": []int{1, 2},
		},
		{
			"cristian": []int{3, 4},
			"jorge": []int{3, 4},
			"felipe": []int{3, 4},
			"samuel": []int{3, 4},
		},
		{
			"cristian": []int{5, 6},
			"jorge": []int{5, 6},
			"felipe": []int{5, 6},
			"samuel": []int{5, 6},
			"jonathan": []int{5, 6},
		},
	}

	output := make(map[string][]int)

	for _, object := range input {
		for key, value := range object {
			output[key] = append(output[key], value...)
		}
	}

	fmt.Println(output)

	response, err := http.Get("http://localhost:3000/users")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		var user User
		a := json.NewDecoder(response.Body).Decode(&user)
		//rr := response.Body
		//data, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(data))

		fmt.Println(a)
		/*
		for _, item := range rr {
			fmt.Println(item)
		}*/

	}


}