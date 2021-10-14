package delivery

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func InitEndpoint() {

	router = httprouter.New()
	router.POST("/rider", RegisterRider)
	router.POST("/driver", RegisterDriver)
	// TODO: router.POST("/cab", RegisterCab)
	router.POST("/cab", RegisterCab)

	log.Fatal(http.ListenAndServe(":8080", router))

}
