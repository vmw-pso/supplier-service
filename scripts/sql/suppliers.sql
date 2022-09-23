SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: suppliers; Type: TABLE; Schema: public; Owner: postgres;
--
CREATE TABLE public.suppliers (
    id SERIAL PRIMARY KEY,
    name CHARACTER VARYING(255),
);

ALTER TABLE public.suppliers OWNER TO postgres;

INSERT INTO "public"."users"("name") VALUES(E'VMware Australian Pty Ltd');
INSERT INTO "public"."users"("name") VALUES(E'Invictus Services');