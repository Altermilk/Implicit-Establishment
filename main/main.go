package main

import (
	csv "lab4/csv"
	difur "lab4/difur"
	plot "lab4/plotting"
	"strconv"
	"sync"
)

func main(){

	wg := sync.WaitGroup{}
	wg.Add(6)
	d := make([]difur.Difur, 0)

	dt := []float64{0.1, 0.01, 0.001}
	h := []float64{0.1, 0.01}

	f := func(i,j int){
		dif := difur.MakeDifur(h[j], dt[i], 1e-3)
		dif.Title = "h = " + strconv.FormatFloat(dif.H, 'f', 3, 64) + ", dt = " + strconv.FormatFloat(dif.Dt, 'f', 3, 64)
		dif.ImplicitEstablish()
		d = append(d, dif)
		csv.MakeCSV(dif.U, dif.Title)
		plot.BuildPlot(dif.U, len(dif.U), dif.Nx, dif.Title)
		wg.Done()
	}

	for i := range dt{
		for j := range h{
			go f(i, j)
		}
	}
		
	wg.Wait()
	csv.CompareResults(d)
}



