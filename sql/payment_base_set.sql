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

 Date: 05/25/2017 11:06:04 AM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `payment_base_set`
-- ----------------------------
DROP TABLE IF EXISTS `payment_base_set`;
CREATE TABLE `payment_base_set` (
  `set_sign_use` int(11) NOT NULL DEFAULT '0' COMMENT '是否验证签名',
  `set_sign_key` varchar(255) NOT NULL DEFAULT '' COMMENT '签名密钥',
  `set_pay_state` varchar(255) NOT NULL DEFAULT '' COMMENT '支付状态',
  `set_currency` varchar(255) NOT NULL DEFAULT '' COMMENT '允许货币',
  `set_currency_info` varchar(255) NOT NULL DEFAULT '' COMMENT '货币信息'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='配置信息表';

SET FOREIGN_KEY_CHECKS = 1;
