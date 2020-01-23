SET FOREIGN_KEY_CHECKS=0;
TRUNCATE TABLE `user_interest`;
TRUNCATE TABLE `interest`;
TRUNCATE TABLE `user`;
TRUNCATE TABLE `city`;
SET FOREIGN_KEY_CHECKS=1;

CALL generate_cities(1000);
CALL generate_interests();
CALL generate_users(100, 1000);