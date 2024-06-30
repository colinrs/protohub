CREATE TABLE `user_table` (
                              `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                              `identity` varchar(36) NOT NULL,

                              `name` varchar(60) NOT NULL,
                              `password` varchar(32) NOT NULL,
                              `email` varchar(100)  NOT NULL,

                              `created_at` timestamp NOT NULL,
                              `updated_at` timestamp NOT NULL,
                              `deleted_at` timestamp DEFAULT NULL,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
