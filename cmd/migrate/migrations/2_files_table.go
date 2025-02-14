package migrations

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
)

func init() {
	up := []string{`
		CREATE TABLE IF NOT EXISTS files (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			user_id UUID NOT NULL,
			filename VARCHAR(255) NOT NULL,
			folder_path TEXT DEFAULT '',
			filepath TEXT NOT NULL,
			filesize BIGINT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);
		CREATE INDEX idx_files_folder_path ON files(folder_path);
	`}

	down := []string{`
		DROP TABLE IF EXISTS files CASCADE;
	`}

	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Println("Creating files table with folder_path column...")
		for _, q := range up {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		fmt.Println("✅ files table created successfully!")
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Println("Dropping files table...")
		for _, q := range down {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		fmt.Println("✅ files table dropped successfully!")
		return nil
	})
}
