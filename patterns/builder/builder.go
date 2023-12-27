package main

import "fmt"

type Computer struct {
	CPU       string
	Memory    string
	HardDrive string
	GPU       string
}
type ComputerBuilder struct {
	computer Computer
}

func NewComputerBuilder() *ComputerBuilder {
	return &ComputerBuilder{computer: Computer{}}
}

func (b *ComputerBuilder) SetCPU(cpu string) *ComputerBuilder {
	b.computer.CPU = cpu
	return b
}

func (b *ComputerBuilder) SetMemory(memory string) *ComputerBuilder {
	b.computer.Memory = memory
	return b
}

func (b *ComputerBuilder) SetHardDrive(hardDrive string) *ComputerBuilder {
	b.computer.HardDrive = hardDrive
	return b
}

func (b *ComputerBuilder) SetGPU(gpu string) *ComputerBuilder {
	b.computer.GPU = gpu
	return b
}

func (b *ComputerBuilder) Build() Computer {
	return b.computer
}
func main() {
	builder := NewComputerBuilder()
	gamingPC := builder.SetCPU("Intel Core i9").SetMemory("32GB").SetHardDrive("1TB SSD").SetGPU("NVIDIA RTX 3080").Build()
	officePC := builder.SetCPU("Intel Core i5").SetMemory("16GB").SetHardDrive("512GB SSD").Build()
	fmt.Println(gamingPC)
	fmt.Println(officePC)
}
