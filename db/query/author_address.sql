-- name: CreateAuthorAddress :one
insert into author_address (
    address_line_one,
    city,
    state,
    country,
    author_id
) values (
    $1,$2,$3,$4,$5
) RETURNING *;