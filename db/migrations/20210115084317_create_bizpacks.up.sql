CREATE TABLE IF NOT EXISTS bizpacks (
  id BIGINT UNSIGNED auto_increment PRIMARY KEY,
  user_id BIGINT UNSIGNED NOT NULL,
  industry VARCHAR(255),
  scale INT,
  title VARCHAR(255),
  category_id BIGINT UNSIGNED,
  description TEXT,
  unit_price INT,
  duration INT,
  is_public BOOLEAN NOT NULL DEFAULT 0,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME,
  CONSTRAINT fk_bizpacks_on_user_id
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT fk_bizpacks_on_category_id
    FOREIGN KEY (category_id)
    REFERENCES categories(id)
    ON DELETE SET NULL
    ON UPDATE CASCADE
);