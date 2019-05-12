package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	response1, _ := http.Get("http://httpbin.org/anything")
	content, _ := ioutil.ReadAll(response1.Body)
	fmt.Println(string(content))

	response2, _ := http.Post("http://httpbin.org/anything","application/json", bytes.NewReader([]byte(`{"name":"golang"'}`)))
	content2, _ := ioutil.ReadAll(response2.Body)
	fmt.Println(string(content2))

	client := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*3) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(5 * time.Second)) //设置发送接收数据超时
				return c, nil
			},
		},
	}
	request, _ := http.NewRequest(http.MethodGet, "http://httpbin.org/anything", nil)
	A, _ := client.Get("http://httpbin.org/anything")
	B, _ := client.Do(request)
	contentA, _ := ioutil.ReadAll(A.Body)
	contentB, _ := ioutil.ReadAll(B.Body)
	fmt.Println(string(contentA))
	fmt.Println(string(contentB))
	StartServer()
}


type My struct {

}

func (H My) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("Hello Again"))
}

func middleWare( next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("[Web-Server]: %s | %s", request.Host + request.RequestURI, time.Now().Format(time.RFC850))
		next.ServeHTTP(writer, request)
	}
}

func say(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello "))
}

func StartServer(){
	http.HandleFunc("/", middleWare(say))
	var a My
	http.HandleFunc("/hello", middleWare(a.ServeHTTP))
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/2", middleWare(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello Golang 2"))
	}))
	serverMux.HandleFunc("/hello/2", middleWare(a.ServeHTTP))
	go func() {
		log.Fatal(http.ListenAndServe(":8081", serverMux))
	}()

	server := &http.Server{
		Addr:":9090",
		Handler:serverMux ,
	}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	select {}
}

type self struct {

}

func (s self) ServeHTTP(writer http.ResponseWriter, req *http.Request){}