INSERT INTO workflows_db.applications (
    id,
    app_name,
    creator_id,
    concurrency_stamp,
    date_added,
    date_updated
)
VALUES (
    :id,
    :app_name,
    :creator_id,
    :concurrency_stamp,
    :date_added,
    :date_updated
)
RETURNING
    id