BEGIN;

CREATE TABLE applications (
    id UUID PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    creator_id VARCHAR(150) NOT NULL,
    date_added TIMESTAMP NOT NULL,
    date_updated TIMESTAMP NOT NULL
);

CREATE TABLE workflows (
    id UUID PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    application_id UUID NOT NULL,
    setting_is_full_page_recog BOOLEAN NOT NULL,
    setting_skip_enhancement BOOLEAN NOT NULL,
    setting_expect_diff_images BOOLEAN NOT NULL,
    date_added TIMESTAMP NOT NULL,
    date_updated TIMESTAMP NOT NULL
);

CREATE TABLE document_templates (
    id UUID PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    width REAL NOT NULL,
    height REAL NOT NULL,
    image BYTEA NOT NULL,
    workflow_id UUID NOT NULL,
    date_added TIMESTAMP NOT NULL,
    date_updated TIMESTAMP NOT NULL
);

CREATE TABLE template_fields (
    id UUID PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    width REAL NOT NULL,
    height REAL NOT NULL,
    x_position REAL NOT NULL,
    y_position REAL NOT NULL,
    expected_value VARCHAR(255),
    is_identifying BOOLEAN NOT NULL,
    date_added TIMESTAMP NOT NULL,
    date_updated TIMESTAMP NOT NULL
);

CREATE TABLE application_users (
    id UUID PRIMARY KEY,
    application_id UUID NOT NULL,
    user_id VARCHAR(150) NOT NULL,
    role VARCHAR(50) NOT NULL
);

COMMIT;