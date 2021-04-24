-- CREA TABLA TIPOS_DOCUMENTOS
-- Table: public.tipos_documentos

-- DROP TABLE public.tipos_documentos;

CREATE TABLE public.tipos_documentos
(
    tido_id integer NOT NULL,
    tido_descripcion character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT tipo_documento_pkey PRIMARY KEY (tido_id)
)

TABLESPACE pg_default;

ALTER TABLE public.tipos_documentos
    OWNER to postgres;

COMMENT ON COLUMN public.tipos_documentos.tido_id
    IS 'Tipo de documento';

COMMENT ON COLUMN public.tipos_documentos.tido_descripcion
    IS 'Descripción del tipo de documento';

-- CREA TABLA USUARIOS 
-- Table: public.usuarios

-- DROP TABLE public.usuarios;

CREATE TABLE public.usuarios
(
    user_id integer NOT NULL,
    user_nombres character varying(255) COLLATE pg_catalog."default" NOT NULL,
    "user.apellidos" character varying(255) COLLATE pg_catalog."default" NOT NULL,
    user_nrodocumento integer NOT NULL,
    user_email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    user_password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    user_tipodocumento integer NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (user_id),
    CONSTRAINT fk_tipos_documentos FOREIGN KEY (user_tipodocumento)
        REFERENCES public.tipos_documentos (tido_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE public.usuarios
    OWNER to postgres;

COMMENT ON COLUMN public.usuarios.user_id
    IS 'LLave primaria de Usuarios';

COMMENT ON COLUMN public.usuarios.user_nombres
    IS 'Nombre completo del Usuario';

COMMENT ON COLUMN public.usuarios."user.apellidos"
    IS 'Apellidos de Usuario';

COMMENT ON COLUMN public.usuarios.user_nrodocumento
    IS 'Numero de documento de usuario';

COMMENT ON COLUMN public.usuarios.user_email
    IS 'Correo electronico del Usuario para el login';

COMMENT ON COLUMN public.usuarios.user_password
    IS 'Password de usuario para login';

COMMENT ON COLUMN public.usuarios.user_tipodocumento
    IS 'Tipo de documento de identificacion del usuario';

-- CREA TABLA PRODUCTOS
-- -- Table: public.productos

-- DROP TABLE public.productos;

CREATE TABLE public.productos
(
    producto_id integer NOT NULL,
    producto_nombre character varying(255) COLLATE pg_catalog."default" NOT NULL,
    producto_cantidad integer,
    producto_usercreacion integer NOT NULL,
    producto_fechcreacion timestamp(0) without time zone NOT NULL,
    producto_usermodificacion integer NOT NULL,
    "producto_fechamodificacion" timestamp(0) without time zone NOT NULL,
    CONSTRAINT producto_pkey PRIMARY KEY (producto_id),
    CONSTRAINT fk_user_creacion FOREIGN KEY (producto_usercreacion)
        REFERENCES public.usuarios (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT fk_user_modificacion FOREIGN KEY (producto_usermodificacion)
        REFERENCES public.usuarios (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE public.productos
    OWNER to postgres;

COMMENT ON COLUMN public.productos.producto_id
    IS 'id del producto';

COMMENT ON COLUMN public.productos.producto_nombre
    IS 'Name del producto';

COMMENT ON COLUMN public.productos.producto_cantidad
    IS 'Cantidad del producto';

COMMENT ON COLUMN public.productos.producto_usercreacion
    IS 'El usuario crea un producto';

COMMENT ON COLUMN public.productos.producto_fechcreacion
    IS 'Fecha de creación del producto';

COMMENT ON COLUMN public.productos.producto_usermodificacion
    IS 'Fecha de  modificación del producto';

COMMENT ON COLUMN public.productos."producto_fechamodificacion"
    IS 'Fecha de modificación del producto';