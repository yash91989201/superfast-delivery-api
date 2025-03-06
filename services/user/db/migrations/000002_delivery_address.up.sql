CREATE TABLE IF NOT EXISTS `delivery_address`(
  `id` VARCHAR(36) PRIMARY KEY,
  `receiver_name` VARCHAR(256) NOT NULL,
  `receiver_phone` VARCHAR(10) NOT NULL,
  `address_alias` ENUM("home", "work", "hotel", "other") DEFAULT("home") NOT NULL,
  `other_alias` VARCHAR(36) DEFAULT NULL,
  `longitude` DOUBLE NOT NULL,
  `latitude` DOUBLE NOT NULL,
  `location` POINT NOT NULL,
  `address` VARCHAR(512) NOT NULL,
  `nearby_landmark` VARCHAR(256) DEFAULT NULL,
  `delivery_instruction` VARCHAR(256) DEFAULT NULL,
  `auth_id` VARCHAR(36) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL
);

-- Trigger to update location on insert
CREATE TRIGGER insert_location
BEFORE INSERT ON `delivery_address` 
FOR EACH ROW
BEGIN
  SET NEW.location = ST_GeomFromText(CONCAT('POINT(', NEW.longitude, ' ', NEW.latitude, ')'));
END;

-- Trigger to update location on update
CREATE TRIGGER update_location
BEFORE UPDATE ON `delivery_address` 
FOR EACH ROW
BEGIN
  SET NEW.location = ST_GeomFromText(CONCAT('POINT(', NEW.longitude, ' ', NEW.latitude, ')'));
END;

