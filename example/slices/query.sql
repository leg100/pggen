-- name: GetNilTime :one
SELECT pggen.arg('data')::timestamp;

-- name: GetTimes :many
SELECT *
FROM unnest(pggen.arg('data')::timestamp WITH TIME ZONE[]);

-- name: GetNilTimes :many
SELECT *
FROM unnest(pggen.arg('data')::timestamp[]);

-- name: GetCustomIntervals :many
SELECT *
FROM unnest(pggen.arg('data')::interval[]);

-- name: GetCustomDates :many
SELECT *
FROM unnest(pggen.arg('data')::date[]);
