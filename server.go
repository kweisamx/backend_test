package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

type ClientInfo struct {
	IP             string
	ResponseNumber int
	LastTime       int
}

func checkIPLimit(clients *[]ClientInfo, m *sync.Mutex) {
	for {
		for i, _ := range *clients {
			(*m).Lock()
			if (*clients)[i].LastTime > 0 {
				(*clients)[i].LastTime -= 1
			} else if (*clients)[i].LastTime == 0 && (*clients)[i].IP != "" {
				(*clients)[i].IP = ""
			}
			(*m).Unlock()
		}
		time.Sleep(time.Second)
	}
}

func search(ip string, clients []ClientInfo) int {
	for i, v := range clients {
		if v.IP == ip {
			return i
		}
	}
	return -1
}

func increment(ip string, clients *[]ClientInfo) int {
	var Index int = search(ip, (*clients))
	if Index == -1 {
		client := ClientInfo{ip, 1, 59}
		(*clients) = append(*(clients), client)
		return 1
	} else {
		(*clients)[Index].ResponseNumber += 1
		fmt.Println((*clients)[Index].IP, (*clients)[Index].ResponseNumber)
		return (*clients)[Index].ResponseNumber
	}
}

func handler(clients *[]ClientInfo, m *sync.Mutex) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		(*m).Lock()
		count := increment(ip, clients)
		(*m).Unlock()
		if count > 60 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "error")
		} else {
			fmt.Fprint(w, count)
		}
	})
}

func main() {

	var ClientInfoList []ClientInfo = make([]ClientInfo, 10)
	var mutex sync.Mutex

	go checkIPLimit(&ClientInfoList, &mutex)
	fmt.Println("Server start... ")
	http.HandleFunc("/", handler(&ClientInfoList, &mutex))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
