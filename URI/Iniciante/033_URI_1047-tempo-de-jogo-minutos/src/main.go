package main

	import (
		io "fmt"
		sys "os"
	)

	func main(){
		var (hi, mi, hf, mf, conth, contm int)

		io.Scanf("%d %d %d %d", &hi, &mi, &hf, &mf)

		conth, contm = 23 + (hf - hi), 60 + (mf - mi)

		if contm >= 60 {
			conth++
			contm = contm - 60
		}
		if(conth >= 24 )&& (contm != 0){
			conth = conth - 24
		}

		io.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", conth, contm);
		sys.Exit(0)
	}