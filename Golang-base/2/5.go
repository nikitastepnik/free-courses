package main

type s struct {
	On    bool
	Ammo  int
	Power int
}

func (str *s) Shoot() bool {
	if str.On == false {

		return false
	} else if str.Ammo > 0 {
		str.Ammo -= 1

		return true
	}

	return false
}

func (str *s) RideBike() bool {
	if str.On == false {
		return false
	} else if str.Power > 0 {
		str.Power -= 1

		return true
	}

	return false
}
