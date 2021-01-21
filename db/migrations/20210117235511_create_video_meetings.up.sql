CREATE TABLE IF NOT EXISTS video_meetings (
  id BIGINT UNSIGNED auto_increment PRIMARY KEY,
  name VARCHAR(255),
  url VARCHAR(255),
  started_at VARCHAR(255) NOT NULL,
  transaction_id BIGINT UNSIGNED NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME,
  CONSTRAINT fk_video_meetings_on_transaction_id
    FOREIGN KEY (transaction_id)
    REFERENCES transactions(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);