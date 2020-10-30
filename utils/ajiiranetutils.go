package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"ajiiranetservice/constants"
	"ajiiranetservice/vo"

	"github.com/dmnlk/stringUtils"
)

// TextToRequestJSON :
func TextToRequestJSON(reqBody string) *vo.AJRequestBodyVO {

	ajRequestBodyVO := &vo.AJRequestBodyVO{}

	// Default Accepted Language
	ajRequestBodyVO.AcceptedLanguage = constants.EnIN

	if stringUtils.IsNotEmpty(reqBody) {
		for _, req := range strings.Split(reqBody, "\n") {
			if strings.HasPrefix(req, "CREATE") || strings.HasPrefix(req, "FETCH") || strings.HasPrefix(req, "MODIFY") {
				typeCounter := 0
				for _, val := range strings.Split(req, " ") {
					if stringUtils.IsNotEmpty(val) {
						if typeCounter == 1 {
							ajRequestBodyVO.Path = val
							break
						}
						ajRequestBodyVO.RequsetType = val
						typeCounter++
					}
				}

			} else if strings.HasPrefix(req, "content-type") {
				ajRequestBodyVO.ContentType = req
			} else if strings.HasPrefix(req, "{") && strings.HasSuffix(req, "}") {
				ajRequestBodyVO.RequestBody = []byte(req)
			}

		}
	}
	return ajRequestBodyVO
}

// RespondError makes the error response with payload as json format
func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, vo.ErrorJSON{ErrorMessage: message})
}

// RespondJSON : Return a response as JSON
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	if status == http.StatusInternalServerError {
		var errorVO *vo.ErrorVO
		tmpErrorVO, ok := payload.(vo.ErrorVO)
		if ok {
			errorVO = &tmpErrorVO
		} else {
			errorVO, ok = payload.(*vo.ErrorVO)
		}
		if ok && errorVO != nil {
			correlationIDString := w.Header().Get(constants.AjiiraNetCorrelationID)
			errorVO.DebugID = correlationIDString
			payload = errorVO
		}
	}
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// FindEmptyError :
func FindEmptyError(val, field string, errorDetails []*vo.ErrorDetailVO) []*vo.ErrorDetailVO {
	if errorDetails == nil {
		errorDetails = make([]*vo.ErrorDetailVO, 0)
	}
	if stringUtils.IsEmpty(val) {
		errorDetails = append(errorDetails, &vo.ErrorDetailVO{Field: field, Issue: "Empty value"})
	}
	return errorDetails
}

// IsSliceContains :
func IsSliceContains(source []string, target string) bool {
	for _, val := range source {
		if val == target {
			return true
		}
	}
	return false
}

func printPath(visited []string) {

	for _, node := range visited {
		fmt.Print(node)
		fmt.Print(" ")
	}
	fmt.Println()

}
