package main

import "fmt"

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
	//Workers := make(map[int]Worker)
	//
	//Workers[1] = boss
	//Workers[2] = director
	//Workers[3] = manager
	//Workers[4] = plankton
	//Workers[5] = staff
	//
	//for _, data := range Workers {
	//	fmt.Println(GetWorkerData(data))
	//}
	//a := byte('2')
	fmt.Println(15)
}

func do(i interface{}) {
	switch val := i.(type) {
	case int:
		if v, ok := i.(int); ok {
			fmt.Println(v)
		} else {
			fmt.Println("Cannot change %T to int", val)
		}
	case string:
		if v, ok := i.(string); ok {
			fmt.Println(v)
		} else {
			fmt.Println("Cannot change %T to string", val)
		}
	case float32:
		if v, ok := i.(float32); ok {
			fmt.Println(v)
		} else {
			fmt.Println("Cannot change %T to float32", val)
		}
	}

}
