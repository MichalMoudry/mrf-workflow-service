BEGIN;

CREATE TABLE applications (
    id UUID PRIMARY KEY,
    app_name VARCHAR(200) NOT NULL,
    creator_id VARCHAR(150) NOT NULL,
    concurrency_stamp UUID NOT NULL,
    date_added TIMESTAMP NOT NULL,
    date_updated TIMESTAMP NOT NULL
);

CREATE TABLE workflows (
    id UUID PRIMARY KEY,
    workflow_name VARCHAR(200) NOT NULL,
    application_id UUID NOT NULL,
    setting_is_full_page_recog BOOLEAN NOT NULL,
    setting_skip_enhancement BOOLEAN NOT NULL,
    setting_expect_diff_images BOOLEAN NOT NULL,
    concurrency_stamp UUID NOT NULL,
    date_added TIMESTAMP NOT NULL,
    date_updated TIMESTAMP NOT NULL,
    CONSTRAINT fk_application
        FOREIGN KEY(application_id)
            REFERENCES applications(id)
            ON DELETE CASCADE
);

CREATE TABLE document_templates (
    id UUID PRIMARY KEY,
    template_name VARCHAR(200) NOT NULL,
    width REAL NOT NULL,
    height REAL NOT NULL,
    image BYTEA NOT NULL,
    workflow_id UUID NOT NULL,
    concurrency_stamp UUID NOT NULL,
    date_added TIMESTAMP NOT NULL,
    date_updated TIMESTAMP NOT NULL,
    CONSTRAINT fk_workflow
        FOREIGN KEY(workflow_id)
            REFERENCES workflows(id)
            ON DELETE CASCADE
);

CREATE TABLE template_fields (
    id UUID PRIMARY KEY,
    field_name VARCHAR(200) NOT NULL,
    width REAL NOT NULL,
    height REAL NOT NULL,
    x_position REAL NOT NULL,
    y_position REAL NOT NULL,
    expected_value VARCHAR(255),
    is_identifying BOOLEAN NOT NULL,
    template_id UUID NOT NULL,
    concurrency_stamp UUID NOT NULL,
    date_added TIMESTAMP NOT NULL,
    date_updated TIMESTAMP NOT NULL,
    CONSTRAINT fk_template
        FOREIGN KEY(template_id)
            REFERENCES document_templates(id)
            ON DELETE CASCADE
);

CREATE TABLE application_users (
    id UUID PRIMARY KEY,
    application_id UUID NOT NULL,
    user_id VARCHAR(150) NOT NULL,
    user_role VARCHAR(50) NOT NULL,
    CONSTRAINT fk_application
        FOREIGN KEY(application_id)
            REFERENCES applications(id)
            ON DELETE CASCADE
);

COMMIT;