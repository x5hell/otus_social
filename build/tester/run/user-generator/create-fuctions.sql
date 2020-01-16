DELIMITER $$

DROP FUNCTION IF EXISTS num_to_sex;
CREATE FUNCTION num_to_sex (num INT) RETURNS VARCHAR(6) 
DETERMINISTIC
BEGIN
    DECLARE result VARCHAR(6);
    CASE num
        WHEN 1 THEN SET result = 'male';
        WHEN 2 THEN SET result = 'female';
    END CASE;
    RETURN result;
END$$

DROP FUNCTION IF EXISTS sex_generator;
CREATE FUNCTION sex_generator () RETURNS VARCHAR(6)
NOT DETERMINISTIC
BEGIN
    DECLARE num INT;
    DECLARE result VARCHAR(6);
    SELECT CEIL(RAND()*3) INTO num;
    CASE num
        WHEN 3 THEN SET result = null;
        ELSE SET result = num_to_sex(num);
    END CASE;
    RETURN result;
END$$

DROP FUNCTION IF EXISTS male_lastname;
CREATE FUNCTION male_lastname () RETURNS VARCHAR(25)
NOT DETERMINISTIC
BEGIN
    DECLARE num INT;
    DECLARE result VARCHAR(50);
    SELECT CEIL(RAND()*100) INTO num;
    CASE num
        WHEN 1 THEN SET result = 'Пономаренко';
        WHEN 2 THEN SET result = 'Красинец';
        WHEN 3 THEN SET result = 'Фёдоров';
        WHEN 4 THEN SET result = 'Пахомов';
        WHEN 5 THEN SET result = 'Петрив';
        WHEN 6 THEN SET result = 'Осипов';
        WHEN 7 THEN SET result = 'Щербак';
        WHEN 8 THEN SET result = 'Кудряшов';
        WHEN 9 THEN SET result = 'Гриневская';
        WHEN 10 THEN SET result = 'Рожков';
        WHEN 11 THEN SET result = 'Негода';
        WHEN 12 THEN SET result = 'Гелетей';
        WHEN 13 THEN SET result = 'Петровский';
        WHEN 14 THEN SET result = 'Туров';
        WHEN 15 THEN SET result = 'Бачей';
        WHEN 16 THEN SET result = 'Кулагин';
        WHEN 17 THEN SET result = 'Саксаганский';
        WHEN 18 THEN SET result = 'Лукин';
        WHEN 19 THEN SET result = 'Елисеев';
        WHEN 20 THEN SET result = 'Токар';
        WHEN 21 THEN SET result = 'Котов';
        WHEN 22 THEN SET result = 'Терентьев';
        WHEN 23 THEN SET result = 'Щукин';
        WHEN 24 THEN SET result = 'Дзюба';
        WHEN 25 THEN SET result = 'Мирный';
        WHEN 26 THEN SET result = 'Передрий';
        WHEN 27 THEN SET result = 'Миронов';
        WHEN 28 THEN SET result = 'Лукин';
        WHEN 29 THEN SET result = 'Ерёменко';
        WHEN 30 THEN SET result = 'Желиба';
        WHEN 31 THEN SET result = 'Журавлёв';
        WHEN 32 THEN SET result = 'Одинцов';
        WHEN 33 THEN SET result = 'Полищук';
        WHEN 34 THEN SET result = 'Забужко';
        WHEN 35 THEN SET result = 'Яворивский';
        WHEN 36 THEN SET result = 'Гордеев';
        WHEN 37 THEN SET result = 'Линник';
        WHEN 38 THEN SET result = 'Скоропадский';
        WHEN 39 THEN SET result = 'Третьяков';
        WHEN 40 THEN SET result = 'Рябов';
        WHEN 41 THEN SET result = 'Мартынов';
        WHEN 42 THEN SET result = 'Буров';
        WHEN 43 THEN SET result = 'Сысоев';
        WHEN 44 THEN SET result = 'Зимин';
        WHEN 45 THEN SET result = 'Ильин';
        WHEN 46 THEN SET result = 'Таранец';
        WHEN 47 THEN SET result = 'Данилов';
        WHEN 48 THEN SET result = 'Павлов';
        WHEN 49 THEN SET result = 'Пархоменко';
        WHEN 50 THEN SET result = 'Моисеев';
        WHEN 51 THEN SET result = 'Фадеев';
        WHEN 52 THEN SET result = 'Кузнецов';
        WHEN 53 THEN SET result = 'Жуков';
        WHEN 54 THEN SET result = 'Лазарев';
        WHEN 55 THEN SET result = 'Миклашевский';
        WHEN 56 THEN SET result = 'Лыткин';
        WHEN 57 THEN SET result = 'Котовский';
        WHEN 58 THEN SET result = 'Тихонов';
        WHEN 59 THEN SET result = 'Беляков';
        WHEN 60 THEN SET result = 'Галкин';
        WHEN 61 THEN SET result = 'Гончар';
        WHEN 62 THEN SET result = 'Яловой';
        WHEN 63 THEN SET result = 'Исаев';
        WHEN 64 THEN SET result = 'Петрик';
        WHEN 65 THEN SET result = 'Селиверстов';
        WHEN 66 THEN SET result = 'Гордеев';
        WHEN 67 THEN SET result = 'Пестов';
        WHEN 68 THEN SET result = 'Ткаченко';
        WHEN 69 THEN SET result = 'Темченко';
        WHEN 70 THEN SET result = 'Бондаренко';
        WHEN 71 THEN SET result = 'Савельев';
        WHEN 72 THEN SET result = 'Козлов';
        WHEN 73 THEN SET result = 'Мясников';
        WHEN 74 THEN SET result = 'Трублаевский';
        WHEN 75 THEN SET result = 'Королёв';
        WHEN 76 THEN SET result = 'Денисов';
        WHEN 77 THEN SET result = 'Цветков';
        WHEN 78 THEN SET result = 'Шкраба';
        WHEN 79 THEN SET result = 'Павленко';
        WHEN 80 THEN SET result = 'Борисенко';
        WHEN 81 THEN SET result = 'Иващенко';
        WHEN 82 THEN SET result = 'Гелетей';
        WHEN 83 THEN SET result = 'Павлив';
        WHEN 84 THEN SET result = 'Сорокин';
        WHEN 85 THEN SET result = 'Масловский';
        WHEN 86 THEN SET result = 'Кулибаба';
        WHEN 87 THEN SET result = 'Мишин';
        WHEN 88 THEN SET result = 'Терещенко';
        WHEN 89 THEN SET result = 'Палий';
        WHEN 90 THEN SET result = 'Колобов';
        WHEN 91 THEN SET result = 'Трясило';
        WHEN 92 THEN SET result = 'Майборода';
        WHEN 93 THEN SET result = 'Иващенко';
        WHEN 94 THEN SET result = 'Юдин';
        WHEN 95 THEN SET result = 'Рыбаков';
        WHEN 96 THEN SET result = 'Рогов';
        WHEN 97 THEN SET result = 'Кулишенко';
        WHEN 98 THEN SET result = 'Казаков';
        WHEN 99 THEN SET result = 'Елисеев';
        WHEN 100 THEN SET result = 'Евсеев';
    END CASE;
    RETURN result;
END $$

DROP FUNCTION IF EXISTS male_firstname;
CREATE FUNCTION male_firstname() RETURNS VARCHAR(25)
NOT DETERMINISTIC
BEGIN
    DECLARE num INT;
    DECLARE result VARCHAR(50);
    SELECT CEIL(RAND()*100) INTO num;
    CASE num
        WHEN 1 THEN SET result = 'Платон';
        WHEN 2 THEN SET result = 'Устин';
        WHEN 3 THEN SET result = 'Тимур';
        WHEN 4 THEN SET result = 'Михаил';
        WHEN 5 THEN SET result = 'Глеб';
        WHEN 6 THEN SET result = 'Шарль';
        WHEN 7 THEN SET result = 'Устин';
        WHEN 8 THEN SET result = 'Еремей';
        WHEN 9 THEN SET result = 'Эдуард';
        WHEN 10 THEN SET result = 'Герасим';
        WHEN 11 THEN SET result = 'Ананий';
        WHEN 12 THEN SET result = 'Жерар';
        WHEN 13 THEN SET result = 'Владлен';
        WHEN 14 THEN SET result = 'Павел';
        WHEN 15 THEN SET result = 'Ефим';
        WHEN 16 THEN SET result = 'Пётр';
        WHEN 17 THEN SET result = 'Устин';
        WHEN 18 THEN SET result = 'Чарльз';
        WHEN 19 THEN SET result = 'Остап';
        WHEN 20 THEN SET result = 'Конрад';
        WHEN 21 THEN SET result = 'Ян';
        WHEN 22 THEN SET result = 'Шерлок';
        WHEN 23 THEN SET result = 'Любомир';
        WHEN 24 THEN SET result = 'Артур';
        WHEN 25 THEN SET result = 'Гавриил';
        WHEN 26 THEN SET result = 'Валерий';
        WHEN 27 THEN SET result = 'Ярослав';
        WHEN 28 THEN SET result = 'Чарльз';
        WHEN 29 THEN SET result = 'Жигер';
        WHEN 30 THEN SET result = 'Жерар';
        WHEN 31 THEN SET result = 'Устин';
        WHEN 32 THEN SET result = 'Богдан';
        WHEN 33 THEN SET result = 'Харитон';
        WHEN 34 THEN SET result = 'Марк';
        WHEN 35 THEN SET result = 'Филипп';
        WHEN 36 THEN SET result = 'Георгий';
        WHEN 37 THEN SET result = 'Иван';
        WHEN 38 THEN SET result = 'Юлиан';
        WHEN 39 THEN SET result = 'Харитон';
        WHEN 40 THEN SET result = 'Тимур';
        WHEN 41 THEN SET result = 'Шерлок';
        WHEN 42 THEN SET result = 'Тит';
        WHEN 43 THEN SET result = 'Матвей';
        WHEN 44 THEN SET result = 'Кристиан';
        WHEN 45 THEN SET result = 'Ярослав';
        WHEN 46 THEN SET result = 'Донат';
        WHEN 47 THEN SET result = 'Цицерон';
        WHEN 48 THEN SET result = 'Динар';
        WHEN 49 THEN SET result = 'Цезарь';
        WHEN 50 THEN SET result = 'Марк';
        WHEN 51 THEN SET result = 'Шарль';
        WHEN 52 THEN SET result = 'Устин';
        WHEN 53 THEN SET result = 'Кузьма';
        WHEN 54 THEN SET result = 'Рафаил';
        WHEN 55 THEN SET result = 'Ефрем';
        WHEN 56 THEN SET result = 'Йоган';
        WHEN 57 THEN SET result = 'Алексей';
        WHEN 58 THEN SET result = 'Роберт';
        WHEN 59 THEN SET result = 'Афанасий';
        WHEN 60 THEN SET result = 'Мирослав';
        WHEN 61 THEN SET result = 'Нестор';
        WHEN 62 THEN SET result = 'Святослав';
        WHEN 63 THEN SET result = 'Казбек';
        WHEN 64 THEN SET result = 'Йоханес';
        WHEN 65 THEN SET result = 'Эдуард';
        WHEN 66 THEN SET result = 'Сергей';
        WHEN 67 THEN SET result = 'Илья';
        WHEN 68 THEN SET result = 'Семён';
        WHEN 69 THEN SET result = 'Святослав';
        WHEN 70 THEN SET result = 'Яромир';
        WHEN 71 THEN SET result = 'Ефрем';
        WHEN 72 THEN SET result = 'Игнат';
        WHEN 73 THEN SET result = 'Константин';
        WHEN 74 THEN SET result = 'Денис';
        WHEN 75 THEN SET result = 'Николай';
        WHEN 76 THEN SET result = 'Цефас';
        WHEN 77 THEN SET result = 'Василий';
        WHEN 78 THEN SET result = 'Ярослав';
        WHEN 79 THEN SET result = 'Антон';
        WHEN 80 THEN SET result = 'Константин';
        WHEN 81 THEN SET result = 'Шамиль';
        WHEN 82 THEN SET result = 'Йоханес';
        WHEN 83 THEN SET result = 'Герасим';
        WHEN 84 THEN SET result = 'Устин';
        WHEN 85 THEN SET result = 'Радислав';
        WHEN 86 THEN SET result = 'Борис';
        WHEN 87 THEN SET result = 'Казбек';
        WHEN 88 THEN SET result = 'Марат';
        WHEN 89 THEN SET result = 'Марк';
        WHEN 90 THEN SET result = 'Чеслав';
        WHEN 91 THEN SET result = 'Юрий';
        WHEN 92 THEN SET result = 'Тимофей';
        WHEN 93 THEN SET result = 'Корнелий';
        WHEN 94 THEN SET result = 'Матвей';
        WHEN 95 THEN SET result = 'Фёдор';
        WHEN 96 THEN SET result = 'Цицерон';
        WHEN 97 THEN SET result = 'Юлий';
        WHEN 98 THEN SET result = 'Спартак';
        WHEN 99 THEN SET result = 'Лев';
        WHEN 100 THEN SET result = 'Алмаз';
    END CASE;
    RETURN result;
END$$

DROP FUNCTION IF EXISTS female_lastname;
CREATE FUNCTION female_lastname () RETURNS VARCHAR(25)
NOT DETERMINISTIC
BEGIN
    DECLARE num INT;
    DECLARE result VARCHAR(25);
    SELECT CEIL(RAND()*100) INTO num;
    CASE num
        WHEN 1 THEN SET result = 'Ширяева';
        WHEN 2 THEN SET result = 'Ялова';
        WHEN 3 THEN SET result = 'Селезнёва';
        WHEN 4 THEN SET result = 'Кириленко';
        WHEN 5 THEN SET result = 'Блинова';
        WHEN 6 THEN SET result = 'Сорокина';
        WHEN 7 THEN SET result = 'Михеева';
        WHEN 8 THEN SET result = 'Моисеева';
        WHEN 9 THEN SET result = 'Колесникова';
        WHEN 10 THEN SET result = 'Петрова';
        WHEN 11 THEN SET result = 'Шарапова';
        WHEN 12 THEN SET result = 'Владимирова';
        WHEN 13 THEN SET result = 'Давыдова';
        WHEN 14 THEN SET result = 'Данилова';
        WHEN 15 THEN SET result = 'Андреева';
        WHEN 16 THEN SET result = 'Громова';
        WHEN 17 THEN SET result = 'Тетерина';
        WHEN 18 THEN SET result = 'Капустина';
        WHEN 19 THEN SET result = 'Смирнова';
        WHEN 20 THEN SET result = 'Красинец';
        WHEN 21 THEN SET result = 'Кириллова';
        WHEN 22 THEN SET result = 'Гамула';
        WHEN 23 THEN SET result = 'Давыдова';
        WHEN 24 THEN SET result = 'Потапова';
        WHEN 25 THEN SET result = 'Ялова';
        WHEN 26 THEN SET result = 'Богданова';
        WHEN 27 THEN SET result = 'Игнатова';
        WHEN 28 THEN SET result = 'Пархоменко';
        WHEN 29 THEN SET result = 'Соловьёва';
        WHEN 30 THEN SET result = 'Рябова';
        WHEN 31 THEN SET result = 'Предыбайло';
        WHEN 32 THEN SET result = 'Белова';
        WHEN 33 THEN SET result = 'Фролова';
        WHEN 34 THEN SET result = 'Никифорова';
        WHEN 35 THEN SET result = 'Фокина';
        WHEN 36 THEN SET result = 'Наумова';
        WHEN 37 THEN SET result = 'Сирко';
        WHEN 38 THEN SET result = 'Иващенко';
        WHEN 39 THEN SET result = 'Цушко';
        WHEN 40 THEN SET result = 'Бачей';
        WHEN 41 THEN SET result = 'Бондаренко';
        WHEN 42 THEN SET result = 'Коломоец';
        WHEN 43 THEN SET result = 'Васильева';
        WHEN 44 THEN SET result = 'Кириллова';
        WHEN 45 THEN SET result = 'Яковенко';
        WHEN 46 THEN SET result = 'Петрова';
        WHEN 47 THEN SET result = 'Лобанова';
        WHEN 48 THEN SET result = 'Данилова';
        WHEN 49 THEN SET result = 'Ткаченко';
        WHEN 50 THEN SET result = 'Повалий';
        WHEN 51 THEN SET result = 'Лукашенко';
        WHEN 52 THEN SET result = 'Андрусейко';
        WHEN 53 THEN SET result = 'Субботина';
        WHEN 54 THEN SET result = 'Яковенко';
        WHEN 55 THEN SET result = 'Бирюкова';
        WHEN 56 THEN SET result = 'Петухова';
        WHEN 57 THEN SET result = 'Никитина';
        WHEN 58 THEN SET result = 'Моисеенко';
        WHEN 59 THEN SET result = 'Осипова';
        WHEN 60 THEN SET result = 'Мирна';
        WHEN 61 THEN SET result = 'Матвеева';
        WHEN 62 THEN SET result = 'Повалий';
        WHEN 63 THEN SET result = 'Котова';
        WHEN 64 THEN SET result = 'Терентьева';
        WHEN 65 THEN SET result = 'Семочко';
        WHEN 66 THEN SET result = 'Белова';
        WHEN 67 THEN SET result = 'Зыкова';
        WHEN 68 THEN SET result = 'Цушко';
        WHEN 69 THEN SET result = 'Маслова';
        WHEN 70 THEN SET result = 'Третьякова';
        WHEN 71 THEN SET result = 'Евсеева';
        WHEN 72 THEN SET result = 'Несвитайло';
        WHEN 73 THEN SET result = 'Власова';
        WHEN 74 THEN SET result = 'Кобзар';
        WHEN 75 THEN SET result = 'Борисенко';
        WHEN 76 THEN SET result = 'Гурьева';
        WHEN 77 THEN SET result = 'Громова';
        WHEN 78 THEN SET result = 'Гайчук';
        WHEN 79 THEN SET result = 'Навальна';
        WHEN 80 THEN SET result = 'Шуфрич';
        WHEN 81 THEN SET result = 'Зуева';
        WHEN 82 THEN SET result = 'Михеева';
        WHEN 83 THEN SET result = 'Герасимова';
        WHEN 84 THEN SET result = 'Антонова';
        WHEN 85 THEN SET result = 'Ершова';
        WHEN 86 THEN SET result = 'Шилова';
        WHEN 87 THEN SET result = 'Мартынова';
        WHEN 88 THEN SET result = 'Кулагина';
        WHEN 89 THEN SET result = 'Анисимова';
        WHEN 90 THEN SET result = 'Толочко';
        WHEN 91 THEN SET result = 'Рябова';
        WHEN 92 THEN SET result = 'Якушева';
        WHEN 93 THEN SET result = 'Колобова';
        WHEN 94 THEN SET result = 'Федотова';
        WHEN 95 THEN SET result = 'Русакова';
        WHEN 96 THEN SET result = 'Алексеева';
        WHEN 97 THEN SET result = 'Хитрук';
        WHEN 98 THEN SET result = 'Константинова';
        WHEN 99 THEN SET result = 'Одинцова';
        WHEN 100 THEN SET result = 'Рыжих';
    END CASE;
    RETURN result;
END$$

DROP FUNCTION IF EXISTS female_firstname;
CREATE FUNCTION female_firstname () RETURNS VARCHAR(25)
NOT DETERMINISTIC
BEGIN
    DECLARE num INT;
    DECLARE result VARCHAR(25);
    SELECT CEIL(RAND()*100) INTO num;
    CASE num
        WHEN 1 THEN SET result = 'Таисия';
        WHEN 2 THEN SET result = 'Георгина';
        WHEN 3 THEN SET result = 'Хильда';
        WHEN 4 THEN SET result = 'Чара';
        WHEN 5 THEN SET result = 'Зоя';
        WHEN 6 THEN SET result = 'София';
        WHEN 7 THEN SET result = 'Альбина';
        WHEN 8 THEN SET result = 'Божена';
        WHEN 9 THEN SET result = 'Эрика';
        WHEN 10 THEN SET result = 'Цезария';
        WHEN 11 THEN SET result = 'Злата';
        WHEN 12 THEN SET result = 'Ядвига';
        WHEN 13 THEN SET result = 'Зинаида';
        WHEN 14 THEN SET result = 'Хильда';
        WHEN 15 THEN SET result = 'Анфиса';
        WHEN 16 THEN SET result = 'Бронислава';
        WHEN 17 THEN SET result = 'Варвара';
        WHEN 18 THEN SET result = 'Элоиза';
        WHEN 19 THEN SET result = 'Доминика';
        WHEN 20 THEN SET result = 'Марта';
        WHEN 21 THEN SET result = 'Раиса';
        WHEN 22 THEN SET result = 'Божена';
        WHEN 23 THEN SET result = 'Люся';
        WHEN 24 THEN SET result = 'Дарья';
        WHEN 25 THEN SET result = 'Шарлота';
        WHEN 26 THEN SET result = 'Капитолина';
        WHEN 27 THEN SET result = 'Ульяна';
        WHEN 28 THEN SET result = 'Мальвина';
        WHEN 29 THEN SET result = 'Софья';
        WHEN 30 THEN SET result = 'Ульяна';
        WHEN 31 THEN SET result = 'Лада';
        WHEN 32 THEN SET result = 'Доминика';
        WHEN 33 THEN SET result = 'Юнона';
        WHEN 34 THEN SET result = 'Надежда';
        WHEN 35 THEN SET result = 'Шарлота';
        WHEN 36 THEN SET result = 'Йоони';
        WHEN 37 THEN SET result = 'Изольда';
        WHEN 38 THEN SET result = 'Чара';
        WHEN 39 THEN SET result = 'Христина';
        WHEN 40 THEN SET result = 'Гюзель';
        WHEN 41 THEN SET result = 'Чилита';
        WHEN 42 THEN SET result = 'Флорентина';
        WHEN 43 THEN SET result = 'Нелли';
        WHEN 44 THEN SET result = 'Ева';
        WHEN 45 THEN SET result = 'Шанетта';
        WHEN 46 THEN SET result = 'Рада';
        WHEN 47 THEN SET result = 'Октябрина';
        WHEN 48 THEN SET result = 'Дарья';
        WHEN 49 THEN SET result = 'Светлана';
        WHEN 50 THEN SET result = 'Устинья';
        WHEN 51 THEN SET result = 'Чеслава';
        WHEN 52 THEN SET result = 'Жанна';
        WHEN 53 THEN SET result = 'Наталья';
        WHEN 54 THEN SET result = 'Юнона';
        WHEN 55 THEN SET result = 'Ника';
        WHEN 56 THEN SET result = 'Лариса';
        WHEN 57 THEN SET result = 'Рада';
        WHEN 58 THEN SET result = 'Зоя';
        WHEN 59 THEN SET result = 'Октябрина';
        WHEN 60 THEN SET result = 'Софья';
        WHEN 61 THEN SET result = 'Федосья';
        WHEN 62 THEN SET result = 'Чечилия';
        WHEN 63 THEN SET result = 'Галина';
        WHEN 64 THEN SET result = 'Светлана';
        WHEN 65 THEN SET result = 'Дарья';
        WHEN 66 THEN SET result = 'Дарья';
        WHEN 67 THEN SET result = 'Цилла';
        WHEN 68 THEN SET result = 'Эльмира';
        WHEN 69 THEN SET result = 'Белла';
        WHEN 70 THEN SET result = 'Прасковья';
        WHEN 71 THEN SET result = 'Йоони';
        WHEN 72 THEN SET result = 'Фаина';
        WHEN 73 THEN SET result = 'Любовь';
        WHEN 74 THEN SET result = 'Мария';
        WHEN 75 THEN SET result = 'Лариса';
        WHEN 76 THEN SET result = 'Рада';
        WHEN 77 THEN SET result = 'Софья';
        WHEN 78 THEN SET result = 'Пелагея';
        WHEN 79 THEN SET result = 'Ева';
        WHEN 80 THEN SET result = 'Владлена';
        WHEN 81 THEN SET result = 'Глория';
        WHEN 82 THEN SET result = 'Антонина';
        WHEN 83 THEN SET result = 'Клара';
        WHEN 84 THEN SET result = 'Пелагея';
        WHEN 85 THEN SET result = 'Георгина';
        WHEN 86 THEN SET result = 'Лариса';
        WHEN 87 THEN SET result = 'Ядвига';
        WHEN 88 THEN SET result = 'Нонна';
        WHEN 89 THEN SET result = 'Зоя';
        WHEN 90 THEN SET result = 'Вера';
        WHEN 91 THEN SET result = 'Дарья';
        WHEN 92 THEN SET result = 'Фаина';
        WHEN 93 THEN SET result = 'Инесса';
        WHEN 94 THEN SET result = 'Лада';
        WHEN 95 THEN SET result = 'Глафира';
        WHEN 96 THEN SET result = 'Чечилия';
        WHEN 97 THEN SET result = 'Рената';
        WHEN 98 THEN SET result = 'Йосифа';
        WHEN 99 THEN SET result = 'Цилла';
        WHEN 100 THEN SET result = 'Елизавета';
    END CASE;
    RETURN result;
END$$

DROP FUNCTION IF EXISTS firstname_generator;
CREATE FUNCTION firstname_generator (sex VARCHAR(6)) RETURNS VARCHAR(6)
NOT DETERMINISTIC
BEGIN
    DECLARE num INT;
    DECLARE result VARCHAR(6);
    CASE sex
        WHEN 'male' THEN SELECT male_firstname() INTO result;
        WHEN 'female' THEN SELECT female_firstname() INTO result;
    END CASE;
    RETURN result;
END$$

DROP FUNCTION IF EXISTS lastname_generator;
CREATE FUNCTION lastname_generator (sex VARCHAR(6)) RETURNS VARCHAR(6)
NOT DETERMINISTIC
BEGIN
    DECLARE num INT;
    DECLARE result VARCHAR(6);
    CASE sex
        WHEN 'male' THEN SELECT male_lastname() INTO result;
        WHEN 'female' THEN SELECT female_lastname() INTO result;
    END CASE;
    RETURN result;
END$$

DELIMITER ;