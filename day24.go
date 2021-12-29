package main

func day24a(gifts []int, b bool) int {
	type takeres struct {
		picks  []int
		remain []int
		prod   int
	}

	nrRooms := 3
	if b {
		nrRooms = 4
	}
	weight := Sum(gifts) / nrRooms
	minGifts := 2
	minprod := 0

	var take func(int, int, []int, []int, int, int, chan takeres)
	take = func(n int, idx int, picks, gifts []int, w int, prd int, chn chan takeres) {
		if n == 0 {
			if w != weight {
				return
			}
			tr := takeres{
				picks:  make([]int, len(picks)),
				remain: gifts,
				prod:   prd,
			}
			copy(tr.picks, picks)
			chn <- tr
		}
		for i := range gifts[idx:] {
			nw := w + gifts[idx+i]
			if nw > weight {
				continue
			}
			nprd := gifts[idx+i]
			if len(picks) > 0 {
				nprd *= prd
			}
			if minprod != 0 && prd >= minprod {
				continue
			}
			ngifts := make([]int, len(gifts)-1)
			copy(ngifts, gifts[:idx+i])
			copy(ngifts[idx+i:], gifts[idx+i+1:])
			take(n-1, i, append(picks, gifts[idx+i]), ngifts, nw, nprd, chn)
		}
	}
	taker := func(n int, gifts []int) chan takeres {
		chn := make(chan takeres)
		go func() {
			defer close(chn)
			take(n, 0, nil, gifts, 0, 0, chn)
		}()
		return chn
	}

	minlen := len(gifts)
	for n1 := minGifts; n1 <= minlen; n1++ {
		for g1 := range taker(n1, gifts) {
			if len(g1.picks) < minlen {
				minprod = g1.prod
				minlen = len(g1.picks)
				continue
			}
			if len(g1.picks) == minlen {
				if g1.prod < minprod {
					minprod = g1.prod
				}
			}
		}
	}
	return minprod
}
