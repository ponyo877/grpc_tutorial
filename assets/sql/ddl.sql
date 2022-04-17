CREATE TABLE app_db.tutorial_tbl(
    place_id    VARCHAR(26) NOT NULL, 
    logo_url_id VARCHAR(26),
    color_code  VARCHAR(7),
    updated_by  VARCHAR(26) NOT NULL,
    updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    created_by  VARCHAR(26) NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (place_id)
);