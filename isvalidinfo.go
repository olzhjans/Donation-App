package main

import (
	"fmt"
	"net/mail"
	"regexp"
)

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isValidPhoneNumber(phoneNumber string) bool {
	// Паттерн для казахстанского телефонного номера
	// Формат: +7XXXXXXXXXX или 8XXXXXXXXXX
	pattern := `^(\+7|8)\d{10}$`

	// Создаем регулярное выражение
	re := regexp.MustCompile(pattern)

	// Проверяем соответствие
	return re.MatchString(phoneNumber)
}

func isValidPassword(password string) bool {
	if len(password) >= 8 {
		return true
	}
	fmt.Println("Invalid password")
	return false
}
