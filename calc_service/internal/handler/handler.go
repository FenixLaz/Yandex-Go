package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"calc_service/internal/calculator"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req Request
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{ "error": "Invalid request format" }`, http.StatusUnprocessableEntity)
		return
	}

	result, err := calculator.Calculate(req.Expression)
	if err != nil {
		http.Error(w, `{ "error": "Expression is not valid" }`, http.StatusUnprocessableEntity)
		return
	}

	response := Response{Result: formatFloat(result)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


func formatFloat(value float64) string {
	return fmt.Sprintf("%.2f", value)
}
