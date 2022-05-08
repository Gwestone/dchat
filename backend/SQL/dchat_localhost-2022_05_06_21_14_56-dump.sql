--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2
-- Dumped by pg_dump version 14.2

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
-- Name: main; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA main;


ALTER SCHEMA main OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: messages; Type: TABLE; Schema: main; Owner: postgres
--

CREATE TABLE main.messages (
    "Id" integer NOT NULL,
    "From" text NOT NULL,
    "To" text NOT NULL,
    "Message" text NOT NULL,
    "Date" integer NOT NULL
);


ALTER TABLE main.messages OWNER TO postgres;

--
-- Name: messages_Id_seq; Type: SEQUENCE; Schema: main; Owner: postgres
--

CREATE SEQUENCE main."messages_Id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE main."messages_Id_seq" OWNER TO postgres;

--
-- Name: messages_Id_seq; Type: SEQUENCE OWNED BY; Schema: main; Owner: postgres
--

ALTER SEQUENCE main."messages_Id_seq" OWNED BY main.messages."Id";


--
-- Name: users; Type: TABLE; Schema: main; Owner: postgres
--

CREATE TABLE main.users (
    "Id" integer NOT NULL,
    "Username" text NOT NULL,
    "Password" text NOT NULL,
    "UserId" text NOT NULL
);


ALTER TABLE main.users OWNER TO postgres;

--
-- Name: users_Id_seq; Type: SEQUENCE; Schema: main; Owner: postgres
--

CREATE SEQUENCE main."users_Id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE main."users_Id_seq" OWNER TO postgres;

--
-- Name: users_Id_seq; Type: SEQUENCE OWNED BY; Schema: main; Owner: postgres
--

ALTER SEQUENCE main."users_Id_seq" OWNED BY main.users."Id";


--
-- Name: messages Id; Type: DEFAULT; Schema: main; Owner: postgres
--

ALTER TABLE ONLY main.messages ALTER COLUMN "Id" SET DEFAULT nextval('main."messages_Id_seq"'::regclass);


--
-- Name: users Id; Type: DEFAULT; Schema: main; Owner: postgres
--

ALTER TABLE ONLY main.users ALTER COLUMN "Id" SET DEFAULT nextval('main."users_Id_seq"'::regclass);


--
-- Data for Name: messages; Type: TABLE DATA; Schema: main; Owner: postgres
--

COPY main.messages ("Id", "From", "To", "Message", "Date") FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: main; Owner: postgres
--

COPY main.users ("Id", "Username", "Password", "UserId") FROM stdin;

--
-- Name: messages_Id_seq; Type: SEQUENCE SET; Schema: main; Owner: postgres
--

SELECT pg_catalog.setval('main."messages_Id_seq"', 5, true);


--
-- Name: users_Id_seq; Type: SEQUENCE SET; Schema: main; Owner: postgres
--

SELECT pg_catalog.setval('main."users_Id_seq"', 6, true);


--
-- Name: messages messages_pk; Type: CONSTRAINT; Schema: main; Owner: postgres
--

ALTER TABLE ONLY main.messages
    ADD CONSTRAINT messages_pk PRIMARY KEY ("Id");


--
-- Name: users users_pk; Type: CONSTRAINT; Schema: main; Owner: postgres
--

ALTER TABLE ONLY main.users
    ADD CONSTRAINT users_pk PRIMARY KEY ("Id");


--
-- Name: messages_id_uindex; Type: INDEX; Schema: main; Owner: postgres
--

CREATE UNIQUE INDEX messages_id_uindex ON main.messages USING btree ("Id");


--
-- Name: users_id_uindex; Type: INDEX; Schema: main; Owner: postgres
--

CREATE UNIQUE INDEX users_id_uindex ON main.users USING btree ("Id");


--
-- Name: users_userid_uindex; Type: INDEX; Schema: main; Owner: postgres
--

CREATE UNIQUE INDEX users_userid_uindex ON main.users USING btree ("UserId");


--
-- Name: users_username_uindex; Type: INDEX; Schema: main; Owner: postgres
--

CREATE UNIQUE INDEX users_username_uindex ON main.users USING btree ("Username");


--
-- Name: messages messages_users_username_fk; Type: FK CONSTRAINT; Schema: main; Owner: postgres
--

ALTER TABLE ONLY main.messages
    ADD CONSTRAINT messages_users_username_fk FOREIGN KEY ("From") REFERENCES main.users("Username") ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: messages messages_users_username_fk_2; Type: FK CONSTRAINT; Schema: main; Owner: postgres
--

ALTER TABLE ONLY main.messages
    ADD CONSTRAINT messages_users_username_fk_2 FOREIGN KEY ("To") REFERENCES main.users("Username") ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

