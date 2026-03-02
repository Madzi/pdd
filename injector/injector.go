package injector

import (
	"fmt"
	"pdd/parser"
	"pdd/utils"
	"strings"
)

func ProcessFile(content string, path string) (string, bool) {
	lines := strings.Split(content, "\n")
	changed := false

	for i, line := range lines {
		m, ok := parser.ParseLine(line)
		if !ok {
			continue
		}

		// Если ID уже есть в формате [123], пропускаем модификацию файла
		if m.Timestamp != "" {
			continue
		}

		newID := utils.Timestamp()
		// Собираем строку обратно:
		// Префикс (код/отступы/комменты) + Тип + : + [ID] + Описание
		lines[i] = fmt.Sprintf("%s%s: [%s] %s", m.Prefix, m.Type, newID, m.Description)
		changed = true
	}
	return strings.Join(lines, "\n"), changed
}
