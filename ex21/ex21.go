package main

func main() {

}

type Target interface { // Интерфейс Target, описывающий целевой интерфейс (тот интерфейс с которым наша система хотела бы работать)
	Request() string
}

type Adaptee struct { //тип, который наша система должна адаптировать под себя
}

func (a *Adaptee) SpecificRequest() string {
	return "Request"
}

func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

type Adapter struct { // Адаптер реализующий целевой интерфейс
	*Adaptee
}

func (a *Adapter) Request() string {
	return a.SpecificRequest()
}
