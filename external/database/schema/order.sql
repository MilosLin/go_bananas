CREATE TABLE IF NOT EXISTS `test_db`.`order` (
  `o_id`          bigint(20) AUTO_INCREMENT NOT NULL COMMENT '流水號',
  `user_id`       varchar(20) NOT NULL COMMENT '使用者id',
  `order_time`    timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '訂單時間',
  `money`         decimal(14,4) NOT NULL DEFAULT '0' COMMENT '訂單金額',
  `remark`        text COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '備註',
  `update_time`   TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最近一次異動時間',
  `create_time`   TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '資料建立時間',
  PRIMARY KEY (`o_id`,`order_time`)
) ENGINE = InnoDB
  CHARACTER SET = utf8
  COLLATE = utf8_unicode_ci;

ALTER TABLE `test_db`.`order` ADD INDEX `user_id` (`user_id`);
