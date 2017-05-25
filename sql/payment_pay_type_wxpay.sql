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

 Date: 05/25/2017 11:06:44 AM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `payment_pay_type_wxpay`
-- ----------------------------
DROP TABLE IF EXISTS `payment_pay_type_wxpay`;
CREATE TABLE `payment_pay_type_wxpay` (
  `pay_type_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '支付手段ID',
  `pay_type_title` varchar(255) NOT NULL DEFAULT '' COMMENT '方式标题',
  `wxpay_gateway_url` varchar(255) NOT NULL DEFAULT '' COMMENT '微信请求网关',
  `wxpay_app_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'APP ID',
  `wxpay_mch_id` varchar(255) NOT NULL DEFAULT '' COMMENT '微信商户号',
  `wxpay_key` varchar(255) NOT NULL DEFAULT '' COMMENT 'APP KEY',
  `wxpay_app_secret` varchar(255) NOT NULL DEFAULT '' COMMENT 'APP SECRET',
  `wxpay_allow_style` varchar(255) NOT NULL DEFAULT '' COMMENT '支持的支付形式',
  `wxpay_client_cert` varchar(255) NOT NULL DEFAULT '' COMMENT '微信CERT证书',
  `wxpay_client_key` varchar(255) NOT NULL DEFAULT '' COMMENT '微信KEY证书',
  `create_time` varchar(32) NOT NULL DEFAULT '' COMMENT '支付手段添加时间',
  `last_edit_time` varchar(32) NOT NULL DEFAULT '' COMMENT '支付手段最后编辑时间',
  PRIMARY KEY (`pay_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='微信信息表';

SET FOREIGN_KEY_CHECKS = 1;
