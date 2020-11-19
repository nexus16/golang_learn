CREATE TABLE IF NOT EXISTS Member_Skill (
    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
    member_id BIGINT NOT NULL,
    skill_id BIGINT NOT NULL,
    UNIQUE (member_id,skill_id)
);