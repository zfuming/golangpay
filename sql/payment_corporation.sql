/*
 Navicat Premium Data Transfer

 Source Server         : db
 Source Server Type    : MySQL
 Source Server Version : 50718
 Source Host           : 10.254.254.254
 Source Database       : payment

 Target Server Type    : MySQL
 Target Server Version : 50718
 File Encoding         : utf-8

 Date: 05/25/2017 11:06:16 AM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `payment_corporation`
-- ----------------------------
DROP TABLE IF EXISTS `payment_corporation`;
CREATE TABLE `payment_corporation` (
  `corporation_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '法人ID',
  `corporation_code` varchar(32) NOT NULL DEFAULT '' COMMENT '法人代码',
  `corporation_name` varchar(32) NOT NULL DEFAULT '' COMMENT '法人名称',
  `pay_types` varchar(255) NOT NULL DEFAULT '' COMMENT '支付手段',
  `create_time` varchar(255) NOT NULL DEFAULT '' COMMENT '法人添加时间',
  `last_edit_time` varchar(255) NOT NULL DEFAULT '' COMMENT '法人最后编辑时间',
  PRIMARY KEY (`corporation_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='法人管理表';

SET FOREIGN_KEY_CHECKS = 1;
