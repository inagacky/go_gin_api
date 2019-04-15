DROP DATABASE IF EXISTS sample;
CREATE DATABASE sample;
USE sample;

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,     /* ID */
  `first_name` varchar(255) DEFAULT NULL,        /* 名 */
  `last_name` varchar(255) DEFAULT NULL,         /* 姓 */
  `email` varchar(255) NOT NULL,                 /* メールアドレス */
  `status` tinyint(4) NOT NULL,                  /* ステータス */
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8
;

INSERT INTO user (first_name, last_name, email, status) VALUES ("テスト", "太郎", "test@gmail.com", 1);
INSERT INTO user (first_name, last_name, email, status) VALUES ("テスト", "二郎", "test2@gmail.com", 1);
INSERT INTO user (first_name, last_name, email, status) VALUES ("テスト", "三郎", "test3@gmail.com", 1)
