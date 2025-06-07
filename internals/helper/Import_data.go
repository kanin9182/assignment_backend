package helper

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func RunSQLFilesInFolder(dsn, folderPath string) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	files, err := os.ReadDir(folderPath)
	if err != nil {
		return fmt.Errorf("error reading folder: %w", err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}
		sqlFilePath := filepath.Join(folderPath, file.Name())
		sqlFile, err := os.ReadFile(sqlFilePath)
		if err != nil {
			return fmt.Errorf("error reading SQL file %s: %w", file.Name(), err)
		}

		queries := strings.Split(string(sqlFile), ";")
		tx, err := db.Begin()
		if err != nil {
			return fmt.Errorf("error starting transaction for %s: %w", file.Name(), err)
		}

		for i, query := range queries {
			query = strings.TrimSpace(query)
			if query == "" || strings.HasPrefix(query, "--") || strings.HasPrefix(query, "#") {
				continue
			}
			if _, err := tx.Exec(query); err != nil {
				tx.Rollback()
				return fmt.Errorf("error executing query %d in %s: %w", i+1, file.Name(), err)
			}
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("error committing transaction for %s: %w", file.Name(), err)
		}
	}

	return nil
}
