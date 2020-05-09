package main

import "fmt"

type ID int

func main() {
	//boss := Boss{
	//	Employee: Employee{
	//		Person: Person{
	//			Name: "Andrew",
	//		},
	//		Rank: "Founder of a company",
	//	},
	//}
	//
	//director := Director{
	//	Employee: Employee{
	//		Person: Person{
	//			Name: "Artem",
	//		},
	//		Rank: "Director of a company",
	//	},
	//}
	//
	//manager := Manager{
	//	Employee: Employee{
	//		Person: Person{
	//			Name: "Myroslava",
	//		},
	//		Rank: "Manager in IT department",
	//	},
	//}
	//
	//plankton := OfficePlankton{
	//	Employee: Employee{
	//		Person: Person{
	//			Name: "Jarik",
	//		},
	//		Rank: "Clerk",
	//	},
	//}
	//
	//staff := OfficeStaff{
	//	Employee: Employee{
	//		Person: Person{
	//			Name: "Nastia",
	//		},
	//		Rank: "Senior System Administrator",
	//	},
	//}
	//
	//Workers := make(map[ID]interface{})
	//
	//Workers[1] = boss
	//Workers[2] = director
	//Workers[3] = manager
	//Workers[4] = plankton
	//Workers[5] = staff
	//Workers[6] = 342
	//
	//WhatTypeIs(Workers)
	fmt.Println(CheckCardNumber("4561261212345467"))
	fmt.Println(CalcDistsance("GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT"))
}

func WhatTypeIs(workers map[ID]interface{}) {
	for _, val := range workers {
		switch v := val.(type) {
		case Boss:
			fmt.Println("I'm BOSS")
		case Director:
			fmt.Println("I'm DIRECTOR")
		case Manager:
			fmt.Println("I'm MANAGER")
		case OfficePlankton:
			fmt.Println("I'm PLANKTON")
		case OfficeStaff:
			fmt.Println("I'm STAFF")
		default:
			fmt.Printf("I'm now working here! I'm %v\n", v)
		}
	}
}
