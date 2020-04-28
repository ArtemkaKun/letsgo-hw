package main

import "time"

type Person struct {
	Name    string
	Surname string
	Age     uint8
}

type Employee struct {
	PersonInfo Person
	Rank       string
	WorkNumber string
}

type Boss struct {
	EmployeeInfo  Employee
	Secretary     Employee `json:"secretary"`
	InvestorsList []Person `json:"investorsList"`
}

type Director struct {
	EmployeeInfo     Employee
	ManagersList     []Employee `json:"managersList"`
	LastListFromBoss string     `json:"lastListFromBoss"`
}

type Manager struct {
	EmployeeInfo            Employee
	ConnectedTeamsList      map[string][]Employee `json:"connectedTeamsList"`
	PresentationForDirector error                 `json:"presentationForDirector"`
}

type OfficePlankton struct {
	EmployeeInfo   Employee
	WorkComputerID uint16        `json:"workComputerID"`
	WorkHours      time.Duration `json:"workHours"`
}

type OfficeStaff struct {
	EmployeeInfo Employee
	OfficeID     uint8    `json:"officeID"`
	WorkTools    []string `json:"workTools"`
}
