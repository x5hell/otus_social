INSERT INTO `server_id` (`id`) VALUES (1);

INSERT INTO city
    (`id`, `name`)
VALUES
    (1, "Неелово"),
    (2, "Горелово"),
    (3, "Неурожайка"),
    (4, "Косяковка"),
    (5, "Козявкино");
    
INSERT INTO interest
    (`id`, `name`)
VALUES
    (1, "вышивание крестиком"),
    (2, "крав-мага"),
    (3, "игра в го"),
    (4, "хор мальчиков");

INSERT INTO `user`
    (`id`, `login`, `password`, `first_name`, `last_name`, `age`, `sex`, `city_id`)
VALUES
    (1, 'grebnev', 'e10adc3949ba59abbe56e057f20f883e', 'Гребневский', 'Евсей', 16, 'male', 1),
    (2, 'unona', 'e10adc3949ba59abbe56e057f20f883e', 'Герасимова', 'Юнона', 30, 'female', 3),
    (3, 'sofia', 'e10adc3949ba59abbe56e057f20f883e', 'Потапова', 'София', 35, 'female', null),
    (4, 'triasilo', 'e10adc3949ba59abbe56e057f20f883e', 'Трясило', 'Дмитрий', 38, 'male', 5),
    (5, 'rymar', 'e10adc3949ba59abbe56e057f20f883e', 'Рымар', 'Рафаил', 32, 'male', 2),
    (6, 'dominica', 'e10adc3949ba59abbe56e057f20f883e', 'Исакова', 'Доминика', 57, 'female', 1),
    (7, 'cheslav', 'e10adc3949ba59abbe56e057f20f883e', 'Грабчак', 'Чеслав', 36, 'male', null),
    (8, 'marta', 'e10adc3949ba59abbe56e057f20f883e', 'Меркушева', 'Марта', 37, 'female', 4),
    (9, 'kseniya', 'e10adc3949ba59abbe56e057f20f883e', 'Никонова', 'Ксения', 32, 'female', 5),
    (10, 'olesya', 'e10adc3949ba59abbe56e057f20f883e', 'Михайлова', 'Олеся', 39, 'female', 2);

INSERT INTO `user_interest`
    (`user_id`, `interest_id`)
VALUES
    ('1', '2'),
    ('1', '4'),
    ('2', '1'),
    ('3', '1'),
    ('4', '1'),
    ('4', '2'),
    ('4', '3'),
    ('5', '4'),
    ('5', '2'),
    ('7', '1'),
    ('7', '4'),
    ('9', '4'),
    ('9', '3'),
    ('9', '1'),
    ('10', '3'),
    ('10', '1');