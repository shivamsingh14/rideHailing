package delivery

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/julienschmidt/httprouter"
// )

// type HttpServer struct {
// 	router *httprouter.Router
// }

// type HttpParams interface {
// 	ByName(string) string
// }

// type HttpResponseBody []byte

// type HttpHandler func(http.Request, HttpParams) (HttpResponseBody, error)

// func NewHttpServer() HttpServer {
// 	router := httprouter.New()
// 	return HttpServer{
// 		router: router,
// 	}
// }

// func routeHandler(path string, httpHandler HttpHandler, server HttpServer) httprouter.Handle {
// 	routeHandler := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

// 	}
// 	return routeHandler
// }

// func (s HttpServer) POST(path string, httpHandler HttpHandler) {

// 	fmt.Println("bp4")

// 	routeHandler := routeHandler(path, httpHandler, s)
// 	s.router.POST(path, routeHandler)
// }
