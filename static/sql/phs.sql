
-- root 用户下
create user 'photoShare'@'localhost' identified by '123456';

create database photoShare;

grant all on photoShare.* to 'photoShare'@'localhost'



-- photoShare 用户下

-- Table structure for phs_users
-- ----------------------------
DROP TABLE IF EXISTS `phs_user`;
CREATE TABLE `phs_user` (
  `userid` bigint(20) NOT NULL,
  `profile_id` bigint(20) DEFAULT NULL,
  `username` varchar(15) DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) DEFAULT NULL COMMENT '密码',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态1正常，2屏蔽',
  PRIMARY KEY (`userid`),
  KEY `INDEX_US` (`username`,`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户主表';



-- Table structure for phs_user_profile
DROP TABLE IF EXISTS `phs_user_profile`;
CREATE TABLE `phs_user_profile` (
  `userid` bigint(20) NOT NULL,
  `realname` varchar(15) DEFAULT NULL COMMENT '姓名',
  `sex` tinyint(1) DEFAULT '1' COMMENT '1男2女',
  `birth` varchar(15) DEFAULT NULL,
  `email` varchar(30) DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(15) DEFAULT NULL COMMENT '手机',
  `address` varchar(100) DEFAULT NULL COMMENT '地址',
  `positionid` bigint(20) DEFAULT NULL COMMENT '职位id',
  `lognum` int(10) DEFAULT '0' COMMENT '登录次数',
  `ip` varchar(15) DEFAULT NULL COMMENT '最近登录IP',
  `lasted` int(10) DEFAULT NULL COMMENT '最近登录时间',
  PRIMARY KEY(`userid`),
  KEY `INDEX_USP` (`realname`,`sex`,`lasted`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户详细信息表';




