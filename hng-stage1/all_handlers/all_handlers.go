package all_handlers

import (
	"encoding/json"
	"github.com/cozyCodr/hng-stage0/types"
	"github.com/cozyCodr/hng-stage0/utils"
	"net/http"
	"strconv"
)

func ClassifyNumberHandler(w http.ResponseWriter, r *http.Request) {
	// Handle CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Get number parameter
	numberStr := r.URL.Query().Get("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(types.ErrorResponse{Number: numberStr, Error: true})
		if err != nil {
			return
		}
		return
	}

	// Classify number
	isPrime := utils.IsPrime(number)
	isPerfect := utils.IsPerfect(number)
	isArmstrong := utils.IsArmstrong(number)
	isOdd := number%2 != 0
	var properties []string
	if isArmstrong {
		properties = append(properties, "armstrong")
	}
	if isOdd {
		properties = append(properties, "odd")
	} else {
		properties = append(properties, "even")
	}
	digitSum := utils.DigitSum(number)

	// Get fun fact
	funFact, err := utils.GetFunFact(number)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create response
	response := types.Response{
		Number:     number,
		IsPrime:    isPrime,
		IsPerfect:  isPerfect,
		Properties: properties,
		DigitSum:   digitSum,
		FunFact:    funFact,
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
