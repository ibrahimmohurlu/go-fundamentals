package quote // this is the quote package

import "math/rand"

const quote_length int = 5 // this constant is not exported const name doesnt start with capital letter.

func getRandomIndex() int { // this function is not exported, function name doesnt start with capital letter.
	return rand.Intn(5)
}

func RandomQuote() string { // this function is exported, function name start with capital letter.
	var quotes = [quote_length]string{
		"Be yourself; everyone else is already taken. \n -Oscar Wilde",
		"So many books, so little time. \n -Frank Zappa",
		"A room without books is like a body without a soul.\n -Marcus Tullius Cicero",
		"You only live once, but if you do it right, once is enough.\n -Mae West",
		"Eğer bir gün benim sözlerim bilimle ters düşerse bilimi seçin.\n -M.Kemal Atatürk",
	}

	return quotes[getRandomIndex()]
}
