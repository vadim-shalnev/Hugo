package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)
	t := ""
	var tick int

	for {
		select {
		case <-ticker.C:
			currentTime := time.Now()

			t = fmt.Sprintf(currentTime.Format("2006-01-02 15:04:05"))
			tick++

			writeToIndexMD(t, tick)
		}
	}
}

func writeToIndexMD(t string, tick int) {
	// Путь к файлу index.md
	mdFilePath := "/home/vubuntu/hugo/hugoproxy/hugo/content/tasks/time/_index.md"

	// Форматируем содержимое файла
	content := fmt.Sprintf(`---
menu:
    before:
        name: tasks
        weight: 5
title: Обновление данных в реальном времени
---

# Задача: Обновление данных в реальном времени

Напишите воркер, который будет обновлять данные в реальном времени, на текущей странице.
Текст данной задачи менять нельзя, только время и счетчик.

Файл данной страницы: "/app/static/tasks/_index.md"

Должен меняться счетчик и время:

Текущее время: %s

Счетчик: %d

## Критерии приемки:
- [ ] Воркер должен обновлять данные каждые 5 секунд
- [ ] Счетчик должен увеличиваться на 1 каждые 5 секунд
- [ ] Время должно обновляться каждые 5 секунд
`, t, tick)

	// Открываем файл на запись
	file, err := os.OpenFile(mdFilePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	// Записываем новое содержимое в файл
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}

	fmt.Println("Файл успешно обновлен.")
}
