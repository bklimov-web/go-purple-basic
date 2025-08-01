package main

import "fmt"

func main() {
	bookmarks := map[string]string{}

loop:
	for {
		menuItem := getUserInput()

		switch menuItem {
		case "1":
			showBookmarks(bookmarks)
		case "2":
			bookmarks = addBookmark(bookmarks)
		case "3":
			bookmarks = deleteBookmark(bookmarks)
		case "4":
			break loop
		}
	}
}

func getUserInput() string {
	var userChoice string
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	fmt.Scan(&userChoice)

	return userChoice
}

func showBookmarks(bookmarks map[string]string) {
	if len(bookmarks) == 0 {
		fmt.Println("У вас нет закладок")
	} else {
		for name, url := range bookmarks {
			fmt.Printf("%v - %v\n", name, url)
		}
	}
}

func addBookmark(bookmarks map[string]string) map[string]string {
	var name string
	var url string

	fmt.Print("Введите название закладки: ")
	fmt.Scan(&name)

	fmt.Print("Введите адрес: ")
	fmt.Scan(&url)

	bookmarks[name] = url

	return bookmarks
}

func deleteBookmark(bookmarks map[string]string) map[string]string {
	var name string

	fmt.Print("Введите название закладки для удаления: ")
	fmt.Scan(&name)

	if _, exists := bookmarks[name]; exists {
		delete(bookmarks, name)
		fmt.Printf("Закладка %v удалена\n", name)
	} else {
		fmt.Printf("Закладки %v не было\n", name)
	}

	return bookmarks
}
