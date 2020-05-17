package main

import (
	"log"
)

func main() {
	boss := Boss{
		Employee: Employee{
			Person: Person{
				Name: "Andrew",
			},
			Rank: "Founder of a company",
		},
	}

	director := Director{
		Employee: Employee{
			Person: Person{
				Name: "Artem",
			},
			Rank: "Director of a company",
		},
	}

	manager := Manager{
		Employee: Employee{
			Person: Person{
				Name: "Myroslava",
			},
			Rank: "Manager in IT department",
		},
	}

	plankton := OfficePlankton{
		Employee: Employee{
			Person: Person{
				Name: "Jarik",
			},
			Rank: "Clerk",
		},
	}

	staff := OfficeStaff{
		Employee: Employee{
			Person: Person{
				Name: "Nastia",
			},
			Rank: "Senior System Administrator",
		},
	}

	InsertNewWorker(boss)
	InsertNewWorker(director)
	InsertNewWorker(manager)
	InsertNewWorker(plankton)
	InsertNewWorker(staff)

	log.Println(WorkersCache)

	GetUserInfo(0)
	GetUserInfo(4)
	GetUserInfo(7)

	err := UpdateWorkerInfo(0, staff)
	if err != nil {
		log.Println(err)
	}

	log.Println(WorkersCache)

	err = DeleteWorker(3)
	if err != nil {
		log.Println(err)
	}

	log.Println(WorkersCache)

	for _, worker := range WorkersCache {
		InsertWorkerInDB(worker)
	}

	log.Println(GetWorkerDataFromDB(2))
	UpdateWorkerDataDB(0, Worker{0, boss})
}

func GetUserInfo(id ID) {
	info, err := GetUserInfoById(id)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(info)
	}
}