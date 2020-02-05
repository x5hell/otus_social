ALTER TABLE `city`
    CHANGE COLUMN `id` `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    ADD PRIMARY KEY (`id`);

ALTER TABLE `interest`
    CHANGE COLUMN `id` `id` INT UNSIGNED NOT NULL AUTO_INCREMENT ,
    ADD PRIMARY KEY (`id`);

ALTER TABLE `user`
    CHANGE COLUMN `id` `id` INT UNSIGNED NOT NULL AUTO_INCREMENT ,
    ADD PRIMARY KEY (`id`);

ALTER TABLE `user_interest`
    ADD PRIMARY KEY (`user_id`, `interest_id`);

ALTER TABLE `user_interest`
ADD INDEX `user_interest_interest_id_idx` (`interest_id` ASC) VISIBLE;

ALTER TABLE `user_interest`
ADD CONSTRAINT `user_interest_user_id`
  FOREIGN KEY (`user_id`)
  REFERENCES `user`(`id`)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION,
ADD CONSTRAINT `user_interest_interest_id`
  FOREIGN KEY (`interest_id`)
  REFERENCES `interest`(`id`)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION;

ALTER TABLE `user`
ADD INDEX `user_city_id_idx` (`city_id` ASC) VISIBLE;
ALTER TABLE `user`
ADD CONSTRAINT `user_city_id`
  FOREIGN KEY (`city_id`)
  REFERENCES `city`(`id`)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION;