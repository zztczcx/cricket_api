-- Build Initial Table Schemas
CREATE TABLE players (
    id     INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(2500) NOT NULL,
    career_start_year INT,
    career_end_year INT,
    matches INT,
    inns INT,
    not_outs INT,
    runs INT,
    highest_scores INT,
    average FLOAT,
    faced_balls INT,
    strike_rate FLOAT,
    score_hundreds INT,
    score_fiftys INT,
    score_zeros INT,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME NULL,
    PRIMARY KEY(id)
);
