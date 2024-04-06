CREATE TABLE `products` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `code` VARCHAR(255),
  `name` VARCHAR(255),
  `unit` VARCHAR(255),
  `category_id` BIGINT,
  `stock` INT,
  `purchase_price` BIGINT,
  `selling_price` BIGINT,
  `description` VARCHAR(255),
  `photo` VARCHAR(255),
  PRIMARY KEY (id)
)