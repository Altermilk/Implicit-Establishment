package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	difur "lab4/difur"
)

func MakeCSV(u [][]float64, filename string) {
	f, err := os.Create(filename + ".csv")

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)

	for i := 0; i < len(u); i++ {
		tmp := make([]string, len(u[0]))
		for j := 0; j < len(u[0]); j++ {
			tmp[j] = strconv.FormatFloat(u[i][j], 'f', 2, 64)
		}
		if err := w.Write(tmp); err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()
	f.Close()
	fmt.Println("CSV: \"", f.Name(), "\" was created")
}

func CompareResults(d []difur.Difur){

	filename := "Result"
	f, err := os.Create(filename + ".csv")

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	x := []float64{ 0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1}
	tmp := make([]string, len(x))
	for i := range x{
		tmp[i] = FF(x[i])
	}
	w.Write(tmp)
	for _, D := range d{
		title := ">> h = " + FF(D.H) + " dt = " + FF(D.Dt)
		w.Write([]string{title})
		n := len(D.U) - 1
		// for n  := range D.U{
			// if n == 0{
			// 	continue
			// }
			w.Write([]string{"n = " +  strconv.Itoa(n)})
			step := 1
			var tmp []string
			if D.H == 0.01{
				step = 10
			}
			for i:=0;i < D.Nx ; i += step{
				tmp = append(tmp, FF(D.U[n][i]))
			}
			w.Write(tmp)
		// }
	}

	w.Flush()
	f.Close()
	fmt.Println("CSV: \"", f.Name(), "\" was created")
}

func FF(x float64) string{
	return strconv.FormatFloat(x, 'f', 2, 64)
}