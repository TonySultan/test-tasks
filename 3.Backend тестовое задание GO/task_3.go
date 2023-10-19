package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector()

	file, err := os.Create("Top Instagram Influencers in Russia.csv")
	if err != nil {
		log.Fatalf("Ошибка при создании файла CSV: %v", err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	headers := []string{"Рейтинг", "Имя пользователя", "Ник", "Категория", "Страна", "Подписчики", "Средние лайки", "Средние комментарии"}
	csvWriter.Write(headers)

	c.OnHTML(".ranking-card .table .row", func(e *colly.HTMLElement) {

		position := e.ChildText(".row-cell.rank > span")
		username := e.ChildText(".contributor .contributor__title")
		nikname := e.ChildText(".contributor .contributor__name-content")
		category := e.ChildText(".category")
		country := e.ChildText(".audience")
		followers := e.ChildText(".subscribers")
		avgLikes := e.ChildText(".authentic")
		avgComments := e.ChildText(".engagement")

		data := []string{position, username, nikname, category, country, followers, avgLikes, avgComments}
		csvWriter.Write(data)

	})

	url := "https://hypeauditor.com/top-instagram-all-russia/"
	err = c.Visit(url)
	if err != nil {
		log.Fatalf("Ошибка при посещении страницы: %v", err)
	}

	fmt.Printf("Парсинг завершен.\n")
}
