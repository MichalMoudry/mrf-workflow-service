SELECT
    w.id,
    w.workflow_name,
    w.setting_is_full_page_recog,
    w.setting_skip_enhancement,
    w.setting_expect_diff_images,
    w.concurrency_stamp,
    w.date_added,
    w.date_updated
FROM
    workflows_db.workflows as w
WHERE
    w.application_id = $1