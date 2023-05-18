package db

import (
	"database/sql"
	"time"

	"github.com/coreos/go-semver/semver"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MigrateStatement interface {
	Exec(db *gorm.DB) error
}

type MigratorInterface interface {
	GetVersion(moduleName string) semver.Version
	SetVersion(moduleName string, version semver.Version)
	MigrateWithVersion(moduleName string, version semver.Version, f func(migrator MigratorInterface, currentVersion semver.Version))

	MigrateModel(model interface{})

	DB() *gorm.DB
	SqlDB() *sql.DB

	Exec(stmt MigrateStatement) error
	MustExec(stmt MigrateStatement)
}

type migrator struct {
	db       *gorm.DB
	versions map[string]semver.Version
}

func NewMigrator(db *gorm.DB) MigratorInterface {
	var VersionRecords = new(VersionRecord)
	if !db.Migrator().HasTable(VersionRecords) {
		if err := db.AutoMigrate(VersionRecords); err != nil {
			logrus.Fatal(err)
		}
		return &migrator{db: db, versions: map[string]semver.Version{}}
	}

	var versionRecords []*VersionRecord
	err := db.Model(VersionRecords).Find(&versionRecords).Error
	if err != nil {
		logrus.Fatal(err)
	}
	versions := make(map[string]semver.Version, len(versionRecords))
	for _, record := range versionRecords {
		versions[record.ModuleName] = *semver.New(record.Version)
	}
	return &migrator{db: db, versions: versions}
}

func (s *migrator) GetVersion(moduleName string) semver.Version {
	return s.versions[moduleName]
}

func (s *migrator) SetVersion(moduleName string, version semver.Version) {
	record := &VersionRecord{
		ModuleName: moduleName,
		Version:    version.String(),
		UpdatedAt:  time.Now(),
	}
	err := s.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "module_name"}},
		DoUpdates: clause.AssignmentColumns([]string{"version", "updated_at"}),
	}).Create(&record).Error
	if err != nil {
		logrus.Fatal(err)
	}
	s.versions[moduleName] = version
}

func (s *migrator) MigrateModel(model interface{}) {
	if err := s.db.AutoMigrate(model); err != nil {
		logrus.Fatal(err)
	}
}

func (s *migrator) Exec(stmt MigrateStatement) error {
	return stmt.Exec(s.DB())
}

func (s *migrator) MustExec(stmt MigrateStatement) {
	if err := s.Exec(stmt); err != nil {
		logrus.Fatal(err)
	}
}

func (s *migrator) DB() *gorm.DB {
	return s.db
}

func (s *migrator) SqlDB() *sql.DB {
	db, err := s.db.DB()
	if err != nil {
		logrus.Fatal(err)
	}
	return db
}

func (s *migrator) MigrateWithVersion(moduleName string, version semver.Version, f func(migrator MigratorInterface, currentVersion semver.Version)) {
	currentVersion := s.GetVersion(moduleName)
	if !currentVersion.LessThan(version) {
		return
	}
	f(s, currentVersion)
	s.SetVersion(moduleName, version)
}
