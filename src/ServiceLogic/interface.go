package ServiceLogic

import "net/http"

type ServiceHandler interface {
	HandleGet(http.ResponseWriter, *http.Request)
	HandlePost(http.ResponseWriter, *http.Request)
}
