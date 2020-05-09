package main

import "time"

type Worker interface {
	GetName() string
	GetRank() string
}

func GetWorkerData(w Worker) (name, rank string) {
	name = w.GetName()
	rank = w.GetRank()
	return
}

type Person struct {
	Name    string
	Surname string
	Age     uint8
}

func (obj Person) GetName() string {
	return obj.Name
}

func (obj Person) GetPersonData() (FCs string) {
	FCs = obj.Name + " " + obj.Surname
	return
}

func (obj Person) IsPoraNaPensiju() (pora bool) {
	if obj.Age >= 60 {
		pora = true
	}
	return
}

func (obj Person) BirthDayToday() {
	obj.Age++
}

type Employee struct {
	Person
	Rank       string
	WorkNumber string
}

func (obj Employee) GetRank() string {
	return obj.Rank
}

type Boss struct {
	Employee
	Secretary     Employee `json:"secretary"`
	InvestorsList []Person `json:"investorsList"`
}

type Director struct {
	Employee
	ManagersList     []Employee `json:"managersList"`
	LastListFromBoss string     `json:"lastListFromBoss"`
}

type Manager struct {
	Employee
	ConnectedTeamsList      map[string][]Employee `json:"connectedTeamsList"`
	PresentationForDirector error                 `json:"presentationForDirector"`
}

type OfficePlankton struct {
	Employee
	WorkComputerID uint16        `json:"workComputerID"`
	WorkHours      time.Duration `json:"workHours"`
}

type OfficeStaff struct {
	Employee
	OfficeID  uint8    `json:"officeID"`
	WorkTools []string `json:"workTools"`
}
