package gosrv

import "cyberpull.com/gotk/v2"

type Request interface {
	Data
}

type pRequest struct {
	pData

	UUID    string `json:"uuid" binding:"required"`
	Method  string `json:"method" binding:"required"`
	Channel string `json:"channel" binding:"required"`
}

// ======================

func newRequest(method, channel string) (req *pRequest, err error) {
	reqUUID, err := gotk.UUID()

	if err != nil {
		return
	}

	req = &pRequest{
		UUID:    reqUUID,
		Method:  method,
		Channel: channel,
	}

	return
}
