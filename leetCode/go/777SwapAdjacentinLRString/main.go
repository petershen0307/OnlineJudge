package q777

//        _   _   _    _
// start |R|X|X|L|R|XR|X|L
// end   |X|R|L|X|X|RR|L|X
//        ‾   ‾   ‾    ‾
// if there are shifting will pair appear RX(XL) and XR(LX)
// start R  X| X  L
// end   X  R| L  X

func canTransform(start string, end string) bool {
	L := 0
	R := 0
	ok := false
	XX := byte('X')
	LL := byte('L')
	RR := byte('R')
	for i := 0; i < len(start); i++ {
		// if go left/(right) shifting, it can't appear 'R'/('L') in the sequence
		if (start[i] == LL && R != 0) || (start[i] == RR && L != 0) {
			break
		}
		// break when L or R be negative
		// we should get positive first when shift will occur
		if L < 0 || R < 0 {
			break
		}
		// check there is an X in start
		if start[i] == XX {
			ok = true
		}
		// check left shift
		if start[i] == XX && end[i] == LL {
			L++
			continue
		}
		if start[i] == LL && end[i] == XX {
			L--
			continue
		}

		// check right shift
		if start[i] == RR && end[i] == XX {
			R++
			continue
		}
		if start[i] == XX && end[i] == RR {
			R--
			continue
		}
	}
	return 0 == L && 0 == R && ok
}
