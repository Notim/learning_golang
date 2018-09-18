package main

	import (
		io "fmt"
		sys "os"
	)

	func main(){
		var hi, hf, cont int

		io.Scanf("%d %d", &hi, &hf)
		cont = (24 - hi) + hf

		if(cont > 24){
			cont = cont - 24
		}

		io.Printf("O JOGO DUROU %d HORA(S)\n", cont);
		sys.Exit(0)
	}