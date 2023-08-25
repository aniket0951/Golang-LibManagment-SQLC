CREATE TABLE "author" (
  "id" bigserial PRIMARY KEY,
  "author_name" varchar UNIQUE NOT NULL,
  "author_address" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE TABLE "book" (
  "id" bigserial PRIMARY KEY,
  "book_name" varchar UNIQUE NOT NULL,
  "book_desc" varchar NOT NULL,
  "author_id" bigint NOT NULL,
  for_this varchar,
  "publish_date" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "book_managment" (
  "id" bigserial PRIMARY KEY,
  "book_id" bigint NOT NULL,
  "total_quantity" int NOT NULL,
  "total_in_lab" int NOT NULL,
  "total_out_lab" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "author" ("author_name");

CREATE INDEX ON "book" ("book_name");

CREATE INDEX ON "book_managment" ("book_id");

ALTER TABLE "book" ADD FOREIGN KEY ("author_id") REFERENCES "author" ("id");

ALTER TABLE "book_managment" ADD FOREIGN KEY ("book_id") REFERENCES "book" ("id");