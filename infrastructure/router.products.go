package infrastructure

import (
	v1Product "backend_crudgo/domain/products/aplication/v1"
	"backend_crudgo/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi"
)

//RoutesProducts aa
func RoutesProducts(conn *database.Data) http.Handler {
	router := chi.NewRouter()
	products := v1Product.NewProductHandler(conn) //domain
	router.Mount("/products", routesProduct(products))
	return router
}

//Router user
func routesProduct(handler *v1Product.ProductRouter) http.Handler {
	router := chi.NewRouter()
	router.Post("/", handler.CreateProductHandler)
	// router.Post("/productAll", handler.)
	return router
}
