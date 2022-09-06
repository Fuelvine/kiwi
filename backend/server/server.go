package server

import (
	"fmt"
	"github.com/Fuelvine/kiwi/backend/constants"
	"github.com/iMeisa/errortrace"
	"net/http"
)

func StartServer() errortrace.ErrorTrace {

	trace := errortrace.NewTrace(nil)

	srv := http.Server{
		Addr:    constants.ServerPort,
		Handler: routes(),
	}

	go func() {
		fmt.Println("Starting server on port", constants.ServerPort)
		if err := srv.ListenAndServe(); err != nil {
			trace = errortrace.NewTrace(err)
		}
	}()

	return trace
}
