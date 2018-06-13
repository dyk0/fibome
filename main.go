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
	//Create the list of Fibonocci numbers of size n
	list := []*big.Int{big.NewInt(0)}
	for i := 1; i < n; i++ {
		list = append(list, fiboget(i))
	}
	return list
}
func fiboget(n int) *big.Int {
	//Calculate the fibonacci number at position N
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
	fmt.Fprint(w, "Welcome to FiboWho!\nTo unlock the secret of FiboWho\nEnter a Positive Number\n\n\nExample: ", r.Host, "/5\n")
	fmt.Fprint(w, fibolist(5))
}

func fib(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Convert the string to integer, if negative, float will raise 400
	i, _ := strconv.Atoi(p.ByName("size"))
	if i <= 0 {
		http.Error(w, "Invalid Entry! Please enter a positive integer", 400)
	} else {
		fmt.Fprint(w, fibolist(i))
	}
}

func main() {
	//Create the HTTP router
	router := httprouter.New()
	//Default / route to return a friendly message
	router.GET("/", index)
	//Take the paramter size to return fibonacci sequence of size
	router.GET("/:size", fib)

	log.Fatal(http.ListenAndServe(":8000", router))
}
