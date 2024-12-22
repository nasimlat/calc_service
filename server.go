package main

import (
	"encoding/json"
	"net/http"
    "fmt"
	"github.com/nasimlat/calc_service/pkg/calc"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result"`
	Error string `json:"error"`
}

func caclculateHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	var res Response

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		res.Error = "Invalid request"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	result, err := calc.Calc(req.Expression)
	if err != nil {
		res.Error = "Expression is not valid"
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Result = fmt.Sprintf("%f", result)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}


func main() {

	http.HandleFunc("/api/v1/calculate", caclculateHandler)
	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("internal server error: %v\n", err)
	}
}