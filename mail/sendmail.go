package mail

import (
	"flag"
	"github.com/golang/glog"
	"gopkg.in/gomail.v2"
	"log"
)

func SendMail(to, header, text string) {
	var err error
	err = flag.Set("logtostderr", "false") // Логировать в stderr (консоль) (false для записи в файл)
	if err != nil {
		log.Fatal(err)
	}
	err = flag.Set("stderrthreshold", "FATAL") // Устанавливаем порог для вывода ошибок в stderr
	if err != nil {
		log.Fatal(err)
	}
	err = flag.Set("log_dir", "C:/golang/logs/") // Указываем директорию для сохранения логов
	if err != nil {
		log.Fatal(err)
	}
	flag.Parse()
	defer glog.Flush()

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
	if err = d.DialAndSend(m); err != nil {
		glog.Fatal(err)
	}
}
