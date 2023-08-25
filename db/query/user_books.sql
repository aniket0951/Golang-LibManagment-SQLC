-- name: AddUserBooks :one
INSERT INTO user_books (
    user_id,
    book_id,
    purchase_date
) VALUES (
    $1,$2,$3
) RETURNING *;

-- name: UpdateUserBooks :one
UPDATE user_books
SET book_return_date = $3
WHERE user_id = $1 and book_id=$2
RETURNING *;