package Controllers

import "fmt"

type sampleWeight struct {
	id     int
	weight int
}

func RunFuncl() {
	ObjOne := []sampleWeight{
		sampleWeight{
			id:     1,
			weight: 50,
		},
		sampleWeight{
			id:     3,
			weight: 60,
		},
		sampleWeight{
			id:     5,
			weight: 70,
		},
	}
	ObjTwo := []sampleWeight{
		sampleWeight{
			id:     2,
			weight: 30,
		},
		sampleWeight{
			id:     4,
			weight: 90,
		},
		sampleWeight{
			id:     6,
			weight: 140,
		},
	}
	ObjOne = append(ObjOne, ObjTwo...)
	fmt.Println("ObjOne: ", ObjOne)
	ShortByWeight(ObjOne)

}

func ShortByWeight(objOne []sampleWeight) {
	smpWt := objOne
	ln := len(smpWt)
	for i := 0; i < ln-1; i++ {
		for j := 0; j < ln-i-1; j++ {
			if smpWt[j].weight > smpWt[j+1].weight {
				smpWt[j].weight, smpWt[j+1].weight = smpWt[j+1].weight, smpWt[j].weight
			}
		}
	}
	fmt.Println("ShortByWeight: ", smpWt)

}
