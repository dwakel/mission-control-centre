package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"mission-control-center/interfaces"
	"mission-control-center/models"
	"mission-control-center/response"
	"net/http"
	"strconv"
)

type ApplicationsController struct {
	logger  *log.Logger
	appRepo *interfaces.IApplicationRepository
}

func NewApplicationsController(logger *log.Logger, appRepo *interfaces.IApplicationRepository) *ApplicationsController {
	return &ApplicationsController{
		logger,
		appRepo,
	}
}

// swagger:route GET Application
// Submits token
// Sample request:
//
//
//responses:
//	200: Success
//	500: If the request processing fails due to an exception

// Returns a redirect to Single Application
func (this *ApplicationsController) Create(rw http.ResponseWriter, r *http.Request) {
	var request models.Application
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		this.logger.Println("Invalid request ", err)
		rw.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(response.Response{Success: true, Message: "Invalid request"})
		rw.Write(resp)
		return
	}

	err = (*this.appRepo).Create(request)

	if err != nil {
		this.logger.Println("Invalid request ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(response.Response{Success: true, Message: "Unexpected error occurred"})
		rw.Write(resp)
		return
	}

	this.logger.Println("Journey Successful")

	rw.WriteHeader(http.StatusCreated)
	resp, _ := json.Marshal(response.Response{Success: true, Message: "Success"})
	rw.Write(resp)
	return
}

func (this *ApplicationsController) Get(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		this.logger.Println("Invalid request id is missing in parameters")
		rw.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(response.Response{Success: true, Message: "Invalid request id is missing in parameters"})
		rw.Write(resp)
		return
	}
	applicationId, err := strconv.Atoi(id)
	if err != nil {
		this.logger.Println("Invalid id parameter")
		rw.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(response.Response{Success: true, Message: "Invalid request id parameter"})
		rw.Write(resp)
		return
	}
	application, err := (*this.appRepo).Get(int64(applicationId))

	if err != nil {
		this.logger.Println("Invalid request ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(response.Response{Success: true, Message: "Unexpected error occurred"})
		rw.Write(resp)
		return
	}

	this.logger.Println("successfully fetch application data")

	rw.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(response.Response{Success: true, Message: "Success", Data: application})
	rw.Write(resp)
	return
}

func (this *ApplicationsController) List(rw http.ResponseWriter, r *http.Request) {
	application, err := (*this.appRepo).List()

	if err != nil {
		this.logger.Println("Invalid request ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(response.Response{Success: true, Message: "Unexpected error occurred"})
		rw.Write(resp)
		return
	}

	this.logger.Println("successfully fetch application data")

	rw.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(response.Response{Success: true, Message: "Success", Data: application})
	rw.Write(resp)
	return
}

func (this *ApplicationsController) GetMetadata(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		this.logger.Println("Invalid request id is missing in parameters")
		rw.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(response.Response{Success: true, Message: "Invalid request id is missing in parameters"})
		rw.Write(resp)
		return
	}
	applicationId, err := strconv.Atoi(id)
	if err != nil {
		this.logger.Println("Invalid id parameter")
		rw.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(response.Response{Success: true, Message: "Invalid request id parameter"})
		rw.Write(resp)
		return
	}
	application, err := (*this.appRepo).Get(int64(applicationId))

	if err != nil {
		this.logger.Println("Invalid request ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(response.Response{Success: true, Message: "Unexpected error occurred"})
		rw.Write(resp)
		return
	}

	this.logger.Println("successfully fetch application data")

	rw.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(response.Response{Success: true, Message: "Success", Data: application})
	rw.Write(resp)
	return
}
