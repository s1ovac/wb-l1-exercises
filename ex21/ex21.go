package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.
func main() {
	xml := &XMLDocument{} // Сервис, который работает только с XML
	client := &Client{}   // Клиент, который прочитает XML файл
	client.SendXML(xml)
	json := &JSONDocument{} // Сервис, который работает только с JSON
	jsonAdapter := &JSONAdaptee{
		XMLtoJSON: json,
	}
	client.SendXML(jsonAdapter) // Теперь наш тип удовлетворяет интерфейсу Sender
}

// Интерфейс сендер отправляет документ в XML формате
type Sender interface {
	SendXML()
}

// Сервис, который отправляет данные в XML формате
type XMLDocument struct {
}

func (x *XMLDocument) SendXML() {
	fmt.Println("Send document in XML format")
}

// Сервис, который отправляет данные в JSON формате
type JSONDocument struct {
}

func (j *JSONDocument) SendJSON() {
	fmt.Println("Send document in JSON format")
}

// Адаптер, который переводит xml в json формат
type JSONAdaptee struct {
	XMLtoJSON *JSONDocument
}

func (j *JSONAdaptee) SendXML() {
	fmt.Println("Converting XML to JSON...")
	j.XMLtoJSON.SendJSON()
}

type Client struct {
}

func (c *Client) SendXML(sender Sender) {
	fmt.Println("Reading XML format...")
	sender.SendXML()
}
