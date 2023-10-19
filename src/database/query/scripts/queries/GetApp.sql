SELECT
    a.id,
    a.app_name,
    a.creator_id,
    a.concurrency_stamp,
    a.date_added,
    a.date_updated
FROM
    workflows_db.applications as a
WHERE
    a.id = $1