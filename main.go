package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"social-network/database"
)

func main() {
	db.StartDB()
	router := mux.NewRouter()

	// Serve static files
	// for _, dir := range []string{"static", "js", "assets"} {
	// 	path := "/" + dir + "/"
	// 	mux.PathPrefix(path).Handler(http.StripPrefix(path, http.FileServer(http.Dir("frontend/public/"+dir))))
	// }
	// Catch-all route to serve index.html for all other routes
	// router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./frontend/public/static/index.html")
	// })
	// http.Handle("/", router)

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":9000", router)
}
