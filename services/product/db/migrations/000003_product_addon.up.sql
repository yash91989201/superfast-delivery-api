CREATE TABLE `product_addon` (
  `id` VARCHAR(36) PRIMARY KEY,
  `addon_name` VARCHAR(128),
  `addon_price` DECIMAL(10, 2),
  `product_id` VARCHAR(36) NOT NULL,
  CONSTRAINT fk_product_addon_product FOREIGN KEY (`product_id`) REFERENCES `product`(id) ON DELETE CASCADE
);
