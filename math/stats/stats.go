package stats

func Avg(vals[]float64) (float64) {
	var total float64 = 0
	for _,val := range vals {
		total+=val
	}
	return total/float64(len(vals))
}