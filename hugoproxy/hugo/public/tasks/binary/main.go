package main

import (
	"fmt"
	avl "github.com/emirpasic/gods/trees/avltree"
	"time"
)

func main() {
	// Создание нового AVL-дерева
	tree := avl.NewWithStringComparator()

	// Запуск воркера для добавления элементов каждые 5 секунд
	go func() {
		for i := 0; ; i++ {
			// Добавление элемента в дерево
			tree.Put(fmt.Sprintf("key%d", i), i)

			// Если дерево достигло 100 элементов, создаем новое дерево
			if i%100 == 0 && i != 0 {
				tree.Clear()
				fmt.Println("Создано новое дерево с 5 элементами.")
				for j := i - 5; j < i; j++ {
					tree.Put(fmt.Sprintf("key%d", j), j)
				}
			}

			// Вывод сбалансированного дерева
			fmt.Println("Сбалансированное дерево:")
			fmt.Println(tree)
			fmt.Println()

			time.Sleep(5 * time.Second)
		}
	}()

	select {}
}
