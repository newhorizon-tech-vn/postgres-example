-- public.category definition

-- Drop table

-- DROP TABLE public.category;

CREATE TABLE public.category (
	id bigserial NOT NULL,
	"name" text NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT category_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_category_deleted_at ON public.category USING btree (deleted_at);


-- public.factory definition

-- Drop table

-- DROP TABLE public.factory;

CREATE TABLE public.factory (
	id bigserial NOT NULL,
	"name" text NULL,
	address text NULL,
	CONSTRAINT factory_pkey PRIMARY KEY (id)
);


-- public.item definition

-- Drop table

-- DROP TABLE public.item;

CREATE TABLE public.item (
	id bigserial NOT NULL,
	product_id int8 NULL,
	"name" text NULL,
	descripton text NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT item_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_item_deleted_at ON public.item USING btree (deleted_at);


-- public.product definition

-- Drop table

-- DROP TABLE public.product;

CREATE TABLE public.product (
	id bigserial NOT NULL,
	category_id int8 NULL,
	"name" text NULL,
	price int8 NULL,
	CONSTRAINT product_pkey PRIMARY KEY (id)
);


-- public.product_factories definition

-- Drop table

-- DROP TABLE public.product_factories;

CREATE TABLE public.product_factories (
	id bigserial NOT NULL,
	product_id int8 NULL,
	factory_id int8 NULL,
	price int8 NULL,
	CONSTRAINT product_factories_pkey PRIMARY KEY (id)
);


-- public.workshop definition

-- Drop table

-- DROP TABLE public.workshop;

CREATE TABLE public.workshop (
	id bigserial NOT NULL,
	factory_id int8 NULL,
	"name" text NULL,
	CONSTRAINT workshop_pkey PRIMARY KEY (id)
);