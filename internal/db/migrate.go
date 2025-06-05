package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
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

// RunMigrations used to run a migrations
func (migration Migration) RunMigrations() {
	slog.Info("Migrations started from ", "directory", migration.directoryName)
	startTime := time.Now()
	defer func() {
		slog.Info("Migrations complete, total time taken ", "time", time.Since(startTime))
	}()

	// dbVersion is the currently active database migration version
	dbVersion, dirty, err := migration.m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		slog.Error(err.Error())
	}

	// localVersion is the local migration version
	localVersion, err := migration.MigrationLocalVersion()
	if err != nil {
		slog.Error(err.Error())
	}

	if dbVersion > uint(localVersion) {
		slog.Error(fmt.Sprintf("Your database migration %d is ahead of local migration %d, you might need to rollback a few migrations", dbVersion, localVersion))
	}
	if dbVersion < uint(localVersion) && dirty {
		slog.Error(fmt.Sprintf("Your currently active database migration %d is dirty, you might need to rollback it and then deploy again.", dbVersion))
	}

	err = migration.m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			return
		}

		dbVersion, _, err2 := migration.m.Version()
		if err2 != nil {
			slog.Error(err2.Error())
		}

		slog.Error(fmt.Sprintf("Migration failed with error %s, current active database migration is %d", err, dbVersion))
	}
}

// MigrationsUp used to make migrations up
func (migration Migration) MigrationsUp() {
	err := migration.m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			slog.Error("No new migrations to apply")
			return
		}
		slog.Error("*******" + err.Error())
		return
	}
	migration.MigrationVersion()
	slog.Info("Migration up completed")
}

// MigrationsDown used to make migrations down
func (migration Migration) MigrationsDown() {
	err := migration.m.Down()
	if err != nil {
		if err == migrate.ErrNoChange {
			slog.Info("No migrations to revert")
			return
		}

		slog.Error(err.Error())
		return
	}
	migration.MigrationVersion()
	slog.Info("Migration down completed")
}

// ForceVersion forces the migration to a specific version
func (migration Migration) ForceVersion(version int) {
	err := migration.m.Force(version)
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info(fmt.Sprintf("Migration force version %d complete", version))
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

// MigrationLocalVersion gets the latest migration version from local file system
func (migration Migration) MigrationLocalVersion() (localversion int, err error) {
	localDIRFileVersions, err := getMigrationVersionsFromDir(migration.directoryName)
	if err != nil {
		return 0, fmt.Errorf("can't get files information from local file system: %w", err)
	}

	if len(localDIRFileVersions) == 0 {
		slog.Warn("no migration files found in local file system")
		return 0, nil
	}

	slog.Info(fmt.Sprintf("latest migration version from local file system: %d", localDIRFileVersions[0]))
	return localDIRFileVersions[0], nil
}

func getMigrationVersionsFromDir(dir string) ([]int, error) {
	return []int{}, nil
}

// GoToSpecificVersion migrates to a specific version
func (migration Migration) GoToSpecificVersion(version uint) (err error) {
	localDIRFileVersions, err := getMigrationVersionsFromDir(migration.directoryName)
	if err != nil {
		return fmt.Errorf("can't get files information from local file system: %w", err)
	}

	if len(localDIRFileVersions) == 0 {
		slog.Warn("no migration files found in local file system, hence migration not required")
		return nil
	}

	dbversion, dirty, err := migration.m.Version()
	if err != nil {
		if err == migrate.ErrNilVersion {
			slog.Info("no migration found, initializing DB with latest migration")
			err = migration.m.Migrate(version)
			if err != migrate.ErrNoChange {
				return err
			}

			slog.Info(fmt.Sprintf("database successfully initialized with migration: %d", version))
			return nil
		}
		return err
	}

	// if the database is in dirty state, we pick the previous successfully executed migration
	// and force the database to that version
	if dirty {
		index, err := getIndexOfSlice(localDIRFileVersions, int(dbversion))
		if err != nil {
			return errors.New("database version corresponding file not found in local file system")
		}

		if len(localDIRFileVersions) <= index+1 {
			return errors.New("previous successfully executed migration not found in local file system")
		}
		forceMigrateVersion := localDIRFileVersions[index+1]

		err = migration.m.Force(forceMigrateVersion)
		if err != nil {
			return err
		}
	}

	err = migration.m.Migrate(version)
	if err != migrate.ErrNoChange {
		return err
	}

	slog.Info(fmt.Sprintf("database successfully migrated to version: %d", version))
	return nil
}

func getIndexOfSlice(slice []int, value int) (int, error) {
	return 0, nil
}

func main() {
	// Setup config
	cfg, err := config.LoadAppConfig()
	if err != nil {
		slog.Error("error loading app config", "error", err)
		return
	}

	if len(os.Args) < 2 {
		slog.Error("Missing action argument. Use 'up' or 'down' or 'create.")
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
		migration.MigrationsUp()
	case "down":
		migration.MigrationsDown()
	case "create":
		migration.CreateMigrationFile(os.Args[2])
	default:
		slog.Info("Invalid action. Use 'up' or 'down'.")
	}
}
