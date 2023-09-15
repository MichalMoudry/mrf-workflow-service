INSERT INTO applications (
    id,
    app_name,
    creator_id,
    date_added,
    date_updated
)
VALUES (
    :id,
    :app_name,
    :creator_id,
    :date_added,
    :date_updated
)
RETURNING
    id