package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Tyulenb/Pennywise600/assembler"
	"github.com/Tyulenb/Pennywise600/pennywise600"
)

func main() {
    args := os.Args
    path := "program.txt"
    debugMode := false
    if len(args) >= 2 {
        path = args[1]
        if len(args) > 2 && args[2] == "d" {
            debugMode = true
        }
    }else {
        fmt.Println("FORMAT cmd.go 'path to your program' 'd (optionaly for debug)'\n"+
        "go run cmd.go program.txt\ngo run cmd.go program.txt d (for debug)")
        return
    }
    p := pennywise600.NewPennywise600()
    code, err := assembler.Assemble(path)
    if err != nil {
        log.Fatalf("Error during code assembling: %v", err)
    }
    p.Load(code)
    if debugMode {
        Debug(p)
    }else {
        Run(p)
    }
}

func Run(p *pennywise600.Pennywise600) {
    for i := 0; i < 1000; i++ {
        p.EmulateCycle()
    }
    mem := p.GetMem()
    fmt.Println("MEM[0:10]",mem[0:10])
}

func Debug(p *pennywise600.Pennywise600) {
    fmt.Println("s - step for next command\ne - to exit")
    var arg string = "s"
    commands := map[uint32]string{
        0x0: "NOP",
        0x1: "LTM",
        0x2: "MTR",
        0x3: "RTR",
        0x4: "SUB",
        0x5: "JUMP_LESS",
        0x6: "MTRK",
        0x7: "RTMK",
        0x8: "JMP",       
        0x9: "SUM",     
    }
    for true { 
        fmt.Scanln(&arg)
        if arg == "\r" {
            arg = "s"
        }
        switch(arg) {
        case "s":
            fmt.Println("PC:", p.GetPc())
            fmt.Println("Cur command", commands[p.GetCurCommand()>>28])
            p.EmulateCycle()
            fmt.Println("Cycle Results:")
            mem := p.GetMem()
            fmt.Println("MEM[0:10]",mem[0:10])
            fmt.Println("REGS:", p.RF)
        case "e":
            return
        }
    }
}
