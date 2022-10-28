package utslice

func CompareTwoSlice[T int | int64 | string | float32 | float64](sliceOne []T, sliceTwo []T) (union []T, originLeft []T, originRight []T) {
	var slicePopulation = make(map[T]map[string]bool)

	for _, so := range sliceOne {
		if _, ok := slicePopulation[so]; !ok {
			var popul = make(map[string]bool)
			slicePopulation[so] = popul
		}

		if _, ok := slicePopulation[so]["one"]; !ok {
			slicePopulation[so]["one"] = true
		}

		if _, ok := slicePopulation[so]["two"]; !ok {
			slicePopulation[so]["two"] = false
		}

		for _, st := range sliceTwo {
			if _, ok := slicePopulation[st]; !ok {
				var popul = make(map[string]bool)
				slicePopulation[st] = popul
			}

			if _, ok := slicePopulation[st]["two"]; !ok {
				slicePopulation[st]["two"] = true
			}

			if _, ok := slicePopulation[st]["one"]; !ok {
				slicePopulation[st]["one"] = false
			}

			if st == so {
				if _, ok := slicePopulation[st]["one"]; !ok {
					slicePopulation[st]["one"] = true
				} else {
					slicePopulation[st]["one"] = true
				}
			}
		}
	}

	for i, sp := range slicePopulation {
		if sp["one"] {
			if sp["two"] {
				union = append(union, i)
			} else {
				originLeft = append(originLeft, i)
			}
		} else {
			if sp["two"] {
				originRight = append(originRight, i)
			}
		}
	}

	return
}
