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
	router.POST("/cab", CabRepository)
	router.GET("/rider/trips", RiderTripHistory)
	router.GET("/driver/trips", DriverTripHistory)
	router.POST("/ride", BookRide)
	router.PATCH("/ride/complete", CompleteTrip)
	router.PATCH("/cab/location", UpdateCabLocation)

	log.Fatal(http.ListenAndServe(":8080", router))

}
