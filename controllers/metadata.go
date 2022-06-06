package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"mission-control-center/interfaces"
	"mission-control-center/models"
	"mission-control-center/response"
	"net/http"
	"strconv"
)

type MetadataController struct {
	logger      *log.Logger
	metaRepo    *interfaces.IMetadataRepository
	metaService *interfaces.IMetadataService
}

func NewMetadataController(logger *log.Logger, metaRepo *interfaces.IMetadataRepository, metaService *interfaces.IMetadataService) *MetadataController {
	return &MetadataController{
		logger,
		metaRepo,
		metaService,
	}
}

// swagger:route GET Metadata
// Submits token
// Sample request:
//
//
//responses:
//	200: Success
//	500: If the request processing fails due to an exception

// Returns a redirect to Single Application
func (this *MetadataController) Create(rw http.ResponseWriter, r *http.Request) {
	var request models.Metadata
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		this.logger.Println("Invalid request ", err)
		rw.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(response.Response{Success: true, Message: "Invalid request"})
		rw.Write(resp)
		return
	}

	err = (*this.metaRepo).Create(request)

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
func (this *MetadataController) Update(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		this.logger.Println("Invalid request id is missing in parameters")
		rw.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(response.Response{Success: false, Message: "Invalid request id is missing in parameters"})
		rw.Write(resp)
		return
	}
	metadataId, err := strconv.Atoi(id)
	if err != nil {
		this.logger.Println("Invalid id parameter")
		rw.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(response.Response{Success: false, Message: "Invalid request id parameter"})
		rw.Write(resp)
		return
	}
	var request models.Metadata
	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		this.logger.Println("Invalid request ", err)
		rw.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(response.Response{Success: false, Message: "Invalid request"})
		rw.Write(resp)
		return
	}
	err = (*this.metaService).UpdateTransaction(int64(metadataId), request)

	if err != nil {
		this.logger.Println("Failed to update request ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(response.Response{Success: false, Message: "Unexpected error occurred"})
		rw.Write(resp)
		return
	}

	this.logger.Println("Journey Successful")

	rw.WriteHeader(http.StatusCreated)
	resp, _ := json.Marshal(response.Response{Success: true, Message: "Success"})
	rw.Write(resp)
	return
}

func (this *MetadataController) Get(rw http.ResponseWriter, r *http.Request) {
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
	application, err := (*this.metaRepo).Get(int64(applicationId))

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

func (this *MetadataController) List(rw http.ResponseWriter, r *http.Request) {
	application, err := (*this.metaRepo).List()

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

func (this *MetadataController) GetHistory(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		this.logger.Println("Invalid request id is missing in parameters")
		rw.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(response.Response{Success: true, Message: "Invalid request id is missing in parameters"})
		rw.Write(resp)
		return
	}
	metadataId, err := strconv.Atoi(id)
	if err != nil {
		this.logger.Println("Invalid id parameter")
		rw.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(response.Response{Success: true, Message: "Invalid request id parameter"})
		rw.Write(resp)
		return
	}

	limitString := r.URL.Query().Get("limit")
	this.logger.Println(fmt.Sprintf("Limit: %s", limitString))
	//todo: Implement logging into mongodb
	if limitString == "" {
		this.logger.Println("No limit was provided, assignin default limit")
		limitString = "-1"
	}

	limit, err := strconv.Atoi(id)
	if err != nil {
		this.logger.Println("Invalid limit parameter")
		rw.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(response.Response{Success: true, Message: "Invalid request id parameter"})
		rw.Write(resp)
		return
	}

	history, err := (*this.metaRepo).GetMetadataHistory(int64(metadataId), int32(limit))

	if err != nil {
		this.logger.Println("Invalid request ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(response.Response{Success: true, Message: "Unexpected error occurred"})
		rw.Write(resp)
		return
	}

	this.logger.Println("successfully fetch application data")

	rw.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(response.Response{Success: true, Message: "Success", Data: history})
	rw.Write(resp)
	return
}
