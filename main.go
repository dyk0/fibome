package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func fibolist(n int) []*big.Int {
	/*
		Fibonocci numbers sequence generator of size n
		Args:
		  n(int)           - The size of sequence
		Returns:
		  list([]*big.Int) - Sequence of Fibonocci numbers
	*/
	list := []*big.Int{big.NewInt(0)}
	for i := 1; i < n; i++ {
		list = append(list, fiboget(i))
	}
	return list
}
func fiboget(n int) *big.Int {
	/*
		Fibonocci number calculator for position n
		Args:
		  n(int)          - The position of the number to calculate
		Returns:
		  fn[n](*big.Int) - The calculated Fibonocci number
	*/
	fn := make(map[int]*big.Int)
	for i := 0; i <= n; i++ {
		var f = big.NewInt(0)
		if i <= 2 {
			f.SetUint64(1)
		} else {
			f = f.Add(fn[i-1], fn[i-2])
		}
		fn[i] = f
	}
	return fn[n]
}
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	/*
		HTTP Index for default route
		Returns:
			Welcome to FiboWho!
			To Unlock the secrets...enter a number


			Example: http://localhost:8000/5
			Response: [0 1 1 2 3]
	*/
	log.Println(r.Method, r.URL.Path, r.UserAgent())
	usage := fmt.Sprintf("Welcome to FiboWho!\nTo Unlock the secrets...enter a number\n\n\nExample: http://%s/5\nResponse: ", r.Host)
	fmt.Fprint(w, usage)
	sequence := fmt.Sprintf("%v\n", fibolist(5))
	fmt.Fprint(w, sequence)
}

func fib(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	/*
		The Fibster, Fibonator! This HTTP handler parses user's input
		Convert the string to integer, if negative, float will raise 400
	*/
	i, _ := strconv.Atoi(p.ByName("size"))
	if i <= 0 {
		http.Error(w, "Invalid Entry! Please enter a positive integer", 400)
		log.Println(r.Method, r.URL.Path, r.UserAgent())
	} else {
		sequence := fmt.Sprintf("%v\n", fibolist(i))
		fmt.Fprint(w, sequence)
		log.Println(r.Method, r.URL.Path, r.UserAgent())
	}
}

func Api() *httprouter.Router {
	/*
		Create the HTTP router

		Routes:
		  GET /      - Default route, returns friendly Index
		  GET /:size - Retrieve Fibonocci sequence of :size(int)
	*/
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/:size", fib)
	return router
}
func main() {
	//Start up the Server
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.LUTC)
	log.Println("Fibome started!")
	err := http.ListenAndServe(":8000", Api())
	log.Fatal(err)
}
