package middleware

import (
	"fmt"
	"net/http"
)

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// ここで認証の処理を行う
		fmt.Println("Authentificated")

		next.ServeHTTP(w, r)
	})

}
