package migrations

import (
	"log/slog"
	"ugly-friend/models"

	"gorm.io/gorm"
)

// MigrateTables - функция для миграции моделей и проверки существования таблиц
func MigrateTables(db *gorm.DB, log *slog.Logger) error {
	log.Info("🔍 Initialize db migration...")

	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Debt{},
	}

	for _, model := range modelsToMigrate {
		if err := MigrateModel(db, log, model); err != nil {
			return err
		}
	}

	return nil
}

func MigrateModel(db *gorm.DB, log *slog.Logger, model interface{}) error {
	stmt := &gorm.Statement{DB: db}
	err := stmt.Parse(model)

	if err != nil {
		log.Error("💔 Model parsing error:", slog.String("error", err.Error()))
		return err
	}
	tableName := stmt.Schema.Table

	if !db.Migrator().HasTable(model) {
		log.Info("⚠️ Table doesn't exist. Execute migration...", slog.String("table", tableName))
		if err := db.AutoMigrate(model); err != nil {
			log.Error("💔 Migration failed", slog.String("table", tableName))
			return err
		}
		log.Info("🤍 Migration completed", slog.String("table", tableName))
	}
	return nil
}
