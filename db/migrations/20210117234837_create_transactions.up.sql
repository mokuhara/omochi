CREATE TABLE IF NOT EXISTS transactions (
  id BIGINT UNSIGNED auto_increment PRIMARY KEY,
  client_user_id BIGINT UNSIGNED NOT NULL,
  specialist_user_id BIGINT UNSIGNED NOT NULL,
  bizpack_id BIGINT UNSIGNED,
  issue_id BIGINT UNSIGNED,
  status VARCHAR(255),
  title VARCHAR(255),
  category_id BIGINT UNSIGNED,
  description VARCHAR(255),
  unit_price INT,
  duration INT,
  client_acceptance BIGINT UNSIGNED,
  specialist_acceptance BIGINT UNSIGNED,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME,
  CONSTRAINT fk_transactions_on_bizpack_id
    FOREIGN KEY (bizpack_id)
    REFERENCES bizpacks(id)
    ON DELETE SET NULL
    ON UPDATE CASCADE,
  CONSTRAINT fk_transactions_on_issue_id
    FOREIGN KEY (issue_id)
    REFERENCES issues(id)
    ON DELETE SET NULL
    ON UPDATE CASCADE
);

