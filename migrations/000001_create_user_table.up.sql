CREATE TABLE "users" (
     "id" bigserial PRIMARY KEY,
     "name" VARCHAR NOT NULL,
     "email" VARCHAR UNIQUE,
     "password" VARCHAR,
     "created_at" TIMESTAMP,
     "updated_at" TIMESTAMP
);