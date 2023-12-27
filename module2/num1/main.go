package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	// Получение точного времени с NTP сервера
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при получении времени:", err)
		os.Exit(1)
	}

	// Вывод полученного времени
	fmt.Printf("Точное время: %s\n", ntpTime.Format(time.RFC3339))
}
