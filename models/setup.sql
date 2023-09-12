CREATE DATABASE IF NOT EXISTS chitchat DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;

DROP TABLE cc_tag
DROP TABLE cc_hot_list

CREATE TABLE `cc_tag`(
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '热搜标签id',
    `name` varchar(50) unique NOT NULL COMMENT '热搜标签名称',
    `sort` tinyint(3) DEFAULT NULL COMMENT '热搜标签排序',
    `source_key` varchar(50) unique NOT NULL COMMENT '热搜标签类型',
    `icon_color` varchar(50) DEFAULT NULL COMMENT '热搜标签图标颜色',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '热搜标签创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `name_UNIQUE` (`name`) 
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='热搜标签表';

CREATE TABLE `cc_hot_list`(
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '热搜清单id',
    `tag_id` int(11) NOT NULL COMMENT '热搜标签id',
    `title` varchar(200) NOT NULL COMMENT '热搜标题',
    `link` varchar(300) NOT NULL COMMENT '热搜地址',
    `extra` varchar(50) DEFAULT NULL COMMENT '额外信息',
    PRIMARY KEY (`id`),
    KEY `FK_tag` (`tag_id`), 
    CONSTRAINT `FK_tag` FOREIGN KEY (`tag_id`) REFERENCES `cc_tag` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='热搜清单表';