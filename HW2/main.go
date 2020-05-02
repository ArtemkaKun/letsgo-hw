package main

import (
	"fmt"
	"time"
)

func main() {
	//Way to create pointers #1
	bossPointer := new(Boss)
	directorPointer := new(Director)
	managerPointer := new(Manager)

	fmt.Println(bossPointer, "\n", directorPointer, "\n", managerPointer)

	//Way to create pointers #2
	var bossOrigin Boss
	var directorOrigin Director
	var managerOrigin Manager

	boss := &bossOrigin
	director := &directorOrigin
	manager := &managerOrigin

	fmt.Println(boss, "\n", director, "\n", manager)

	Paul := OfficePlankton{
		Employee:       new(Employee),
		WorkComputerID: 1,
		WorkHours:      8 * time.Hour,
	}

	Paul.Person = new(Person)
	Paul.Name = "Paul"
	Paul.Surname = "Wazowski"
	Paul.Age = 59

	fmt.Println(Paul.GetPersonData())
	fmt.Println(fmt.Sprintf("Go to pensija? %v", Paul.IsPoraNaPensiju()))
	Paul.BirthDayToday()
	fmt.Println(fmt.Sprintf("Go to pensija? %v", Paul.IsPoraNaPensiju()))

	panic("Panic here!")
}
