// Package v1 contains the full set of handler functions and routes
// supported by the v1 web api.
package loader

import (
	"github.com/emreodabas/pod-loader/pkg/web"
	"go.uber.org/zap"
)

// Config contains all the mandatory systems required by controller.
type Config struct {
	Log *zap.SugaredLogger
}

// Routes binds all the version 1 routes.
func Routes(app *web.App, cfg Config) {
	//TODO
	const version = "v1"

	/*	pkg.Handle(http.MethodGet, version, "/users/token", ugh.Token)
		pkg.Handle(http.MethodGet, version, "/users/:page/:rows", ugh.Query, authen, admin)
		pkg.Handle(http.MethodGet, version, "/users/:id", ugh.QueryByID, authen)
		pkg.Handle(http.MethodPost, version, "/users", ugh.Create, authen, admin)
		pkg.Handle(http.MethodPut, version, "/users/:id", ugh.Update, authen, admin)
		pkg.Handle(http.MethodDelete, version, "/users/:id", ugh.Delete, authen, admin)

		// Register product and sale endpoints.
		pgh := productgrp.Handlers{
			Product: product.NewCore(cfg.Log, cfg.DB),
		}
		pkg.Handle(http.MethodGet, version, "/products/:page/:rows", pgh.Query, authen)
		pkg.Handle(http.MethodGet, version, "/products/:id", pgh.QueryByID, authen)
		pkg.Handle(http.MethodPost, version, "/products", pgh.Create, authen)
		pkg.Handle(http.MethodPut, version, "/products/:id", pgh.Update, authen)
		pkg.Handle(http.MethodDelete, version, "/products/:id", pgh.Delete, authen)
	*/
}
