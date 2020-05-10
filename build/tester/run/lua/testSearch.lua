local socket = require'socket'

local char_to_hex = function(c)
  return string.format("%%%02X", string.byte(c))
end

local function urlencode(url)
  if url == nil then
    return
  end
  url = url:gsub("\n", "\r\n")
  url = url:gsub("([^%w ])", char_to_hex)
  url = url:gsub(" ", "+")
  return url
end

firstnamePart={
  "Пла","Уст","Тим","Мих","Гле","Шар","Уст","Ере","Эду","Гер",
  "Ана","Жер","Вла","Пав","Ефи","Пёт","Уст","Чар","Ост","Кон",
  "Ян" ,"Шер","Люб","Арт","Гав","Вал","Яро","Чар","Жиг","Жер",
  "Уст","Бог","Хар","Мар","Фил","Гео","Ива","Юли","Хар","Тим",
  "Шер","Тит","Мат","Кри","Яро","Дон","Циц","Дин","Цез","Мар",
  "Шар","Уст","Куз","Раф","Ефр","Йог","Але","Роб","Афа","Мир",
  "Нес","Свя","Каз","Йох","Эду","Сер","Иль","Сем","Свя","Яро",
  "Ефр","Игн","Кон","Ден","Ник","Цеф","Вас","Яро","Ант","Кон",
  "Шам","Йох","Гер","Уст","Рад","Бор","Каз","Мар","Мар","Чес",
  "Юри","Тим","Кор","Мат","Фёд","Циц","Юли","Спа","Лев","Алм"
}
lastnamePart={
  "Пон","Кра","Фёд","Пах","Пет","Оси","Щер","Куд","Гри","Рож",
  "Нег","Гел","Пет","Тур","Бач","Кул","Сак","Лук","Ели","Ток",
  "Кот","Тер","Щук","Дзю","Мир","Пер","Мир","Лук","Ерё","Жел",
  "Жур","Оди","Пол","Заб","Яво","Гор","Лин","Ско","Тре","Ряб",
  "Мар","Бур","Сыс","Зим","Иль","Тар","Дан","Пав","Пар","Мои",
  "Фад","Куз","Жук","Лаз","Мик","Лыт","Кот","Тих","Бел","Гал",
  "Гон","Яло","Иса","Пет","Сел","Гор","Пес","Тка","Тем","Бон",
  "Сав","Коз","Мяс","Тру","Кор","Ден","Цве","Шкр","Пав","Бор",
  "Ива","Гел","Пав","Сор","Мас","Кул","Миш","Тер","Пал","Кол",
  "Тря","Май","Ива","Юди","Рыб","Рог","Кул","Каз","Ели","Евс"
}

requests = {}

init = function()
    local socketTime = socket.gettime()*10000
    math.randomseed(socketTime)
    for i = 1, 1000, 1 do
        local randomFirstNamePartIndex = math.random(1,100)
        local randomFirstNamePart = firstnamePart[randomFirstNamePartIndex]
        local randomLastnamePartIndex = math.random(1,100)
        local randomLastnamePart = lastnamePart[randomLastnamePartIndex]
        local path = '/search?first-name=' ..
                   urlencode(randomFirstNamePart) ..
                   '&last-name=' ..
                   urlencode(randomLastnamePart)
        requests[i] = wrk.format("GET", path)
    end
end

request = function()
    local socketTime = socket.gettime()*10000
    math.randomseed(socketTime)
    randomRequest = math.random(1,1000)
    return requests[randomRequest]
end