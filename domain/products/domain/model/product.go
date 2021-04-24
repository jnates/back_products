package model

import "time"

//Product type struct
type Product struct {
	ProductoID                string    `json:"producto_id,omitempty"`
	ProductoNombre            string    `json:"producto_nombre,omitempty"`
	ProductoCantidad          int       `json:"producto_cantidad,omitempty"`
	ProductoUsercreacion      int       `json:"producto_usercreacion,omitempty"`
	ProductoFechCreacion      time.Time `json:"-"`
	ProductoUserModificacion  int       `json:"producto_usermodificacion,omitempty"`
	ProductoFechaModificacion time.Time `json:"-"`
}
