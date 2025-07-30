-- name: ListFiles :many
SELECT * FROM files
ORDER BY name;