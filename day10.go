package main

func day10a(input string, repeat int) string {
	items := []byte(input)
	for n := 1; n <= repeat; n++ {
		var result []byte
		count := 1
		for i, c := range items[1:] {
			if items[i] == c {
				count++
			} else {
				if count > 9 {
					panic("too high!")
				}
				result = append(result, byte(count)+'0', items[i])
				count = 1
			}
		}
		items = append(result, byte(count)+'0', items[len(items)-1])
	}
	return string(items)
}
