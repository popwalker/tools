package main

import (
	"strconv"
	"fmt"
	"github.com/jeffail/tunny"
	"time"
	"math/rand"
	"log"
	"net/http"
)

func main() {
	http.Handle("/runWorker", http.HandlerFunc(myhandle))
	http.ListenAndServe(":1234", nil)
}

func myhandle(w http.ResponseWriter, r *http.Request) {
	go run()
	w.Write([]byte("running some worker"))
}

func run() {
	var urls []string
	// make some url list
	for i := 1; i <= 100; i++ {
		tmp := "url_" + strconv.Itoa(i) + ".html"
		urls = append(urls, tmp)
	}

	// create a pool with a handle func: send
	pool, _ := tunny.CreatePool(10, send).Open()

	//defer pool.Close()

	// circulate url list and async send to worker,handle response with func: callback
	for _, url := range urls {
		pool.SendWorkAsync(url, callback)
	}
}

// process with the data
func send(object interface{}) interface{} {
	input := object.(string)
	sleepNumber := rand.Intn(5)
	fmt.Println(time.Now().String() + "got data " + input + ", and will sleep " + strconv.Itoa(sleepNumber) + " second")
	time.Sleep(time.Duration(sleepNumber) * time.Second)
	output := input
	return output
}

// process with the response
func callback(resp interface{}, err error) {
	res := resp.(string)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("got response data:" + res)
}
