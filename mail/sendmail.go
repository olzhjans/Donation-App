package mail

import "gopkg.in/gomail.v2"

func SendMail(to, header, text string) {
	// Создаем новое сообщение
	m := gomail.NewMessage()

	// Задаем заголовок письма
	m.SetHeader("From", "zhunuskanov_olzhas@mail.ru")
	m.SetHeader("To", to)
	m.SetHeader("Subject", header)

	// Задаем тело письма
	m.SetBody("text/plain", text)

	// Устанавливаем параметры для SMTP сервера
	d := gomail.NewDialer("smtp.mail.ru", 465, "zhunuskanov_olzhas@mail.ru", "AtdcWzHhjfr1he7qEu7E")

	// Отправляем письмо
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
