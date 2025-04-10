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
-- Name: object_types; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.object_types (
    id integer NOT NULL,
    type text NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.object_types OWNER TO postgres;

--
-- Name: object_types_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.object_types_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.object_types_id_seq OWNER TO postgres;

--
-- Name: object_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.object_types_id_seq OWNED BY public.object_types.id;


--
-- Name: object_types id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.object_types ALTER COLUMN id SET DEFAULT nextval('public.object_types_id_seq'::regclass);


--
-- Data for Name: object_types; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.object_types (id, type, created_at, updated_at) FROM stdin;
1	House	2025-02-04 17:00:48.051018	2025-02-04 17:00:48.051018
2	Office	2025-02-04 17:00:48.051018	2025-02-04 17:00:48.051018
3	Apartment	2025-02-04 17:00:48.051018	2025-02-04 17:00:48.051018
\.


--
-- Name: object_types_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.object_types_id_seq', 3, true);


--
-- Name: object_types object_types_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.object_types
    ADD CONSTRAINT object_types_name_key UNIQUE (type);


--
-- Name: object_types object_types_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.object_types
    ADD CONSTRAINT object_types_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

