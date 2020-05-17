package main

import "time"

type Worker struct {
	Id ID `bson:"_id"`
	WorkerInfo interface{}
}

type Person struct {
	Name    string
	Surname string
	Age     uint8
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
