CREATE TABLE `file_content_table` (
                        `id` int(11) unsigned NOT NULL AUTO_INCREMENT,

                        `file_content` mediumtext default null,
                        `creator` varchar(100) NOT NULL,

                        `created_at` timestamp NOT NULL,
                        `updated_at` timestamp NOT NULL,
                        `deleted_at` timestamp DEFAULT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
