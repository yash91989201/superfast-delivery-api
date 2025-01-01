CREATE TABLE `product` (
  `id` VARCHAR(36) PRIMARY KEY,
  `shop_id` VARCHAR(36) NOT NULL,
  `name` VARCHAR(256) NOT NULL,
  `description` TEXT,
  `product_type` ENUM('food', 'grocery', 'medicine') NOT NULL,
  `price` DECIMAL(10, 2) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP
);
