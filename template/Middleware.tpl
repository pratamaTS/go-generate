package main

// var middlewareMap = map[string]mux.MiddlewareFunc{
// 	"auth": Auth(),
// }

// func HandleMiddleware(h http.Handler, middlewares []string) mux.MiddlewareFunc {
// 	m := make([]mux.MiddlewareFunc, 0)
// 	for _, v := range middlewares {
// 		datum, ok := middlewareMap[v]
// 		if !ok {
// 			continue
// 		}
// 		m = append(m, datum)
// 	}

// 	return func(next http.Handler) http.Handler {
// 		return HandleJojoMiddleware(next, m...)
// 	}
// }
