package util

import (
	"backend/pkg/model"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func GenerateSessionToken() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("Error generating session token: %v", err)
	}
	return hex.EncodeToString(b)
}

func ImageParser(w http.ResponseWriter, r *http.Request, regData *model.RegistrationData) {

    // 2. Extract the image file from the form data
    file, _, err := r.FormFile("avatar") 
    if err != nil {
        http.Error(w, "Error retrieving file from form data: "+err.Error(), http.StatusBadRequest)
        return
    }
    defer file.Close()
    // Define the relative path to the images directory
    imagePath := filepath.Join(".", "pkg", "db", "images", regData.Username+".jpg")
    out, err := os.Create(imagePath)
    if err != nil {
        http.Error(w, "Error creating file: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer out.Close()

    _, err = io.Copy(out, file)
    if err != nil {
        http.Error(w, "Error saving file: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // 4. Get the path of the saved image
    // 5. Replace the regData.AvatarURL with the path of the saved image
    regData.AvatarURL = imagePath
	
}