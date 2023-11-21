package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	/*if len(os.Args) != 2 {
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("erro, não e numero valido")
		os.Exit(1)
	}
	fmt.Println("numero convertido", n)

	/*pessoa := &funcionarios.Pessoa{
		nome:    "Marcio",
		idade:   44,
		salario: 100,
	}

	salFuncs := make(map[string]int)
	salFuncs["glaucio"] = 10
	salFuncs["marcio"] = 11

	sal, exists := salFuncs["marcio"]
	fmt.Println(sal, exists)
	totalSal := len(salFuncs)
	fmt.Println(totalSal)

	/*salarios := []int{}
	//salarios := make([]int, 5)

	for i := 0; i < 10; i++ {
		salarios = append(salarios, 100+i)
	}
	for _, salario := range salarios {
		fmt.Println(salario)
	}

	/*pessoa2 := new(Pessoa)
	pessoa2.nome = "pessoa2"
	pessoa2.idade = 22
	pessoa2.salario = 200
	pessoa := &Pessoa{
		nome:    "Marcio",
		idade:   44,
		salario: 100,

	}
	pessoa.addSalaryPessoa(10)
	fmt.Println(pessoa.salario)

	/*name, salario := "Marcio", 100
	setName(name)
	newSalary, bonus := addSalary(salario, 10)
	fmt.Println("Novo Salario", newSalary)
	fmt.Println("Bonus", bonus)
	*/

}

func setName(name string) {
	fmt.Println("Hello", name)
}

func addSalary(valorAtual int, bonus int) (int, int) {
	return valorAtual + bonus, bonus
}

func getName() string {
	return "Marcio"
}
