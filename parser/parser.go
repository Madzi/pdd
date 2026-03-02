package parser

import "regexp"

// Группа 1: Всё, что до TODO: [1772488711993091361] или FIXME:
// Группа 2: Тип (TODO или FIXME)
// Группа 3: ID в скобках [123] (опционально)
// Группа 4: Описание (всё, что после)
var Re = regexp.MustCompile(`(?i)^(.*?)(TODO|FIXME):(?:\s*\[(\d+)\])?\s*(.*)$`)

type Match struct {
	Prefix      string
	Type        string
	Timestamp   string
	Description string
}

func ParseLine(line string) (Match, bool) {
	m := Re.FindStringSubmatch(line)
	if m == nil {
		return Match{}, false
	}
	// Если после TODO/FIXME нет двоеточия (например, "TODO :"),
	// регулярка выше сработает только если есть двоеточие.
	return Match{
		Prefix:      m[1],
		Type:        m[2],
		Timestamp:   m[3],
		Description: m[4],
	}, true
}
