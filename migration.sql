CREATE TABLE public.users (
	id serial4 NOT NULL,
	name varchar NULL,
	phone varchar NULL,
	CONSTRAINT users_pk PRIMARY KEY (id)
);

CREATE TABLE "order" (
	id serial4 NOT NULL,
	product varchar NOT NULL,
	qty int8 NOT NULL,
	price int4 NOT NULL,
	CONSTRAINT order_pk PRIMARY KEY (id)
);