CREATE TABLE IF NOT EXISTS `announcement_pages_old_version` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `page_name` varchar(100) DEFAULT NULL,
  `can_close` tinyint(1) DEFAULT '0',
  `type` varchar(100) DEFAULT NULL,
  `start_date` datetime DEFAULT NULL,
  `end_date` datetime DEFAULT NULL,
  `title` varchar(100) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `priority` int DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=127 DEFAULT CHARSET=utf8mb3;

SELECT * FROM announcement_pages_old_version ap;

SELECT count(1) FROM announcement_pages_old_version ap;

DELETE FROM announcement_pages_old_version;