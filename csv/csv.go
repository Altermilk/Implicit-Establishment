package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
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