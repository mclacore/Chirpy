-- name: PostChirp :one
INSERT INTO chirps (id, created_at, updated_at, body, user_id)
VALUES (
	gen_random_uuid(),
	NOW(),
	NOW(),
	$1,
	$2
)
RETURNING *;

-- name: GetChirps :many
select *
from chirps
order by created_at asc
;

-- name: GetChirp :one
select *
from chirps
where id = $1
;

-- name: DeleteAllChirps :exec
delete from chirps
;

