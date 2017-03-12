
-- root 用户下
create user 'photoShare'@'localhost' identified by '123456';

create database photoShare;

grant all on photoShare.* to 'photoShare'@'localhost';



-- photoShare 用户下

-- Table structure for phs_users
-- ----------------------------
DROP TABLE IF EXISTS `phs_user`;
CREATE TABLE `phs_user` (
  `userid` bigint(20) NOT NULL,
  `profile_id` bigint(20) DEFAULT NULL,
  `username` varchar(15) DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) DEFAULT NULL COMMENT '密码',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态1正常未登录，2屏蔽,3正常已登录',
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
  `position` varchar(255) DEFAULT NULL COMMENT '职位id,多职位以|间隔',
  `lognum` int(10) DEFAULT '0' COMMENT '登录次数',
  `ip` varchar(15) DEFAULT NULL COMMENT '最近登录IP',
  `lasted` bigint(20) DEFAULT NULL COMMENT '最近登录时间',
  PRIMARY KEY(`userid`),
  KEY `INDEX_USP` (`realname`,`sex`,`lasted`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户详细信息表';


-- Table structure for phs_position
DROP TABLE IF EXISTS `phs_position`;
CREATE TABLE `phs_position`(
  `positionid` bigint(20) NOT NULL,
  `positionname` varchar(255) NOT NULL COMMENT '职位描述',
  `createtime` varchar(14) NOT NULL COMMENT '创建时间',
  PRIMARY KEY(`positionid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='职位描述表';

-- Table structure for phs_message
DROP TABLE IF EXISTS `phs_message`;
CREATE TABLE `phs_message`(
  `messageid` varchar(30) NOT NULL,
  `senderid` bigint(20) NOT NULL COMMENT '发送人id',
  `receiverid` bigint(20) NOT NULL COMMENT '接受人id',
  `createtime` varchar(14) NOT NULL COMMENT '信息创建时间',
  `content` varchar(300) COMMENT '信息内容',
  PRIMARY KEY(`messageId`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='信息表';


-- Table structure for phs_notice
DROP TABLE IF EXISTS `phs_notice`;
CREATE TABLE `phs_notice`(
  `noticeid` varchar(30) NOT NULL, 
  `receivertype` bigint(20) NOT NULL COMMENT '接受人种类:{1:所有用户...}',
  `createtime` varchar(14) NOT NULL COMMENT '公告创建时间',
  `status` tinyint(1) NOT NULL COMMENT '状态:0:置顶 1:正常 2:删除',
  `title` varchar(100) NOT NULL COMMENT '标题',
  `content` varchar(300) COMMENT '公告内容',
  PRIMARY KEY(`noticeid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='公告表';

