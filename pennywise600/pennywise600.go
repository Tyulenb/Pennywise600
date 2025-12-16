package pennywise600 

import "log"

// struct of processor "Pennywise600"
type Pennywise600 struct {
	//memory of commands
	cmd_mem [1024]uint32
	//main memory
	mem [1024]uint16
	//registers
	RF [16]uint16
	//program counter
	pc uint16
	//current command
	cmd_reg uint32
}

var commandMap = make(map[uint16]func())

func NewPennywise600() *Pennywise600 {
	//initialization
	p := &Pennywise600{
		pc:      0,
		cmd_reg: 0,
	}
    p.RF[1] = 1
	initCommandTable(p)
	return p
}

func (p *Pennywise600) EmulateCycle() {
    if int(p.pc) >= len(p.cmd_mem) {
        return
    }
	p.cmd_reg = p.cmd_mem[p.pc]
	p.pc += 1
    commandMap[uint16((p.cmd_reg&0xF0000000)>>28)]()
}

// DO NOTHING
func (p *Pennywise600) NOP() {
}

// Loading literal to main memory
func (p *Pennywise600) LTM() {
	adr_m := (p.cmd_reg & 0x0FFC0000) >> 18
	literal := (p.cmd_reg & 0x0003FF00) >> 8
	p.mem[adr_m] = uint16(literal)
}

// Load data from memory to register
func (p *Pennywise600) MTR() {
	adr_m := (p.cmd_reg & 0x0FFC0000) >> 18
	adr_r := (p.cmd_reg & 0x0003C000) >> 14
	p.RF[adr_r] = p.mem[adr_m]
}

// Load data from register to register
func (p *Pennywise600) RTR() {
	adr_r1 := (p.cmd_reg & 0x0F000000) >> 24
	adr_r2 := (p.cmd_reg & 0x00F00000) >> 20
	p.RF[adr_r1] = p.RF[adr_r2]
}

// Load data to register from memory by address written in register
func (p *Pennywise600) MTRK() {
	adr_r1 := (p.cmd_reg & 0x0F000000) >> 24
	adr_r2 := (p.cmd_reg & 0x00F00000) >> 20
	p.RF[adr_r1] = p.mem[p.RF[adr_r2]]
}

// Load data to memory by address written in register from register
func (p *Pennywise600) RTMK() {
	adr_r1 := (p.cmd_reg & 0x0F000000) >> 24
	adr_r2 := (p.cmd_reg & 0x00F00000) >> 20
	p.mem[p.RF[adr_r1]] = p.RF[adr_r2]
}

// SUBTRACTION
func (p *Pennywise600) SUB() {
	adr_r1 := (p.cmd_reg & 0x0F000000) >> 24
	adr_r2 := (p.cmd_reg & 0x00F00000) >> 20
	adr_r3 := (p.cmd_reg & 0x000F0000) >> 16
	p.RF[adr_r3] = p.RF[adr_r1] - p.RF[adr_r2]
}

// SUM
func (p *Pennywise600) SUM() {
	adr_r1 := (p.cmd_reg & 0x0F000000) >> 24
	adr_r2 := (p.cmd_reg & 0x00F00000) >> 20
	adr_r3 := (p.cmd_reg & 0x000F0000) >> 16
	p.RF[adr_r3] = p.RF[adr_r1] + p.RF[adr_r2]
}

// Jump to another command by condition
func (p *Pennywise600) JUMP_LESS() {
	adr_r1 := (p.cmd_reg & 0x0F000000) >> 24
	adr_r2 := (p.cmd_reg & 0x00F00000) >> 20
	adr_to_jump := (p.cmd_reg & 0x000FFC00) >> 10
	if p.RF[adr_r1] >= p.RF[adr_r2] {
		p.pc = uint16(adr_to_jump)
	}
}

// Jump to another command
func (p *Pennywise600) JMP() {
	adr_to_jump := (p.cmd_reg & 0x0FFC0000) >> 18
	p.pc = uint16(adr_to_jump)
}

func (p *Pennywise600) Load(code []uint32) {
    if len(code) > len(p.cmd_mem) {
        log.Fatal("Program is too long")
    }
    for i := range code {
        p.cmd_mem[i] = code[i]
    }
}

//SOME DEBUG PURPOSE FUNCTIONS
func (p *Pennywise600) GetMem() [1024]uint16 {
    return p.mem
}
func (p *Pennywise600) GetPc() uint16 {
    return p.pc
}
func (p *Pennywise600) GetCurCommand() uint32 {
    return p.cmd_mem[p.pc]
}


func initCommandTable(p *Pennywise600) {
	commandMap[0x0] = p.NOP
	commandMap[0x1] = p.LTM
	commandMap[0x2] = p.MTR
	commandMap[0x3] = p.RTR
	commandMap[0x4] = p.SUB
	commandMap[0x5] = p.JUMP_LESS
	commandMap[0x6] = p.MTRK
	commandMap[0x7] = p.RTMK
	commandMap[0x8] = p.JMP
	commandMap[0x9] = p.SUM
}

