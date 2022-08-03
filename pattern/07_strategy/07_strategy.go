package main

/*
Паттерн Стратегия используется, когда есть семейство некоторых схожих алгоритмов, которые часто
изменяются или расширяются. Тогда, согласно паттерну Стратегия, каждый из этих алгоритмов помещается в свой
собственный класс (их можно назвать конкретными стратегиями), эти классы реализуют один и тот же интерфейс.
Затем некоторый основной класс, вместо того, чтобы самому реализовывать алгоритм, будет ссылаться на один
из классов-стратегий и делегировать реализацию алгоритма этому классу.
Плюсы:
- изолирует код алгоритмов от основного класса, что позволяет не трогать код основного класса при изменении
или добавлении новых алгоритмов;
- так как классы-стратегии реализуют один интерфейс, с помощью сеттера можно изменять
используемый в данный момент алгоритм;
- вместо наследования используется композиция.
Минусы:
- из-за дополнительных классов усложняется код программы;
- клиент должен знать различия между стратегиями, чтобы использовать их в коде.
*/

func processOrder(product string, payment Payment) {
	// ... implementation
	err := payment.Pay()
	if err != nil {
		return
	}
}

// ----

type Payment interface {
	Pay() error
}

// ----

type cardPayment struct {
	cardNumber, cvv string
}

func NewCardPayment(cardNumber, cvv string) Payment {
	return &cardPayment{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

func (p *cardPayment) Pay() error {
	// ... implementation
	return nil
}

type payPalPayment struct {
}

func NewPayPalPayment() Payment {
	return &payPalPayment{}
}

func (p *payPalPayment) Pay() error {
	// ... implementation
	return nil
}

type qiwiPayment struct {
}

func NewQIWIPayment() Payment {
	return &qiwiPayment{}
}

func (p *qiwiPayment) Pay() error {
	// ... implementation
	return nil
}

func main() {
	product := "vehicle"
	payWay := 3

	var payment Payment
	switch payWay {
	case 1:
		payment = NewCardPayment("12345", "12345")
	case 2:
		payment = NewPayPalPayment()
	case 3:
		payment = NewQIWIPayment()
	}

	processOrder(product, payment)
}