package main

type Hero struct {
	On    bool
	Ammo  int
	Power int
}

func (hero *Hero) Shoot() bool {
	if !hero.On {
		return false
	}

	hasAmmo := hero.Ammo > 0
	if hasAmmo {
		hero.Ammo--
	}
	return hasAmmo
}

func (hero *Hero) RideBike() bool {
	if !hero.On {
		return false
	}

	hasPower := hero.Power > 0
	if hasPower {
		hero.Power--
	}
	return hasPower
}

func main() {
}
