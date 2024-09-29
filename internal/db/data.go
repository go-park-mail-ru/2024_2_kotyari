package db

// Инициализация продуктов в базе данных
var productsData = map[string]Product{
	"1": {
		CurrentPrice:     "2473",
		OldPrice:         "7632",
		Discount:         "77",
		Image:            "src/assets/img/static/1.webp",
		Description:      "Кроссовки ASICS Модная обувь.",
		ShortDescription: "Кроссовки ASICS.",
		URL:              "/catalog/product/1",
		Currency:         "Р",
	},
	"2": {
		CurrentPrice:     "8099",
		OldPrice:         "14990",
		Discount:         "46",
		Description:      "Lydsto Робот-пылесос G1, белый.",
		ShortDescription: "Lydsto Робот-пылесос G1, белый.",
		Image:            "src/assets/img/static/2.webp",
		URL:              "/catalog/product/2",
		Currency:         "Р",
	},
	"3": {
		CurrentPrice:     "31 513",
		Description:      "Посудомоечная машина встраиваемая Kuppersberg GSM 4574 (Модификация 2024 года).",
		ShortDescription: "Посудомоечная машина встраиваемая.",
		Image:            "src/assets/img/static/3.webp",
		URL:              "/catalog/product/3",
		Currency:         "Р",
	},
	"4": {
		CurrentPrice:     "7999",
		Description:      "Капсулы для стирки Персил Power Caps Color 4в1, 14 шт х 4 упаковки, 56 шт, для цветного белья.",
		ShortDescription: "Капсулы для стирки Персил Power Caps Color 4в1",
		Image:            "src/assets/img/static/4.webp",
		URL:              "/catalog/product/4",
		Currency:         "Р",
	},
	"5": {
		CurrentPrice:     "127 977",
		OldPrice:         "174 444",
		Discount:         "30",
		Description:      "Vivo Смартфон X100 Ultra CN 16/512 ГБ, черный.",
		ShortDescription: "Vivo Смартфон X100 Ultra CN 16/512 ГБ, черный.",
		Image:            "src/assets/img/static/5.webp",
		URL:              "/catalog/product/5",
		Currency:         "Р",
	},
	"6": {
		CurrentPrice:     "12999",
		Description:      "Лонгслив Cave",
		ShortDescription: "Лонгслив Cave",
		Image:            "src/assets/img/static/6.webp",
		URL:              "/catalog/product/6",
		Currency:         "Р",
	},
	"7": {
		CurrentPrice:     "999",
		Description:      "Кабель для зарядки USB-C длиной 1.5 метра.",
		ShortDescription: "Кабель для зарядки USB-C, 1.5 метра.",
		Image:            "src/assets/img/static/7.webp",
		URL:              "/catalog/product/7",
		Currency:         "Р",
	},
	"8": {
		CurrentPrice:     "344",
		OldPrice:         "572",
		Discount:         "39",
		Description:      "Туалетная бумага Papia Bali flower белая 3 слоя 12 рулонов",
		ShortDescription: "Туалетная бумага Papia Bali flower белая 3 слоя 12 рулонов",
		Image:            "src/assets/img/static/8.webp",
		URL:              "/catalog/product/8",
		Currency:         "Р",
	},
	"9": {
		CurrentPrice: "19999",
		OldPrice:     "24999",
		Discount:     "20",
		Description: "Acer Extensa 15 EX215-54-510N Ноутбук 15.6\", Intel Core " +
			"i7-1135G7, RAM 8 ГБ, SSD 512 ГБ, Intel UHD Graphics,",
		ShortDescription: "Ноутбук Acer Extensa 15",
		Image:            "src/assets/img/static/9.webp",
		URL:              "/catalog/product/9",
		Currency:         "Р",
	},
	"10": {
		CurrentPrice:     "799",
		Description:      "Чехол для смартфона из мягкого силикона.",
		ShortDescription: "Чехол для смартфона, силикон.",
		Image:            "src/assets/img/static/10.webp",
		URL:              "/catalog/product/10",
		Currency:         "Р",
	},
	"11": {
		CurrentPrice:     "11999",
		Description:      "Зеркальный фотоаппарат Canon EOS 80D",
		ShortDescription: "Зеркальный фотоаппарат Canon EOS 80D",
		Image:            "src/assets/img/static/11.webp",
		URL:              "/catalog/product/11",
		Currency:         "Р",
	},
	"12": {
		CurrentPrice:     "2999",
		Description:      "Беспроводная Gsou V4 портативная (мобильная) музыкальная колонка 5 Вт",
		ShortDescription: "Bluetooth колонка беспроводная Gsou V4",
		Image:            "src/assets/img/static/12.webp",
		URL:              "/catalog/product/12",
		Currency:         "Р",
	},
	"13": {
		CurrentPrice:     "49999",
		Description:      "Игровая консоль PlayStation 5 Blu-Ray Edition",
		ShortDescription: "Игровая консоль PlayStation",
		Image:            "src/assets/img/static/13.webp",
		URL:              "/catalog/product/13",
		Currency:         "Р",
	},
	"14": {
		CurrentPrice:     "1199",
		Description:      "Микрофон для компьютера игровой для стрима Vita Musica",
		ShortDescription: "Микрофон для компьютера игровой для стрима Vita Musica",
		Image:            "src/assets/img/static/14.webp",
		URL:              "/catalog/product/14",
		Currency:         "Р",
	},
	"15": {
		CurrentPrice:     "6150",
		Description:      "Microsoft Геймпад Xbox Series, Bluetooth, белый",
		ShortDescription: "Microsoft Геймпад Xbox Series",
		Image:            "src/assets/img/static/15.webp",
		URL:              "/catalog/product/15",
		Currency:         "Р",
	},
	"16": {
		CurrentPrice:     "2299",
		Description:      "Xiaomi Внешний аккумулятор беспроводная зарядка, 10000 мАч, черный",
		ShortDescription: "Внешний аккумулятор Xiaomi 10000 мАч",
		Image:            "src/assets/img/static/16.webp",
		URL:              "/catalog/product/16",
		Currency:         "Р",
	},
	"17": {
		CurrentPrice:     "19999",
		Description:      "Портативный проектор Wanbo Projector T6R Max",
		ShortDescription: "Портативный проектор Wanbo Projector T6R Max",
		Image:            "src/assets/img/static/17.webp",
		URL:              "/catalog/product/17",
		Currency:         "Р",
	},
	"18": {
		CurrentPrice:     "2640",
		Description:      "Графический планшет с 8192 уровнями нажатия и стилусом.",
		ShortDescription: "Графический планшет со стилусом.",
		Image:            "src/assets/img/static/18.webp",
		URL:              "/catalog/product/18",
		Currency:         "Р",
	},
	"19": {
		CurrentPrice:     "57000",
		OldPrice:         "7632",
		Discount:         "31",
		Description:      "Автономный VR шлем очки виртуальной реальности Oculus Quest 3 128 GB (Meta Quest)",
		ShortDescription: "Очки виртуальной реальности Oculus Quest 3 128 GB",
		Image:            "src/assets/img/static/19.webp",
		URL:              "/catalog/product/19",
		Currency:         "Р",
	},
	"20": {
		CurrentPrice:     "11358",
		Description:      "Видеорегистратор Fujida Zoom Smart S WiFi",
		ShortDescription: "Видеорегистратор Fujida Zoom Smart S WiFi",
		Image:            "src/assets/img/static/20.webp",
		URL:              "/catalog/product/20",
		Currency:         "Р",
	},
	"21": {
		CurrentPrice:     "2473",
		OldPrice:         "7632",
		Discount:         "67",
		Image:            "src/assets/img/static/21.webp",
		Description:      "Встраиваемый электрический духовой шкаф Indesit IBFTE 2430 BL, черный",
		ShortDescription: "Встраиваемый электрический духовой шкаф Indesit IBFTE",
		URL:              "/catalog/product/21",
		Currency:         "Р",
	},
	"22": {
		CurrentPrice:     "7799",
		OldPrice:         "12999",
		Discount:         "40",
		Description:      "Hartens 24\" Монитор HTM24C165, черный",
		ShortDescription: "Hartens 24\" Монитор HTM24C165",
		Image:            "src/assets/img/static/22.webp",
		URL:              "/catalog/product/22",
		Currency:         "Р",
	},
	"23": {
		CurrentPrice:     "3790",
		Description:      "Сухой корм Whiskas® для кошек «Подушечки с паштетом, Аппетитный обед с говядиной», 13.8кг",
		ShortDescription: "Сухой корм Whiskas® для кошек.",
		Image:            "src/assets/img/static/23.webp",
		URL:              "/catalog/product/23",
		Currency:         "Р",
	},
	"24": {
		CurrentPrice:     "1325",
		Description:      "JOONIES Premium Soft Подгузники, размер M (6-11 кг), 58 шт.",
		ShortDescription: "JOONIES Premium Soft Подгузники.",
		Image:            "src/assets/img/static/24.webp",
		URL:              "/catalog/product/24",
		Currency:         "Р",
	},
	"25": {
		CurrentPrice:     "742",
		OldPrice:         "1240",
		Discount:         "40",
		Description:      "Влажный корм SHEBA НАТУРАЛЬНАЯ КОЛЛЕКЦИЯ для кошек, утка с добавлением яблок 28шт x 75г",
		ShortDescription: "Влажный корм SHEBA НАТУРАЛЬНАЯ КОЛЛЕКЦИЯ для кошек.",
		Image:            "src/assets/img/static/25.webp",
		URL:              "/catalog/product/25",
		Currency:         "Р",
	},
	"26": {
		CurrentPrice:     "47999",
		Description:      "Стайлер DYSON HS05 Long Prussian Blue",
		ShortDescription: "Стайлер DYSON HS05",
		Image:            "src/assets/img/static/26.webp",
		URL:              "/catalog/product/26",
		Currency:         "Р",
	},
	"27": {
		CurrentPrice:     "9724",
		Description:      "Беспроводной пылесос Tefal Air Force Light TY6545RH, черный",
		ShortDescription: "Беспроводной пылесос Tefal Air Force",
		Image:            "src/assets/img/static/27.webp",
		URL:              "/catalog/product/27",
		Currency:         "Р",
	},
	"28": {
		CurrentPrice:     "628",
		Description:      "Футболка ELIZA Art Хит",
		ShortDescription: "Футболка ELIZA Art Хит",
		Image:            "src/assets/img/static/28.webp",
		URL:              "/catalog/product/28",
		Currency:         "Р",
	},
	"29": {
		CurrentPrice:     "16454",
		OldPrice:         "21999",
		Discount:         "25",
		Description:      "HUAWEI Смартфон nova Y91 8/128 ГБ, черный",
		ShortDescription: "HUAWEI Смартфон nova Y91 8/128 ГБ",
		Image:            "src/assets/img/static/29.webp",
		URL:              "/catalog/product/29",
		Currency:         "Р",
	},
	"30": {
		CurrentPrice: "5569",
		OldPrice:     "7199",
		Discount:     "22",
		Description: "Паровой утюг Tefal Easygliss Plus FV5715E0, с " +
			"автоотключением, защитой от накипи, большим резервуаром для воды, автоматической настройкой пара, 2400 Вт",
		ShortDescription: "Паровой утюг Tefal Easygliss Plus FV5715E0",
		Image:            "src/assets/img/static/30.webp",
		URL:              "/catalog/product/30",
		Currency:         "Р",
	},
	"31": {
		CurrentPrice:     "10080",
		Description:      "Вертикальный пылесос TY6545RH",
		ShortDescription: "Вертикальный пылесос TY6545RH",
		Image:            "src/assets/img/static/31.webp",
		URL:              "/catalog/product/31",
		Currency:         "Р",
	},
	"32": {
		CurrentPrice:     "4349",
		Description:      "Тостер с функцией размораживания и подогрева Tefal Express Metal TT365031",
		ShortDescription: "Тостер Tefal Express Metal",
		Image:            "src/assets/img/static/32.webp",
		URL:              "/catalog/product/32",
		Currency:         "Р",
	},
	"33": {
		CurrentPrice:     "30171",
		OldPrice:         "80863",
		Discount:         "63",
		Description:      "Xiaomi 11Ultra Global 8/256 ГБ, белый",
		ShortDescription: "Смартфон Xiaomi 11 Ultra",
		Image:            "src/assets/img/static/33.webp",
		URL:              "/catalog/product/33",
		Currency:         "Р",
	},
	"34": {
		CurrentPrice: "22887",
		Description: "Simfer духовой шкаф встраиваемый / 5 режимов работы, верхний и нижний нагрев, " +
			"конвекция / таймер + часы / объем 58 литров / 2-ое стекло дверцы / подсветка / противень " +
			"/ хромированная решетка ",
		ShortDescription: "Simfer духовой шкаф встраиваемый",
		Image:            "src/assets/img/static/34.webp",
		URL:              "/catalog/product/34",
		Currency:         "Р",
	},
	"35": {
		CurrentPrice:     "9999",
		Description:      "Наушники Marshall Major IV, черные",
		ShortDescription: "Наушники Marshall Major IV",
		Image:            "src/assets/img/static/35.webp",
		URL:              "/catalog/product/35",
		Currency:         "Р",
	},
	"36": {
		CurrentPrice:     "2594",
		Description:      "Xiaomi беспроводной паровой утюг Lofans Iron YD-012V, фиолетовый.",
		ShortDescription: "Xiaomi беспроводной паровой утюг",
		Image:            "src/assets/img/static/36.webp",
		URL:              "/catalog/product/36",
		Currency:         "Р",
	},
	"37": {
		CurrentPrice:     "8912",
		Description:      "SHOWJI Смартфон LLLS19 Pro-WE-01 Global 16/512 ГБ, белый, прозрачный",
		ShortDescription: "SHOWJI Смартфон LLLS19",
		Image:            "src/assets/img/static/37.webp",
		URL:              "/catalog/product/37",
		Currency:         "Р",
	},
	"38": {
		CurrentPrice:     "18699",
		OldPrice:         "24999",
		Discount:         "27",
		Description:      "HUAWEI Умные часы GT 3 Pro, 46mm, Black",
		ShortDescription: "HUAWEI Умные часы",
		Image:            "src/assets/img/static/38.webp",
		URL:              "/catalog/product/38",
		Currency:         "Р",
	},
	"39": {
		CurrentPrice:     "43",
		Description:      "Смесь кисломолочная Агуша 2 3.4% 200мл/204г с 6 месяцев\n",
		ShortDescription: "Смесь кисломолочная Агуша",
		Image:            "src/assets/img/static/39.webp",
		URL:              "/catalog/product/39",
		Currency:         "Р",
	},
	"40": {
		CurrentPrice: "3329",
		OldPrice:     "6500",
		Discount:     "49",
		Description: "Насадки для электрической зубной щетки Philips Sonicare ProResult HX6014/07, " +
			"для эффективного удаления налёта, 4 шт",
		ShortDescription: "Насадки для электрической зубной щетки Philips Sonicare",
		Image:            "src/assets/img/static/40.webp",
		URL:              "/catalog/product/40",
		Currency:         "Р",
	},
}

var usersData = map[string]User{
	"user@example.com": {
		Username: "Goshanchik",
		Password: "Password123@",
	},
	"user1@example.com": {
		Username: "Igorechik",
		Password: "Password124@",
	},
}
