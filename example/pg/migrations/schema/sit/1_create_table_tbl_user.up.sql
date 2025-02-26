CREATE TABLE tbl_user (
    id VARCHAR(36) NOT NULL,
    user_ref_id VARCHAR(36) NOT NULL,
    name VARCHAR(100) NOT NULL,
    image_profile VARCHAR(255),
    max_transfer_per_day FLOAT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    created_by_id VARCHAR(36) NOT NULL,
    created_by VARCHAR(200) NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    updated_by_id VARCHAR(36) NOT NULL,
    updated_by VARCHAR(200) NOT NULL,
    CONSTRAINT tbl_user_pkey PRIMARY KEY (id, updated_at)
);

CREATE INDEX idx_user_ref_id ON tbl_user USING BTREE (user_ref_id);