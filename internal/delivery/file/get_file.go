package file

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

func (fd *FilesDelivery) GetImage(w http.ResponseWriter, r *http.Request) {
	imageName := mux.Vars(r)["name"]

	log.Printf(" [ FilesDelivery.GetImage ] Getting image %s", imageName)
	file, err := fd.repo.GetFile(imageName)
	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Println("Error closing file")
		}
	}(file)

	contentType := "image/" + strings.TrimPrefix(filepath.Ext(imageName), ".")
	w.Header().Set("Content-Type", contentType)
	http.ServeFile(w, r, file.Name())
}
