package generator

import (
	"math/rand"
	"strings"
)

const CityForms = 4

func City() string {
	form := rand.Intn(CityForms)
	var cityParts []string
	placeType := generatePlaceType()
	cityPart := generateCityPart(form)
	cityName := generateCityName(form)
	if len(cityPart) > 0 {
		cityParts = []string{placeType, cityPart, cityName}
	} else {
		cityParts = []string{placeType, cityName}
	}
	return strings.Join(cityParts, " ")
}

func generatePlaceType() string {
	num := rand.Intn(13)
	switch num {
		case 1: return "город"
		case 2: return "село"
		case 3: return "деревня"
		case 4: return "посёлок"
		case 5: return "аул"
		case 6: return "станица"
		case 7: return "г"
		case 8: return "с"
		case 9: return "д"
		case 10: return "п"
		case 11: return "а"
		case 12: return "с"
		default: return ""
	}
}

func generateCityPart(form int) string {
	num := rand.Intn(11)
	var forms [CityForms]string
	switch num {
		case 1: forms = [CityForms]string{"Большие", "Большая", "Большой", "Большое"}
		case 2: forms = [CityForms]string{"Малые", "Малая", "Малый", "Малое"}
		case 3: forms = [CityForms]string{"Верхние", "Верхняя", "Верхний", "Верхнее"}
		case 4: forms = [CityForms]string{"Нижние", "Нижняя", "Нижний", "Нижнее"}
		case 5: forms = [CityForms]string{"Весёлые", "Весёлая", "Весёлый", "Весёлое"}
		case 6: forms = [CityForms]string{"Святые", "Святая", "Святой", "Святое"}
		case 7: forms = [CityForms]string{"Великие", "Великая", "Великий", "Великое"}
		case 8: forms = [CityForms]string{"Старые", "Старая", "Старый", "Старое"}
		case 9: forms = [CityForms]string{"Новые", "Новая", "Новый", "Новое"}
		default:
			return ""
	}
	return forms[form]
}

func generateCityName(form int) string {
	num := rand.Intn(26)
	var forms [CityForms]string
	switch num {
		case 1: forms = [CityForms]string{"Пупсы","Пысса","Куяш","Лохово"}
		case 2: forms = [CityForms]string{"Дешевки","Баклань","Сисковский","Струйкино"}
		case 3: forms = [CityForms]string{"Кокаиновые горы","Куриловка","Гадюшник","Овнище"}
		case 4: forms = [CityForms]string{"Ломки","Балда","Крыжополь","Дно"}
		case 5: forms = [CityForms]string{"Черви","Засосная","Бугор","Трусово"}
		case 6: forms = [CityForms]string{"Блювиничи","Звероножка","Бобрик","Ширяево"}
		case 7: forms = [CityForms]string{"Чуваки","Вобля","Усох","Новопозорново"}
		case 8: forms = [CityForms]string{"Блохи","Мусорка","Лох","Зачатье"}
		case 9: forms = [CityForms]string{"Козлы","Бухловка","Мухосранск","Дураково"}
		case 10: forms = [CityForms]string{"Опухлики","Коноплянка","Пупс","Муходоево"}
		case 11: forms = [CityForms]string{"Сувалки","Мухоудёровка","Червь","Хреново"}
		case 12: forms = [CityForms]string{"Гробы","Кончинка","Чувак","Бухалово"}
		case 13: forms = [CityForms]string{"Хачики","Ушмары","Козёл","Жабино"}
		case 14: forms = [CityForms]string{"Ишаки","Лапша","Опухлик","Кончинино"}
		case 15: forms = [CityForms]string{"Лужи","Щель","Гроб","Голодранкино"}
		case 16: forms = [CityForms]string{"Мочилки","Дешевка","Хачик","Хотелово"}
		case 17: forms = [CityForms]string{"Лобки","Ломка","Ишак","Бухалово"}
		case 18: forms = [CityForms]string{"Пупки","Блоха","Лобок","Лобково"}
		case 19: forms = [CityForms]string{"Кобеляки","Сувалка","Пупок","Какино"}
		case 20: forms = [CityForms]string{"Бздюли","Лужа","Кобеляка","Отхожее"}
		case 21: forms = [CityForms]string{"Бобрики","Мочилка","Трус","Хренище "}
		case 22: forms = [CityForms]string{"Мусорки","Бздюля","Дурак","Матюково"}
		case 23: forms = [CityForms]string{"Кончинки","Жаба","Хрен","Пьянкино"}
		case 24: forms = [CityForms]string{"Бугры","Голодранка","Свин","Свинорье"}
		case 25: forms = [CityForms]string{"Лохи","Хотелка","Мухоед","Матюково"}
	}
	return forms[form]
}