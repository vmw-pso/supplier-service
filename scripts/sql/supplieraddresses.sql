SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: supplieraddresses; Type: TABLE; Schema: public; Owner: postgres;
--
CREATE TABLE public.supplier_addresses (
    id SERIAL PRIMARY KEY,
    supplier_id INTEGER NOT NULL,
    address_id INTEGER NOT NULL,
    UNIQUE (supplier_id, address_id)
);

ALTER TABLE public.suppliers OWNER TO postgres;

INSERT INTO "public"."addresses"("building", "unit_floor", "street_number", "street_name", "city", "zip_or_postcode", "state_or_territory", "country") 
VALUES(E'Nishi Building', E'Level 9', E'2', E'Phillip Law Street', E'Canberra', E'2601', E'ACT', E'Australia');
INSERT INTO "public"."addresses"("street_number", "street_name", "city", "zip_or_postcode", "state_or_territory", "country") 
VALUES(E'2', E'Limestone Drive', E'Jerrabomberra', E'2619', E'NSW', E'Australia');