package util

import (
	"backend/pkg/model"
	"crypto/rand"
	"encoding/hex"
	"fmt"
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

func ImageSave(w http.ResponseWriter, r *http.Request, regData *model.RegistrationData) {

	// 2. Extract the image file from the form data
	file, _, err := r.FormFile("avatar")
	if err != nil {
		fmt.Println("No image uploaded on register")
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
	// TODO: Change this to the actual server URL
	regData.AvatarURL = "http://localhost:8080/images/" + regData.Username + ".jpg"

}

func GetSessionToken(r *http.Request) string {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return ""
		}
		log.Fatalf("Error getting session token: %v", err)
	}
	return cookie.Value
}
