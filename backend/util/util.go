package util

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Data interface {
}

func GenerateSessionToken() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("Error generating session token: %v", err)
	}
	return hex.EncodeToString(b)
}

func ImageSave(w http.ResponseWriter, r *http.Request, key string, action string) {
	file, _, err := r.FormFile("image")
	if err != nil {
		return
	}
	defer file.Close()
	var imagePath string
	switch action {
	case "register":
		imagePath = filepath.Join(".", "pkg", "db", "images", key+".jpg")
		// v is now of type *model.RegistrationData
	case "post":
		imagePath = filepath.Join(".", "pkg", "db", "images", "posts", key+".jpg")
		// Handle PostData
		// v is now of type *model.PostData
	case "comment":
		imagePath = filepath.Join(".", "pkg", "db", "images", "comments", key+".jpg")
	// Handle CommentData
	// v is now of type *model.CommentData
	case "group":
		imagePath = filepath.Join(".", "pkg", "db", "images", "groups", key+".jpg")
	default:
		// Handle other types
	}

	// Define the relative path to the images directo
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
	// 5. Replace the regData.AvatarURL with the path of the saved imag

	// 4. Get the path of the saved image
	// 5. Replace the regData.AvatarURL with the path of the saved image
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
