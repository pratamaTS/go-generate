package template

// func Auth() mux.MiddlewareFunc {
// 	return func(h http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			jojoJWT, isValid, isExpired := jwt.ParseJojoJWT(r)
// 			if !isValid || isExpired {
// 				response.JSON(w, errors.New("bad credential").Error(), http.StatusUnauthorized)
// 				return
// 			}

// 			h.ServeHTTP(w, addItemToContext(r, jwtKey, jojoJWT))
// 		})
// 	}
// }
