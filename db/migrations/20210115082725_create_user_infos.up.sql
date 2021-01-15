CREATE TABLE IF NOT EXISTS user_infos (
  id BIGINT UNSIGNED auto_increment PRIMARY KEY,
  user_id BIGINT UNSIGNED NOT NULL,
  name VARCHAR(255) NOT NULL,
  kana VARCHAR(255) NOT NULL,
  phone VARCHAR(255),
  company_name VARCHAR(255),
  department VARCHAR(255),
  position VARCHAR(255),
  company_phone VARCHAR(255),
  motivation INT,
  support_request BOOLEAN NOT NULL,
  consent BOOLEAN NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  CONSTRAINT fk_user_infos_on_user_id
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);