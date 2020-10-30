package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"ajiiranetservice/constants"
	"ajiiranetservice/utils"
	"ajiiranetservice/vo"
)

// AJIIRANetHandler :
type AJIIRANetHandler int

// ProcessData :
func (ajiiraNetHandler *AJIIRANetHandler) ProcessData(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())

		return
	}

	ajRequestBodyVO := utils.TextToRequestJSON(string(b))

	if ajRequestBodyVO.Path == "/devices" && ajRequestBodyVO.RequsetType == constants.CREATE {
		ajiiraNetHandler.CreateDevice(w, r, ajRequestBodyVO)
	} else if ajRequestBodyVO.Path == "/connections" && ajRequestBodyVO.RequsetType == constants.CREATE {
		ajiiraNetHandler.CreateConnection(w, r, ajRequestBodyVO)
	} else if ajRequestBodyVO.Path == "/devices" && ajRequestBodyVO.RequsetType == constants.FETCH {
		ajiiraNetHandler.FetchDevices(w, r, ajRequestBodyVO)
	} else if strings.HasPrefix(ajRequestBodyVO.Path, "/devices/") && ajRequestBodyVO.RequsetType == constants.MODIFY {
		ajiiraNetHandler.UpdateStrength(w, r, ajRequestBodyVO)
	} else if strings.HasPrefix(ajRequestBodyVO.Path, "/info-routes?") && ajRequestBodyVO.RequsetType == constants.FETCH {
		ajiiraNetHandler.FetchRouteInfo(w, r, ajRequestBodyVO)
	} else {
		utils.RespondJSON(w, http.StatusBadRequest, &vo.AJIRAErrorMsg{Error: getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorInvalidRequest)})
	}

}

// CreateDevice :
func (ajiiraNetHandler *AJIIRANetHandler) CreateDevice(w http.ResponseWriter, r *http.Request, ajRequestBodyVO *vo.AJRequestBodyVO) {

	var processor AJIIRANetProcessor
	successMsg, ajiraError := processor.CreateDevice(ajRequestBodyVO)
	if ajiraError != nil {
		utils.RespondJSON(w, ajiraError.HTTPStatus, ajiraError)
		return
	}

	utils.RespondJSON(w, http.StatusCreated, successMsg)

}

// CreateConnection :
func (ajiiraNetHandler *AJIIRANetHandler) CreateConnection(w http.ResponseWriter, r *http.Request, ajRequestBodyVO *vo.AJRequestBodyVO) {

	var processor AJIIRANetProcessor
	successMsg, errorVO := processor.CreateConnections(ajRequestBodyVO)
	if errorVO != nil {
		utils.RespondJSON(w, http.StatusBadRequest, errorVO)
		return
	}

	utils.RespondJSON(w, successMsg.HTTPStatus, successMsg)

}

// FetchDevices :
func (ajiiraNetHandler *AJIIRANetHandler) FetchDevices(w http.ResponseWriter, r *http.Request, ajRequestBodyVO *vo.AJRequestBodyVO) {

	var processor AJIIRANetProcessor
	devices := processor.FetchDevices()

	utils.RespondJSON(w, http.StatusOK, devices)

}

// UpdateStrength :
func (ajiiraNetHandler *AJIIRANetHandler) UpdateStrength(w http.ResponseWriter, r *http.Request, ajRequestBodyVO *vo.AJRequestBodyVO) {

	var processor AJIIRANetProcessor
	successMsg, ajiraError := processor.UpdateStrength(ajRequestBodyVO)
	if ajiraError != nil {
		utils.RespondJSON(w, ajiraError.HTTPStatus, ajiraError)
		return
	}

	utils.RespondJSON(w, successMsg.HTTPStatus, successMsg)

}

// FetchRouteInfo :
func (ajiiraNetHandler *AJIIRANetHandler) FetchRouteInfo(w http.ResponseWriter, r *http.Request, ajRequestBodyVO *vo.AJRequestBodyVO) {

	var processor AJIIRANetProcessor
	successMsg, ajiraError := processor.FetchRouteInfo(ajRequestBodyVO)
	if ajiraError != nil {
		utils.RespondJSON(w, ajiraError.HTTPStatus, ajiraError)
		return
	}

	utils.RespondJSON(w, successMsg.HTTPStatus, successMsg)

}
