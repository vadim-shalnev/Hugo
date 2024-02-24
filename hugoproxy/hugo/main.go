package main

import (
	"fmt"
	avl "github.com/emirpasic/gods/trees/avltree"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"math/rand"
	"os"
	"time"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)
	t := ""
	var tick int
	go binary()
	go graf()
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

func graf() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			nodesCount := rand.Intn(26) + 5 // От 5 до 30 элементов

			g := simple.NewUndirectedGraph()
			nodes := make([]graph.Node, nodesCount)

			// Создание узлов графа
			for i := 0; i < nodesCount; i++ {
				nodes[i] = g.NewNode()
				g.AddNode(nodes[i])
			}

			// Создание ребер между узлами
			for i := 0; i < nodesCount-1; i++ {
				g.SetEdge(g.NewEdge(nodes[i], nodes[i+1]))
			}
			var gr string
			for _, n := range nodes {
				edges := g.From(n.ID())

				for edges.Next() {

					gr += fmt.Sprintf("%d -> %d\n", n.ID(), edges.Node().ID())

				}
			}
			mdgraf(gr)

		}
	}
}
func mdgraf(n string) {
	mdFilePath := "/home/vubuntu/hugo/hugoproxy/hugo/content/tasks/graph.md"
	content := fmt.Sprintf(`---
menu:
    after:
        name: graph
        weight: 1
title: Построение графа
---

# Построение графа

Нужно написать воркер, который будет строить граф на текущей странице, каждые 5 секунд
От 5 до 30 элементов, случайным образом. Все ноды графа должны быть связаны.

----
%s`, n)
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
}
func binary() {

	tree := avl.NewWithStringComparator()

	for i := 0; ; i++ {
		// Добавление элемента в дерево
		tree.Put(fmt.Sprintf("key%d", i), i)

		// Если дерево достигло 100 элементов, создаем новое дерево
		if i == 100 || i == 0 {
			tree.Clear()
			fmt.Println("Создано новое дерево с 5 элементами.")
			for j := i - 5; j < i; j++ {
				tree.Put(fmt.Sprintf("key%d", j), j)
			}
		}
		mdThree(tree)

	}

}
func mdThree(tree *avl.Tree) {
	mdFilePath := "/home/vubuntu/hugo/hugoproxy/hugo/content/tasks/binary.md"
	content := fmt.Sprintf(`---
menu:
    after:
        name: binary_tree
        weight: 2
title: Построение сбалансированного бинарного дерева
---

# Задача построить сбалансированное бинарное дерево
Используя AVL дерево, постройте сбалансированное бинарное дерево, на текущей странице.

Нужно написать воркер, который стартует дерево с 5 элементов, и каждые 5 секунд добавляет новый элемент в дерево.

Каждые 5 секунд на странице появляется актуальная версия, сбалансированного дерева.

При вставке нового элемента, в дерево, нужно перестраивать дерево, чтобы оно оставалось сбалансированным.

Как только дерево достигнет 100 элементов, генерируется новое дерево с 5 элементами.

-%s`, tree)
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
	time.Sleep(5 * time.Second)
}

func writeToIndexMD(t string, tick int) {
	// Путь к файлу index.md
	mdFilePath := "/home/vubuntu/hugo/hugoproxy/hugo/content/tasks/_index.md"

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
