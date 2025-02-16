CREATE TABLE `rewards` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `brand` longtext,
  `currency` longtext,
  `denomination` float DEFAULT NULL,
  PRIMARY KEY (`id`)
);