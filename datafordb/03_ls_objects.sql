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
-- Name: ls_objects; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ls_objects (
    id integer NOT NULL,
    num_ls bigint NOT NULL,
    object_type_id integer NOT NULL,
    address_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.ls_objects OWNER TO postgres;

--
-- Name: ls_objects_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ls_objects_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ls_objects_id_seq OWNER TO postgres;

--
-- Name: ls_objects_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ls_objects_id_seq OWNED BY public.ls_objects.id;


--
-- Name: ls_objects id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ls_objects ALTER COLUMN id SET DEFAULT nextval('public.ls_objects_id_seq'::regclass);


--
-- Data for Name: ls_objects; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ls_objects (id, num_ls, object_type_id, address_id, created_at, updated_at) FROM stdin;
1	11111	1	1	2025-02-04 17:18:55.669342	2025-02-04 17:18:55.669342
2	11112	1	2	2025-02-04 17:18:55.669342	2025-02-04 17:18:55.669342
3	11113	2	3	2025-02-04 17:18:55.669342	2025-02-04 17:18:55.669342
4	11114	1	4	2025-02-04 17:18:55.669342	2025-02-04 17:18:55.669342
5	11115	2	5	2025-02-04 17:18:55.669342	2025-02-04 17:18:55.669342
6	11116	1	6	2025-02-04 17:18:55.669342	2025-02-04 17:18:55.669342
7	11117	3	7	2025-02-04 17:18:55.669342	2025-02-04 17:18:55.669342
8	11118	3	8	2025-02-04 17:18:55.669342	2025-02-04 17:18:55.669342
\.


--
-- Name: ls_objects_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ls_objects_id_seq', 8, true);


--
-- Name: ls_objects ls_objects_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ls_objects
    ADD CONSTRAINT ls_objects_pkey PRIMARY KEY (id);


--
-- Name: ls_objects ls_objects_address_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ls_objects
    ADD CONSTRAINT ls_objects_address_id_fkey FOREIGN KEY (address_id) REFERENCES public.addresses(id);


--
-- Name: ls_objects ls_objects_object_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ls_objects
    ADD CONSTRAINT ls_objects_object_type_id_fkey FOREIGN KEY (object_type_id) REFERENCES public.object_types(id);


--
-- PostgreSQL database dump complete
--

