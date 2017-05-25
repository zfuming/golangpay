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

 Date: 05/25/2017 11:06:23 AM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `payment_pay_order`
-- ----------------------------
DROP TABLE IF EXISTS `payment_pay_order`;
CREATE TABLE `payment_pay_order` (
  `pay_order_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '订单ID',
  `corporation_code` varchar(32) NOT NULL DEFAULT '' COMMENT '法人代号',
  `pay_type_code` varchar(32) NOT NULL DEFAULT '' COMMENT '支付手段代号',
  `pay_type_style` varchar(32) NOT NULL DEFAULT '' COMMENT '支付方式代号',
  `pay_currency` varchar(32) NOT NULL DEFAULT '' COMMENT '支付货币',
  `total_amount` varchar(32) NOT NULL DEFAULT '' COMMENT '订单总金额',
  `pay_amount` varchar(32) NOT NULL DEFAULT '' COMMENT '支付金额',
  `trade_goods_list` varchar(255) NOT NULL DEFAULT '' COMMENT '商品列表',
  `discount_amount` varchar(32) NOT NULL DEFAULT '' COMMENT '优惠金额',
  `discount_list` varchar(255) NOT NULL DEFAULT '' COMMENT '优惠详情',
  `pay_buyer_user` varchar(32) NOT NULL DEFAULT '' COMMENT '支付账户',
  `pay_sale_user` varchar(32) NOT NULL DEFAULT '' COMMENT '收款账户',
  `pay_state` int(11) NOT NULL DEFAULT '0' COMMENT '支付状态 0未支付，1已支付，2已退款，3已完成',
  `merchant_no` varchar(32) NOT NULL DEFAULT '' COMMENT '订单号',
  `tp_trade_no` varchar(32) NOT NULL DEFAULT '' COMMENT '第三方流水号',
  `tp_refund_no` varchar(32) NOT NULL DEFAULT '' COMMENT '第三方退款号',
  `refund_reason` varchar(255) NOT NULL DEFAULT '' COMMENT '退款原因',
  `create_time` varchar(32) NOT NULL DEFAULT '' COMMENT '订单创建时间',
  `payed_time` varchar(32) NOT NULL DEFAULT '' COMMENT '支付时间',
  `refund_time` varchar(32) NOT NULL DEFAULT '' COMMENT '退款时间',
  PRIMARY KEY (`pay_order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='支付订单表';

SET FOREIGN_KEY_CHECKS = 1;
