package cmd

import (
	"os"
	"strings" // Добавили импорт
	"pdd/injector"
	"pdd/model"
	"pdd/parser"
	"pdd/scanner"
	"pdd/store"
	"pdd/utils"
)

func Scan() {
	oldPuzzles, _ := store.Load()
	activeInCode := make(map[string]model.Puzzle)

	scanner.ScanFiles(".", func(path string) error {
		// Пропускаем сам файл базы данных, чтобы не было рекурсии
		if !utils.IsTextFile(path) || strings.HasSuffix(path, store.FileName) {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}

		content := string(data)
		// Убрали неиспользуемую ids, оставили только флаг changed
		newContent, changed := injector.ProcessFile(content, path)

		if changed {
			os.WriteFile(path, []byte(newContent), 0644)
		}

		for _, line := range strings.Split(newContent, "\n") {
			if m, ok := parser.ParseLine(line); ok {
				activeInCode[m.Timestamp] = model.Puzzle{
					Type:        m.Type,
					Timestamp:   m.Timestamp,
					Description: m.Description,
					File:        path,
					Status:      "OPEN",
				}
			}
		}
		return nil
	})

	// Сливаем состояние: старое помечаем CLOSED, новое (из кода) — OPEN
	resultMap := make(map[string]model.Puzzle)
	for _, old := range oldPuzzles.Items {
		old.Status = "CLOSED"
		resultMap[old.Timestamp] = old
	}
	for id, p := range activeInCode {
		resultMap[id] = p
	}

	var finalPuzzles []model.Puzzle
	for _, p := range resultMap {
		finalPuzzles = append(finalPuzzles, p)
	}

	_ = store.Save(model.Puzzles{Items: finalPuzzles})
}
