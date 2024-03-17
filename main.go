package main

import (
	"awesomeProject1/api"
	"awesomeProject1/chat"
	"sync"
)

func main() {
	var Wg sync.WaitGroup
	Wg.Add(2)
	//schedule.DeactivateNeedIfExpired()
	//schedule.ChargeOffBySubscription()
	api.LaunchApiServer()
	chat.LaunchChatServer()
	Wg.Wait() // Ожидаем завершения работы всех серверов
}
