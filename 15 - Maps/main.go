package main

func main() {
	groupCity := map[int][]string{
		10:   {"Село", "Деревня", "ПГТ"},  // города с населением 10-99 тыс. человек
		100:  {"Город", "Станица"},        // города с населением 100-999 тыс. человек
		1000: {"Мегаполис", "Миллионник"}, // города с населением 1000 тыс. человек и более
	}

	cityPopulation := map[string]int{
		"Город":     101,
		"Станица":   102,
		"Село":      103,
		"Мегаполис": 104,
	}

	for city, _ := range cityPopulation {
		var exists bool

		for _, gCity := range groupCity[100] {
			if city == gCity {
				exists = true
				break
			}
		}

		if !exists {
			delete(cityPopulation, city)
		}
	}
}
