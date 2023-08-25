
CREATE TABLE "user_books" (
 	"id" bigserial PRIMARY KEY,
 	"user_id" bigint NOT NULL,
 	"book_id" bigint NOT NULL,
 	"created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "user_books" ("user_id");

ALTER TABLE "user_books" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "user_books" ADD FOREIGN KEY ("book_id") REFERENCES "book" ("id");

