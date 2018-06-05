package stats

//Data Used with the Container
type Data struct {
	Timestamp    int64
	ResponseTime float64
	StatusCode   int
	ErrorMessage string
}

//FlattenedValues Represents a summary of a given set of Data{}
//throught a time range
type FlattenedValues struct {
	ResponseTimeStdDev float64
	ResponseTimeMean   float64
	ResponseTimeMin    float64
	ResponseTimeMax    float64
	StatusCodeCount    map[int]int
	ReqPerSec          float64
}
