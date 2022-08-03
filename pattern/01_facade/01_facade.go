package main

import (
	"fmt"
	"log"
)

/*
Фасад — это структурный паттерн проектирования,
который предоставляет простой интерфейс к сложной системе классов, библиотеке или фреймворку.
«Фасад» — некоторый объект,
аккумулирующий в себе высокоуровневый набор операций для работы с некоторой сложной подсистемой.
Клиент, при этом, не лишается более низкоуровневого доступа к классам подсистемы.
Фасад упрощает выполнение некоторых операций с подсистемой, но не навязывает клиенту свое использование.
Плюсами такого подхода является:
- Упрощение работы клиента с подсистемой - меньше кода, меньше ошибок, быстрее разработка.
- Уменьшении зависимости от подсистемы - проще внести изменения, проще тестировать.
- Упрощение внешней документации - упрощение работы с подсистемой для клиента - проще клиентская документация
Минусами подхода является:
- Требуется дополнительная реализация необходимых интерфейсов - дополнительная разработка.
- Нужно хорошо продумать реализуемый набор интерфейсов для клиента, чтобы вся функциональность, ему
необходимая, была у него доступна (при доработках подсистемы нужно поддерживать и фасад).
- Фасад может стать божествееным объектом
*/

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{balance: 0}
}

func (w *wallet) debitBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
}

type user struct {
	name string
}

func newUser(name string) *user {
	return &user{name: name}
}

func (u *user) checkUser(userName string) error {
	if u.name != userName {
		return fmt.Errorf("user name is incorrect")
	}
	fmt.Println("User verified")
	return nil
}

type walletFacade struct {
	wallet *wallet
	user   *user
}

func newWalletFacade(user string) *walletFacade {
	return &walletFacade{wallet: newWallet(), user: newUser(user)}
}

func (f *walletFacade) addMoney(user string, money int) error {
	err := f.user.checkUser(user)
	if err != nil {
		return err
	}
	f.wallet.debitBalance(money)
	fmt.Println(f.wallet.balance)
	return nil
}

func main() {
	facade := newWalletFacade("Bob")
	err := facade.addMoney("Bob", 1250)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}