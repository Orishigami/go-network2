package work2

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Work2() {
	r1 := mux.NewRouter()
	r1.HandleFunc("/cal/{number1}/plus/{number2}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		number1, _ := strconv.Atoi(vars["number1"])
		number2, _ := strconv.Atoi(vars["number2"])
		sum := number1 + number2

		fmt.Fprintf(w, "%d + %d = %d", number1, number2, sum)
	})

	http.ListenAndServe(":80", nil)
}
