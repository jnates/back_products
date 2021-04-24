package response

//ProductCreateResponse to create product
type ProductCreateResponse struct {
	Message string `json:"message,omitempty"`
}

//ProductALLResponse a
type ProductALLResponse struct {
	ProductoNombre       string `json:"producto_nombre,omitempty"`
	Cantidad             string `json:"cantidad,omitempty"`
	ProductoUsercreacion string `json:"producto_usercreacion,omitempty"`
}
