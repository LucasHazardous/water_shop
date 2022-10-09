package main


type water struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Origin string `json:"origin"`
	Type   string `json:"type"`
	Price  int    `json:"price"`
	Amount int    `json:"amount"`
}

var water_menu = []water{
	{Id: 0, Name: "Giga Chad Water", Origin: "Fiji", Type: "mineral", Price: 7, Amount: 10},
	{Id: 1, Name: "Big Chungus Drink", Origin: "Japan", Type: "sparkling", Price: 20, Amount: 5},
}

func main() {

}
