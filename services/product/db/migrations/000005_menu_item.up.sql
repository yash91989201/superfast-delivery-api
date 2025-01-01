CREATE TABLE `menu_item` (
  `id` VARCHAR(36) PRIMARY KEY,
  `position` INT DEFAULT 0,
  `menu_id` VARCHAR(36) NOT NULL,
  `product_id` VARCHAR(36) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT fk_menu_item_menu FOREIGN KEY (`menu_id`) REFERENCES `restaurant_menu`(id) ON DELETE CASCADE,
  CONSTRAINT fk_menu_item_product FOREIGN KEY (`product_id`) REFERENCES `product`(id) ON DELETE CASCADE
);
