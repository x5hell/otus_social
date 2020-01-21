SET GLOBAL log_bin_trust_function_creators = 1;
DELIMITER $$


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

CREATE FUNCTION male_last_name () RETURNS VARCHAR(25)
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

CREATE FUNCTION male_first_name() RETURNS VARCHAR(25)
NOT DETERMINISTIC
BEGIN
    DECLARE num INT;
    DECLARE result VARCHAR(25);
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

CREATE FUNCTION female_last_name () RETURNS VARCHAR(25)
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

CREATE FUNCTION female_first_name () RETURNS VARCHAR(25)
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

CREATE FUNCTION first_name_generator (sex VARCHAR(6)) RETURNS VARCHAR(6)
NOT DETERMINISTIC
BEGIN
    DECLARE result VARCHAR(6);
    CASE sex
        WHEN 'male' THEN SELECT male_first_name() INTO result;
        WHEN 'female' THEN SELECT female_first_name() INTO result;
    END CASE;
    RETURN result;
END$$

CREATE FUNCTION last_name_generator (sex VARCHAR(6)) RETURNS VARCHAR(6)
NOT DETERMINISTIC
BEGIN
    DECLARE result VARCHAR(6);
    CASE sex
        WHEN 'male' THEN SELECT male_last_name() INTO result;
        WHEN 'female' THEN SELECT female_last_name() INTO result;
    END CASE;
    RETURN result;
END$$

CREATE FUNCTION generate_place_type () RETURNS VARCHAR(10)
    NOT DETERMINISTIC
BEGIN
    DECLARE num INT;
    DECLARE result VARCHAR(10);
    SELECT CEIL(RAND()*12) INTO num;
    CASE num
        WHEN 1 THEN SET result = 'город';
        WHEN 2 THEN SET result = 'село';
        WHEN 3 THEN SET result = 'деревня';
        WHEN 4 THEN SET result = 'посёлок';
        WHEN 5 THEN SET result = 'аул';
        WHEN 6 THEN SET result = 'станица';
        WHEN 7 THEN SET result = 'г';
        WHEN 8 THEN SET result = 'с';
        WHEN 9 THEN SET result = 'д';
        WHEN 10 THEN SET result = 'п';
        WHEN 11 THEN SET result = 'а';
        WHEN 12 THEN SET result = 'с';
        END CASE;
    RETURN result;
END$$

CREATE FUNCTION generate_city_name_form_1 () RETURNS VARCHAR(25)
    NOT DETERMINISTIC
BEGIN
    DECLARE num INT;
    DECLARE result VARCHAR(25);
    SELECT CEIL(RAND()*25) INTO num;
    CASE num
        WHEN 1 THEN SET result = 'Пупсы';
        WHEN 2 THEN SET result = 'Дешевки';
        WHEN 3 THEN SET result = 'Кокаиновые горы';
        WHEN 4 THEN SET result = 'Ломки';
        WHEN 5 THEN SET result = 'Черви';
        WHEN 6 THEN SET result = 'Блювиничи';
        WHEN 7 THEN SET result = 'Чуваки';
        WHEN 8 THEN SET result = 'Блохи';
        WHEN 9 THEN SET result = 'Козлы';
        WHEN 10 THEN SET result = 'Опухлики';
        WHEN 11 THEN SET result = 'Сувалки';
        WHEN 12 THEN SET result = 'Гробы';
        WHEN 13 THEN SET result = 'Хачики';
        WHEN 14 THEN SET result = 'Ишаки';
        WHEN 15 THEN SET result = 'Лужи';
        WHEN 16 THEN SET result = 'Мочилки';
        WHEN 17 THEN SET result = 'Лобки';
        WHEN 18 THEN SET result = 'Пупки';
        WHEN 19 THEN SET result = 'Кобеляки';
        WHEN 20 THEN SET result = 'Бздюли';
        WHEN 21 THEN SET result = 'Бобрики';
        WHEN 22 THEN SET result = 'Мусорки';
        WHEN 23 THEN SET result = 'Кончинки';
        WHEN 24 THEN SET result = 'Бугры';
        WHEN 25 THEN SET result = 'Лохи';
        END CASE;
    RETURN result;
END$$

CREATE FUNCTION generate_city_part_form_1 () RETURNS VARCHAR(10)
    NOT DETERMINISTIC
BEGIN
    DECLARE num INT;
    DECLARE result VARCHAR(10);
    SELECT CEIL(RAND()*15) INTO num;
    CASE num
        WHEN 1 THEN SET result = 'Большие ';
        WHEN 2 THEN SET result = 'Малые ';
        WHEN 3 THEN SET result = 'Верхние ';
        WHEN 4 THEN SET result = 'Нижние ';
        WHEN 5 THEN SET result = 'Весёлые ';
        WHEN 6 THEN SET result = 'Святые ';
        WHEN 7 THEN SET result = 'Великие ';
        WHEN 8 THEN SET result = 'Старые ';
        WHEN 9 THEN SET result = 'Новые ';
        ELSE SET result = '';
        END CASE;
    RETURN result;
END$$

CREATE FUNCTION generate_city_name_form_2 () RETURNS VARCHAR(25)
    NOT DETERMINISTIC
BEGIN
    DECLARE num INT;
    DECLARE result VARCHAR(25);
    SELECT CEIL(RAND()*25) INTO num;
    CASE num
        WHEN 1 THEN SET result = 'Пысса';
        WHEN 2 THEN SET result = 'Баклань';
        WHEN 3 THEN SET result = 'Куриловка';
        WHEN 4 THEN SET result = 'Балда';
        WHEN 5 THEN SET result = 'Засосная';
        WHEN 6 THEN SET result = 'Звероножка';
        WHEN 7 THEN SET result = 'Вобля';
        WHEN 8 THEN SET result = 'Мусорка';
        WHEN 9 THEN SET result = 'Бухловка';
        WHEN 10 THEN SET result = 'Коноплянка';
        WHEN 11 THEN SET result = 'Мухоудёровка';
        WHEN 12 THEN SET result = 'Кончинка';
        WHEN 13 THEN SET result = 'Ушмары';
        WHEN 14 THEN SET result = 'Лапша';
        WHEN 15 THEN SET result = 'Щель';
        WHEN 16 THEN SET result = 'Дешевка';
        WHEN 17 THEN SET result = 'Ломка';
        WHEN 18 THEN SET result = 'Блоха';
        WHEN 19 THEN SET result = 'Сувалка';
        WHEN 20 THEN SET result = 'Лужа';
        WHEN 21 THEN SET result = 'Мочилка';
        WHEN 22 THEN SET result = 'Бздюля';
        WHEN 23 THEN SET result = 'Жаба';
        WHEN 24 THEN SET result = 'Голодранка';
        WHEN 25 THEN SET result = 'Хотелка';
        END CASE;
    RETURN result;
END$$

CREATE FUNCTION generate_city_part_form_2 () RETURNS VARCHAR(10)
    NOT DETERMINISTIC
BEGIN
    DECLARE num INT;
    DECLARE result VARCHAR(10);
    SELECT CEIL(RAND()*15) INTO num;
    CASE num
        WHEN 1 THEN SET result = 'Большая ';
        WHEN 2 THEN SET result = 'Малая ';
        WHEN 3 THEN SET result = 'Верхняя ';
        WHEN 4 THEN SET result = 'Нижняя ';
        WHEN 5 THEN SET result = 'Весёлая ';
        WHEN 6 THEN SET result = 'Святая ';
        WHEN 7 THEN SET result = 'Великая ';
        WHEN 8 THEN SET result = 'Старая ';
        WHEN 9 THEN SET result = 'Новая ';
        ELSE SET result = '';
        END CASE;
    RETURN result;
END$$

CREATE FUNCTION generate_city_name_form_3 () RETURNS VARCHAR(25)
    NOT DETERMINISTIC
BEGIN
    DECLARE num INT;
    DECLARE result VARCHAR(25);
    SELECT CEIL(RAND()*25) INTO num;
    CASE num
        WHEN 1 THEN SET result = 'Куяш';
        WHEN 2 THEN SET result = 'Сисковский';
        WHEN 3 THEN SET result = 'Гадюшник';
        WHEN 4 THEN SET result = 'Крыжополь';
        WHEN 5 THEN SET result = 'Бугор';
        WHEN 6 THEN SET result = 'Бобрик';
        WHEN 7 THEN SET result = 'Усох';
        WHEN 8 THEN SET result = 'Лох';
        WHEN 9 THEN SET result = 'Мухосранск';
        WHEN 10 THEN SET result = 'Пупс';
        WHEN 11 THEN SET result = 'Червь';
        WHEN 12 THEN SET result = 'Чувак';
        WHEN 13 THEN SET result = 'Козёл';
        WHEN 14 THEN SET result = 'Опухлик';
        WHEN 15 THEN SET result = 'Гроб';
        WHEN 16 THEN SET result = 'Хачик';
        WHEN 17 THEN SET result = 'Ишак';
        WHEN 18 THEN SET result = 'Лобок';
        WHEN 19 THEN SET result = 'Пупок';
        WHEN 20 THEN SET result = 'Кобеляка';
        WHEN 21 THEN SET result = 'Трус';
        WHEN 22 THEN SET result = 'Дурак';
        WHEN 23 THEN SET result = 'Хрен';
        WHEN 24 THEN SET result = 'Свин';
        WHEN 25 THEN SET result = 'Мухоед';
        END CASE;
    RETURN result;
END$$

CREATE FUNCTION generate_city_part_form_3 () RETURNS VARCHAR(10)
    NOT DETERMINISTIC
BEGIN
    DECLARE num INT;
    DECLARE result VARCHAR(10);
    SELECT CEIL(RAND()*15) INTO num;
    CASE num
        WHEN 1 THEN SET result = 'Большой ';
        WHEN 2 THEN SET result = 'Малый ';
        WHEN 3 THEN SET result = 'Верхний ';
        WHEN 4 THEN SET result = 'Нижний ';
        WHEN 5 THEN SET result = 'Весёлый ';
        WHEN 6 THEN SET result = 'Святой ';
        WHEN 7 THEN SET result = 'Великий ';
        WHEN 8 THEN SET result = 'Старый ';
        WHEN 9 THEN SET result = 'Новый ';
        ELSE SET result = '';
        END CASE;
    RETURN result;
END$$

CREATE FUNCTION generate_city_name_form_4 () RETURNS VARCHAR(25)
    NOT DETERMINISTIC
BEGIN
    DECLARE num INT;
    DECLARE result VARCHAR(25);
    SELECT CEIL(RAND()*25) INTO num;
    CASE num
        WHEN 1 THEN SET result = 'Лохово';
        WHEN 2 THEN SET result = 'Струйкино';
        WHEN 3 THEN SET result = 'Овнище';
        WHEN 4 THEN SET result = 'Дно';
        WHEN 5 THEN SET result = 'Трусово';
        WHEN 6 THEN SET result = 'Ширяево';
        WHEN 7 THEN SET result = 'Новопозорново';
        WHEN 8 THEN SET result = 'Зачатье';
        WHEN 9 THEN SET result = 'Дураково';
        WHEN 10 THEN SET result = 'Муходоево';
        WHEN 11 THEN SET result = 'Хреново';
        WHEN 12 THEN SET result = 'Бухалово';
        WHEN 13 THEN SET result = 'Жабино';
        WHEN 14 THEN SET result = 'Кончинино';
        WHEN 15 THEN SET result = 'Голодранкино';
        WHEN 16 THEN SET result = 'Хотелово';
        WHEN 17 THEN SET result = 'Бухалово';
        WHEN 18 THEN SET result = 'Лобково';
        WHEN 19 THEN SET result = 'Какино';
        WHEN 20 THEN SET result = 'Отхожее';
        WHEN 21 THEN SET result = 'Хренище ';
        WHEN 22 THEN SET result = 'Матюково';
        WHEN 23 THEN SET result = 'Пьянкино';
        WHEN 24 THEN SET result = 'Свинорье';
        WHEN 25 THEN SET result = 'Матюково';
        END CASE;
    RETURN result;
END$$

CREATE FUNCTION generate_city_part_form_4 () RETURNS VARCHAR(10)
    NOT DETERMINISTIC
BEGIN
    DECLARE num INT;
    DECLARE result VARCHAR(10);
    SELECT CEIL(RAND()*15) INTO num;
    CASE num
        WHEN 1 THEN SET result = 'Большое ';
        WHEN 2 THEN SET result = 'Малое ';
        WHEN 3 THEN SET result = 'Верхнее ';
        WHEN 4 THEN SET result = 'Нижнее ';
        WHEN 5 THEN SET result = 'Весёлое ';
        WHEN 6 THEN SET result = 'Святое ';
        WHEN 7 THEN SET result = 'Великое ';
        WHEN 8 THEN SET result = 'Старое ';
        WHEN 9 THEN SET result = 'Новое ';
        ELSE SET result = '';
        END CASE;
    RETURN result;
END$$

CREATE FUNCTION generate_city_name () RETURNS VARCHAR(50)
    NOT DETERMINISTIC
BEGIN
    DECLARE form INT;
    DECLARE city_name VARCHAR(25);
    DECLARE result VARCHAR(50);
    SELECT CEIL(RAND()*4) INTO form;
    CASE form
        WHEN 1 THEN
            SELECT CONCAT(generate_city_part_form_1(),  generate_city_name_form_1()) INTO city_name;
        WHEN 2 THEN
            SELECT CONCAT(generate_city_part_form_2(),  generate_city_name_form_2()) INTO city_name;
        WHEN 3 THEN
            SELECT CONCAT(generate_city_part_form_3(),  generate_city_name_form_3()) INTO city_name;
        WHEN 4 THEN
            SELECT CONCAT(generate_city_part_form_4(),  generate_city_name_form_4()) INTO city_name;
        END CASE;
    SELECT CONCAT(generate_place_type(), ' ', city_name)  INTO result;
    RETURN result;
END$$

CREATE PROCEDURE generate_cities (quantity INT)
BEGIN
    DECLARE step INT;
    SET step = 1;
    WHILE step <= quantity DO
        INSERT INTO city (`id`, `name`) VALUES (step, generate_city_name());
        SET step = step + 1;
    END WHILE;
END$$

CALL generate_cities(100000)$$

DROP FUNCTION IF EXISTS num_to_sex$$
DROP FUNCTION IF EXISTS sex_generator$$
DROP FUNCTION IF EXISTS male_last_name$$
DROP FUNCTION IF EXISTS male_first_name$$
DROP FUNCTION IF EXISTS female_last_name$$
DROP FUNCTION IF EXISTS female_first_name$$
DROP FUNCTION IF EXISTS first_name_generator$$
DROP FUNCTION IF EXISTS last_name_generator$$
DROP FUNCTION IF EXISTS generate_place_type$$
DROP FUNCTION IF EXISTS generate_city_name_form_1$$
DROP FUNCTION IF EXISTS generate_city_part_form_1$$
DROP FUNCTION IF EXISTS generate_city_name_form_2$$
DROP FUNCTION IF EXISTS generate_city_part_form_2$$
DROP FUNCTION IF EXISTS generate_city_name_form_3$$
DROP FUNCTION IF EXISTS generate_city_part_form_3$$
DROP FUNCTION IF EXISTS generate_city_name_form_4$$
DROP FUNCTION IF EXISTS generate_city_part_form_4$$
DROP FUNCTION IF EXISTS generate_city_name;
DROP PROCEDURE IF EXISTS generate_cities;

DELIMITER ;

