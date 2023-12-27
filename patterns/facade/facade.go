package main

import "fmt"

// Процессор
type CPU struct{}

func (c *CPU) Start() {
	fmt.Println("CPU is starting")
}

// Память
type Memory struct{}

func (m *Memory) Load() {
	fmt.Println("Memory is loading data")
}

// Жесткий диск
type HardDrive struct{}

func (h *HardDrive) Read() {
	fmt.Println("Reading data from disk")
}

// Компьютерный фасад
type ComputerFacade struct {
	cpu       CPU
	memory    Memory
	hardDrive HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       CPU{},
		memory:    Memory{},
		hardDrive: HardDrive{},
	}
}

func (f *ComputerFacade) Start() {
	f.cpu.Start()
	f.memory.Load()
	f.hardDrive.Read()
}

func main() {
	computer := NewComputerFacade()
	computer.Start()
}
