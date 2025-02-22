package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func fetchAttendanceData() {
	resp, err := http.Get("http://192.235.252.10/get_attendance")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Data dari server:", string(body))
}

func main() {
	fetchAttendanceData()
}
