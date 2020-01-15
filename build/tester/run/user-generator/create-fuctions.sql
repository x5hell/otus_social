DELIMITER //

CREATE FUNCTION sex_generator () RETURNS VARCHAR(6)
BEGIN
    DECLARE num INT;
    DECLARE result VARCHAR(6);
    SELECT CEIL(RAND()*3) INTO num;
    CASE num
        WHEN 1 THEN SET result = 'male';
        WHEN 2 THEN SET result = 'female';
        ELSE SET result = null;
    END CASE;
END
