package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


type todo struct {
	UserID int `json:"userId"`
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func Get()  {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	var todo todo
	json.Unmarshal(bodyBytes, &todo)
	fmt.Println(todo)
}

func Post()  {
	_todo := todo{1,1,"Marketing Strategies", false}
	bodyBytes, _ := json.Marshal(_todo)

	response, err := http.Post("https://jsonplaceholder.typicode.com/todos","application/json;charset=utf-8",bytes.NewBuffer(bodyBytes))
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	bodyBytes, _ = ioutil.ReadAll(response.Body)
	var todoRes todo 
	json.Unmarshal(bodyBytes, &todoRes)
	fmt.Println(todoRes)
}