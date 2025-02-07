package all_handlers

import (
	"encoding/json"
	"github.com/cozyCodr/hng-stage0/types"
	"log"
	"net/http"
	"time"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	response := &types.Response{
		Email:           "brightl.dev@gmail.com",
		CurrentDateTime: time.Now().Format(time.RFC3339),
		GithubUrl:       "https://github.com/cozyCodr/hng-stage0",
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
