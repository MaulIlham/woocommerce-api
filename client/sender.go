package client

import (
	"net/http"
	"woocommerce-api/request"
)

// Sender interface
type Sender interface {
	Send(req request.Request) (resp *http.Response, err error)
}
