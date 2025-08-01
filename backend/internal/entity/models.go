package entity

type Car struct {
    ID     string `json:"id"`
    Image  string `json:"image"`
    Title  string `json:"title"`
    Text   string `json:"text"`
    Prices []int  `json:"prices"`
    Brand  string `json:"brand"`
}

type Order struct {
    Car   string `json:"car"`
    Name  string `json:"name"`
    Phone string `json:"phone"`
}

var Cars = []Car{
	{
		ID:     "1",
		Image:  "/images/cars/1.png",
		Title:  "BMW M4 Competition",
		Text:   "Идеальный баланс скорости и стиля. BMW M4 Competition — 510 л.с. и драйв, созданный для покорения трасс и городских улиц.",
		Prices: []int{1450, 1300, 1100},
		Brand:  "BMW",
	},
	{
		ID:     "2",
		Image:  "/images/cars/2.png",
		Title:  "BMW M5",
		Text:   "Бизнес-класс с душой гонщика. BMW M5: 600 л.с., интеллектуальный полный привод и комфорт для самых требовательных.",
		Prices: []int{1600, 1450, 1250},
		Brand:  "BMW",
	},
	{
		ID:     "3",
		Image:  "/images/cars/3.png",
		Title:  "Lamborghini Huracan Spyder Green",
		Text:   "Воплощение скорости и страсти. Зеленый Lamborghini Huracan Spyder — мощь 640 л.с. и открытый верх для ярких приключений.",
		Prices: []int{3200, 2900, 2600},
		Brand:  "Lamborghini",
	},
	{
		ID:     "4",
		Image:  "/images/cars/4.png",
		Title:  "Ferrari F8 Spider",
		Text:   "Мечта на колесах. Ferrari F8 Spider: 720 л.с., аэродинамика F1 и открытая кабина для тех, кто живет на полной скорости.",
		Prices: []int{3500, 3200, 2900},
		Brand:  "Ferrari",
	},
	{
		ID:     "5",
		Image:  "/images/cars/5.png",
		Title:  "Porsche 911 Targa 4S Yellow",
		Text:   "Элегантная мощь в ярком цвете. Porsche 911 Targa 4S: полный привод, 450 л.с. и уникальный стиль Targa",
		Prices: []int{1800, 1650, 1450},
		Brand:  "Porsche",
	},
	{
		ID:     "6",
		Image:  "/images/cars/6.png",
		Title:  "Mercedes SL 55 AMG",
		Text:   "Классика спортивного шика. Mercedes SL 55 AMG: роскошь, кабриолет и мощь 469 л.с. для незабываемых поездок.",
		Prices: []int{1700, 1550, 1350},
		Brand:  "Mercedes",
	},
	{
		ID:     "7",
		Image:  "/images/cars/7.png",
		Title:  "BMW Z4",
		Text:   "Компактный кабриолет для стильных путешествий. BMW Z4: идеальный союз маневренности, мощности и элегантного дизайна.",
		Prices: []int{1200, 1100, 950},
		Brand:  "BMW",
	},
	{
		ID:     "8",
		Image:  "/images/cars/8.png",
		Title:  "Mercedes C180 Convertible",
		Text:   "Идеальный выбор для солнечного дня. Mercedes C180 Convertible: кабриолет для легкого и комфортного вождения.",
		Prices: []int{1000, 900, 800},
		Brand:  "Mercedes",
	},
	{
		ID:     "9",
		Image:  "/images/cars/9.png",
		Title:  "Chevrolet Corvette Orange",
		Text:   "Яркий, мощный, незабываемый. Chevrolet Corvette в оранжевом цвете: 495 л.с. и стиль, который говорит сам за себя.",
		Prices: []int{1400, 1250, 1100},
		Brand:  "Chevrolet",
	},
	{
		ID:     "10",
		Image:  "/images/cars/10.png",
		Title:  "Audi R8 Blue",
		Text:   "Суперкар, созданный для влюбленных в скорость. Audi R8 Blue: полный привод, 570 л.с. и потрясающий дизайн.",
		Prices: []int{2000, 1850, 1600},
		Brand:  "Audi",
	},
	{
		ID:     "11",
		Image:  "/images/cars/11.png",
		Title:  "Chevrolet Corvette White",
		Text:   "Классическая мощь в изысканном цвете. Chevrolet Corvette White: мощность, динамика и стиль для настоящих ценителей.",
		Prices: []int{1350, 1200, 1000},
		Brand:  "Chevrolet",
	},
	{
		ID:     "12",
		Image:  "/images/cars/12.png",
		Title:  "Ford Mustang Convertible Black",
		Text:   "Легенда в открытом формате. Ford Mustang Convertible Black: 450 л.с., культовый стиль и свобода под открытым небом.",
		Prices: []int{1250, 1150, 1000},
		Brand:  "Ford",
	},
}
