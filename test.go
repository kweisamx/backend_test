//test.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func reqeust_test() {
	for i := 1; i <= 70; i++ {
		res, err := http.Get("http://127.0.0.1:8080")
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var expected string
		if i <= 60 {
			expected = strconv.Itoa(i)
		} else {
			expected = "error"
		}
		if string(resBody) != expected {
			log.Fatal("handler returned unexpected body: got %v want %v",
				string(resBody), expected)
		}
	}
}

func main() {
	reqeust_test()
	fmt.Println("First Request is OK, wait 60 seconds")
	time.Sleep(time.Minute) // wait for 60 sec

	reqeust_test()
	fmt.Println("Second Reqeust is OK, test finish")

}
