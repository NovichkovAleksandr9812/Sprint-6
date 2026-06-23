package handlers

import (
	"net/http"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"time"
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func ReturnHandle(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("../index.html")
	if err != nil {
		http.Error(w, "Error" + err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err = tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error" + err.Error(), http.StatusInternalServerError)
		return
	}
}

func ConvertHandle(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error" + err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error" + err.Error(), http.StatusInternalServerError)
		return
	}

	fileContent := string(fileData)

	conv := service.DefinOfContent(fileContent)
	

	filename := time.Now().Format("2006-01-02_15-04-05") + filepath.Ext(header.Filename)
	outFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Error" + err.Error(), http.StatusInternalServerError)
		return
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, strings.NewReader(conv))
	if err != nil {
		http.Error(w, "Error" + err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, conv)
}
