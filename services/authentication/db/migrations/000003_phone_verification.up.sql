CREATE TABLE IF NOT EXISTS `phone_verification`(
  `token` VARCHAR(8) PRIMARY KEY,
  `phone` VARCHAR(10) UNIQUE NOT NULL,
  `expires_at` TIMESTAMP NOT NULL
);

