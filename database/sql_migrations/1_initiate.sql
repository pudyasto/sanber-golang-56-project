-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE public.canvasser
(
    id bigserial NOT NULL,
    code character varying(10) NOT NULL,
    name character varying(255) NOT NULL,
    phone character varying(20),
    username character varying(50) NOT NULL,
    password character varying(255) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT uq_canvasser_code UNIQUE (code)
);

CREATE TABLE public.item
(
    id bigserial NOT NULL,
    code character varying(10) NOT NULL,
    name character varying(255) NOT NULL,
    price numeric(18,2) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT uq_item_code UNIQUE (code)
);

CREATE TABLE public.customer
(
    id bigserial NOT NULL,
    code character varying(10) NOT NULL,
    name character varying(255) NOT NULL,
    address character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT uq_customer_code UNIQUE (code)
);

CREATE TABLE public.stock
(
    id bigserial NOT NULL,
    item_id bigint NOT NULL,
    canvasser_id bigint NOT NULL,
    qty bigint NOT NULL DEFAULT 0,
    PRIMARY KEY (id),
    CONSTRAINT uq_stock_item_canvasser UNIQUE (item_id, canvasser_id),
    CONSTRAINT fk_stock_item FOREIGN KEY (item_id)
        REFERENCES public.item (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT fk_stock_canvasser FOREIGN KEY (canvasser_id)
        REFERENCES public.canvasser (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

CREATE TABLE public.trn_sales
(
    id bigserial NOT NULL,
    customer_id bigint NOT NULL,
    canvasser_id bigint NOT NULL,
    code character varying(10) NOT NULL,
    date_sales timestamp NOT NULL,
    description character varying(255) NOT NULL,
    subtotal numeric(18,2) NOT NULL,
    discount numeric(18,2) NOT NULL,
    total numeric(18,2) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_stock_customer FOREIGN KEY (customer_id)
        REFERENCES public.customer (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT fk_stock_canvasser FOREIGN KEY (canvasser_id)
        REFERENCES public.canvasser (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

CREATE TABLE public.trn_sales_detail
(
    id bigserial NOT NULL,
    trn_sales_id bigint NOT NULL,
    item_id bigint NOT NULL,
    qty bigint NOT NULL DEFAULT 0,
    subtotal numeric(18,2) NOT NULL,
    total numeric(18,2) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_stock_trn_sales FOREIGN KEY (trn_sales_id)
        REFERENCES public.trn_sales (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT fk_stock_item FOREIGN KEY (item_id)
        REFERENCES public.item (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

-- +migrate StatementEnd