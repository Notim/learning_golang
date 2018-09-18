package main
	import io "fmt"

	func main(){
		var opcao1, opcao2, opcao3 string

		io.Scanf("%s\n%s\n%s", &opcao1, &opcao2, &opcao3)

		if(opcao1 == "vertebrado"){
			if(opcao2 == "ave"){
				if(opcao3 == "carnivoro"){
					io.Printf("aguia\n")
				}else{
					io.Printf("pomba\n")
				}
			}else{
				if(opcao3 == "onivoro"){
					io.Printf("homem\n")
				}else{
					io.Printf("vaca\n")
				}
			}
		}else{
			if(opcao2 == "inseto"){
				if(opcao3 == "hematofago"){
					io.Printf("pulga\n")
				}else{
					io.Printf("lagarta\n")
				}
			}else{
				if(opcao3 == "hematofago"){
					io.Printf("sanguessuga\n")
				}else{
					io.Printf("minhoca\n")
				}
			}
		}
	}