package main

import "fmt"

//struct of processor "Pennywise600"
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

func NewPennywise600() *Pennywise600 {
    //TO DO initialization
    return &Pennywise600{}
}

func (p *Pennywise600) EmulateCycle() {
    p.cmd_reg = p.cmd_mem[p.pc]
    p.pc += 1
    //to do command routing
    if ((p.cmd_reg&0xF0000000)>>28 == 0x3) {}
}

//DO NOTHING
func (p *Pennywise600) NOP() {
}

//Loading literal to main memory
func (p *Pennywise600) LTM() {
    adr_m := (p.cmd_reg&0x0FFC0000) >> 18
    literal := (p.cmd_reg&0x0003FF00) >> 8
    p.mem[adr_m] = uint16(literal)
}

//Load data from memory to register
func (p *Pennywise600) MTR() {
    adr_m := (p.cmd_reg&0x0FFC0000) >> 18
    adr_r := (p.cmd_reg&0x0003C000) >> 14
    p.RF[adr_r] = p.mem[adr_m]
}

//Load data from register to register
func (p *Pennywise600) RTR() {
    adr_r1 := (p.cmd_reg&0x0F000000) >> 24
    adr_r2 := (p.cmd_reg&0x00F00000) >> 20
    p.RF[adr_r1] = p.RF[adr_r2]
}

//Load data to register from memory by address written in register
func (p *Pennywise600) MTRK () {
    adr_r1 := (p.cmd_reg&0x0F000000) >> 24
    adr_r2 := (p.cmd_reg&0x00F00000) >> 20
    p.RF[adr_r1] = p.mem[p.RF[adr_r2]]
}

//Load data to memory by address written in register from register
func (p *Pennywise600) RTMK () {
    adr_r1 := (p.cmd_reg&0x0F000000) >> 24
    adr_r2 := (p.cmd_reg&0x00F00000) >> 20
    p.mem[p.RF[adr_r1]] = p.RF[adr_r2]
}

//SUBTRACTION
func (p *Pennywise600) SUB() {
    adr_r1 := (p.cmd_reg&0x0F000000) >> 24
    adr_r2 := (p.cmd_reg&0x00F00000) >> 20
    adr_r3 := (p.cmd_reg&0x000F0000) >> 16 
    p.RF[adr_r3] = p.RF[adr_r1] - p.RF[adr_r2]
}

//SUM
func (p *Pennywise600) SUM() {
    adr_r1 := (p.cmd_reg&0x0F000000) >> 24
    adr_r2 := (p.cmd_reg&0x00F00000) >> 20
    adr_r3 := (p.cmd_reg&0x000F0000) >> 16 
    p.RF[adr_r3] = p.RF[adr_r1] + p.RF[adr_r2]
}

//Jump to another command by condition
func (p *Pennywise600) JUMP_LESS() {
    adr_r1 := (p.cmd_reg&0x0F000000) >> 24
    adr_r2 := (p.cmd_reg&0x00F00000) >> 20
    adr_to_jump := (p.cmd_reg&0x000FFC00) >> 10
    if p.RF[adr_r1] < p.RF[adr_r2] {
        p.pc = p.pc+1
    }else {
        p.pc = uint16(adr_to_jump)
    }
}

//Jump to another command
func (p *Pennywise600) JMP() {
    adr_to_jump := (p.cmd_reg&0x0FFC0000) >> 18
    p.pc = uint16(adr_to_jump)
}

func main(){
}
