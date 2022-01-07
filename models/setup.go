package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	dbConn "github.com/phommata/appointment-scheduling/adapter/gorm"
	"github.com/phommata/appointment-scheduling/config"
	"log"
	"time"
)

func InitDb(conf *config.Conf) *gorm.DB {
	var DB, err = dbConn.New(conf)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connection Established")
	//DB.DropTableIfExists(&Service{}, &Version{})
	//DB.AutoMigrate(&Service{}, &Version{})

	DB.DropTableIfExists(&Appointment{})
	DB.AutoMigrate(&Appointment{})

	//services := []Service{
	//	{
	//		Name: "Security",
	//		Description: "Lorem ipsum dolor",
	//		Versions: []Version{
	//			{
	//				Version: "1.0.0",
	//			},
	//			{
	//				Version: "1.0.1",
	//			},
	//			{
	//				Version: "1.0.2",
	//			},
	//		},
	//	},
	//	{
	//		Name: "Security",
	//		Description: "Lorem ipsum dolor",
	//		Versions: []Version{
	//			{
	//				Version: "1.0.0",
	//			},
	//			{
	//				Version: "1.0.1",
	//			},
	//			{
	//				Version: "1.0.2",
	//			},
	//		},
	//	},
	//	{
	//		Name: "Reporting",
	//		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor...",
	//		Versions: []Version{
	//			{
	//				Version: "1.0.0",
	//			},
	//			{
	//				Version: "1.0.1",
	//			},
	//			{
	//				Version: "1.0.2",
	//			},
	//		},
	//	},
	//	{
	//		Name: "Priority Services",
	//		Versions: []Version{
	//			{
	//				Version: "1.0.0",
	//			},
	//			{
	//				Version: "1.0.1",
	//			},
	//			{
	//				Version: "1.0.2",
	//			},
	//		},
	//	},
	//	{
	//		Name: "Notifications",
	//		Versions: []Version{
	//			{
	//				Version: "1.0.0",
	//			},
	//			{
	//				Version: "1.0.1",
	//			},
	//			{
	//				Version: "1.0.2",
	//			},
	//		},
	//	},
	//	{
	//		Name: "Notifications",
	//		Versions: []Version{
	//			{
	//				Version: "1.0.0",
	//			},
	//			{
	//				Version: "1.0.1",
	//			},
	//			{
	//				Version: "1.0.2",
	//			},
	//		},
	//	},
	//	{
	//		Name: "FX Rates International...",
	//		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor",
	//		Versions: []Version{
	//			{
	//				Version: "1.0.0",
	//			},
	//			{
	//				Version: "1.0.1",
	//			},
	//			{
	//				Version: "1.0.2",
	//			},
	//		},
	//	},
	//	{
	//		Name: "FX Rates International",
	//		Description: "Lorem ipsum dolor",
	//		Versions: []Version{
	//			{
	//				Version: "1.0.0",
	//			},
	//			{
	//				Version: "1.0.1",
	//			},
	//			{
	//				Version: "1.0.2",
	//			},
	//		},
	//	},
	//	{
	//		Name: "Contact Us",
	//		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing",
	//		Versions: []Version{
	//			{
	//				Version: "1.0.0",
	//			},
	//			{
	//				Version: "1.0.1",
	//			},
	//			{
	//				Version: "1.0.2",
	//			},
	//		},
	//	},
	//	{
	//		Name: "Contact Us",
	//		Description:"Lorem ipsum dolor sit amet, consectetur adipiscing",
	//		Versions: []Version{
	//			{
	//				Version: "1.0.0",
	//			},
	//			{
	//				Version: "1.0.1",
	//			},
	//			{
	//				Version: "1.0.2",
	//			},
	//		},
	//	},
	//	{
	//		Name: "Collect Monday",
	//		Versions: []Version{
	//			{
	//				Version: "1.0.0",
	//			},
	//			{
	//				Version: "1.0.1",
	//			},
	//			{
	//				Version: "1.0.2",
	//			},
	//		},
	//	},
	//	{
	//		Name: "Locate Us",
	//		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor...",
	//		Versions: []Version{
	//			{
	//				Version: "1.0.0",
	//			},
	//			{
	//				Version: "1.0.1",
	//			},
	//			{
	//				Version: "1.0.2",
	//			},
	//		},
	//	},
	//}

	appointments := []Appointment{
		{
			TrainerID: 1,
			UserID: 1,
			StartsAt: time.Date(2019, 01, 25, 9, 0, 00, 0, time.UTC),
			EndsAt: time.Date(2019, 01, 25, 9, 30, 00, 0, time.UTC),
		},
	}

	//for _, service := range services {
	//	DB.Create(&service)
	//}

	for _, appointment := range appointments {
		DB.Create(&appointment)
	}

	return DB
}