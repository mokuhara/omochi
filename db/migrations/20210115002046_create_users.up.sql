CREATE TABLE IF NOT EXISTS users (
  id BIGINT UNSIGNED auto_increment PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  password text NOT NULL,
  type INTEGER UNSIGNED,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME
);