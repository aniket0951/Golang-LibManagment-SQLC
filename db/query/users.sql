-- name: CreateUser :one
INSERT INTO users (
    user_name,
    user_email,
    user_password
) VALUES (
    $1,$2,$3
) RETURNING *;

-- name: GetUser :one 
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users SET user_name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;


-- name: GetUserWithBooks :many
SELECT * FROM users as u
INNER JOIN user_books as ub 
ON u.id = ub.user_id
INNER JOIN book as b
ON ub.book_id = b.id
RIGHT JOIN author as a 
ON b.author_id = a.id
LEFT JOIN  author_address as ad 
ON a.id = ad.author_id
WHERE u.id = $1;