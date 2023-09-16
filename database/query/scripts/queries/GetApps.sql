SELECT
    a.id,
    a.app_name,
    a.creator_id,
    a.date_added
FROM
    applications as a
WHERE
    a.creator_id = $1