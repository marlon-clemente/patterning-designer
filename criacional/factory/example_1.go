package main

import "fmt"

// Veiculo é a interface comum para todos os tipos de veículos
type Veiculo interface {
	Dirigir() string
}

// Carro é um tipo concreto de veículo
type Carro struct{}

func (c Carro) Dirigir() string {
	return "Dirigindo um carro"
}

// Moto é outro tipo concreto de veículo
type Moto struct{}

func (m Moto) Dirigir() string {
	return "Pilotando uma moto"
}

// Factory Method que cria um veículo com base no tipo desejado
func NovoVeiculo(tipo string) Veiculo {
	if tipo == "carro" {
		return Carro{}
	} else if tipo == "moto" {
		return Moto{}
	}
	return nil
}

func main() {
	// Criando um carro
	veiculo1 := NovoVeiculo("carro")
	fmt.Println(veiculo1.Dirigir())

	// Criando uma moto
	veiculo2 := NovoVeiculo("moto")
	fmt.Println(veiculo2.Dirigir())
}
