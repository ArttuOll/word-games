-- name: GetRandomWord :one
SELECT * FROM word ORDER BY RANDOM() LIMIT 1;
