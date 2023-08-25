-- name: CreateBookManagment :one
INSERT INTO book_managment(
    book_id,
    total_quantity,
    total_in_lab,
    total_out_lab
) VALUES (
    $1,$2,$3,$4
) RETURNING *;

-- name: GetBookManagment :one 
SELECT * FROM book_managment
WHERE id = $1 LIMIT 1;

-- name: GetBookManagmentByBookID :one
SELECT * FROM book_managment
WHERE book_id = $1 LIMIT 1;

-- name: GetBookManagments :many
SELECT * FROM book_managment
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateBookManagment :one 
UPDATE book_managment
SET total_quantity = $2, total_in_lab=$3, total_out_lab = $4
WHERE id = $1
RETURNING *;


-- name: DeleteBookManagment :exec
DELETE FROM book_managment
WHERE id = $1;

-- name: UpdateUserBookManagment :one
UPDATE book_managment
SET total_in_lab = total_in_lab + 1,
total_out_lab = total_out_lab - 1
WHERE book_id = $1 AND 
total_quantity > total_out_lab
RETURNING *;