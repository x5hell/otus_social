USE `social`;

CREATE TABLE IF NOT EXISTS city (
    `id` INT(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(80) NOT NULL
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS `user` (
    `id` INT(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `login` VARCHAR(40) NOT NULL,
    `password` VARCHAR(32) NOT NULL,
    `firstname` VARCHAR(50) NOT NULL,
    `lastname` VARCHAR(50) NOT NULL,
    `age` TINYINT NOT NULL,
    `sex` ENUM('male', 'female') NULL,
    `city_id` INT(11) UNSIGNED NULL,
    FOREIGN KEY (`city_id`) REFERENCES `city`(`id`)
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS interest (
    `id` INT(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(80) NOT NULL
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS user_interest (
    `user_id` INT(11) UNSIGNED NOT NULL,
    `interest_id` INT(11) UNSIGNED NOT NULL,
    PRIMARY KEY (`user_id`, `interest_id`),
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`),
    FOREIGN KEY (`interest_id`) REFERENCES `interest`(`id`)
) ENGINE=INNODB;