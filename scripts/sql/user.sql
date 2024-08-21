CREATE TABLE `user_table` (
                              `id` int(11) unsigned NOT NULL AUTO_INCREMENT,

                              `user_name` varchar(60) NOT NULL,
                              `password` varchar(32) NOT NULL,
                              `email` varchar(100)  NOT NULL,
                              `mobile` varchar(100)  NOT NULL,
                              `user_status`  int(11) unsigned default 1,
                              `description` varchar(500)  default '',
                              `created_at` timestamp default CURRENT_TIMESTAMP,
                              `updated_at` timestamp default CURRENT_TIMESTAMP,
                              `deleted_at` timestamp DEFAULT NULL,

                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;



CREATE TABLE `user_roles_table` (
                              `id` int(11) unsigned NOT NULL AUTO_INCREMENT,

                              `role_code` varchar(24) NOT NULL,
                              `created_at` timestamp default CURRENT_TIMESTAMP,
                              `updated_at` timestamp default CURRENT_TIMESTAMP,
                              `deleted_at` timestamp DEFAULT NULL,

                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;


CREATE TABLE `user_project_table` (
                                    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,

                                    `project_id`  int(11) unsigned NOT NULL,
                                    `created_at` timestamp default CURRENT_TIMESTAMP,
                                    `updated_at` timestamp default CURRENT_TIMESTAMP,
                                    `deleted_at` timestamp DEFAULT NULL,

                                    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;



CREATE TABLE `user_token_table` (
                                      `id` int(11) unsigned NOT NULL AUTO_INCREMENT,

                                      `token` varchar(128) NOT NULL,
                                      `expired_at` int(11) unsigned NOT NULL,
                                      `created_at` timestamp default CURRENT_TIMESTAMP,
                                      `updated_at` timestamp default CURRENT_TIMESTAMP,
                                      `deleted_at` timestamp DEFAULT NULL,
                                      PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;