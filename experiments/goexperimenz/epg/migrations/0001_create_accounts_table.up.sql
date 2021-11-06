CREATE TABLE public.accounts
(
    id bigserial,
    name character varying(1024) NOT NULL,
    email character varying(1024) NOT NULL,
    imap_server character varying(1024) NOT NULL,
    imap_port numeric NOT NULL,
    use_ssl boolean,
    password character varying(1024) NOT NULL,
    PRIMARY KEY (id)
)
