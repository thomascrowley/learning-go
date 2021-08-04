package main

import "fmt"

func main() {

	var publisher, writer, artist, title string
	var year, pageNumber int
	var grade float32

	title = "Mr. GotToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5

	printDetails(title, writer, artist, publisher, year, pageNumber, grade)

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	publisher = "DizzyBooks Publishing Inc."
	year = 2013
	pageNumber = 160
	grade = 9.0

	printDetails(title, writer, artist, publisher, year, pageNumber, grade)

}

func printDetails(title string, writer string, artist string, publisher string, year int, pageNumber int, grade float32) {

	fmt.Println(title, "written by", writer, "drawn by", artist, ". Published in", year, "by", publisher, ". It has", pageNumber, "pages and a score of", grade)
}
