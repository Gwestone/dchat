--
-- PostgreSQL database cluster dump
--

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Drop databases (except postgres and template1)
--

DROP DATABASE dchat;




--
-- Drop roles
--

DROP ROLE postgres;


--
-- Roles
--

CREATE ROLE postgres;
ALTER ROLE postgres WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN REPLICATION BYPASSRLS PASSWORD 'SCRAM-SHA-256$4096:qt/eIIZcM1mDKVJOA0/Rjw==$1RyiVlPlr/chLwwpA3XyVbFk4/bK3QjtlZjWVw9usxE=:nHP3SwZ9fgmEN9+QQTfA5iSX9D+ls9KvjOcO+ai11Hw=';






--
-- Databases
--

--
-- Database "template1" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3 (Debian 14.3-1.pgdg110+1)
-- Dumped by pg_dump version 14.3 (Debian 14.3-1.pgdg110+1)

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

UPDATE pg_catalog.pg_database SET datistemplate = false WHERE datname = 'template1';
DROP DATABASE template1;
--
-- Name: template1; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE template1 WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE template1 OWNER TO postgres;

\connect template1

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
-- Name: DATABASE template1; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE template1 IS 'default template for new databases';


--
-- Name: template1; Type: DATABASE PROPERTIES; Schema: -; Owner: postgres
--

ALTER DATABASE template1 IS_TEMPLATE = true;


\connect template1

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
-- Name: DATABASE template1; Type: ACL; Schema: -; Owner: postgres
--

REVOKE CONNECT,TEMPORARY ON DATABASE template1 FROM PUBLIC;
GRANT CONNECT ON DATABASE template1 TO PUBLIC;


--
-- PostgreSQL database dump complete
--

--
-- Database "dchat" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3 (Debian 14.3-1.pgdg110+1)
-- Dumped by pg_dump version 14.3 (Debian 14.3-1.pgdg110+1)

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
-- Name: dchat; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE dchat WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE dchat OWNER TO postgres;

\connect dchat

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
    "MessageText" text NOT NULL,
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

COPY main.messages ("Id", "From", "To", "MessageText", "Date") FROM stdin;
1	user2	user1	hello world	1654612850
2	user1	user2	hello user1	1654612868
3	user4	user1	1	1654613403
4	user2	Password1	adnauwd	1654619678
5	gwestone	Password1	1312313	1654619698
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: main; Owner: postgres
--

COPY main.users ("Id", "Username", "Password", "UserId") FROM stdin;
1	user1	$2a$14$W1YRwgHCBkrkGjZEcA1MDueABPDkfTaUacSvaRUMBcefa9rDSzy96	f5806ddc-8f0c-43dd-b394-d6b41a49b835
2	user2	$2a$14$Boe6/BsL1k4sHkwndaxHdODivA8AvwBM0oTgGOwglikHIxZY6Fb7y	6679c58e-2808-4720-8704-4bf9a335ec5c
3	user3	$2a$14$Zt15BqzK6MUWMIc32gJZTOjr0oEFt9Rd9afYAC.O8FV077KGTmZW6	b97a685d-d215-4a98-99eb-2083a18cee10
5	user4	$2a$14$snkhmAmor2cGkmTkeiReXutbeHwxsL7SQdphDh4T1PqYFndyyLxFq	b18a3bd8-8494-42f1-a137-2a626c586e09
6	gwestone	$2a$14$faAsou1FFN6.QAbWFxAzPe8rbC.FOK11iU5O2XlADh5G1W0gFohk.	ec052ad1-6a2a-4189-8915-c9b423495b75
8	adwaddwad	$2a$14$y2HvD12Kxq33aZEeRpaib.3H.j53LcFIlAMuFrg8DGV5qzH9HD8TS	11c6935f-9b99-41d1-8fb3-1babbd873b1a
9	111	$2a$14$N2DYAxTiuD90VJFYt.wDYu3jj4luqFkg2X7PFlwUvaX3fk.xGTGU6	15bf21bb-c1aa-4070-b7a3-ee9c4d933f3b
10	Password1	$2a$14$0WA5DXpt85LHW1owywvz9u3SA4B/y.lhHePvi.jFD4jKKLO0M5bJa	fd4f3659-4941-49b5-847a-f3fc9301c0c4
11	poson1o	$2a$14$zx3jtaYvH9BkZD1QJ0XUBOsNsX1NcO4XME7tvstLong2NS33Lejde	ff02eb46-04e9-4650-ba81-99fc34756f17
12	122	$2a$14$PXlZPhCuVeDFJVVmKM0G7.RGDYi1X3cmfW/g8mhpfF6eI/DgNeZ9O	4adceead-b0bd-4826-a456-4bc656c9668a
\.


--
-- Name: messages_Id_seq; Type: SEQUENCE SET; Schema: main; Owner: postgres
--

SELECT pg_catalog.setval('main."messages_Id_seq"', 5, true);


--
-- Name: users_Id_seq; Type: SEQUENCE SET; Schema: main; Owner: postgres
--

SELECT pg_catalog.setval('main."users_Id_seq"', 12, true);


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
-- Name: users_password_uindex; Type: INDEX; Schema: main; Owner: postgres
--

CREATE UNIQUE INDEX users_password_uindex ON main.users USING btree ("Password");


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
-- Name: messages messages_users_username_fk_; Type: FK CONSTRAINT; Schema: main; Owner: postgres
--

ALTER TABLE ONLY main.messages
    ADD CONSTRAINT messages_users_username_fk_ FOREIGN KEY ("To") REFERENCES main.users("Username") ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

--
-- Database "postgres" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3 (Debian 14.3-1.pgdg110+1)
-- Dumped by pg_dump version 14.3 (Debian 14.3-1.pgdg110+1)

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

DROP DATABASE postgres;
--
-- Name: postgres; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE postgres OWNER TO postgres;

\connect postgres

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
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database cluster dump complete
--

