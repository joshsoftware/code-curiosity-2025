package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joshsoftware/code-curiosity-2025/internal/config"
)

var (
	// mainMigrationsDIR defines the directory where all migration files are located
	mainMigrationsDIR = "./internal/db/migrations"

	// mainMigrationFilesPath defines path for migration files
	mainMigrationFilesPath = "file://" + mainMigrationsDIR
)

// Migration used to define migrations
type Migration struct {
	m             *migrate.Migrate
	directoryName string
	filesPath     string
}

// InitMainDBMigrations used to initialize migrations
func InitMainDBMigrations(config config.AppConfig) (migration Migration, er error) {
	var dbConnection string = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

	migration.directoryName = mainMigrationsDIR
	migration.filesPath = mainMigrationFilesPath

	migration.m, er = migrate.New(migration.filesPath, dbConnection)
	// if err == migrate.ErrNoChange {

	return
}

func (migration Migration) MigrationsUpAll(){
	err := migration.m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			slog.Error("No new migrations to apply")
			return
		}
		slog.Error("An error occurred while making migrations up", "error", err)
		return
	}
	slog.Info("Current migration version:", "version", migration.MigrationVersion())
	slog.Info("Migration up completed")
}

func (migration Migration) MigrationsUpWithSteps(steps int){
	if err := migration.m.Steps(steps); err != nil {
		if err == migrate.ErrNoChange {
			slog.Error("No new migrations to apply")
			return
		}

		slog.Error("An error occurred while making migrations up", "error", err)
		return
    }

	slog.Info("Current migration version:", "version", migration.MigrationVersion())
	slog.Info("Migration up completed")
}

// MigrationsUp used to make migrations up
func (migration Migration) MigrationsUp(steps string) {
	if len(steps) == 0 {
		migration.MigrationsUpAll()
	} else {
		stepsCnt, err := strconv.Atoi(steps)
		if err != nil {
			slog.Error("Failed to parse steps argument to integer", "error", err)
			return
		}

		migration.MigrationsUpWithSteps(stepsCnt)
	}
}

func (migration Migration) MigrationsDownAll() {
	err := migration.m.Down()
	if err != nil {
		if err == migrate.ErrNoChange {
			slog.Info("No migrations to revert")
			return
		}

		slog.Error("An error occurred while making migrations down", "error", err)
		return
	}
	slog.Info("Current migration version:", "version", migration.MigrationVersion())
	slog.Info("Migration down completed")
}

func (migration Migration) MigrationsDownWithSteps(steps int) {
	if err := migration.m.Steps(-1 * steps); err != nil {
		if err == migrate.ErrNoChange {
			slog.Error("No migrations to revert")
			return
		}

		slog.Error("An error occurred while making migrations down", "error", err)
		return
    }

	slog.Info("Current migration version:", "version", migration.MigrationVersion())
	slog.Info("Migration down completed")
}

// MigrationsDown used to make migrations down
func (migration Migration) MigrationsDown(steps string) {
	if len(steps) == 0 {
		migration.MigrationsDownAll()
	} else {
		stepsCnt, err := strconv.Atoi(steps)
		if err != nil {
			slog.Error("Failed to parse steps argument to integer", "error", err)
			return
		}

		migration.MigrationsDownWithSteps(stepsCnt)
	}
}

// CreateMigrationFile creates new migration files
func (migration Migration) CreateMigrationFile(filename string) (err error) {
	if len(filename) == 0 {
		return errors.New("filename is not provided")
	}

	timeStamp := time.Now().Unix()
	upMigrationFilePath := fmt.Sprintf("%s/%d_%s.up.sql", migration.directoryName, timeStamp, filename)
	downMigrationFilePath := fmt.Sprintf("%s/%d_%s.down.sql", migration.directoryName, timeStamp, filename)

	defer func() {
		if err != nil {
			os.Remove(upMigrationFilePath)
			os.Remove(downMigrationFilePath)
		}
	}()

	err = createFile(upMigrationFilePath)
	if err != nil {
		return
	}

	slog.Info(fmt.Sprintf("created %s\n", upMigrationFilePath))

	err = createFile(downMigrationFilePath)
	if err != nil {
		return
	}

	slog.Info(fmt.Sprintf("created %s\n", downMigrationFilePath))
	return
}

// createFile used to create a file with specified name of versioning
func createFile(filename string) (err error) {
	f, err := os.Create(filename)
	if err != nil {
		return
	}

	err = f.Close()
	return
}

// MigrationVersion prints the current migration version
func (migration Migration) MigrationVersion() (err error) {
	version, dirty, err := migration.m.Version()
	if err != nil {
		return
	}

	slog.Info(fmt.Sprintf("version: %v, dirty: %v", version, dirty))
	return
}

func main() {
	// Setup config
	cfg, err := config.LoadAppConfig()
	if err != nil {
		slog.Error("error loading app config", "error", err)
		return
	}

	if len(os.Args) < 2 {
		slog.Error("Missing action argument. Use 'up' or 'down' or 'create'.")
		os.Exit(1)
	}

	migration, err := InitMainDBMigrations(cfg)
	if err != nil {
		slog.Error("Error initializing migrations:", "Error", err.Error())
		return
	}

	action := os.Args[1]
	switch action {
	case "up":
		migration.MigrationsUp(os.Args[2])
	case "down":
		migration.MigrationsDown(os.Args[2])
	case "create":
		migration.CreateMigrationFile(os.Args[2])
	default:
		slog.Info("Invalid action. Use 'up' or 'down' or 'create'.")
	}
}
