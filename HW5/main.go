package main

import (
	"fmt"
	"sync"
)

var directors = DirectorCache{Directors: make(map[int]Director)}
var workers = WorkersCache{Workers: make(map[int]Employee)}

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

	var persons []Person

	persons = append(persons, ToType(boss.Employee.Person))
	persons = append(persons, ToType(director.Employee.Person))
	persons = append(persons, ToType(manager.Employee.Person))
	persons = append(persons, ToType(plankton.Employee.Person))
	persons = append(persons, ToType(staff.Employee.Person))

	fmt.Println(persons)

	for i := 0; i < 5; i++ {
		directors.Directors[i] = director
	}

	for i := 0; i < 5; i++ {
		workers.Workers[i] = staff.Employee
	}

	var wg sync.WaitGroup

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go GetOneDirector(i, &wg)
	}
	wg.Wait()

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go GetOneWorker(i, &wg)
	}
	wg.Wait()
}

func GetOneDirector(id int, wg *sync.WaitGroup) {
	v, err := directors.GetDirector(id)
	if err != nil {
		fmt.Println(err)
		wg.Done()
	}

	fmt.Println(v)
	wg.Done()
}

func GetOneWorker(id int, wg *sync.WaitGroup) {
	v, err := workers.GetCache(id)
	if err != nil {
		fmt.Println(err)
		wg.Done()
	}

	fmt.Println(v)
	wg.Done()
}

func ToType(worker interface{}) Person {
	if v, ok := worker.(Person); ok {
		return v
	}

	fmt.Println("Error")
	return Person{}
}
