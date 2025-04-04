--
-- PostgreSQL database dump
--

-- Dumped from database version 17.2 (Debian 17.2-1.pgdg120+1)
-- Dumped by pg_dump version 17.2 (Debian 17.2-1.pgdg120+1)

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
-- Name: addresses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.addresses (
    id integer NOT NULL,
    address text NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.addresses OWNER TO postgres;

--
-- Name: addresses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.addresses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.addresses_id_seq OWNER TO postgres;

--
-- Name: addresses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.addresses_id_seq OWNED BY public.addresses.id;


--
-- Name: addresses id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.addresses ALTER COLUMN id SET DEFAULT nextval('public.addresses_id_seq'::regclass);


--
-- Data for Name: addresses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.addresses (id, address, created_at, updated_at) FROM stdin;
2	789 Oak St, Springfield, IL 62703	2025-02-04 17:08:09.949158	2025-02-04 17:08:09.949158
3	101 Pine St, Springfield, IL 62704	2025-02-04 17:08:09.949158	2025-02-04 17:08:09.949158
4	202 Maple St, Springfield, IL 62705	2025-02-04 17:08:09.949158	2025-02-04 17:08:09.949158
5	303 Cedar St, Springfield, IL 62706	2025-02-04 17:08:09.949158	2025-02-04 17:08:09.949158
6	404 Birch St, Springfield, IL 62707	2025-02-04 17:08:09.949158	2025-02-04 17:08:09.949158
7	505 Walnut St, Springfield, IL 62708	2025-02-04 17:08:09.949158	2025-02-04 17:08:09.949158
8	606 Cherry St, Springfield, IL 62709	2025-02-04 17:08:09.949158	2025-02-04 17:08:09.949158
1	4561 Elm St, Springfield, IL 62702	2025-02-04 17:08:09.949158	2025-02-04 17:08:09.949158
\.


--
-- Name: addresses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.addresses_id_seq', 8, true);


--
-- Name: addresses addresses_address_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.addresses
    ADD CONSTRAINT addresses_address_key UNIQUE (address);


--
-- Name: addresses addresses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.addresses
    ADD CONSTRAINT addresses_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

