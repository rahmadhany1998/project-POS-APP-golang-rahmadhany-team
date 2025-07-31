--
-- PostgreSQL database dump
--

-- Dumped from database version 17.5
-- Dumped by pg_dump version 17.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    icon text,
    name character varying(100),
    description text
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.categories_id_seq OWNER TO postgres;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: login_tokens; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.login_tokens (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id bigint NOT NULL,
    token character varying(100) NOT NULL
);


ALTER TABLE public.login_tokens OWNER TO postgres;

--
-- Name: login_tokens_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.login_tokens_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.login_tokens_id_seq OWNER TO postgres;

--
-- Name: login_tokens_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.login_tokens_id_seq OWNED BY public.login_tokens.id;


--
-- Name: order_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_items (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    order_id bigint,
    product_id bigint,
    quantity bigint,
    price bigint
);


ALTER TABLE public.order_items OWNER TO postgres;

--
-- Name: order_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.order_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.order_items_id_seq OWNER TO postgres;

--
-- Name: order_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.order_items_id_seq OWNED BY public.order_items.id;


--
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    customer text,
    table_id bigint,
    subtotal numeric,
    tax numeric,
    total numeric,
    status text,
    method text
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- Name: orders_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.orders_id_seq OWNER TO postgres;

--
-- Name: orders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;


--
-- Name: password_reset_tokens; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.password_reset_tokens (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id bigint NOT NULL,
    otp_code character(4) NOT NULL,
    is_used boolean DEFAULT false,
    expires_at timestamp with time zone NOT NULL
);


ALTER TABLE public.password_reset_tokens OWNER TO postgres;

--
-- Name: password_reset_tokens_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.password_reset_tokens_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.password_reset_tokens_id_seq OWNER TO postgres;

--
-- Name: password_reset_tokens_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.password_reset_tokens_id_seq OWNED BY public.password_reset_tokens.id;


--
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    photo text,
    name character varying(100),
    item_code character varying(50),
    stock bigint DEFAULT 0,
    category_id bigint,
    price bigint,
    available boolean,
    quantity bigint,
    unit text,
    status text,
    retail_price bigint
);


ALTER TABLE public.products OWNER TO postgres;

--
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.products_id_seq OWNER TO postgres;

--
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- Name: tables; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tables (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(50),
    status character varying(20)
);


ALTER TABLE public.tables OWNER TO postgres;

--
-- Name: tables_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tables_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tables_id_seq OWNER TO postgres;

--
-- Name: tables_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tables_id_seq OWNED BY public.tables.id;


--
-- Name: user_accesses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_accesses (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id bigint,
    module character varying(100)
);


ALTER TABLE public.user_accesses OWNER TO postgres;

--
-- Name: user_accesses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_accesses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_accesses_id_seq OWNER TO postgres;

--
-- Name: user_accesses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_accesses_id_seq OWNED BY public.user_accesses.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(100) NOT NULL,
    role character varying(50) NOT NULL,
    photo text,
    phone character varying(20),
    address text,
    salary numeric,
    dob timestamp with time zone,
    shift_start text,
    shift_end text,
    detail text,
    is_active boolean DEFAULT true
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: login_tokens id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.login_tokens ALTER COLUMN id SET DEFAULT nextval('public.login_tokens_id_seq'::regclass);


--
-- Name: order_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items ALTER COLUMN id SET DEFAULT nextval('public.order_items_id_seq'::regclass);


--
-- Name: orders id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);


--
-- Name: password_reset_tokens id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.password_reset_tokens ALTER COLUMN id SET DEFAULT nextval('public.password_reset_tokens_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- Name: tables id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tables ALTER COLUMN id SET DEFAULT nextval('public.tables_id_seq'::regclass);


--
-- Name: user_accesses id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_accesses ALTER COLUMN id SET DEFAULT nextval('public.user_accesses_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (id, created_at, updated_at, deleted_at, icon, name, description) FROM stdin;
1	2025-07-01 00:00:00+08	2025-07-01 00:00:00+08	\N	https://example.com/icons/icon_1.png	Category 1	Description 1
2	2025-07-01 00:00:00+08	2025-07-01 00:00:00+08	\N	https://example.com/icons/icon_2.png	Category 2	Description 2
3	2025-07-01 00:00:00+08	2025-07-01 00:00:00+08	\N	https://example.com/icons/icon_3.png	Category 3	Description 3
4	2025-07-01 00:00:00+08	2025-07-01 00:00:00+08	\N	https://example.com/icons/icon_4.png	Category 4	Description 4
5	2025-07-01 00:00:00+08	2025-07-01 00:00:00+08	\N	https://example.com/icons/icon_5.png	Category 5	Description 5
\.


--
-- Data for Name: login_tokens; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.login_tokens (id, created_at, updated_at, deleted_at, user_id, token) FROM stdin;
1	2025-07-27 10:54:08.397142+08	2025-07-27 10:54:08.397142+08	\N	1	0a4816a2-afd4-4b4d-82bb-37be3bd103ba
2	2025-07-30 20:07:05.354855+08	2025-07-30 20:07:05.354855+08	\N	1	9569b3e4-3ce2-41b7-a193-f148c17f420b
3	2025-07-31 19:21:20.83313+08	2025-07-31 19:21:20.83313+08	\N	1	1a2ba2c4-6159-4502-b5e4-987512d229e2
\.


--
-- Data for Name: order_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.order_items (id, created_at, updated_at, deleted_at, order_id, product_id, quantity, price) FROM stdin;
1	\N	\N	\N	3	8	3	17043
2	\N	\N	\N	7	9	1	12430
3	\N	\N	\N	2	2	3	10605
4	\N	\N	\N	3	9	1	12430
5	\N	\N	\N	6	10	3	12961
6	\N	\N	\N	3	7	1	19790
7	\N	\N	\N	1	5	2	18156
8	\N	\N	\N	1	6	1	14593
9	\N	\N	\N	4	2	2	10605
10	\N	\N	\N	9	7	3	19790
11	\N	\N	\N	3	4	1	12724
12	\N	\N	\N	3	7	1	19790
13	\N	\N	\N	3	6	2	14593
14	\N	\N	\N	4	5	1	18156
15	\N	\N	\N	2	7	1	19790
16	\N	\N	\N	8	4	1	12724
17	\N	\N	\N	8	6	2	14593
18	\N	\N	\N	4	4	1	12724
19	\N	\N	\N	4	7	2	19790
20	\N	\N	\N	5	2	2	10605
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (id, created_at, updated_at, deleted_at, customer, table_id, subtotal, tax, total, status, method) FROM stdin;
1	2025-07-08 01:15:14.611855+08	2025-07-08 01:15:14.611855+08	\N	Customer 1	9	50905	2545.25	53450.25	confirmed	debit
2	2025-07-01 01:15:14.61187+08	2025-07-01 01:15:14.61187+08	\N	Customer 2	9	51605	2580.25	54185.25	cancelled	debit
3	2025-07-24 01:15:14.611878+08	2025-07-24 01:15:14.611878+08	\N	Customer 3	3	145049	7252.450000000001	152301.45	cancelled	cash
4	2025-06-29 01:15:14.611885+08	2025-06-29 01:15:14.611885+08	\N	Customer 4	2	91670	4583.5	96253.5	awaited	debit
5	2025-07-23 01:15:14.611892+08	2025-07-23 01:15:14.611892+08	\N	Customer 5	5	21210	1060.5	22270.5	cancelled	debit
6	2025-07-21 01:15:14.611899+08	2025-07-21 01:15:14.611899+08	\N	Customer 6	6	38883	1944.15	40827.15	confirmed	debit
7	2025-07-07 01:15:14.611906+08	2025-07-07 01:15:14.611906+08	\N	Customer 7	5	12430	621.5	13051.5	awaited	e-wallet
8	2025-07-19 01:15:14.611914+08	2025-07-19 01:15:14.611914+08	\N	Customer 8	1	41910	2095.5	44005.5	confirmed	debit
9	2025-07-14 01:15:14.611923+08	2025-07-14 01:15:14.611923+08	\N	Customer 9	5	59370	2968.5	62338.5	confirmed	cash
10	2025-07-17 01:15:14.61193+08	2025-07-17 01:15:14.61193+08	\N	Customer 10	3	0	0.0	0.0	awaited	e-wallet
\.


--
-- Data for Name: password_reset_tokens; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.password_reset_tokens (id, created_at, updated_at, deleted_at, user_id, otp_code, is_used, expires_at) FROM stdin;
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.products (id, created_at, updated_at, deleted_at, photo, name, item_code, stock, category_id, price, available, quantity, unit, status, retail_price) FROM stdin;
1	2025-07-27 01:15:14.611601+08	2025-07-27 01:15:14.61161+08	\N	https://example.com/photo1.jpg	Air Mineral	ITEM001	17	3	16990	t	1	pcs	active	13701
2	2025-07-27 01:15:14.611619+08	2025-07-27 01:15:14.611621+08	\N	https://example.com/photo2.jpg	Air Mineral	ITEM002	86	1	10605	t	2	pcs	active	8044
3	2025-07-27 01:15:14.611628+08	2025-07-27 01:15:14.61163+08	\N	https://example.com/photo3.jpg	Jus Mangga	ITEM003	82	1	16677	t	5	pcs	active	6964
4	2025-07-27 01:15:14.611637+08	2025-07-27 01:15:14.611639+08	\N	https://example.com/photo4.jpg	Kopi Hitam	ITEM004	20	2	12724	t	5	pcs	active	5651
5	2025-07-27 01:15:14.61165+08	2025-07-27 01:15:14.611654+08	\N	https://example.com/photo5.jpg	Kopi Hitam	ITEM005	43	1	18156	t	3	pcs	active	13565
6	2025-07-27 01:15:14.611665+08	2025-07-27 01:15:14.611667+08	\N	https://example.com/photo6.jpg	Mie Goreng	ITEM006	26	3	14593	t	3	pcs	active	9351
7	2025-07-27 01:15:14.611674+08	2025-07-27 01:15:14.611676+08	\N	https://example.com/photo7.jpg	Jus Mangga	ITEM007	19	1	19790	t	4	pcs	active	10180
8	2025-07-27 01:15:14.611692+08	2025-07-27 01:15:14.611696+08	\N	https://example.com/photo8.jpg	Kopi Hitam	ITEM008	19	3	17043	t	2	pcs	active	14224
9	2025-07-27 01:15:14.611704+08	2025-07-27 01:15:14.611707+08	\N	https://example.com/photo9.jpg	Kopi Hitam	ITEM009	54	1	12430	t	2	pcs	active	9345
10	2025-07-27 01:15:14.611714+08	2025-07-27 01:15:14.611716+08	\N	https://example.com/photo10.jpg	Es Teh	ITEM010	66	3	12961	t	3	pcs	active	9669
\.


--
-- Data for Name: tables; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tables (id, created_at, updated_at, deleted_at, name, status) FROM stdin;
1	2025-07-01 00:00:00+08	2025-07-01 00:00:00+08	\N	Table-1	available
2	2025-07-01 00:00:00+08	2025-07-01 00:00:00+08	\N	Table-2	available
3	2025-07-01 00:00:00+08	2025-07-01 00:00:00+08	\N	Table-3	available
4	2025-07-01 00:00:00+08	2025-07-01 00:00:00+08	\N	Table-4	available
5	2025-07-01 00:00:00+08	2025-07-01 00:00:00+08	\N	Table-5	available
6	2025-07-01 00:00:00+08	2025-07-01 00:00:00+08	\N	Table 6	available
9	2025-07-01 00:00:00+08	2025-07-01 00:00:00+08	\N	Table 9	available
\.


--
-- Data for Name: user_accesses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_accesses (id, created_at, updated_at, deleted_at, user_id, module) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, created_at, updated_at, deleted_at, name, email, password, role, photo, phone, address, salary, dob, shift_start, shift_end, detail, is_active) FROM stdin;
1	2025-07-26 13:49:11.388727+08	2025-07-26 13:49:11.388727+08	\N	Budi Santoso	budi@example.com	$2a$10$4esaDJvbfr8NR.XgZRyPmOesMULnEGgLkBXXN0f2VKTCrQ0Ikktoa	superadmin				0	0001-01-01 06:55:25+06:55:25				t
\.


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_id_seq', 1, false);


--
-- Name: login_tokens_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.login_tokens_id_seq', 3, true);


--
-- Name: order_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_items_id_seq', 1, false);


--
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.orders_id_seq', 1, false);


--
-- Name: password_reset_tokens_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.password_reset_tokens_id_seq', 1, false);


--
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.products_id_seq', 1, false);


--
-- Name: tables_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tables_id_seq', 1, false);


--
-- Name: user_accesses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_accesses_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 12, true);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: login_tokens login_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.login_tokens
    ADD CONSTRAINT login_tokens_pkey PRIMARY KEY (id);


--
-- Name: order_items order_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT order_items_pkey PRIMARY KEY (id);


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- Name: password_reset_tokens password_reset_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.password_reset_tokens
    ADD CONSTRAINT password_reset_tokens_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: tables tables_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tables
    ADD CONSTRAINT tables_pkey PRIMARY KEY (id);


--
-- Name: products uni_products_item_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT uni_products_item_code UNIQUE (item_code);


--
-- Name: tables uni_tables_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tables
    ADD CONSTRAINT uni_tables_name UNIQUE (name);


--
-- Name: user_accesses user_accesses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_accesses
    ADD CONSTRAINT user_accesses_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_login_tokens_token; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_login_tokens_token ON public.login_tokens USING btree (token);


--
-- Name: idx_users_email; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_users_email ON public.users USING btree (email);


--
-- Name: login_tokens fk_login_tokens_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.login_tokens
    ADD CONSTRAINT fk_login_tokens_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: order_items fk_order_items_product; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT fk_order_items_product FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- Name: order_items fk_orders_order_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT fk_orders_order_items FOREIGN KEY (order_id) REFERENCES public.orders(id);


--
-- Name: orders fk_orders_table; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT fk_orders_table FOREIGN KEY (table_id) REFERENCES public.tables(id);


--
-- Name: password_reset_tokens fk_password_reset_tokens_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.password_reset_tokens
    ADD CONSTRAINT fk_password_reset_tokens_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: products fk_products_category; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_products_category FOREIGN KEY (category_id) REFERENCES public.categories(id);


--
-- Name: user_accesses fk_user_accesses_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_accesses
    ADD CONSTRAINT fk_user_accesses_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

