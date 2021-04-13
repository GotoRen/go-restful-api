CREATE DATABASE IF NOT EXISTS `sample_db`;

CREATE TABLE IF NOT EXISTS `sample_db`.`users` (
  `id`   INT          NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(256) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 

INSERT INTO users(name) VALUE('金村美玖');
INSERT INTO users(name) VALUE('河田陽菜');
INSERT INTO users(name) VALUE('小坂菜緒');
INSERT INTO users(name) VALUE('富田鈴花');
INSERT INTO users(name) VALUE('丹生明里');
INSERT INTO users(name) VALUE('濱岸ひより');
INSERT INTO users(name) VALUE('松田好花');
INSERT INTO users(name) VALUE('宮田愛萌');
INSERT INTO users(name) VALUE('渡邊美穂');
