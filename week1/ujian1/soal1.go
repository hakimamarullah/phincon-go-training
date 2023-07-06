package main

import (
	"fmt"
	"math"
	"strings"
)

type Toy struct{
	Name string
	Brand string
	Sales int64
}

func main(){
	var toys []Toy = make([]Toy, 0)

	toys = append(toys, Toy{"Action Figure", "Hasbro", 5000},
	Toy{"Lego Set", "Lego", 10000},
	Toy{"Barbie Doll", "Mattel", 8000},
	Toy{"Hot Wheels Car", "Mattel", 6000},
	Toy{"Nerf Gun", "Hasbro", 7000},
	Toy{"Play-Doh", "Hasbro", 9000},
	Toy{"Monopoly Board Game", "Hasbro", 4000},
	Toy{"Fisher-Price Baby Toy", "Mattel", 5500},
	Toy{"Transformers Action Figure", "Hasbro", 3000},
	Toy{"Paw Patrol Toy", "Spin Master", 3500},
	Toy{"My Little Pony Doll", "Hasbro", 5000},
	Toy{"LEGO Star Wars Set", "Lego", 7500},
	Toy{"American Girl Doll", "Mattel", 6000},
	Toy{"Marvel Legends Action Figure", "Hasbro", 4000},
	Toy{"FurReal Friends Pet", "Hasbro", 4500},
	Toy{"Hot Wheels Track Set", "Mattel", 5500},
	Toy{"Playmobil Playset", "Playmobil", 3000},
	Toy{"Fisher-Price Little People", "Mattel", 4000},
	Toy{"Nerf N-Strike Blaster", "Hasbro", 6000},
	Toy{"Monster High Doll", "Mattel", 3500},
	Toy{"Power Rangers Action Figure", "Hasbro", 4500},
	Toy{"Thomas & Friends Train Set", "Mattel", 5500},
	Toy{"Play-Doh Kitchen Set", "Hasbro", 7000},
	Toy{"LEGO Technic Set", "Lego", 8000},
	Toy{"Mega Bloks Building Blocks", "Mattel", 4000},
	Toy{"Baby Alive Doll", "Hasbro", 5000},
	Toy{"Uno Card Game", "Mattel", 3000},
	Toy{"Transformers Rescue Bots", "Hasbro", 3500},
	Toy{"Hot Wheels Monster Trucks", "Mattel", 5500},
	Toy{"Fisher-Price Rock-a-Stack", "Mattel", 2500},
	Toy{"Play-Doh Pizza Set", "Hasbro", 3000},
	Toy{"Paw Patrol Playset", "Spin Master", 4000},
	Toy{"LEGO City Set", "Lego", 7000},
	Toy{"Barbie Dreamhouse", "Mattel", 8000},
	Toy{"Nerf Rival Blaster", "Hasbro", 6000},
	Toy{"Hot Wheels Ultimate Garage", "Mattel", 7500},
	Toy{"Playmobil Pirate Ship", "Playmobil", 4000},
	Toy{"Transformers Bumblebee", "Hasbro", 5000},
	Toy{"Fisher-Price Rocking Horse", "Mattel", 3500},
	Toy{"LEGO Friends Set", "Lego", 4500},
	Toy{"Baby Alive Potty Dance", "Hasbro", 5500},
	Toy{"Play-Doh Ice Cream Set", "Hasbro", 6000},
	Toy{"Monster Jam Truck", "Spin Master", 3500},
	Toy{"Hot Wheels Race Track", "Mattel", 4000},
	Toy{"Nerf Fortnite Blaster", "Hasbro", 5000},
	Toy{"Barbie Fashionista Doll", "Mattel", 7000},
	Toy{"LEGO Harry Potter Set", "Lego", 8000},
	Toy{"Fisher-Price Sit-Me-Up", "Mattel", 6000},
	Toy{"Paw Patrol Lookout Tower", "Spin Master", 7500},
	Toy{"Playmobil Police Station", "Playmobil", 4000},
	Toy{"Transformers Optimus Prime", "Hasbro", 5500},
	Toy{"Hot Wheels Color Shifters", "Mattel", 3500},
	Toy{"LEGO Ninjago Set", "Lego", 4500},
	Toy{"FurReal Friends Monkey", "Hasbro", 5500},
	Toy{"Play-Doh Burger Set", "Hasbro", 6000},
	Toy{"Monster Jam Arena", "Spin Master", 4000},
	Toy{"Barbie Dreamtopia Doll", "Mattel", 5000},
	Toy{"Nerf Elite Blaster", "Hasbro", 7000},
	Toy{"Hot Wheels Criss Cross Crash", "Mattel", 8000},
	Toy{"Playmobil Fire Station", "Playmobil", 6000},
	Toy{"Transformers Megatron", "Hasbro", 3500},
	Toy{"Fisher-Price Laugh & Learn", "Mattel", 4500},
	Toy{"LEGO Creator Set", "Lego", 5500},
	Toy{"Baby Alive Sweet Spoonfuls", "Hasbro", 4000},
	Toy{"Play-Doh Doctor Drill 'n Fill", "Hasbro", 5000},
	Toy{"Paw Patrol Adventure Bay", "Spin Master", 7000},
	Toy{"Hot Wheels City Garage", "Mattel", 8000},
	Toy{"LEGO Classic Set", "Lego", 6000},
	Toy{"Barbie Skipper Babysitters Inc.", "Mattel", 7500},
	Toy{"FurReal Friends Puppy", "Hasbro", 4000},
	Toy{"Play-Doh Kitchen Creations", "Hasbro", 5500},
	Toy{"Monster Jam Monster Truck", "Spin Master", 3500},
	Toy{"Hot Wheels Mega Hauler", "Mattel", 4500},
	Toy{"Nerf Modulus Blaster", "Hasbro", 5500},
	Toy{"Barbie Dream Camper", "Mattel", 6000},
	Toy{"LEGO Architecture Set", "Lego", 7000},
	Toy{"Fisher-Price Rock 'n Play", "Mattel", 8000},
	Toy{"Paw Patrol Mighty Pups", "Spin Master", 6000},
	Toy{"Playmobil Zoo Set", "Playmobil", 4000},
	Toy{"Transformers Starscream", "Hasbro", 5000},
	Toy{"Hot Wheels Super Spin Carwash", "Mattel", 7500},
	Toy{"LEGO Duplo Set", "Lego", 4000},
	Toy{"Baby Alive Super Snacks", "Hasbro", 5000},
	Toy{"Play-Doh Wheels Gravel Yard", "Hasbro", 6000},
	Toy{"Monster Jam Grave Digger", "Spin Master", 7000},
	Toy{"Hot Wheels Track Builder", "Mattel", 8000})

	fmt.Println(countsales(toys))
	fmt.Println(vocalname(toys))
	rankSelling(toys)
}

func countsales(toys []Toy) map[string]int {
	results := make(map[string]int, 1)

	for _, item:= range toys {
		results[item.Brand] += int(item.Sales)
	}
	return results
}

func rankSelling(toys []Toy){
	results := make(map[string]int, 1)

	
	max := 0
	min := math.MaxInt
	topsales := ""
	lowestsales := ""
	for _, item:= range toys {
		results[item.Brand] += 1
	}

	for key := range results {
		if results[key] > max {
			max = results[key]
			topsales = key
		}
	}

	for key := range results {
		if results[key] < min {
			min = results[key]
			lowestsales = key
			
		}
	}
	fmt.Printf("Top sales %s %d \n", topsales, results[topsales])
	fmt.Printf("Lowest sales %s %d \n", lowestsales, results[lowestsales])
	
}

func vocalname(toys []Toy) map[string]int {
	results := make(map[string]int, 1)
	results["a"] = 0
	results["i"] = 0
	results["u"] = 0
	results["e"] = 0
	results["o"] = 0
	
	for _, item := range toys {
		switch {
		case strings.EqualFold("a", item.Name[0:1]):
			results["a"] += 1
		case strings.EqualFold("i", item.Name[0:1]):
			results["i"] += 1
		case strings.EqualFold("u", item.Name[0:1]):
			results["u"] += 1
		case strings.EqualFold("e", item.Name[0:1]):
			results["e"] += 1
		case strings.EqualFold("o", item.Name[0:1]):
			results["o"] += 1
		}
	}
	return results
}