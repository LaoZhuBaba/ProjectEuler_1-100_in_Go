package challenges

import "fmt"

func createSuperList(included, excluded []*[]int) []*[]int {
	returnList := make([]*[]int, 0)
	var skip bool
	for _, i := range included {
		skip = false
		for _, e := range excluded {
			if i == e {
				skip = true
			}
		}
		if skip {
			continue
		}
		returnList = append(returnList, i)
	}
	return returnList
}

func triangle(i int) int {

	return i * (i + 1) / 2
}
func square(i int) int {
	return i * i
}
func pentagon(i int) int {

	return i * (3*i - 1) / 2
}
func hexagon(i int) int {

	return i * (2*i - 1)
}
func heptagon(i int) int {

	return i * (5*i - 3) / 2
}
func octagon(i int) int {

	return i * (3*i - 2)
}

func checkMiddleOverlap(i1, i2 int) bool {
	return (i1-(i2/100))%100 == 0
}
func checkOuterOverlap(i1, i2 int) bool {
	return (i2-(i1/100))%100 == 0
}
func makePolygonalList(floor, ceiling int, f func(int) int) []int {
	squareList := make([]int, 0)
	for count := 1; ; count++ {
		transformed := f(count)
		if transformed < floor {
			continue
		}
		if transformed > ceiling {
			return squareList
		}
		squareList = append(squareList, transformed)
	}
}

func Challenge61() {
	triangleList := makePolygonalList(1000, 9999, triangle)
	squareList := makePolygonalList(1000, 9999, square)
	pentagonList := makePolygonalList(1000, 9999, pentagon)
	hexagonList := makePolygonalList(1000, 9999, hexagon)
	heptagonList := makePolygonalList(1000, 9999, heptagon)
	octagonList := makePolygonalList(1000, 9999, octagon)

	for _, v1 := range triangleList {
		bigList := []*[]int{&triangleList, &squareList, &pentagonList, &hexagonList, &heptagonList, &octagonList}
		excludeList := []*[]int{&triangleList}
		//
		for _, p1 := range createSuperList(bigList, excludeList) {
			for _, v2 := range *p1 {
				if checkMiddleOverlap(v1, v2) {
					// fmt.Printf("v1, v2: %d %d\n", v1, v2)
					excludeList = append(excludeList, p1)
					//
					for _, p2 := range createSuperList(bigList, excludeList) {
						for _, v3 := range *p2 {
							if checkMiddleOverlap(v2, v3) {
								// fmt.Printf("v1, v2, v3: %d %d %d\n", v1, v2, v3)
								excludeList = append(excludeList, p2)
								//
								for _, p3 := range createSuperList(bigList, excludeList) {
									for _, v4 := range *p3 {
										if checkMiddleOverlap(v3, v4) {
											// fmt.Printf("v1, v2, v3, v4: %d %d %d %d\n", v1, v2, v3, v4)
											excludeList = append(excludeList, p3)
											//
											for _, p4 := range createSuperList(bigList, excludeList) {
												for _, v5 := range *p4 {
													if checkMiddleOverlap(v4, v5) {
														// fmt.Printf("v1, v2, v3, v4, v5: %d %d %d %d %d\n", v1, v2, v3, v4, v5)
														excludeList = append(excludeList, p4)
														//
														for _, p5 := range createSuperList(bigList, excludeList) {
															for _, v6 := range *p5 {
																if checkMiddleOverlap(v5, v6) {
																	// fmt.Printf("v1, v2, v3, v4, v5, v6: %d %d %d %d %d %d\n", v1, v2, v3, v4, v5, v6)
																	excludeList = append(excludeList, p5)
																	//
																	if checkOuterOverlap(v1, v6) {
																		fmt.Printf("solution: %d\n", v1+v2+v3+v4+v5+v6)
																		return
																	}
																	excludeList = excludeList[:len(excludeList)-1]
																}
															}
														}
														excludeList = excludeList[:len(excludeList)-1]
													}
												}
											}
											excludeList = excludeList[:len(excludeList)-1]
										}
									}
								}
								excludeList = excludeList[:len(excludeList)-1]
							}
						}
					}
					excludeList = excludeList[:len(excludeList)-1]
				}
			}
		}
	}
}
