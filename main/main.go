package main

import (
	"fmt"
	difur "lab4/difur"
	plot "lab4/plotting"
	csv "lab4/csv"
	"strconv"
	"sync"
)

func main(){

	wg := sync.WaitGroup{}
	wg.Add(6)
	d := make([]difur.Difur, 0)

	dt := []float64{0.1, 0.01, 0.001}
	h := []float64{0.1, 0.01}

	for i := range dt{
		for j := range h{
			go func(i, j int){
				dif := difur.MakeDifur(h[j], dt[i], 0.1)
				dif.Title = "h = " + strconv.FormatFloat(dif.H, 'f', 3, 64) + ", dt = " + strconv.FormatFloat(dif.Dt, 'f', 3, 64)
				dif.ImplicitEstablish()
				d = append(d, dif)
				csv.MakeCSV(dif.U, dif.Title)
				plot.BuildPlot(dif.U, len(dif.U), dif.Nx, dif.Title)
				wg.Done()
			}(i,j)
		}
	}
		
	wg.Wait()

	for _, D := range d{
		fmt.Println("\n>> h = ", D.H, " dt = ", D.Dt)
		for n  := range D.U{
			if n == 0{
				continue
			}
			fmt.Println("n = ", n)
			step := 1
			if D.H == 0.01{
				step = 10
			}
			for i:=0;i < D.Nx ; i += step{
				fmt.Print(strconv.FormatFloat(D.U[n][i], 'f', 2, 64), " ")
			}
			fmt.Println()
		}
	}
}

