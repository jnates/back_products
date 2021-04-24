package v1

import (
	"backend_crudgo/domain/products/domain/model"
	repoDomain "backend_crudgo/domain/products/domain/repository"
	"backend_crudgo/domain/products/infrastructure/persistence"
	"backend_crudgo/infrastructure/database"
	"backend_crudgo/infrastructure/middleware"
	"encoding/json"
	"fmt"
	"net/http"
)

//ProductRouter rutas
type ProductRouter struct {
	Repo repoDomain.ProductRepository
}

//NewProductHandler aa
func NewProductHandler(db *database.Data) *ProductRouter {
	return &ProductRouter{
		Repo: persistence.NewProductRepository(db),
	}
}

// CreateProductHandler aa
func (prod *ProductRouter) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	ctx := r.Context()
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		_ = middleware.HTTPError(w, r, http.StatusBadRequest, "bad request", err.Error())
		return
	}
	result, err := prod.Repo.CreateProductHandler(ctx, &product)
	if err != nil {
		_ = middleware.HTTPError(w, r, http.StatusConflict, "conflict", err.Error())
		return
	}
	w.Header().Add("Location", fmt.Sprintf("%s%s", r.URL.String(), result))
	_ = middleware.JSON(w, r, http.StatusCreated, result)
}
