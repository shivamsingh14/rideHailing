package delivery

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func InitEndpoint() {

	router = httprouter.New()
	router.POST("/rider", RegisterRider)
	router.GET("/rider", ListRider)
	router.POST("/driver", RegisterDriver)
	router.GET("/driver", ListDriver)
	router.POST("/cab", CreateCab)
	router.GET("/cab", ListCabs)
	router.POST("/cab/location", UpdateCabLocation)
	router.POST("/cab/assign", AssignCab)
	router.POST("/available-cab", GetAvailableCabs)
	router.GET("/rider/trips", RiderTripHistory)
	router.GET("/driver/trips", DriverTripHistory)
	router.POST("/ride", BookRide)
	router.POST("/ride/complete", CompleteTrip)

	log.Printf("Starting server..")
	log.Fatal(http.ListenAndServe(":8080", router))

}
