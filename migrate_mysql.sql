CREATE DATABASE foobar_m0cchi____duplicate_guard;
USE  foobar_m0cchi____duplicate_guard;
CREATE TABLE `users` (
  `id` INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `name` VARCHAR(60)
);
INSERT INTO `users` (`name`) VALUES ("foo");
INSERT INTO `users` (`name`) VALUES ("bar");
