CREATE TABLE `tk_user` (
     `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
     `account` varchar(40) NOT NULL,
     `password` varchar(80) NOT NULL,
     `created_at` datetime NOT NULL,
     `updated_at` datetime DEFAULT NULL,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';