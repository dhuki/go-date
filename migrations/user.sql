CREATE TABLE public.users (
	id serial PRIMARY KEY,
    username varchar(250),
    "password" text,
	"firstName" varchar(250),
    "lastName" varchar(250),
    "gender" varchar(100),
    "picUrl" text,
    district varchar(200),
    city varchar(200),
    "isPremium" boolean,
	"createdAt" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" timestamp with time zone
);
-- GRANT ALL ON TABLE users TO dhukidwirachman;
-- GRANT ALL ON SEQUENCE users_id_seq TO dhukidwirachman;