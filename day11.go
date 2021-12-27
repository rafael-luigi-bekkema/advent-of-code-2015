package main

func passvalid(pass string) bool {
	return passvalidb([]byte(pass))
}

func passvalidb(pass []byte) bool {
	var straight, dupecount int
	var skipdupe bool
	var lastdupe byte
	for i, c := range []byte(pass[1:]) {
		if c == pass[i]+1 {
			straight++
		} else if straight < 2 {
			straight = 0
		}
		if In(c, 'i', 'o', 'l') {
			return false
		}
		if c == pass[i] && !skipdupe {
			if c != lastdupe {
				dupecount++
				lastdupe = c
			}
			skipdupe = true
			continue
		}
		skipdupe = false
	}
	return dupecount >= 2 && straight >= 2
}

func incpass(pass string) string {
	inp := []byte(pass)
	for {
		for i := len(inp) - 1; i >= 0; i-- {
			if inp[i] != 'z' {
				inp[i]++
				break
			}
			inp[i] = 'a'
		}
		if passvalidb(inp) {
			return string(inp)
		}
	}
}
