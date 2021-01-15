CREATE TABLE IF NOT EXISTS companies (
  id BIGINT UNSIGNED auto_increment PRIMARY KEY,
  name text NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL
);