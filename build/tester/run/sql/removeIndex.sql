ALTER TABLE user_interest
    DROP FOREIGN KEY `user_interest_ibfk_2`,
    DROP FOREIGN KEY `user_interest_ibfk_1`;
ALTER TABLE user_interest
    DROP INDEX `interest_id`;
ALTER TABLE user_interest
    DROP PRIMARY KEY;
ALTER TABLE `user`
    DROP FOREIGN KEY `user_ibfk_1`;
ALTER TABLE `user`
    DROP INDEX `city_id`;
ALTER TABLE `user`
    CHANGE COLUMN `id` `id` INT(10) UNSIGNED NOT NULL,
    DROP PRIMARY KEY;
ALTER TABLE `interest`
    CHANGE COLUMN `id` `id` INT(10) UNSIGNED NOT NULL,
    DROP PRIMARY KEY;
ALTER TABLE `city`
    CHANGE COLUMN `id` `id` INT(10) UNSIGNED NOT NULL ,
    DROP PRIMARY KEY;
