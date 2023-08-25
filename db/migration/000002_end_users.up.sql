CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "user_name" varchar NOT NULL,
    "user_email" varchar UNIQUE NOT NULL,
    "user_password" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "users" ("user_name");
