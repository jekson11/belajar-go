package query

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/VauntDev/tqla"
)

type LoadQuery struct {
	queries  map[string]string
	filePath string
}

func NewLoadQuery(filePath string) (*LoadQuery, error) {
	ql := &LoadQuery{
		queries:  make(map[string]string),
		filePath: filePath,
	}

	if err := ql.load(); err != nil {
		return nil, err
	}

	return ql, nil
}

// Read SQL file
func (ql *LoadQuery) load() error {
	file, err := os.Open(ql.filePath)
	if err != nil {
		return fmt.Errorf("failed to open query file: %w", err)
	}
	defer file.Close()

	var (
		currentName string
		currentSQL  strings.Builder
		scanner     = bufio.NewScanner(file)
	)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "-- ") {
			if currentName != "" {
				ql.queries[currentName] = strings.TrimSpace(currentSQL.String())
				currentSQL.Reset()
			}
			currentName = strings.TrimPrefix(line, "-- ")
			continue
		}

		if currentName != "" {
			currentSQL.WriteString(line + "\n")
		}
	}

	if currentName != "" {
		ql.queries[currentName] = strings.TrimSpace(currentSQL.String())
	}

	return scanner.Err()
}

func (ql *LoadQuery) ExecuteTemplate(name string, data any) (string, []any, error) {
	queryTemplate, ok := ql.Get(name)
	if !ok {
		return "", nil, fmt.Errorf("query %s not found", name)
	}
	t, err := tqla.New(tqla.WithPlaceHolder(tqla.Dollar))
	if err != nil {
		return "", nil, err
	}

	query, args, err := t.Compile(queryTemplate, data)
	if err != nil {
		return "", nil, err
	}

	//tmpl, err := template.New(name).Parse(queryTemplate)
	//if err != nil {
	//	return "", nil, err
	//}

	//var buf bytes.Buffer
	//if err := tmpl.Execute(&buf, data); err != nil {
	//	return "", nil, err
	//}
	//
	//query := buf.String()
	return query, args, nil
}

func (ql *LoadQuery) Get(name string) (string, bool) {
	query, ok := ql.queries[name]
	return query, ok
}
