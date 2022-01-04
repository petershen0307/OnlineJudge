package complementofbase10integer

func bitwiseComplement2(n int) int {
	if n == 0 {
		return 1
	}
	bits := 0
	temp := n
	for ; temp>>bits > 0; bits++ {
	}
	c := (1 << bits) - 1
	return c - n
}

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	mask := 0
	for ; mask < n; mask = (mask << 1) | 1 {
	}
	return mask ^ n
}
