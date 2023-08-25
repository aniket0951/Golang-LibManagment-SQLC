CREATE TABLE "author_address" (
  "id" bigserial PRIMARY KEY,
  "address_line_one" varchar NOT NULL,
  "city" varchar NOT NULL,
  "state" varchar NOT NULL,
  "country" varchar NOT NULL,
  "author_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "author_address" ("author_id");
ALTER TABLE "author_address" ADD FOREIGN KEY ("author_id") REFERENCES "author" ("id");