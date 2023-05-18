package models

import (
	"fmt"
	"github.com/coreos/go-semver/semver"
	"xueyigou_demo/db"
)

const moduleName = "xueyigo"

var version = *semver.New("1.0.0-20230505-alpha")

func SetupDatabase(migrator db.MigratorInterface) {
	currentVersion := migrator.GetVersion(moduleName)
	if !currentVersion.LessThan(version) {
		return
	}
	if err := InitUser(migrator); err != nil {
		panic(err)
	}
	if err := Initaccount(); err != nil {
		panic(err)
	}
	if err := InitCommentL1(); err != nil {
		panic(err)
	}
	if err := InitCommentL2(); err != nil {
		panic(err)
	}
	if err := InitWork(migrator); err != nil {
		panic(err)
	}
	if err := InitWorkMid(); err != nil {
		panic(err)
	}
	if err := InitHonor(); err != nil {
		panic(err)
	}
	if err := InitTask(); err != nil {
		panic(err)
	}
	if err := InitTaskMid(); err != nil {
		panic(err)
	}
	if err := InitWelfare(); err != nil {
		panic(err)
	}
	if err := InitIdentity(); err != nil {
		panic(err)
	}
	if err := InitPeople(); err != nil {
		panic(err)
	}
	if err := InitPeopleMid(); err != nil {
		panic(err)
	}
	if err := InitMessage(); err != nil {
		panic(err)
	}

	if err := InitReport(); err != nil {
		panic(err)
	}

	// if err := InitDefaultPic(); err != nil {
	// 	DefaultPicRecover()
	// }

	if err := InitWelfareMember(); err != nil {
		panic(err)
	}
	if err := InitEvents(); err != nil {
		panic(err)
	}
	if err := InitAdminUser(); err != nil {
		panic(err)
	}
	migrator.SetVersion(moduleName, version)

	fmt.Println("The database is initialized successful.")
}

func InitUser(migrator db.MigratorInterface) error {
	migrator.MigrateModel(&User{})
	migrator.MigrateModel(&UserLables{})
	migrator.MigrateModel(&Useraddress{})
	migrator.MigrateModel(&Account{})
	return nil
}
func InitWelfare() error {
	err := db.DB.AutoMigrate(&Welfare{}, &WelfareHistory{}, &WelfarePictureUrl{})
	return err
}
func InitMessage() error {
	err := db.DB.AutoMigrate(&SystemMessage{}, &InteractiveMessage{}, &OfficialMessage{})
	return err
}
func Initaccount() error {
	var err error
	m := db.DB.Migrator()
	if m.HasTable(&Account{}) {
		return nil
	}
	err = m.CreateTable(&Account{})
	return err
}

func InitCommentL1() error {
	var err error
	err = db.DB.AutoMigrate(&FirstLevelComment{})
	// m := db.DB.Migrator()
	// if m.HasTable(&model.FirstLevelComment{}) {
	// 	return nil
	// }
	// err = m.CreateTable(&model.FirstLevelComment{})
	return err
}

func InitCommentL2() error {
	var err error
	err = db.DB.AutoMigrate(&SecondLevelComment{})
	// m := db.DB.Migrator()
	// if m.HasTable(&model.SecondLevelComment{}) {
	// 	return nil
	// }
	// err = m.CreateTable(&model.SecondLevelComment{})
	return err
}

func InitWork(migrator db.MigratorInterface) error {
	migrator.MigrateModel(WorkModel)
	migrator.MigrateModel(WorkPicturesUrl{})
	migrator.MigrateModel(&WorkType{})
	return nil
}

func InitWorkMid() error {
	err := db.DB.AutoMigrate(&WorkMid{}, &WorkMidPicturesUrl{}, &WorkMidSubjectItem{})
	return err
}

func InitHonor() error {
	err := db.DB.AutoMigrate(&HonorList{})

	return err
}

func InitTask() error {
	err := db.DB.AutoMigrate(&Task{}, &TaskPicturesUrl{}, &TaskSubjectItem{}, &Candidate{})
	return err
}

func InitEvents() error {
	var err error
	err = db.DB.AutoMigrate(&Event{}, &Eventpicurl{})
	return err
}
func InitTaskMid() error {
	err := db.DB.AutoMigrate(&TaskMid{}, &TaskMidPicturesUrl{}, &TaskMidSubjectItem{})
	return err
}

func InitIdentity() error {
	return db.DB.AutoMigrate(&Indentity{})
}

func InitPeople() error {
	return db.DB.AutoMigrate(&People{}, &PeopleSubjectItem{})
}

func InitPeopleMid() error {
	return db.DB.AutoMigrate(&PeopleMid{}, &PeopleMidSubjectItem{})
}

func InitWelfareMember() error {
	err := db.DB.AutoMigrate(&WelfareMember{})
	return err
}

func InitAdminUser() error {
	err := db.DB.AutoMigrate(&SuperUser{})
	return err
}

func InitReport() error {
	err := db.DB.AutoMigrate(&Report{})
	return err
}

func InitDefaultPic() error {
	return db.DB.AutoMigrate(&DefaultPicture{}, &DefaultPictureId2Url{})
}
