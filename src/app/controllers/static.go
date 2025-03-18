package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"path"

	"websocket/src/app/models"
)

func Static(folder string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := models.Error{
			Code:     http.StatusFound,
			Message:  "success",
			Redirect: "",
		}


		filePath := path.Join(folder, r.URL.Path)

		fileinfo, err := os.Stat(filePath)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			response.Code = http.StatusNotFound
			response.Message = "Page not found!"
			json.NewEncoder(w).Encode(response)
			return
		}

		if fileinfo.IsDir() {
			w.Header().Set("Content-Type", "application/json")
			response.Code = http.StatusForbidden
			response.Message = "Forbidden route!"
			json.NewEncoder(w).Encode(response)
			return
		}

		http.ServeFile(w, r, filePath)
	}
}
