--
-- PostgreSQL database dump
--

-- Dumped from database version 14.12 (Homebrew)
-- Dumped by pg_dump version 14.12 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: update_updated_at(); Type: FUNCTION; Schema: public; Owner: maxim
--

CREATE FUNCTION public.update_updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
begin
new.updated_at=now();
return new;
end;
$$;


ALTER FUNCTION public.update_updated_at() OWNER TO maxim;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: statistics; Type: TABLE; Schema: public; Owner: maxim
--

CREATE TABLE public.statistics (
    id integer NOT NULL,
    url_id integer NOT NULL,
    user_id integer,
    ip inet NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone
);


ALTER TABLE public.statistics OWNER TO maxim;

--
-- Name: statistics_id_seq; Type: SEQUENCE; Schema: public; Owner: maxim
--

CREATE SEQUENCE public.statistics_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.statistics_id_seq OWNER TO maxim;

--
-- Name: statistics_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: maxim
--

ALTER SEQUENCE public.statistics_id_seq OWNED BY public.statistics.id;


--
-- Name: urls; Type: TABLE; Schema: public; Owner: maxim
--

CREATE TABLE public.urls (
    id integer NOT NULL,
    code character(4) NOT NULL,
    value character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone
);


ALTER TABLE public.urls OWNER TO maxim;

--
-- Name: urls_id_seq; Type: SEQUENCE; Schema: public; Owner: maxim
--

CREATE SEQUENCE public.urls_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.urls_id_seq OWNER TO maxim;

--
-- Name: urls_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: maxim
--

ALTER SEQUENCE public.urls_id_seq OWNED BY public.urls.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: maxim
--

CREATE TABLE public.users (
    id integer NOT NULL,
    login character varying(50) NOT NULL,
    password character varying(50) NOT NULL,
    email character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO maxim;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: maxim
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO maxim;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: maxim
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: statistics id; Type: DEFAULT; Schema: public; Owner: maxim
--

ALTER TABLE ONLY public.statistics ALTER COLUMN id SET DEFAULT nextval('public.statistics_id_seq'::regclass);


--
-- Name: urls id; Type: DEFAULT; Schema: public; Owner: maxim
--

ALTER TABLE ONLY public.urls ALTER COLUMN id SET DEFAULT nextval('public.urls_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: maxim
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: statistics; Type: TABLE DATA; Schema: public; Owner: maxim
--

COPY public.statistics (id, url_id, user_id, ip, created_at, updated_at) FROM stdin;
1	1	2	127.0.0.1	2024-06-14 15:10:09.968875	\N
\.


--
-- Data for Name: urls; Type: TABLE DATA; Schema: public; Owner: maxim
--

COPY public.urls (id, code, value, created_at, updated_at) FROM stdin;
1	test	https://google.com/?q=test	2024-06-14 14:58:17.806637	\N
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: maxim
--

COPY public.users (id, login, password, email, created_at, updated_at) FROM stdin;
2	login123	password	user@email.com	2024-06-14 14:01:27.248386	2024-06-14 14:52:32.245865
\.


--
-- Name: statistics_id_seq; Type: SEQUENCE SET; Schema: public; Owner: maxim
--

SELECT pg_catalog.setval('public.statistics_id_seq', 3, true);


--
-- Name: urls_id_seq; Type: SEQUENCE SET; Schema: public; Owner: maxim
--

SELECT pg_catalog.setval('public.urls_id_seq', 1, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: maxim
--

SELECT pg_catalog.setval('public.users_id_seq', 2, true);


--
-- Name: urls code; Type: CONSTRAINT; Schema: public; Owner: maxim
--

ALTER TABLE ONLY public.urls
    ADD CONSTRAINT code UNIQUE (code);


--
-- Name: statistics statistics_pkey; Type: CONSTRAINT; Schema: public; Owner: maxim
--

ALTER TABLE ONLY public.statistics
    ADD CONSTRAINT statistics_pkey PRIMARY KEY (id);


--
-- Name: urls urls_pkey; Type: CONSTRAINT; Schema: public; Owner: maxim
--

ALTER TABLE ONLY public.urls
    ADD CONSTRAINT urls_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: maxim
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_login_key; Type: CONSTRAINT; Schema: public; Owner: maxim
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_login_key UNIQUE (login);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: maxim
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: statistics statistics_updated_at; Type: TRIGGER; Schema: public; Owner: maxim
--

CREATE TRIGGER statistics_updated_at BEFORE UPDATE ON public.statistics FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: urls urls_updated_at; Type: TRIGGER; Schema: public; Owner: maxim
--

CREATE TRIGGER urls_updated_at BEFORE UPDATE ON public.urls FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: users users_updated_at; Type: TRIGGER; Schema: public; Owner: maxim
--

CREATE TRIGGER users_updated_at BEFORE UPDATE ON public.users FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: statistics url_id; Type: FK CONSTRAINT; Schema: public; Owner: maxim
--

ALTER TABLE ONLY public.statistics
    ADD CONSTRAINT url_id FOREIGN KEY (url_id) REFERENCES public.urls(id);


--
-- Name: statistics user_id; Type: FK CONSTRAINT; Schema: public; Owner: maxim
--

ALTER TABLE ONLY public.statistics
    ADD CONSTRAINT user_id FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

