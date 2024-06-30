CREATE TABLE `pb_file_table` (
                        `id` int(11) unsigned NOT NULL AUTO_INCREMENT,


                        `project_name` varchar(60) NOT NULL,
                        `service_name` varchar(32) NOT NULL,
                        `branch` varchar(100) NOT NULL,
                        `file_name` varchar(100) NOT NULL,
                        `file_id`  int(11) NOT NULL,
                        `sign`  varchar(255) NOT NULL,
                        `creator` varchar(100) NOT NULL,

                        `created_at` timestamp NOT NULL,
                        `updated_at` timestamp NOT NULL,
                        `deleted_at` timestamp DEFAULT NULL,
                        PRIMARY KEY (`id`),
                        INDEX project_service_idx (`project_name`,`service_name`),
                        INDEX file_id_idx (`file_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
