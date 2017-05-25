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

 Date: 05/25/2017 11:06:29 AM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `payment_pay_type`
-- ----------------------------
DROP TABLE IF EXISTS `payment_pay_type`;
CREATE TABLE `payment_pay_type` (
  `pay_type_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '支付渠道ID',
  `pay_type_name` varchar(255) NOT NULL DEFAULT '' COMMENT '支付渠道名称',
  `pay_type_code` varchar(255) NOT NULL DEFAULT '' COMMENT '支付渠道代码',
  `pay_type_style` varchar(255) NOT NULL DEFAULT '' COMMENT '支付方式列表',
  `create_time` varchar(255) NOT NULL DEFAULT '' COMMENT '添加时间',
  PRIMARY KEY (`pay_type_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='支付手渠道';

-- ----------------------------
--  Records of `payment_pay_type`
-- ----------------------------
BEGIN;
INSERT INTO `payment_pay_type` VALUES ('1', '支付宝', 'alipay', '{\"alipay_trade_pay\":\"支付宝条码支付\",\"alipay_trade_precreate\":\"支付宝扫码支付\",\"alipay_trade_wap_pay\":\"手机网站支付\"}', '1495681443'), ('2', '银联支付', 'unipay', '{\"codepay\":\"银联刷卡支付\",\"qrpay\":\"银联扫码支付\"}', '1495681444'), ('3', '微信支付', 'wxpay', '{\"MICROPAY\":\"微信刷卡支付\",\"NATIVE\":\"微信扫码支付\",\"APP\":\"微信移动支付\",\"JSAPI\":\"微信公众账号支付\"}', '1495681445');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
