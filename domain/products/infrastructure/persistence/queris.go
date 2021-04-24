package persistence

const (
	insertProduct  = "INSERT INTO public.productos (producto_id,producto_nombre,producto_cantidad,producto_usercreacion,producto_fechcreacion,producto_usermodificacion,producto_fechamodificacion) VALUES ($1,$2,$3,$4,'NOW()',$5,'NOW()')"
	GetAllProducts = "SELECT * FROM public.productos"
)
