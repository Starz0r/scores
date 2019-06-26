package routing

import (
	"github.com/spidernest-go/mux"
)

var r *echo.Echo

const ErrGeneric = `{"errno": "404", "message": "Bad Request"}`

func ListenAndServe() {
	r = echo.New()

	v0 := r.Group("/api/v0")
	v0.POST("/score", createScore)
	v0.GET("/score/:id", getScore)
	v0.GET("/score", getScoresFromBoard)
	v0.DELETE("/score/:id", deleteScore)

	r.Start(":5000")
}
