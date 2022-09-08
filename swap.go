package piscine

func Swap(a *int, b *int) {
	nb_a := *a
	nb_b := *b
	*a = nb_b
	*b = nb_a
}
