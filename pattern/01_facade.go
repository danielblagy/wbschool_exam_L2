package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type CPU struct {
	programCounter int
	running        bool
}

func (p *CPU) Freeze() {
	p.running = false
}

func (p *CPU) Jump(position int) {
	p.programCounter = position
}

func (p *CPU) Execute() {
	p.programCounter++
}

type HardDrive struct {
}

func (d HardDrive) Read(lba, size int) []byte {
	return make([]byte, size)
}

type Memory struct {
	data []byte
}

func BootMemory(memory *Memory) {
	memory.data = make([]byte, 512)
}

func (m *Memory) Load(position int) []byte {
	return m.data[position:]
}

// ComputerFacade uses different interdependent classes to perform the booting
type ComputerFacade struct {
	cpu       CPU
	hardDrive HardDrive
	memory    Memory
}

func (cf *ComputerFacade) Start() {
	bootAddress := 8
	bootSector := 128
	sectorSize := 256

	BootMemory(&cf.memory)

	cf.cpu.Freeze()
	copy(cf.hardDrive.Read(bootSector, sectorSize), cf.memory.Load(bootAddress))
	cf.cpu.Jump(bootAddress)
	cf.cpu.Execute()
}
