package persistence

import (
	"backend_crudgo/domain/products/aplication/v1/response"
	"backend_crudgo/domain/products/domain/model"
	repoDomain "backend_crudgo/domain/products/domain/repository"
	"backend_crudgo/infrastructure/database"
	"context"
	"database/sql"
)

type sqlProductRepo struct {
	Conn *database.Data
}

//NewProductRepository aaa
func NewProductRepository(Conn *database.Data) repoDomain.ProductRepository {
	return &sqlProductRepo{
		Conn: Conn,
	}
}

func (sr *sqlProductRepo) CreateProductHandler(ctx context.Context, product *model.Product) (*response.ProductCreateResponse, error) {
	stmt, err := sr.Conn.DB.PrepareContext(ctx, insertProduct)
	if err != nil {
		return &response.ProductCreateResponse{}, err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, &product.ProductoID, &product.ProductoNombre, &product.ProductoCantidad, &product.ProductoUsercreacion, &product.ProductoUserModificacion)
	var idResult string
	err = row.Scan(&idResult)
	if err != sql.ErrNoRows {
		return &response.ProductCreateResponse{}, err
	}
	ProductResponse := response.ProductCreateResponse{
		Message: "product created",
	}
	return &ProductResponse, nil
}

// func (sr *sqlProductRepo) GetAllProducts(ctx context.Context, product *model.Product) (*response.ProductALLResponse, error) {
// 	stmt, err := sr.Conn.DB.PrepareContext(ctx, insertProduct)
// 	if err != nil {
// 		return &response.ProductALLResponse{}, err
// 	}
// 	defer stmt.Close()
// 	row := stmt.QueryRowContext(ctx, &product.ProductoID, &product.ProductoNombre, &product.ProductoCantidad, &product.ProductoUsercreacion, &product.ProductoUserModificacion)
// 	var idResult string
// 	err = row.Scan(&idResult)
// 	if err != sql.ErrNoRows {
// 		return &response.ProductALLResponse{}, err
// 	}
// 	ProductAllResponse :=  [] response.ProductCreateResponse{}
// 	return &[]ProductAllResponse, nil
// }
