package piscine

func UltimateDivMod(a *int, b *int) {
	nb_a := *a
	nb_b := *b
	*a = nb_a / nb_b
	*b = nb_a % nb_b
}
