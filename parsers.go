package gosrv

import (
	"reflect"
	"strings"

	"cyberpull.com/gotk/v2/errors"
)

const (
	separator string = "::"

	errorPrefix    string = "ERROR"
	infoPrefix     string = "INFO"
	requestPrefix  string = "REQUEST"
	responsePrefix string = "RESPONSE"
	updatePrefix   string = "UPDATE"
	statePrefix    string = "STATE"
)

func parse[T any](v string, prefix string) (value T, err error) {
	chunks := strings.SplitN(v, separator, 2)

	if chunks[0] != prefix {
		err = errors.New("Invalid "+prefix, 400)
		return
	}

	return parseValue[T](chunks[1], prefix)
}

func parseValue[T any](v string, prefix string) (value T, err error) {
	vType := reflect.TypeOf(value)

	if vType.Kind() != reflect.Pointer {
		err = errors.New("Value type must be a pointer")
		return
	}

	tmpInfo := reflect.New(vType.Elem()).Interface().(T)

	if err = json.Decode([]byte(v), tmpInfo); err != nil {
		return
	}

	value = tmpInfo

	return
}

func parseError(v string) (data *pData, err error) {
	return parse[*pData](v, errorPrefix)
}

func parseErrorValue(v string) (data *pData, err error) {
	return parseValue[*pData](v, errorPrefix)
}

func parseInfo(v string) (info *Info, err error) {
	return parse[*Info](v, infoPrefix)
}

func parseInfoValue(v string) (info *Info, err error) {
	return parseValue[*Info](v, infoPrefix)
}

func parseRequest(v string) (resp *pRequest, err error) {
	return parse[*pRequest](v, requestPrefix)
}

func parseRequestValue(v string) (resp *pRequest, err error) {
	return parseValue[*pRequest](v, requestPrefix)
}

func parseResponse(v string) (resp *pResponse, err error) {
	return parse[*pResponse](v, responsePrefix)
}

func parseResponseValue(v string) (resp *pResponse, err error) {
	return parseValue[*pResponse](v, responsePrefix)
}

func parseUpdate(v string) (data *pUpdate, err error) {
	return parse[*pUpdate](v, updatePrefix)
}

func parseUpdateValue(v string) (data *pUpdate, err error) {
	return parseValue[*pUpdate](v, updatePrefix)
}
