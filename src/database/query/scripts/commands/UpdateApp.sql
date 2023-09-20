UPDATE
    applications
SET
    app_name = $2,
    date_updated = $3
WHERE
    id = $1