CREATE TABLE `check_accounts`
(
    `list`         varchar(30)                  NOT NULL  COMMENT '流水',
    `money1`       decimal(8,2)                 NULL  COMMENT '本地金额',
    `money2`       decimal(8,2)                 NULL  COMMENT '第三方金额',
    `state1`       int(1) unsigned              NULL  COMMENT '本地状态',
    `state2`       int(1) unsigned              NULL  COMMENT '第三方状态',
    `created_time` timestamp(3)                 NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`list`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='对账表';