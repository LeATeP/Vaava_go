package drop

import (
	// "fmt"
	"math/rand"
	"time"
)

func GenerateDrop() map[string]int {
	rand.Seed(time.Now().Unix())

	chances := map[string][]int{}
	chances["Rock"] = []int{100,50,2}
	chances["Iron_ore"] = []int{100,25,1}
	chances["Mana_crystal"] = []int{100,5,1}

	drop_list := map[string]int{}
	for name, ch := range chances {
		drop := get_drop(ch)

		if drop > 0 {
			drop_list[name] = drop
		}
	}
	return drop_list
}

func get_drop(chance_list []int) int {
	chance_from, drop_chance, drop_amount := chance_list[0], chance_list[1], chance_list[2]
	chance_from = rand.Intn(chance_from)

	if chance_from < drop_chance {
		return rand.Intn(drop_amount) + 1}
	return 0
}