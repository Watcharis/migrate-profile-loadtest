CREATE TABLE `announcement_users_old_version` (
  `uuid` varchar(255) NOT NULL,
  `announcement_page_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`uuid`,`announcement_page_id`),
  KEY `announcement_users_old_version_FK` (`announcement_page_id`),
  CONSTRAINT `announcement_users_old_version_FK` FOREIGN KEY (`announcement_page_id`) REFERENCES `announcement_pages_old_version` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `announcement_users_old_version_FK_1` FOREIGN KEY (`uuid`) REFERENCES `profile_announcements` (`uuid`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

SELECT count(1) FROm announcement_users_old_version ap;

DELETE FROM announcement_users_old_version;