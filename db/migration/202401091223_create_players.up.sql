-- Build Initial Table Schemas
CREATE TABLE players (
    id     BIGINT AUTO_INCREMENT NOT NULL,
    name VARCHAR(2500) NOT NULL,
    career_start_year BIGINT,
    career_end_year BIGINT,
    matches BIGINT,
    inns BIGINT,
    not_outs BIGINT,
    runs BIGINT,
    highest_scores BIGINT,
    average FLOAT,
    faced_balls BIGINT,
    strike_rate FLOAT,
    score_hundreds BIGINT,
    score_fiftys BIGINT,
    score_zeros BIGINT,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME NOT NULL DEFAULT NOW(),
    PRIMARY KEY(id)
);
