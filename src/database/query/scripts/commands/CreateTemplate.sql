INSERT INTO document_templates (
    id,
    template_name,
    width,
    height,
    image,
    workflow_id,
    concurrency_stamp,
    date_added,
    date_updated
)
VALUES (
    :id,
    :template_name,
    :width,
    :height,
    :image,
    :workflow_id,
    :concurrency_stamp,
    :date_added,
    :date_updated
)
RETURNING
    id