CREATE TABLE public.users (
	id serial PRIMARY KEY,
    username varchar(250) not null,
    "password" text not null,
	"firstName" varchar(250) not null default '',
    "lastName" varchar(250) not null default '',
    "gender" varchar(100) not null default '',
    "picUrl" text not null default '',
    district varchar(200) not null default '',
    city varchar(200) not null default '',
    "isPremium" boolean not null default false,
	"createdAt" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" timestamp with time zone
);
-- GRANT ALL ON TABLE users TO dhukidwirachman;
-- GRANT ALL ON SEQUENCE users_id_seq TO dhukidwirachman;