-- name: CreateBook :one
INSERT INTO book (
    book_name,
    book_desc,
    author_id,
    publish_date
) VALUES (
    $1,$2,$3,$4
) RETURNING *;


-- name: GetBook :one
SELECT * FROM book
WHERE id = $1 LIMIT 1;

-- name: GetBooks :many
SELECT * FROM book
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetNewBooks :many
SELECT * FROM book
ORDER BY id ;

-- name: UpdateBook :one
UPDATE book
SET book_desc = $2
WHERE id = $1
RETURNING *;

-- name: BookWithAuthor :one
SELECT * FROM book as b 
INNER JOIN author as a 
ON b.author_id = a.id
WHERE b.id = $1;

-- name: AvailableBooks :many
SELECT 
b.id,
b.book_name,
b.book_desc,
b.author_id,
b.publish_date 
FROM book as b 
INNER JOIN book_managment as bm 
ON b.id = bm.book_id
WHERE bm.total_in_lab > 0
ORDER BY b.id;

-- name: PurchaseBook :one
UPDATE book_managment
SET total_in_lab = total_in_lab - 1,
total_out_lab = total_out_lab + 1
WHERE book_id = $1
RETURNING *;