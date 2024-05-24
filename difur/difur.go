package difur

import (
	"fmt"
	"math"
)

const(
	Nx_end int = 1
)

type Difur struct {
	H float64
	Dt float64
	Nx int
	U [][]float64
	Eps float64
	Title string
}

func Cut(interval int, delta float64) int {
	return (int)(float64(interval) / delta) + 1
}

func (difur *Difur) SetN(){
	difur.Nx = Cut(Nx_end, difur.H)
}

func MakeDifur(H, Dt, eps float64) Difur{
	difur := Difur{
		H:  H,
		Dt: Dt,
		Eps: eps,
	}
	difur.SetN()
	return difur
}

func (d*Difur) AddLine(){
	d.U = append(d.U, make([]float64, d.Nx))
}

func (d *Difur) ImplicitEstablish(){
	d.AddLine()
	for j := 0; j < d.Nx; j++{
		d.U[0][j] = 4 * d.H * float64(j) - 1
	}

	a := - d.Dt/(d.H * d.H)
	b := 1 + d.Dt/d.H + 2*d.Dt/(d.H *d.H)
	c := -d.Dt/d.H - d.Dt/(d.H *d.H)
	n := 0

	finish := false

	for {
		if finish{
			break
		}
		
		d.AddLine()

		alpha := make([]float64, d.Nx - 1)
		beta := make([]float64, d.Nx - 1)

		alpha[0], beta[0] = 0, 0
		d.U[n+1][0] = 0

		for j := 1; j < d.Nx - 1; j++{
			psy := d.U[n][j] + d.Dt * (4*d.H*float64(j) - 1)
			if j > 1{
				alpha[j-1] = -a / (b + c*alpha[j-2])
				beta[j-1] = (psy - c*beta[j-2])/(b + c * alpha[j-2])
			}
			
		}

		d.U[n+1][d.Nx - 1] = 5

		for j := d.Nx - 2; j>0; j--{
			d.U[n+1][j] = alpha[j-1] * d.U[n+1][j+1] + beta[j-1]
		}

		if (NormaCheck(d.U, d.Nx, len(d.U), d.H, d.Eps, d.Title)) || n>5{
			
			finish = true
			break
		}
		n++
	}
}

func NormaCheck(u [][]float64, Nx, Nt int,  h, eps float64, title string) bool{
	var sum float64 = 0
	for j := 0; j<Nx; j++{
		for n:= 1; n<Nt; n++{
			sum += math.Pow(u[n][j] - u[n-1][j], 2)
		}
	}
	
	norm := math.Sqrt(h*sum)
	fmt.Println("-- norm for {",  title, "} = ", norm)
	return norm <= eps

}
