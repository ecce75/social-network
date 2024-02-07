package middleware

// ----- DEPRECATED -----
// import (
// 	"backend/pkg/repository"
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// // CheckAuthMiddleware is a middleware function that checks if the user is authenticated.
// // It checks for the presence of a session token cookie in the request.
// // If the session cookie doesn't exist, it sets isAuthenticated to false and returns an unauthorized error.
// // If there is an error checking the session token, it returns an internal server error.
// // If the session token is valid, it stores the user ID in the request context and calls the next handler.
// func CheckAuthMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		cookie, err := r.Cookie("session_token")
// 		if err != nil {
//             if err == http.ErrNoCookie {
//                 // If the session cookie doesn't exist, set isAuthenticated to false
//                 http.Error(w, "User not authenticated", http.StatusUnauthorized)
//                 return
//             } else {
//                 http.Error(w, "Error checking session token: " + err.Error(), http.StatusInternalServerError)
//                 return
//             }
//         }
// 		type contextKey string

// 		const userIDKey contextKey = "AuthUserID"

// 		userID, err := ConfirmAuthentication(cookie)
// 		if err != nil {
// 			http.Error(w, "Error confirming authentication: "+err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		// Store the user ID in the request context
// 		ctx := context.WithValue(r.Context(), userIDKey, userID)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

// // ConfirmAuthentication checks the validity of a session token stored in a cookie.
// // It takes a pointer to an http.Cookie as input and returns the user ID associated with the session token and an error, if any.
// // If the session token is invalid or expired, it returns an error.
// // The function uses the GetSessionBySessionToken function from the repository package to retrieve the session information from the database.
// func ConfirmAuthentication(cookie *http.Cookie, sessionRepo *repository.SessionRepository) (int, error) {
// 	sessionToken := cookie.Value

// 	session, err := sessionRepo.GetSessionBySessionToken(sessionToken)
// 	if err != nil {
// 		fmt.Println("Could not get session token: ", err)
// 		return 0, err
// 	}
// 	if time.Now().After(session.ExpiresAt) {
// 		return 0, fmt.Errorf("session token expired: %v", session.ExpiresAt)
// 	}
// 	return session.UserID, nil
// }