CREATE TABLE IF NOT EXISTS `email_verification`(
  `token` VARCHAR(8) PRIMARY KEY,
  `email` VARCHAR(256) UNIQUE NOT NULL,
  `expires_at` TIMESTAMP NOT NULL
);
