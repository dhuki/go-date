CREATE TABLE public.user (
	id serial PRIMARY KEY,
	"firstName" varchar(250),
    "lastName" varchar(250),
    email text,
    "picUrl" text,
    district varchar(200)
    city varchar(200)
    "isPremium" boolean,
	"createdAt" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" timestamp with time zone
);

GRANT ALL ON SEQUENCE user TO postgres;
GRANT ALL ON SEQUENCE user_id_seq TO postgres;