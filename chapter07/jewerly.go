package main

import "fmt"

func main() {
	jewerly := make(map[string]float64)
	jewerly["necklace"] = 89.99
	jewerly["earrings"] = 79.99

	clothing := map[string]float64{"pants": 39.99, "shirt": 59.99}

	fmt.Println("Earrings:", jewerly["earrings"])
	fmt.Println("Necklace:", jewerly["necklace"])
	fmt.Println("Shirt:", clothing["shirt"])
	fmt.Println("Pants:", clothing["pants"])
}
