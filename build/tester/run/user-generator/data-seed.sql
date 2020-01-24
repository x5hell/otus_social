SET @city_quantity = 10000;
SET @users_quantity = 5000;
SET @from_age = 16;
SET @to_age = 60;
SET @max_user_interrsts = 10;

SET SQL_MODE='NO_AUTO_VALUE_ON_ZERO';
SET FOREIGN_KEY_CHECKS=0;
SET UNIQUE_CHECKS=0;
SET AUTOCOMMIT=0;

TRUNCATE TABLE `user_interest`;
TRUNCATE TABLE `interest`;
TRUNCATE TABLE `user`;
TRUNCATE TABLE `city`;

CALL generate_cities(@city_quantity);
CALL generate_interests();
CALL generate_users(@users_quantity, @from_age, @to_age, @city_quantity);
CALL generate_users_interests(@users_quantity, @max_user_interrsts);

SET FOREIGN_KEY_CHECKS=1;
SET UNIQUE_CHECKS=1;
COMMIT;