CREATE TABLE `products` (
  `id` CHAR(36) NOT NULL,
  `code` VARCHAR(255) UNIQUE,
  `name` VARCHAR(255),
  `unit` VARCHAR(255),
  `category_id` CHAR(36),
  `stock` INT,
  `purchase_price` BIGINT,
  `selling_price` BIGINT,
  `description` VARCHAR(255),
  `photo` VARCHAR(255),
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (id),
  FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL
)