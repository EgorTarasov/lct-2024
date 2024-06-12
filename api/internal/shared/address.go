package shared // Адрес: Байкальская ул., д.43, стр.1 ; Административный округ (ТП): Восточный ; municipal_district: район Гольяново
import (
	"strings"
)

// PAddress адрес объекта
//
// адрес разбитый на столбцы address_registry для дальнейшего поиска
// L1TYPE: примеры значений,
// _id: "земельный участок",
// count: 192177
// example поселение Сосенское, деревня Летово, Летовская улица, земельный участок 113А
// _id: "дом",
// count: 230797
// _id: "сооружение",
// count: 6306
// Российская Федерация, город Москва, внутригородская территория поселение Марушкинское, деревня Акиньшино, Солнечная улица, сооружение 9, строение 3
// _id: "домовладение",
// count: 2081
// Российская Федерация, город Москва, внутригородская территория муниципальный округ Северное Бутово, бульвар Дмитрия Донского, домовладение 17А
// _id: null,
// count: 4160
// _id: "владение",
// город Москва, Балаклавский проспект, владение 28, строение 1
// count: 55774
// _id: "участок",
// count: 45
// город Москва, поселение Первомайское, деревня Фоминское, Северная улица, участок 107
//
// L2TYPE: примеры значений,
// корпус
//
// L3TYPE: примеры значений,
// l3_types
// сооружение
// город Москва, поселение Десеновское, 5-я Нововатутинская улица, владение 1, сооружение 2
// строение
// Российская Федерация, город Москва, внутригородская территория муниципальный округ Солнцево, Боровский проезд, дом 17, строение 2.
type PAddress struct {
	P7      string `json:"p7"`      // Улица
	L1TYPE  string `json:"L1Type"`  // Тип адреса
	L1Value string `json:"L1Value"` // Номер дома
	L2TYPE  string `json:"L2Type"`  // Тип адреса
	L2Value string `json:"L2Value"` // Строение
	L3TYPE  string `json:"L3Type"`  // Тип адреса
	L3Value string `json:"L3Value"` // Корпус
}

func (p PAddress) String() string {
	return p.P7 + ", " + p.L1Value + ", " + p.L2Value + ", " + p.L3Value
}

var p7StopWords = []string{
	"ул.",
	"пр.",
	"просп.",
	"наб.",
	"шоссе",
	"пл.",
	"аллея",
	"пер.",
}

var l1StopWords = []string{
	"д.",
	"вл.",
	"з/у",
}

// ParseAddress разбивает строку адреса на составляющие
//
// district: район Гольяново, ADM: Восточный, address: Байкальская ул., д.37
// inputStr : Восточный, район Гольяново, Байкальская ул., д.37.
func ParseAddress(input string) PAddress {
	var result PAddress
	parts := strings.Split(input, ", ")
	for i, p := range parts {
		parts[i] = strings.TrimSpace(p)
		// if p contains one of p7StopWords
		if result.P7 == "" {
			for _, sw := range p7StopWords {
				if strings.Contains(p, sw) {
					if strings.Contains(p, ".") {
						splited := strings.Split(p, ".")
						if len(splited[0]) > len(splited[1]) {
							result.P7 = splited[0]
						} else {
							result.P7 = splited[1]
						}
					} else {
						result.P7 = p
					}
					result.P7 = strings.TrimSpace(result.P7)
					break
				}
			}
		}
		if result.L1Value == "" {
			for _, sw := range l1StopWords {
				if strings.Contains(p, sw) {
					if strings.Contains(p, ".") {
						result.L1Value = strings.Split(p, ".")[1]
					} else {
						result.L1Value = p
					}
					result.L1Value = strings.TrimSpace(result.L1Value)
					break
				}
			}
		}
		if result.L2Value == "" {
			if strings.Contains(p, "корп.") {
				result.L2Value = strings.Split(p, ".")[1]
			}
			result.L2Value = strings.TrimSpace(result.L2Value)
		}

		if result.L3Value == "" {
			if strings.Contains(p, "стр.") {
				result.L3Value = strings.Split(p, ".")[1]
			}
			result.L3Value = strings.TrimSpace(result.L3Value)
		}
	}
	return result
}
