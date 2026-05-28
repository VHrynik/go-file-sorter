package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	// Проверяем: передал ли пользователь путь при запуске программы
	// Например: go run main.go "D:\test"
	// os.Args: [0] -> путь к программе, [1] -> первый аргумент
	if len(os.Args) < 2 {
		fmt.Println("Please provide a folder path!")
		return // останавливаем программу
	}

	// Берём путь из аргументов командной строки
	folderPath := os.Args[1]

	// Пытаемся получить информацию о пути
	// info -> информация о файле/папке
	// err  -> ошибка, если что-то пошло не так
	info, err := os.Stat(folderPath)

	// Проверяем: произошла ли ошибка при попытке получить информацию о пути
	if err != nil {

		// Если путь не существует
		if os.IsNotExist(err) {
			fmt.Println("Path doesn't exist!")
		} else {
			// Любая другая ошибка, например: нет доступа, проблемы с файловой системой
			fmt.Println("Error: ", err)
		}
		return
	}

	// Проверяем: является ли путь папкой
	// !info.IsDir() читается как: "если это НЕ папка"
	if !info.IsDir() {
		fmt.Println("This is not a folder!")
		return
	}

	// Читаем содержимое папки
	// entries содержит: файлы, папки
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error reading directory: ", err)
		return
	}

	// Создаём map: extension -> category
	// Например: .png -> images, .mp3 -> audio
	myMap := make(map[string]string)

	myMap[".png"] = "images"
	myMap[".jpg"] = "images"
	myMap[".jpeg"] = "images"
	myMap[".gif"] = "images"
	myMap[".webp"] = "images"
	myMap[".bmp"] = "images"

	myMap[".zip"] = "archives"
	myMap[".rar"] = "archives"
	myMap[".7z"] = "archives"
	myMap[".tar"] = "archives"
	myMap[".gz"] = "archives"

	myMap[".txt"] = "docs"
	myMap[".pdf"] = "docs"
	myMap[".docx"] = "docs"
	myMap[".doc"] = "docs"
	myMap[".xlsx"] = "docs"
	myMap[".xls"] = "docs"
	myMap[".pptx"] = "docs"

	myMap[".mp3"] = "audio"
	myMap[".wav"] = "audio"
	myMap[".flac"] = "audio"
	myMap[".ogg"] = "audio"

	myMap[".mp4"] = "video"
	myMap[".mkv"] = "video"
	myMap[".avi"] = "video"
	myMap[".mov"] = "video"
	myMap[".webm"] = "video"

	// Проходимся по каждому объекту внутри папки
	for _, entry := range entries {
		// Если объект является папкой — пропускаем его
		if entry.IsDir() {
			continue
		}

		// Получаем расширение файла. Например: image.png -> .png
		extension := strings.ToLower(filepath.Ext(entry.Name()))

		// Пытаемся найти категорию для данного extension в map.
		// value -> category
		// ok    -> существует ли такой ключ
		value, ok := myMap[extension]

		// Если extension найден в map
		if ok {
			// Создаём путь будущей папки
			// Например: D:\test\images
			// filepath.Join правильно собирает пути под любую ОС
			categoryPath := filepath.Join(folderPath, value)

			//fmt.Printf("%s → %s\n", entry.Name(), value)
			//fmt.Printf("%s → %s\n", entry.Name(), filepath.Join(folderPath, value))
			// Создаём папку категории
			// MkdirAll: создаёт папку, НЕ падает если папка уже существует
			err := os.MkdirAll(categoryPath, 0755)

			// Если произошла ошибка — завершаем программу
			if err != nil {
				fmt.Println(err)
				return
			}

			sourcePath := filepath.Join(folderPath, entry.Name())
			destDir := filepath.Join(folderPath, value)
			destPath := filepath.Join(destDir, entry.Name())
			err = os.Rename(sourcePath, destPath)

			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("%s → %s\n", entry.Name(), value)
		} else {
			// Если extension неизвестен
			fmt.Printf("%s → unknown\n", entry.Name())
		}
	}
}
