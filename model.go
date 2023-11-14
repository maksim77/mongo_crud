package main

import "fmt"

type Book struct {
	Title  string `bson:"title"`
	Author string `bson:"author"`
	Year   int    `bson:"year"`
	Rating int    `bson:"rating"`
}

func (b Book) String() string {
	return fmt.Sprintf("{\n\tTtile: %s,\n\tAuthor: %s\n\tYear: %d\n\tRating: %d\n}\n", b.Title, b.Author, b.Year, b.Rating)
}

var book Book = Book{
	Title:  "gRPC: запуск и эксплуатация облачных приложений. Go и Java для Docker и Kubernetes",
	Author: "Касун Индрасири",
	Year:   2020,
}

var books []Book = []Book{
	{
		Title:  "Go: идиомы и паттерны проектирования",
		Author: "Боднер Джон",
		Year:   2022,
	},
	{
		Title:  "Высоконагруженные приложения. Программирование, масштабирование, поддержка",
		Author: "Клеппман Мартин",
		Year:   2021,
	},
}
