DROP DATABASE IF EXISTS `bee_test`;

CREATE DATABASE `bee_test`;

USE `bee_test`;

DROP TABLE IF EXISTS `t_user`;

CREATE TABLE `t_user`(
    `u_id` INTEGER AUTO_INCREMENT,
    `username` VARCHAR(50) NOT NULL,
    `password` VARCHAR(50) NOT NULL DEFAULT'1111',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`u_id`)
);

