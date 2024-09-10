package telexgin

import (
	"time"

	"github.com/Micah-Shallom/go-apm-sdk/telex"
	"github.com/gin-gonic/gin"
)

type Options struct {
	Repanic         bool
	WaitForDelivery bool
	Timeout         time.Duration
}

type Handler struct {
	apmClient *telex.Client
	//more clients can be added 

	Options
}

func NewGin(apmClient *telex.Client, options Options) gin.HandlerFunc {
	if options.Timeout == 0 {
		options.Timeout = 5 * time.Second
	}

	handler := &Handler{
		apmClient: apmClient,
		Options:   options,
	}

	return handler.handlegin
}
