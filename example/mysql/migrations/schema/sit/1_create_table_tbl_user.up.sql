CREATE TABLE tbl_user (
    id VARCHAR(36) NOT NULL,
    user_ref_id VARCHAR(36) NOT NULL,
    name VARCHAR(100) NOT NULL,
    image_profile VARCHAR(255) NULL,
    max_transfer_per_day FLOAT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    created_by_id VARCHAR(36) NOT NULL,
    created_by VARCHAR(200) NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    updated_by_id VARCHAR(36) NOT NULL,
    updated_by VARCHAR(200) NOT NULL,
    PRIMARY KEY (id, updated_at)
);