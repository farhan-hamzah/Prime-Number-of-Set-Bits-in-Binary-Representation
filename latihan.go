package main
import "fmt"

func main() {
	var n, m, x, biner, pangkat, sisa, digit1, isPrime, c int
	var b bool
	fmt.Scan(&n, &m)
	for n <= m {
		x = n
		biner = 0      
		pangkat = 1    
		digit1 = 0 // reset di awal iterasi

		for x > 0 {
			sisa = x % 2
			if sisa != 0 {
				biner += sisa * pangkat
			}
			x = x / 2
			pangkat *= 10
		}

		temp := biner // simpan salinan biner agar nilai aslinya tidak rusak
		for temp > 0 {
			if temp%10 == 1 {
				digit1++
			}
			temp = temp / 10
		}

		b = true
		if digit1 < 2 {
			b = false
		} else {
			for c = 2; c*c <= digit1; c++ {
				if digit1%c == 0 {
					b = false
					break
				}
			}
		}

		if b == true {
			isPrime++
		}
		n++
	}
	fmt.Print(isPrime)
}