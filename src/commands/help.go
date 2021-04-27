//package with commands
package commands

// Help возвращает приветственное сообщение
func Help() string {
	return "Поиск по имени: /search или /s \n" +
		"Поиск по автору: /author или /a \n" +
		"Последние 20 книг: /last или /l \n" +
		"Случайная книга: /r \n" +
		"Статистика: /stat \n\n" +
		"------------------------------------\n" +
		"При открытии книг может быть небольшой таймаут из-за большого размера файла."
}
