CREATE TABLE public.relation_users (
	id serial PRIMARY KEY,
	"userId" bigint,
    "candidateId" bigint,
    "relationType" text,
	"createdAt" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" timestamp with time zone,
    CONSTRAINT "relation_user_userId_fkey" FOREIGN KEY("userId") REFERENCES users (id),
    CONSTRAINT "relation_user_candidateId_fkey" FOREIGN KEY("candidateId") REFERENCES users (id)
);
-- GRANT ALL ON TABLE relation_users TO dhukidwirachman;
-- GRANT ALL ON SEQUENCE relation_users_id_seq TO dhukidwirachman;