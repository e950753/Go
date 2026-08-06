package main

import "fmt"

type Film struct {
	Name string
	PublicationDate int
	Genre []string
	Rating float64
}

type Account struct {
	Id int
	Owner string
	Balance float64
}

func main() {
	film := Film{
		Name: "The Matrix",
		PublicationDate: 1999,
		Genre: []string{"Action", "Sci-Fi"},
		Rating: 8.7,
	}
	fmt.Println(film.Name, "(", film.PublicationDate, ") - жанры: ", film.Genre, " - рейтинг: ", film.Rating)

	account := Account{
		Id: 1,
		Owner: "John Doe",
		Balance: 500,
	}
	fmt.Println(account.Balance)
	Deposit(&account, 500)
	fmt.Println(account.Balance)
}

func Deposit(acc *Account, amount float64) {
	acc.Balance += amount
}