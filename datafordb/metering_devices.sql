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
-- Name: metering_devices; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.metering_devices (
    id integer NOT NULL,
    ls_object_id bigint NOT NULL,
    day_night_type text NOT NULL,
    device_number text NOT NULL,
    device_type text NOT NULL,
    device_guid uuid NOT NULL,
    last_metering numeric(10,2) NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.metering_devices OWNER TO postgres;

--
-- Name: metering_devices_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.metering_devices_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.metering_devices_id_seq OWNER TO postgres;

--
-- Name: metering_devices_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.metering_devices_id_seq OWNED BY public.metering_devices.id;


--
-- Name: metering_devices id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.metering_devices ALTER COLUMN id SET DEFAULT nextval('public.metering_devices_id_seq'::regclass);


--
-- Data for Name: metering_devices; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.metering_devices (id, ls_object_id, day_night_type, device_number, device_type, device_guid, last_metering, created_at, updated_at) FROM stdin;
1	1	Day	FZRTW705B5	CE 102 R5	3830bef9-1cc5-11e2-8484-d485645e2c54	4030.00	2025-02-04 17:36:43.200972	2025-02-04 17:36:43.200972
2	1	Night	FZRTW705B5	CE 102 R5	3830bef9-1cc5-11e2-8484-d485645e2c55	2478.49	2025-02-04 17:36:43.200972	2025-02-04 17:36:43.200972
3	2	Common	FZRTW705B9	CE 102 R5	3830bf00-1cc5-11e2-8484-d485645e2c56	4000.67	2025-02-04 17:36:43.200972	2025-02-04 17:36:43.200972
4	3	Day	FZRTW705B6	CE 102 R5	3830bf01-1cc5-11e2-8484-d485645e2c57	3500.12	2025-02-04 17:36:43.200972	2025-02-04 17:36:43.200972
5	3	Night	FZRTW705B6	CE 102 R5	3830bf02-1cc5-11e2-8484-d485645e2c58	1800.50	2025-02-04 17:36:43.200972	2025-02-04 17:36:43.200972
7	5	Day	FZRTW705B8	CE 102 R5	3830bf04-1cc5-11e2-8484-d485645e2c60	4200.75	2025-02-04 17:36:43.200972	2025-02-04 17:36:43.200972
8	5	Night	FZRTW705B8	CE 102 R5	3830bf05-1cc5-11e2-8484-d485645e2c61	2100.25	2025-02-04 17:36:43.200972	2025-02-04 17:36:43.200972
9	6	Common	FZRTW705B10	CE 102 R5	3830bf06-1cc5-11e2-8484-d485645e2c62	3800.33	2025-02-04 17:36:43.200972	2025-02-04 17:36:43.200972
12	8	Common	FZRTW705B33	CE 102 R5	3830bf07-1cc5-11e2-8484-d485645e2c63	6789.88	2025-02-04 17:36:43.200972	2025-02-04 17:36:43.200972
11	7	Night	FZRTW705B11	CE 102 R5	3830bf07-1cc5-11e2-8484-d485645e2c90	1202.88	2025-02-04 17:36:43.200972	2025-02-04 17:36:43.200972
10	7	Day	FZRTW705B11	CE 102 R5	3830bf07-1cc5-11e2-8484-d485645e2c78	4503.88	2025-02-04 17:36:43.200972	2025-02-21 16:21:24.983743
6	4	Common	FZRTW705B7	CE 102 R5	3830bf03-1cc5-11e2-8484-d485645e2c59	5001.00	2025-02-04 17:36:43.200972	2025-02-21 17:50:46.01076
\.


--
-- Name: metering_devices_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.metering_devices_id_seq', 12, true);


--
-- Name: metering_devices metering_devices_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.metering_devices
    ADD CONSTRAINT metering_devices_pkey PRIMARY KEY (id);


--
-- Name: metering_devices metering_devices_ls_object_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.metering_devices
    ADD CONSTRAINT metering_devices_ls_object_id_fkey FOREIGN KEY (ls_object_id) REFERENCES public.ls_objects(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

