CREATE TABLE IF NOT EXISTS issues (
  id BIGINT UNSIGNED auto_increment PRIMARY KEY,
  title VARCHAR(255),
  background VARCHAR(255),
  description VARCHAR(255),
  desired_specialist VARCHAR(255),
  required_item VARCHAR(255),
  client_info VARCHAR(255),
  category_id BIGINT UNSIGNED,
  budget INT,
  recruitment_capacity INT,
  start_at VARCHAR(255),
  end_at VARCHAR(255),
  application_deadline VARCHAR(255) NOT NULL,
  user_id BIGINT UNSIGNED NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME,
  CONSTRAINT fk_issues_on_category_id
    FOREIGN KEY (category_id)
    REFERENCES categories(id)
    ON DELETE SET NULL
    ON UPDATE CASCADE
);