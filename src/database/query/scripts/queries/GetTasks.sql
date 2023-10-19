SELECT
    t.id,
    t.task_name,
    t.content,
    t.description,
    t.date_added
FROM
    workflows_db.workflow_tasks as t
WHERE
    t.group_id = $1