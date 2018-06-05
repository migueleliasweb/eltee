package stats

//NewContainer Configures and returns a Container
func NewContainer() Container {
	return Container{
		rawValues: make(map[int64][]Data),
		// flattenedValues: make(map[int64][]Data),
	}
}

//Container This contains all the data from the stats
//This structure is not thread safe
type Container struct {
	rawValues map[int64][]Data
}

//Insert Inserts data into the Container
func (C *Container) Insert(d Data) {
	// C.rawValues[d.Timestamp] = append(C.perSecondValues[d.Timestamp], d)
}

func (C *Container) FlattenValues() {
	// fiveSecondsInThePast := time.Now().Add(time.Second * -5).Unix()
	//
	// if dataPoints, ok := C.perSecondValues[fiveSecondsInThePast]; ok {
	// 	var responseTimes []float64
	//
	// 	for _, data := range dataPoints {
	// 		responseTimes := append(responseTimes, data.ResponseTime)
	// 	}
	//
	// 	responseTimesStdDev := gonumstat.StdDev(responseTimes, nil)
	//
	// 	fmt.Println(responseTimesStdDev)
	// }
}
