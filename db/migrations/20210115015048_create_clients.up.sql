CREATE TABLE IF NOT EXISTS clients (
  id BIGINT UNSIGNED auto_increment PRIMARY KEY,
  user_id BIGINT UNSIGNED NOT NULL,
  company_id BIGINT UNSIGNED NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME,
  CONSTRAINT fk_clients_on_user_id
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT fk_clients_on_company_id
    FOREIGN KEY (company_id)
    REFERENCES companies(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);