package main

	import (
		io "fmt"
		sys "os"
	)

	func main(){
		var percent, salary, ValueOfFinalSalary float64

		io.Scanf("%f", &salary)

		if(salary > 0) && (salary < 400.01){
			percent = 15
		}else if(salary > 400) && (salary < 800.01){
			percent = 12
		}else if(salary > 800) && (salary < 1200.01){
			percent = 10
		}else if(salary > 1200) && (salary < 2000.01){
			percent = 7
		}else if(salary > 2000){
			percent = 4
		}
		ValueOfFinalSalary = (percent / 100.0) * salary

		io.Printf("Novo salario: %.2f\n",(ValueOfFinalSalary+salary))
		io.Printf("Reajuste ganho: %.2f\n",(ValueOfFinalSalary))
		io.Printf("Em percentual: %.0f %%\n",(percent))

		sys.Exit(0)
	}