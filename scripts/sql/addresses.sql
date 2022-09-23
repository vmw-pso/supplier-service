SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: addressess; Type: TABLE; Schema: public; Owner: postgres;
--
CREATE TABLE public.addresses (
    id SERIAL PRIMARY KEY,
    building CHARACTER VARYING(255),
    unit_floor CHARACTER VARYING(255),
    street_number CHARACTER VARYING(255),
    street_name CHARACTER VARYING(255),
    city CHARACTER VARYING(255),
    zip_or_postcode CHARACTER VARYING(255),
    state_or_territory CHARACTER VARYING(255),
    country CHARACTER VARYING(255),
);

ALTER TABLE public.suppliers OWNER TO postgres;

INSERT INTO "public"."addresses"("building", "unit_floor", "street_number", "street_name", "city", "zip_or_postcode", "state_or_territory", "country") 
VALUES(E'Nishi Building', E'Level 9', E'2', E'Phillip Law Street', E'Canberra', E'2601', E'ACT', E'Australia');
INSERT INTO "public"."addresses"("street_number", "street_name", "city", "zip_or_postcode", "state_or_territory", "country") 
VALUES(E'2', E'Limestone Drive', E'Jerrabomberra', E'2619', E'NSW', E'Australia');