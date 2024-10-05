package main

import (
	"fmt"
	"time"

	"github.com/kuraaa/gitCommit/internal/git" // Импорт внутреннего пакета git
)

func main() {

	for {
		now := time.Now()
		fmt.Printf("Текущее время: %s\n", now)

		git.MakeRandomCommits()

		next := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
		duration := time.Until(next)

		fmt.Printf("Следующий запуск через: %s\n", duration)

		// Ожидаем до следующего дня
		time.Sleep(duration)
	}
}
