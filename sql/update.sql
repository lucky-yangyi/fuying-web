DROP TABLE IF EXISTS `jm_website_news`;
CREATE TABLE `jm_website_news` (
                                   `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增长ID',
                                   `type` tinyint(1) NOT NULL COMMENT '（1:家长故事 2：孩子的故事 3：扶小鹰成长故事 4：公司新闻 5：行业新闻 6：亲子教育）',
                                   `image_url` varchar(255) DEFAULT NULL COMMENT '图片url',
                                   `video_url` varchar(255) DEFAULT NULL COMMENT '视频url',
                                   `title` varchar(255) DEFAULT NULL COMMENT '标题',
                                   `source` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '来源',
                                   `adit` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '责任编辑',
                                   `desc` varchar(255) DEFAULT NULL COMMENT '描述',
                                   `content` text CHARACTER SET utf8mb4 COMMENT '详情',
                                   `status` tinyint(4) DEFAULT '0' COMMENT '状态（-1；删除）',
                                   `create_id` int(11) unsigned NOT NULL COMMENT '创建者',
                                   `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
                                   `update_id` int(11) unsigned NOT NULL COMMENT '修改者',
                                   `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
                                   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='jm_user_setting（扶鹰官网首页配置）';



DROP TABLE IF EXISTS `jm_parent_perception`;
CREATE TABLE `jm_parent_perception` (
                                   `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增长ID',
                                   `image_url` varchar(255) DEFAULT NULL COMMENT '图片url',
                                   `nick_name` varchar(255) DEFAULT NULL COMMENT '标题',
                                   `content` text CHARACTER SET utf8mb4 COMMENT '详情',
                                   `status` tinyint(4) DEFAULT '0' COMMENT '状态（-1；删除）',
                                   `publish_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
                                   `create_id` int(11) unsigned NOT NULL COMMENT '创建者',
                                   `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
                                   `update_id` int(11) unsigned NOT NULL COMMENT '修改者',
                                   `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
                                   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='jm_parent_perception（家长感悟）';
