package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, map[string]string{"error": message})
}

func RenderPage(tmpName string, w http.ResponseWriter, tmp *template.Template, data TemplateData) {
	err := tmp.ExecuteTemplate(w, tmpName, data)
	if err != nil {
		fmt.Println("masuuukl", err.Error())
		//http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type TemplateData struct {
	IsLogin      bool
	Title        string
	Data         map[string]interface{}
	Error        error
	ErrorMessage string
}
