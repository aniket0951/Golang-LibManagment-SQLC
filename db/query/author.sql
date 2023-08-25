-- name: CreateAuthor :one
INSERT INTO author(
    author_name,
    author_address
) VALUES (
    $1,$2
) RETURNING *;

-- name: GetAuthor :one
SELECT * FROM author
WHERE id = $1 LIMIT 1;

-- name: GetAuthors :many
SELECT * FROM author
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAuthor :one
UPDATE author
SET author_name = $2, author_address = $3
WHERE id = $1
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM author
WHERE id = $1;

-- name: GetAuthorWithBooks :many
SELECT * FROM author AS a 
INNER JOIN book AS B 
ON a.id = b.author_id
WHERE a.id = $1;

-- name: GetAuthorWithBooksAndManagment :many
SELECT DISTINCT *  FROM author AS a
INNER JOIN book AS b
ON a.id = b.author_id
LEFT JOIN book_managment AS bm 
ON b.id = bm.book_id
WHERE a.id = $1;
