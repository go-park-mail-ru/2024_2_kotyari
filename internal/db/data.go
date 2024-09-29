package db

// Инициализация продуктов в базе данных
var productsData = map[string]Product{
	"1": {
		CurrentPrice:     "2473",
		OldPrice:         "7632",
		Discount:         "77",
		Image:            "https://sun9-25.userapi.com/impg/dsKDTkLYpWXfVMYj_21Rn7CESXspaL3zrXGF3A/riTPmwVCVaw.jpg?size=750x1000&quality=95&sign=3f49cd35acc30ab4f3dea29e4e0540d6&type=album",
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
		Image:            "https://sun9-27.userapi.com/impg/n4x2LZ7IpCfYgOAYedj3wkDaVS2CF1aATpCVDQ/0D8LB0AiXNs.jpg?size=1000x1000&quality=95&sign=9478ab570b9f6735a2536ec4cabf7777&type=album",
		URL:              "/catalog/product/2",
		Currency:         "Р",
	},
	"3": {
		CurrentPrice:     "31 513",
		Description:      "Посудомоечная машина встраиваемая Kuppersberg GSM 4574 (Модификация 2024 года).",
		ShortDescription: "Посудомоечная машина встраиваемая.",
		Image:            "https://sun9-62.userapi.com/impg/Pn7njR824gsUONsgRhuLCoGhQp1eSwzs21A0JQ/OW-FgCP2ZQU.jpg?size=440x440&quality=95&sign=05605d495f05bf373368e5d31dcf1900&type=album",
		URL:              "/catalog/product/3",
		Currency:         "Р",
	},
	"4": {
		CurrentPrice:     "7999",
		Description:      "Капсулы для стирки Персил Power Caps Color 4в1, 14 шт х 4 упаковки, 56 шт, для цветного белья.",
		ShortDescription: "Капсулы для стирки Персил Power Caps Color 4в1",
		Image:            "https://sun9-68.userapi.com/impg/QYwloSnXfDLLj0jvlgGUjRs3qoQU-y3sfRZ0sA/8JwvXl5S-V4.jpg?size=1000x1000&quality=95&sign=ccaeaff8ef53cdc72133a2406f8c4606&type=album",
		URL:              "/catalog/product/4",
		Currency:         "Р",
	},
	"5": {
		CurrentPrice:     "127 977",
		OldPrice:         "174 444",
		Discount:         "30",
		Description:      "Vivo Смартфон X100 Ultra CN 16/512 ГБ, черный.",
		ShortDescription: "Vivo Смартфон X100 Ultra CN 16/512 ГБ, черный.",
		Image:            "https://sun9-80.userapi.com/impg/VI4aKugj-Q-EmS3xoRU2NR5m3fPHgg_DHRNxaQ/JIJSaE97BPw.jpg?size=540x640&quality=95&sign=18b3c1d6b3f6d65425bfd11fa799d9de&type=album",
		URL:              "/catalog/product/5",
		Currency:         "Р",
	},
	"6": {
		CurrentPrice:     "12999",
		Description:      "Лонгслив Cave",
		ShortDescription: "Лонгслив Cave",
		Image:            "https://sun9-7.userapi.com/impg/H_hFnufaf39UnwPmjsDBfVubEVr0KEVQvivSPQ/hOQU4EEcmEQ.jpg?size=750x1000&quality=95&sign=3316458ee12ef8b56ee560f7d07e2ece&type=album",
		URL:              "/catalog/product/6",
		Currency:         "Р",
	},
	"7": {
		CurrentPrice:     "999",
		Description:      "Кабель для зарядки USB-C длиной 1.5 метра.",
		ShortDescription: "Кабель для зарядки USB-C, 1.5 метра.",
		Image:            "https://sun9-4.userapi.com/impg/WznotV_muk7sFi6WCXXyJtEPZid24LPjn9Ka0w/ucyDoTvM26o.jpg?size=800x800&quality=95&sign=01f4ffea990cdf74780cf078e1d34e25&type=album",
		URL:              "/catalog/product/7",
		Currency:         "Р",
	},
	"8": {
		CurrentPrice:     "344",
		OldPrice:         "572",
		Discount:         "39",
		Description:      "Туалетная бумага Papia Bali flower белая 3 слоя 12 рулонов",
		ShortDescription: "Туалетная бумага Papia Bali flower белая 3 слоя 12 рулонов",
		Image:            "https://sun9-58.userapi.com/impg/mrqp5bYDkHfdp7LeAsK0EiAlO52vpgLt29CsCg/GU6qVL8mHSc.jpg?size=1000x1000&quality=95&sign=8b714f4785fc572475baa3637a9c87e1&type=album",
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
		Image:            "https://sun9-71.userapi.com/impg/3DYwi2Zy4seFbx2NQnRfsST4Z7zVk39YRmAb3Q/DS-UZlT-fK0.jpg?size=1000x1000&quality=95&sign=f11954ea6fbc5cb3fac47ae2dfa700e3&type=album",
		URL:              "/catalog/product/9",
		Currency:         "Р",
	},
	"10": {
		CurrentPrice:     "799",
		Description:      "Чехол для смартфона из мягкого силикона.",
		ShortDescription: "Чехол для смартфона, силикон.",
		Image:            "https://sun9-19.userapi.com/impg/HHsIxJ16kaWMJoUurnCOhMOzLb_cCG6IaEi9ug/pWkjnvIdfho.jpg?size=1000x1000&quality=95&sign=ef2e8d71e4c5c842ad71a0335c6432d0&type=album",
		URL:              "/catalog/product/10",
		Currency:         "Р",
	},
	"11": {
		CurrentPrice:     "11999",
		Description:      "Зеркальный фотоаппарат Canon EOS 80D",
		ShortDescription: "Зеркальный фотоаппарат Canon EOS 80D",
		Image:            "https://sun9-33.userapi.com/impg/LhAqbA8EPqFSzWUQvsS1x8cMnvZDoCxWF9WF9Q/-oeUqfv4YwM.jpg?size=1000x1000&quality=95&sign=7e8ef2774838ae9e64f6da061fd93d72&type=album",
		URL:              "/catalog/product/11",
		Currency:         "Р",
	},
	"12": {
		CurrentPrice:     "2999",
		Description:      "Беспроводная Gsou V4 портативная (мобильная) музыкальная колонка 5 Вт",
		ShortDescription: "Bluetooth колонка беспроводная Gsou V4",
		Image:            "https://sun9-42.userapi.com/impg/BiXNSa0K330IT4MYZB7z5JDVEvA7LoYiN08OYw/VxcltGzt2yE.jpg?size=500x500&quality=95&sign=fb3aab52383db9b607f29a7ba1e52845&type=album",
		URL:              "/catalog/product/12",
		Currency:         "Р",
	},
	"13": {
		CurrentPrice:     "49999",
		Description:      "Игровая консоль PlayStation 5 Blu-Ray Edition",
		ShortDescription: "Игровая консоль PlayStation",
		Image:            "https://sun9-69.userapi.com/impg/HFCSgth1tU6XcOcQOwvdxZEO-t9qYfYfaTuQPw/vdL5sTAdI3Q.jpg?size=1000x1000&quality=95&sign=af12cbd3b6b10b8c15504e904929f37e&type=album",
		URL:              "/catalog/product/13",
		Currency:         "Р",
	},
	"14": {
		CurrentPrice:     "1199",
		Description:      "Микрофон для компьютера игровой для стрима Vita Musica",
		ShortDescription: "Микрофон для компьютера игровой для стрима Vita Musica",
		Image:            "https://sun9-74.userapi.com/impg/7JY7M8-4DS-y-JZxXBTcEfApYoE41IjrjdNneA/Yli42RwvACw.jpg?size=900x900&quality=95&sign=590350111a187a0ea686b675f6521c40&type=album",
		URL:              "/catalog/product/14",
		Currency:         "Р",
	},
	"15": {
		CurrentPrice:     "6150",
		Description:      "Microsoft Геймпад Xbox Series, Bluetooth, белый",
		ShortDescription: "Microsoft Геймпад Xbox Series",
		Image:            "https://sun9-39.userapi.com/impg/pQczyxkuOw7KoQxhWt84Oz8NfDLB7b3RSvN_Gw/FcAe0osO-UU.jpg?size=1000x701&quality=95&sign=d3729a6cc4fababdc2eccfd7ad770e51&type=album",
		URL:              "/catalog/product/15",
		Currency:         "Р",
	},
	"16": {
		CurrentPrice:     "2299",
		Description:      "Xiaomi Внешний аккумулятор беспроводная зарядка, 10000 мАч, черный",
		ShortDescription: "Внешний аккумулятор Xiaomi 10000 мАч",
		Image:            "https://sun9-33.userapi.com/s/v1/ig2/AT0NKZjI4W7JUKLX5GgQyITKTMmUyJ7Ey9ceUgkFVuBZTlcYkSxS5WQnjZNQfx7RUlBPUlTzda_0J0b7T8-JmNsx.jpg?quality=95&as=32x43,48x64,72x96,108x144,160x213,240x320,360x480,480x640,540x720,640x853,720x960,750x1000&from=bu&u=jJQ7wESbKYOuCNKPn74d2PsIdy5oPFZvRSwu5PVdaNA&cs=750x1000",
		URL:              "/catalog/product/16",
		Currency:         "Р",
	},
	"17": {
		CurrentPrice:     "19999",
		Description:      "Портативный проектор Wanbo Projector T6R Max",
		ShortDescription: "Портативный проектор Wanbo Projector T6R Max",
		Image:            "https://sun9-34.userapi.com/impg/VuaBvOlHiCSiW9_Df7d3UyecOiDHF0Iqj5_g5Q/y8iA4Gnykwk.jpg?size=1000x1000&quality=95&sign=163cda1efba31f9ec54f2041c886ff18&type=album",
		URL:              "/catalog/product/17",
		Currency:         "Р",
	},
	"18": {
		CurrentPrice:     "2640",
		Description:      "Графический планшет с 8192 уровнями нажатия и стилусом.",
		ShortDescription: "Графический планшет со стилусом.",
		Image:            "https://sun9-50.userapi.com/impg/lzFA9qKxFDM-iC-KALhjpA7JtASisthbKgnaMg/GAYTNHqKS6c.jpg?size=900x900&quality=95&sign=780096d014fda68c7601d67e0ac4c196&type=album",
		URL:              "/catalog/product/18",
		Currency:         "Р",
	},
	"19": {
		CurrentPrice:     "57000",
		OldPrice:         "7632",
		Discount:         "31",
		Description:      "Автономный VR шлем очки виртуальной реальности Oculus Quest 3 128 GB (Meta Quest)",
		ShortDescription: "Очки виртуальной реальности Oculus Quest 3 128 GB",
		Image:            "https://sun9-80.userapi.com/impg/dOwuQmWMIsYl502R4ld6FXY5iy7A1r193Y549A/39PYI88WbKo.jpg?size=1000x1000&quality=95&sign=2a29b2a93a10dfa21a6856135596cb94&type=album",
		URL:              "/catalog/product/19",
		Currency:         "Р",
	},
	"20": {
		CurrentPrice:     "11358",
		Description:      "Видеорегистратор Fujida Zoom Smart S WiFi",
		ShortDescription: "Видеорегистратор Fujida Zoom Smart S WiFi",
		Image:            "https://sun9-51.userapi.com/impg/tdkql9Kv_rZXehnCqQf1RzXZXjG5jjvD3K4QJA/czzCrT0aEng.jpg?size=1000x1000&quality=95&sign=9850d84551527fb62c999752e8a6ea4c&type=album",
		URL:              "/catalog/product/20",
		Currency:         "Р",
	},
	"21": {
		CurrentPrice:     "2473",
		OldPrice:         "7632",
		Discount:         "67",
		Image:            "https://sun9-3.userapi.com/impg/f1vae26hJFooObPSTkVAzYEU6EtCGEYArbzdlg/L2JGSPHMoFY.jpg?size=1000x1000&quality=95&sign=1cd968a862dd10c702e5592ca00ea0e9&type=album",
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
		Image:            "https://sun9-79.userapi.com/impg/kTNYnPvB6UFxX68yGGSCAdXPMPuXwBvRCD6Gdg/49xoAav931E.jpg?size=1000x1000&quality=95&sign=892cfbf6ccba07936f70a3c461db3505&type=album",
		URL:              "/catalog/product/22",
		Currency:         "Р",
	},
	"23": {
		CurrentPrice:     "3790",
		Description:      "Сухой корм Whiskas® для кошек «Подушечки с паштетом, Аппетитный обед с говядиной», 13.8кг",
		ShortDescription: "Сухой корм Whiskas® для кошек.",
		Image:            "https://sun9-69.userapi.com/impg/kTzGHrJG2EyYujdsjkpzlkJgzVcTCyT5RVa9tw/jNSouP9vQOo.jpg?size=1000x1000&quality=95&sign=38d4a259b1dd65f159135b15bdbffd93&type=album",
		URL:              "/catalog/product/23",
		Currency:         "Р",
	},
	"24": {
		CurrentPrice:     "1325",
		Description:      "JOONIES Premium Soft Подгузники, размер M (6-11 кг), 58 шт.",
		ShortDescription: "JOONIES Premium Soft Подгузники.",
		Image:            "https://sun9-30.userapi.com/impg/z2spci3adSXlCbqRGxZXjoGk1gmfaBOlMk3ZqQ/9ieEx54hcFY.jpg?size=1000x1000&quality=95&sign=6bcfd6d8c15a1ddd651d888bf4320bf2&type=album",
		URL:              "/catalog/product/24",
		Currency:         "Р",
	},
	"25": {
		CurrentPrice:     "742",
		OldPrice:         "1240",
		Discount:         "40",
		Description:      "Влажный корм SHEBA НАТУРАЛЬНАЯ КОЛЛЕКЦИЯ для кошек, утка с добавлением яблок 28шт x 75г",
		ShortDescription: "Влажный корм SHEBA НАТУРАЛЬНАЯ КОЛЛЕКЦИЯ для кошек.",
		Image:            "https://sun9-26.userapi.com/impg/o9XnWBK_Kub-tldUmNrnkLtNewtxsAWLNddd-Q/-kYU1WrY0lk.jpg?size=1000x1000&quality=95&sign=56456d55608a8a7a79340cd9a15accbb&type=album",
		URL:              "/catalog/product/25",
		Currency:         "Р",
	},
	"26": {
		CurrentPrice:     "47999",
		Description:      "Стайлер DYSON HS05 Long Prussian Blue",
		ShortDescription: "Стайлер DYSON HS05",
		Image:            "https://sun9-56.userapi.com/impg/FLJ_yQ82yWBXTsGaJs4Yiz01p_m6by1PaChSVw/lsdvcnVABIA.jpg?size=1000x1000&quality=95&sign=1860695bc21ae5a32773383fa0cd8838&type=album",
		URL:              "/catalog/product/26",
		Currency:         "Р",
	},
	"27": {
		CurrentPrice:     "9724",
		Description:      "Беспроводной пылесос Tefal Air Force Light TY6545RH, черный",
		ShortDescription: "Беспроводной пылесос Tefal Air Force",
		Image:            "https://sun9-29.userapi.com/impg/pXsvbC3xCRWjYbR618CTmCt2iyX3dQi1f9MhYA/7h3Lc2ESuYc.jpg?size=1000x1000&quality=95&sign=19e2300d322da211d5b7474a58558d90&type=album",
		URL:              "/catalog/product/27",
		Currency:         "Р",
	},
	"28": {
		CurrentPrice:     "628",
		Description:      "Футболка ELIZA Art Хит",
		ShortDescription: "Футболка ELIZA Art Хит",
		Image:            "https://sun9-43.userapi.com/impg/50ok-umirdL93cIZgNfQCwPfZHDtFvrdk7hkow/Tiv-4UoE5Oc.jpg?size=750x1000&quality=95&sign=1ddfe851b6bae0dd02fef89ba9502a61&type=album",
		URL:              "/catalog/product/28",
		Currency:         "Р",
	},
	"29": {
		CurrentPrice:     "16454",
		OldPrice:         "21999",
		Discount:         "25",
		Description:      "HUAWEI Смартфон nova Y91 8/128 ГБ, черный",
		ShortDescription: "HUAWEI Смартфон nova Y91 8/128 ГБ",
		Image:            "https://sun9-12.userapi.com/impg/IfuzivqSVqhulxCfKOGKeRsNOYvsHaiULrUuWw/kfKnm1tEMBo.jpg?size=1000x1000&quality=95&sign=776621d5eae21736366f5dd0a2d9067c&type=album",
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
		Image:            "https://sun9-40.userapi.com/impg/eU_KoCoxp35UUuB810HuBejn4_zssqnJlE-sYw/szSAxSTAkv0.jpg?size=1000x1000&quality=95&sign=61b0d0f3e7d6a8e8c0369b9b2d1393f5&type=album",
		URL:              "/catalog/product/30",
		Currency:         "Р",
	},
	"31": {
		CurrentPrice:     "10080",
		Description:      "Вертикальный пылесос TY6545RH",
		ShortDescription: "Вертикальный пылесос TY6545RH",
		Image:            "/src/assets/img/static/31.webp",
		URL:              "https://sun9-19.userapi.com/impg/zaTx2ipyBWkFIr1ZEhtnrU6Ld9qJ33B5r2bm0w/3-HbMG7pvX0.jpg?size=264x956&quality=95&sign=c6462ec73d672001248797e366734b51&type=album",
		Currency:         "Р",
	},
	"32": {
		CurrentPrice:     "4349",
		Description:      "Тостер с функцией размораживания и подогрева Tefal Express Metal TT365031",
		ShortDescription: "Тостер Tefal Express Metal",
		Image:            "https://sun9-3.userapi.com/impg/B4DeQocJwiAYtSiU01PmUxEVQU_642764tZlTA/iLVLgmtZCw4.jpg?size=1000x1000&quality=95&sign=16c7e651460f52fe5ef04041173ed6d6&type=album",
		URL:              "/catalog/product/32",
		Currency:         "Р",
	},
	"33": {
		CurrentPrice:     "30171",
		OldPrice:         "80863",
		Discount:         "63",
		Description:      "Xiaomi 11Ultra Global 8/256 ГБ, белый",
		ShortDescription: "Смартфон Xiaomi 11 Ultra",
		Image:            "https://sun9-62.userapi.com/impg/sxYGT-uIAaQrwDjy6HECOf8V1JkLsq90niOU8g/UF-ak4uy9w0.jpg?size=667x1000&quality=95&sign=e7f7c769935f6d2e51fe1ffc86adc54b&type=album",
		URL:              "/catalog/product/33",
		Currency:         "Р",
	},
	"34": {
		CurrentPrice: "22887",
		Description: "Simfer духовой шкаф встраиваемый / 5 режимов работы, верхний и нижний нагрев, " +
			"конвекция / таймер + часы / объем 58 литров / 2-ое стекло дверцы / подсветка / противень " +
			"/ хромированная решетка ",
		ShortDescription: "Simfer духовой шкаф встраиваемый",
		URL:              "/catalog/product/34",
		Currency:         "Р",
		Image:            "https://sun9-3.userapi.com/impg/f1vae26hJFooObPSTkVAzYEU6EtCGEYArbzdlg/L2JGSPHMoFY.jpg?size=1000x1000&quality=95&sign=1cd968a862dd10c702e5592ca00ea0e9&type=album",
	},
	"35": {
		CurrentPrice:     "9999",
		Description:      "Наушники Marshall Major IV, черные",
		ShortDescription: "Наушники Marshall Major IV",
		Image:            "https://sun9-18.userapi.com/impg/GgOkuiP4To8R32R-dnZIe80D35kwiwqQd7aRjQ/pTJxyoFleqo.jpg?size=300x300&quality=95&sign=2586712b62fad48c450fd6be2d628304&type=album",
		URL:              "/catalog/product/35",
		Currency:         "Р",
	},
	"36": {
		CurrentPrice:     "2594",
		Description:      "Xiaomi беспроводной паровой утюг Lofans Iron YD-012V, фиолетовый.",
		ShortDescription: "Xiaomi беспроводной паровой утюг",
		URL:              "/catalog/product/36",
		Currency:         "Р",
		Image:            "https://sun9-34.userapi.com/impg/7OkGqD3CMQp8Q8DKkjM9jGBq_BHk6-UQ9RQ56Q/Us4pUM9hsdI.jpg?size=1000x1000&quality=95&sign=4a8c83952f69fec80270f59dab54fcec&type=album",
	},
	"37": {
		CurrentPrice:     "8912",
		Description:      "SHOWJI Смартфон LLLS19 Pro-WE-01 Global 16/512 ГБ, белый, прозрачный",
		ShortDescription: "SHOWJI Смартфон LLLS19",
		URL:              "/catalog/product/37",
		Currency:         "Р",
		Image:            "https://sun9-20.userapi.com/impg/VUGPRKZ6CfqeP-zJE_aqllJ8S5RBGQIBfBFBiQ/ii3iTeGN4rA.jpg?size=1000x1000&quality=95&sign=5059a25a2603644ae1fdb30aa753b2ce&type=album",
	},
	"38": {
		Image:            "",
		CurrentPrice:     "18699",
		OldPrice:         "24999",
		Discount:         "27",
		Description:      "HUAWEI Умные часы GT 3 Pro, 46mm, Black",
		ShortDescription: "HUAWEI Умные часы",
		URL:              "/catalog/product/38",
		Currency:         "Р",
	},
	"39": {
		Image:            "https://sun9-42.userapi.com/impg/9pxx3Rmdjcx_eaoogf9CGmoao2djhm_V1oFZZw/Jo-oxGIzLoE.jpg?size=1000x1000&quality=95&sign=8a43e587d2d465f9c53415e85fafa0f4&type=album",
		CurrentPrice:     "43",
		Description:      "Смесь кисломолочная Агуша 2 3.4% 200мл/204г с 6 месяцев\n",
		ShortDescription: "Смесь кисломолочная Агуша",
		URL:              "/catalog/product/39",
		Currency:         "Р",
	},
	"40": {
		Image:        "https://sun9-59.userapi.com/impg/VeDofQAc3DzoNSxkFYohrmq0oCmFV75BuVQe7A/sQMZ-0EQcAU.jpg?size=962x1000&quality=95&sign=72f009a69723db7490d69c796ca5d6c2&type=album",
		CurrentPrice: "3329",
		OldPrice:     "6500",
		Discount:     "49",
		Description: "Насадки для электрической зубной щетки Philips Sonicare ProResult HX6014/07, " +
			"для эффективного удаления налёта, 4 шт",
		ShortDescription: "Насадки для электрической зубной щетки Philips Sonicare",
		URL:              "/catalog/product/40",
		Currency:         "Р",
	},
}

var usersData = map[string]User{
	"user@example.com": {
		Username: "Goshanchik",
		Password: "gbHWrVy4JEmoO06xZa4Z3h/LnkSFl0wzkJNtDXXLmq9pU8LRhOhRQRnZ79AdABaK",
	},
	"user1@example.com": {
		Username: "Igorechik",
		Password: "Password124@",
	},
}
