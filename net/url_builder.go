package net

import "woocommerce-api/request"

// URLBuilder interface
type URLBuilder interface {
	GetURL(req request.Request) string
}
