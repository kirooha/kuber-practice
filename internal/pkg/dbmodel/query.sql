-- name: ListFiles :many
SELECT * FROM files
ORDER BY name;

-- name: AddFile :execresult
INSERT INTO files(name, content)
VALUES(@name, @content);