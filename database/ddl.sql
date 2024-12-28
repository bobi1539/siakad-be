-- siakad.m_parameter definition

USE siakad;

CREATE TABLE `m_parameter` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- siakad.m_parameter_list definition

CREATE TABLE `m_parameter_list` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `parameter_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `parameter_list_parameter_id_foreign` (`parameter_id`),
  CONSTRAINT `parameter_list_parameter_id_foreign` FOREIGN KEY (`parameter_id`) REFERENCES `m_parameter` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;