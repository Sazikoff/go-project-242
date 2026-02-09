package pathsize

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetSize(path string, a bool, r bool) (int64, error) {
	// определяем файл или дир
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	// если не дир, то выводим размер (не имеет значения, видимый или нет)
	if !info.IsDir() {
		fmt.Println(info.Name())
		return info.Size(), nil
	}
	// переводим наш путь в тип, с которым можно что то делать
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var total int64
	var sub_total int64

	// например вынимать содержимое директории (формат не строковый)
	// вот тут начинаем фильтрацию!!!
	for _, entry := range entries {

		// если нет флага "с учетом скрытых файлов и директорий", то пропускаем
		if !a {
			if strings.HasPrefix(entry.Name(), ".") {
				continue
			}
		}

		// Если entry это файл
		if !entry.IsDir() {

			fileInfo, err := entry.Info()
			if err != nil {
				continue
			}

			// если файл удовлетворяет, фиксируем размер
			sub_total = fileInfo.Size()

		} else {

			// если нет флага "рекурсия", то пропускаем
			if !r {
				continue
			}

			// формируем путь до внутренней директории (потому что в fileInfo только последний компонент пути)
			sub_path := filepath.Join(path, entry.Name())
			sub_total, err = GetSize(sub_path, a, r)

			if err != nil {
				// return 0, err
				continue
			}
		}
		total += sub_total

	}

	return total, nil
}
