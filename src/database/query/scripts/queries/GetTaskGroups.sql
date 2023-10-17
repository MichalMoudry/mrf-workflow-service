SELECT
    tg.id,
    tg.group_name,
    tg.date_added
FROM
    workflows_db.workflow_task_groups as tg
WHERE
    tg.workflow_id = $1