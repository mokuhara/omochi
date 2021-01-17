CREATE TABLE IF NOT EXISTS products (
  id BIGINT UNSIGNED auto_increment PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  bizpack_id BIGINT UNSIGNED NOT NULL,
  portfolio_id BIGINT UNSIGNED NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME,
  CONSTRAINT fk_products_on_bizpack_id
    FOREIGN KEY (bizpack_id)
    REFERENCES bizpacks(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT fk_products_on_portfolio_id
    FOREIGN KEY (portfolio_id)
    REFERENCES portfolios(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);