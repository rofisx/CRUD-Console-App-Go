-- CREATE DATABASE laundryChallenge;

CREATE TABLE mst_satuan
(
    satuanid character varying(100) NOT NULL,
    nama character varying(100) NOT NULL,
    indate timestamp without time zone NOT NULL,
    inby character varying(100) NOT NULL,
    updateddate timestamp without time zone,
    updatedby character varying(100),
    CONSTRAINT mst_satuan_pkey PRIMARY KEY (satuanid)
);

CREATE TABLE mst_service
(
    serviceid integer NOT NULL,
    nama character varying(100) NOT NULL,
    satuan character varying(100) NOT NULL,
    harga integer NOT NULL,
    indate timestamp without time zone NOT NULL,
    inby character varying(100) NOT NULL,
    updateddate timestamp without time zone,
    updatedby character varying(100),
    CONSTRAINT mst_service_pkey PRIMARY KEY (serviceid),
    CONSTRAINT satuanidfkey FOREIGN KEY (satuan)
    REFERENCES mst_satuan (satuanid)
);

CREATE TABLE trx_laundry
(
    trsid character varying(100) NOT NULL,
    customer character varying(100) NOT NULL,
    contact bigint NOT NULL,
    totalqty integer NOT NULL,
    totaltagihan integer NOT NULL,
    indate timestamp without time zone NOT NULL,
    inby character varying(100) NOT NULL,
    CONSTRAINT trx_laundry_pkey PRIMARY KEY (trsid)
);

CREATE TABLE trx_laundry_detail(
	detailid serial,
    trsid character varying(100) NOT NULL,
    serviceid integer NOT NULL,
    harga integer NOT NULL,
    qty integer NOT NULL,
    CONSTRAINT trx_laundry_detail_pkey PRIMARY KEY (detailid),
    CONSTRAINT trx_laundry_detail_serviceid_fkey FOREIGN KEY (serviceid)
        REFERENCES public.mst_service (serviceid),
    CONSTRAINT trx_laundry_detail_trsid_fkey FOREIGN KEY (trsid)
        REFERENCES public.trx_laundry (trsid) 
);