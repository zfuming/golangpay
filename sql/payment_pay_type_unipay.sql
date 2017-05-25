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

 Date: 05/25/2017 11:06:39 AM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `payment_pay_type_unipay`
-- ----------------------------
DROP TABLE IF EXISTS `payment_pay_type_unipay`;
CREATE TABLE `payment_pay_type_unipay` (
  `pay_type_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '支付手段ID',
  `pay_type_title` varchar(255) NOT NULL DEFAULT '' COMMENT '方式标题',
  `unipay_allow_style` varchar(255) NOT NULL DEFAULT '' COMMENT '支持的支付形式',
  `unipay_store_id` varchar(255) NOT NULL DEFAULT '' COMMENT '门店ID号',
  `unipay_gateway_url` varchar(255) NOT NULL DEFAULT '' COMMENT '银联请求网关',
  `unipay_back_url` varchar(255) NOT NULL DEFAULT '' COMMENT '返回url',
  `unipay_biz_type` varchar(255) NOT NULL DEFAULT '' COMMENT '业务类型',
  `unipay_mer_id` varchar(255) NOT NULL DEFAULT '' COMMENT '商户号',
  `unipay_sign_cert_pwd` varchar(255) NOT NULL DEFAULT '' COMMENT '签名证书密码',
  `unipay_sign_cert` varchar(255) NOT NULL DEFAULT '' COMMENT '银联CERT证书',
  `unipay_encrypt_cert` varchar(255) NOT NULL DEFAULT '' COMMENT '敏感信息加密证书路径',
  `unipay_middle_cert` varchar(255) NOT NULL DEFAULT '' COMMENT '验签中级证书',
  `unipay_root_cert` varchar(255) NOT NULL DEFAULT '' COMMENT '验签根证书',
  `create_time` varchar(32) NOT NULL DEFAULT '' COMMENT '支付手段添加时间',
  `last_edit_time` varchar(32) NOT NULL DEFAULT '' COMMENT '支付手段最后编辑时间',
  PRIMARY KEY (`pay_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='银联信息表';

SET FOREIGN_KEY_CHECKS = 1;
