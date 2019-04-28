// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"mustrum/server/restapi/operations"
	"mustrum/server/restapi/operations/default_controller"
)

//go:generate swagger generate server --target ../../server --name Mustrum --spec ../swagger.yaml

func configureFlags(api *operations.MustrumAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MustrumAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.DefaultControllerControllersDefaultControllerAllPixelsGetHandler == nil {
		api.DefaultControllerControllersDefaultControllerAllPixelsGetHandler = default_controller.ControllersDefaultControllerAllPixelsGetHandlerFunc(func(params default_controller.ControllersDefaultControllerAllPixelsGetParams) middleware.Responder {
			return middleware.NotImplemented("operation default_controller.ControllersDefaultControllerAllPixelsGet has not yet been implemented")
		})
	}
	if api.DefaultControllerControllersDefaultControllerDotXPutHandler == nil {
		api.DefaultControllerControllersDefaultControllerDotXPutHandler = default_controller.ControllersDefaultControllerDotXPutHandlerFunc(func(params default_controller.ControllersDefaultControllerDotXPutParams) middleware.Responder {
			return middleware.NotImplemented("operation default_controller.ControllersDefaultControllerDotXPut has not yet been implemented")
		})
	}
	if api.DefaultControllerControllersDefaultControllerMonoframePixelmapPutHandler == nil {
		api.DefaultControllerControllersDefaultControllerMonoframePixelmapPutHandler = default_controller.ControllersDefaultControllerMonoframePixelmapPutHandlerFunc(func(params default_controller.ControllersDefaultControllerMonoframePixelmapPutParams) middleware.Responder {
			return middleware.NotImplemented("operation default_controller.ControllersDefaultControllerMonoframePixelmapPut has not yet been implemented")
		})
	}
	if api.DefaultControllerControllersDefaultControllerMonolinePutHandler == nil {
		api.DefaultControllerControllersDefaultControllerMonolinePutHandler = default_controller.ControllersDefaultControllerMonolinePutHandlerFunc(func(params default_controller.ControllersDefaultControllerMonolinePutParams) middleware.Responder {
			return middleware.NotImplemented("operation default_controller.ControllersDefaultControllerMonolinePut has not yet been implemented")
		})
	}
	if api.DefaultControllerControllersDefaultControllerShowGetHandler == nil {
		api.DefaultControllerControllersDefaultControllerShowGetHandler = default_controller.ControllersDefaultControllerShowGetHandlerFunc(func(params default_controller.ControllersDefaultControllerShowGetParams) middleware.Responder {
			return middleware.NotImplemented("operation default_controller.ControllersDefaultControllerShowGet has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
