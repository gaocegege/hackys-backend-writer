package api

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/gaocegege/hackys-backend-writer/cognitive"
	"github.com/gaocegege/hackys-backend-writer/pkg/log"
)

func createPoint(request *restful.Request, response *restful.Response) {
	point := &Point{}
	err := request.ReadEntity(point)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusBadRequest, "Writer Service unable to parse request body")
		return
	}

	log.Infof("Point: %s\n", point)
	go PushPointToDB(point)
	var createPointResponse CreatePointResponse
	createPointResponse.ErrorMessage = ""
	response.WriteHeaderAndEntity(http.StatusAccepted, createPointResponse)
}

func PushPointToDB(point *Point) {
	text := point.Content
	result, err := cognitive.RecognizeText(text)
	if err != nil {
		log.Warnf("Err when push point to DB: %v", err)
		return
	}
	log.Infof("Result: %v", result)
}
