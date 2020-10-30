package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/dmnlk/stringUtils"

	"ajiiranetservice/constants"
	"ajiiranetservice/utils"
	"ajiiranetservice/vo"
)

// AJIIRANetProcessor :
type AJIIRANetProcessor int

// CreateDevice :
func (ajiiraNetProcessor *AJIIRANetProcessor) CreateDevice(ajRequestBodyVO *vo.AJRequestBodyVO) (*vo.AJIRASuccessMsg, *vo.AJIRAErrorMsg) {
	ajiiraDevice := &vo.AJIIRADevice{}
	ajiraSuccessMsg := &vo.AJIRASuccessMsg{HTTPStatus: http.StatusOK}
	ajiraErrorMsg := &vo.AJIRAErrorMsg{HTTPStatus: http.StatusBadRequest}

	err := json.Unmarshal(ajRequestBodyVO.RequestBody, ajiiraDevice)

	if err != nil {
		ajiraErrorMsg.Error = fmt.Sprintf(getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorInvalidMessage))
		return nil, ajiraErrorMsg
	}

	if stringUtils.IsEmpty(ajiiraDevice.Name) || stringUtils.IsEmpty(ajiiraDevice.Type) {
		ajiraErrorMsg.Error = fmt.Sprintf(getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorInvalidMessage))
		return nil, ajiraErrorMsg
	}

	if ajiiraDevice.Type != constants.COMPUTER && ajiiraDevice.Type != constants.REPEATER {
		ajiraErrorMsg.Error = fmt.Sprintf(getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorNotValidDeviceType), ajiiraDevice.Type)
		return nil, ajiraErrorMsg
	}

	if _, ok := ajiiraNetGraph.ConnectedDevices[ajiiraDevice.Name]; ok {
		ajiraErrorMsg.Error = fmt.Sprintf(getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorDeviceAlreadyExist), ajiiraDevice.Name)
		return nil, ajiraErrorMsg
	}

	if ajiiraDevice.Strength == nil {
		ajiiraDevice.Strength = &vo.AJIIRADeviceStrength{Value: 5}
	}

	ajiiraNetGraph.ConnectedDevices[ajiiraDevice.Name] = ajiiraDevice

	ajiraSuccessMsg.Message = fmt.Sprintf(getMessage(ajRequestBodyVO.AcceptedLanguage, constants.MsgDeviceAddSuccess), ajiiraDevice.Name)
	return ajiraSuccessMsg, nil

}

// CreateConnections :
func (ajiiraNetProcessor *AJIIRANetProcessor) CreateConnections(ajRequestBodyVO *vo.AJRequestBodyVO) (*vo.AJIRASuccessMsg, *vo.AJIRAErrorMsg) {
	ajiiraDeviceConnection := &vo.AJIIRADeviceConnection{}
	ajiraSuccessMsg := &vo.AJIRASuccessMsg{HTTPStatus: http.StatusOK}
	ajiraErrorMsg := &vo.AJIRAErrorMsg{HTTPStatus: http.StatusBadRequest}

	err := json.Unmarshal(ajRequestBodyVO.RequestBody, ajiiraDeviceConnection)

	if err != nil {
		ajiraErrorMsg.Error = fmt.Sprintf(getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorInvalidCommandSyntax))
		return nil, ajiraErrorMsg
	}

	if len(ajiiraDeviceConnection.Target) == 0 || stringUtils.IsEmpty(ajiiraDeviceConnection.Source) {
		ajiraErrorMsg.Error = getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorInvalidCommandSyntax)
		return nil, ajiraErrorMsg
	}

	if _, ok := ajiiraNetGraph.ConnectedDevices[ajiiraDeviceConnection.Source]; !ok {
		ajiraErrorMsg.Error = fmt.Sprintf(getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorNodeNotFound), ("'" + ajiiraDeviceConnection.Source + "'"))
		return nil, ajiraErrorMsg
	}

	if len(ajiiraDeviceConnection.Target) > 0 {
		if utils.IsSliceContains(ajiiraDeviceConnection.Target, ajiiraDeviceConnection.Source) {
			ajiraErrorMsg.Error = getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorDeviceConnectItself)
			return nil, ajiraErrorMsg
		}

		connections, _ := ajiiraNetGraph.Connections[ajiiraDeviceConnection.Source]

		for _, targetNode := range ajiiraDeviceConnection.Target {
			if _, ok := ajiiraNetGraph.ConnectedDevices[targetNode]; !ok {
				ajiraErrorMsg.Error = fmt.Sprintf(getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorNodeNotFound), ("'" + targetNode + "'"))
				return nil, ajiraErrorMsg
			}

			if len(connections) > 0 && utils.IsSliceContains(connections, targetNode) {
				ajiraErrorMsg.Error = getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorDeviceConnectAlready)
				return nil, ajiraErrorMsg
			}

		}
	}

	if len(ajiiraDeviceConnection.Target) > 0 {
		for _, target := range ajiiraDeviceConnection.Target {
			ajiiraNetGraph.AddEdge(ajiiraDeviceConnection.Source, target)
		}
	}

	ajiraSuccessMsg.Message = getMessage(ajRequestBodyVO.AcceptedLanguage, constants.MsgDeviceConnectionSuccess)

	return ajiraSuccessMsg, nil

}

// FetchDevices :
func (ajiiraNetProcessor *AJIIRANetProcessor) FetchDevices() *vo.AJIRADevices {
	devices := make([]*vo.AJIIRADeviceVO, 0)
	for _, device := range ajiiraNetGraph.ConnectedDevices {
		devices = append(devices, &vo.AJIIRADeviceVO{Name: device.Name, Type: device.Type})
	}
	ajiraDevices := &vo.AJIRADevices{}
	ajiraDevices.Devices = devices
	return ajiraDevices

}

// UpdateStrength :
func (ajiiraNetProcessor *AJIIRANetProcessor) UpdateStrength(ajRequestBodyVO *vo.AJRequestBodyVO) (*vo.AJIRASuccessMsg, *vo.AJIRAErrorMsg) {

	ajiiraDeviceStrength := &vo.AJIIRADeviceStrength{}

	ajiraSuccessMsg := &vo.AJIRASuccessMsg{HTTPStatus: http.StatusOK}
	ajiraErrorMsg := &vo.AJIRAErrorMsg{HTTPStatus: http.StatusBadRequest}

	nodeName := strings.Replace(ajRequestBodyVO.Path, "/devices/", "", -1)
	nodeName = strings.Replace(nodeName, "/strength", "", -1)

	if _, ok := ajiiraNetGraph.ConnectedDevices[nodeName]; !ok {
		errorMsg := fmt.Sprintf(getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorDeviceNotFound))
		ajiraErrorMsg.Error = errorMsg
		ajiraErrorMsg.HTTPStatus = http.StatusNotFound
		return nil, ajiraErrorMsg
	}

	err := json.Unmarshal(ajRequestBodyVO.RequestBody, ajiiraDeviceStrength)

	if err != nil {
		ajiraErrorMsg.Error = getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorValueMustBeInteger)
		return nil, ajiraErrorMsg
	}

	if stringUtils.IsEmpty(nodeName) {
		ajiraErrorMsg.Error = getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorNotValidNode)
		return nil, ajiraErrorMsg
	}

	val, _ := ajiiraNetGraph.ConnectedDevices[nodeName]

	val.Strength = ajiiraDeviceStrength
	ajiiraNetGraph.ConnectedDevices[nodeName] = val

	ajiraSuccessMsg.Message = fmt.Sprintf(getMessage(ajRequestBodyVO.AcceptedLanguage, constants.MsgStrengthUpdateSuccess))

	return ajiraSuccessMsg, nil

}

// FetchRouteInfo :
func (ajiiraNetProcessor *AJIIRANetProcessor) FetchRouteInfo(ajRequestBodyVO *vo.AJRequestBodyVO) (*vo.AJIRASuccessMsg, *vo.AJIRAErrorMsg) {

	urlpath := strings.TrimSpace(ajRequestBodyVO.Path)
	ajiraSuccessMsg := &vo.AJIRASuccessMsg{HTTPStatus: http.StatusOK}
	ajiraErrorMsg := &vo.AJIRAErrorMsg{HTTPStatus: http.StatusBadRequest}

	u, err := url.Parse(urlpath)
	if err != nil {
		ajiraErrorMsg.Error = getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorInvalidRequest)
		return nil, ajiraErrorMsg
	}

	queryMap, _ := url.ParseQuery(u.RawQuery)

	start, end := "", ""
	if startNode, ok := queryMap["from"]; ok {
		start = startNode[0]
	}

	if endNode, ok := queryMap["to"]; ok {
		end = endNode[0]
	}

	if stringUtils.IsEmpty(start) || stringUtils.IsEmpty(end) {
		ajiraErrorMsg.Error = getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorInvalidRequest)
		return nil, ajiraErrorMsg
	}

	if device, ok := ajiiraNetGraph.ConnectedDevices[start]; ok {
		if device.Type == constants.REPEATER {
			ajiraErrorMsg.Error = getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorRouteCannotCalculateRouter)
			return nil, ajiraErrorMsg
		}
	} else {
		ajiraErrorMsg.Error = fmt.Sprintf(getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorNodeNotFound), ("'" + start + "'"))
		return nil, ajiraErrorMsg
	}

	if device, ok := ajiiraNetGraph.ConnectedDevices[end]; ok {
		if device.Type == constants.REPEATER {
			ajiraErrorMsg.Error = getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorRouteCannotCalculateRouter)
			return nil, ajiraErrorMsg
		}
	} else {
		ajiraErrorMsg.Error = fmt.Sprintf(getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorNodeNotFound), ("'" + end + "'"))
		return nil, ajiraErrorMsg
	}

	// Note : For future to find best path we're using this
	paths := make([][]string, 0)

	if start == end {
		path := fmt.Sprintf(getMessage(ajRequestBodyVO.AcceptedLanguage, constants.MsgRoutes), (start + "->" + start))
		ajiraSuccessMsg.Message = path
		return ajiraSuccessMsg, nil
	}

	paths = ajiiraNetGraph.DepthFirst([]string{start}, end, paths)

	if len(paths) == 0 {
		ajiraErrorMsg.Error = getMessage(ajRequestBodyVO.AcceptedLanguage, constants.ErrorRouteNotFound)
		ajiraErrorMsg.HTTPStatus = http.StatusNotFound
		return nil, ajiraErrorMsg
	}

	path := fmt.Sprintf(getMessage(ajRequestBodyVO.AcceptedLanguage, constants.MsgRoutes), (strings.Join(paths[0], "->")))
	ajiraSuccessMsg.Message = path
	return ajiraSuccessMsg, nil

}

func getMessage(i18n, key string) string {
	if val, ok := message[i18n]; ok {
		if message, ok := val[key]; ok {
			return message
		}
	}
	return ""

}
