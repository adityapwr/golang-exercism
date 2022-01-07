package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		return len(layers) * 2
	} else {
		return len(layers) * time
	}
}

// TODO: define the 'Quantities()' function
func Quantities(items []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, v := range items {
		if v == "noodles" {
			noodles += 50
		} else if v == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce

}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, persons int) []float64 {
	var result []float64
	count := float64(persons) / 2
	for i := range quantities {
		result = append(result, quantities[i]*count)
	}
	return result

}
