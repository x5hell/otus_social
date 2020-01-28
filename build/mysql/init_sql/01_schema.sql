CREATE TABLE IF NOT EXISTS city (
    `id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(80) NOT NULL
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS `user` (
    `id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `login` VARCHAR(40) NOT NULL,
    `password` VARCHAR(32) NOT NULL,
    `first_name` VARCHAR(50) NOT NULL,
    `last_name` VARCHAR(50) NOT NULL,
    `age` TINYINT NOT NULL,
    `sex` ENUM('male', 'female') NULL,
    `city_id` INT UNSIGNED NULL,
    FOREIGN KEY (`city_id`) REFERENCES `city`(`id`)
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS interest (
    `id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(80) NOT NULL
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS user_interest (
    `user_id` INT UNSIGNED NOT NULL,
    `interest_id` INT UNSIGNED NOT NULL,
    PRIMARY KEY (`user_id`, `interest_id`),
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`),
    FOREIGN KEY (`interest_id`) REFERENCES `interest`(`id`)
) ENGINE=INNODB;