package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var commands = map[string]func([]string) (uint32, error){
	"NOP":       asbNOP,       //0x0
	"LTM":       asbLTM,       //0x1
	"MTR":       asbMTR,       //0x2
	"RTR":       asbRTR,       //0x3
	"SUB":       asbSUB,       //0x4
	"JUMP_LESS": asbJUMP_LESS, //0x5
	"MTRK":      asbMTRK,      //0x6
	"RTMK":      asbRTMK,      //0x7
	"JMP":       asbJMP,       //0x8
	"SUM":       asbSUM,       //0x9
}

func sep(c rune) bool {
	return !unicode.IsNumber(c) && !unicode.IsLetter(c) && c != '_'
}

func assemble(path string) ([]uint32, error) {
	assembler, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	machineCode := make([]uint32, 0)

	scanner := bufio.NewScanner(assembler)
	lineNumber := 1
	for scanner.Scan() {
		str := scanner.Text()
		cmdsLine := strings.Split(str, ";")
		for _, cmd := range cmdsLine {
			if cmd == "" {
				continue
			}
			tokens := strings.FieldsFunc(cmd, sep)
			code, err := commands[tokens[0]](tokens)
			if err != nil {
				return nil, fmt.Errorf("Error: %v in line %d", err, lineNumber)
			}
			machineCode = append(machineCode, code)
		}
		lineNumber++
	}

	return machineCode, nil
}

func asbNOP(tokens []string) (uint32, error) {
	if len(tokens) == 1 {
		return 0, nil
	} else {
		return 0, fmt.Errorf("Unexpected amount of operands for NOP command, expected 1, but got %v", len(tokens))
	}
}

func asbLTM(tokens []string) (uint32, error) {
	if len(tokens) == 3 {
		adr_m, err := strconv.ParseUint(tokens[1], 10, 10)
		if err != nil {
			return 0, err
		}
		literal, err := strconv.ParseUint(tokens[2], 10, 10)
		if err != nil {
			return 0, err
		}
		code := uint32((0x1 << 28) | (adr_m << 18) | (literal << 8))
		return code, nil
	} else {
		return 0, fmt.Errorf("Unexpected amount of operands for LTM command, expected 3, but got %v", len(tokens))
	}
}

func asbMTR(tokens []string) (uint32, error) {
	if len(tokens) == 3 {
		adr_m, err := strconv.ParseUint(tokens[1], 10, 10)
		if err != nil {
			return 0, err
		}
		adr_r, err := strconv.ParseUint(tokens[2], 10, 4)
		if err != nil {
			return 0, err
		}
		code := uint32((0x2 << 28) | (adr_m << 18) | (adr_r << 14))
		return code, nil
	} else {
		return 0, fmt.Errorf("Unexpected amount of operands for MTR command, expected 3, but got %v", len(tokens))
	}
}

func asbRTR(tokens []string) (uint32, error) {
	if len(tokens) == 3 {
		adr_m, err := strconv.ParseUint(tokens[1], 10, 4)
		if err != nil {
			return 0, err
		}
		adr_r, err := strconv.ParseUint(tokens[2], 10, 4)
		if err != nil {
			return 0, err
		}
		code := uint32((0x3 << 28) | (adr_m << 24) | (adr_r << 20))
		return code, nil
	} else {
		return 0, fmt.Errorf("Unexpected amount of operands for RTR command, expected 3, but got %v", len(tokens))
	}
}

func asbSUB(tokens []string) (uint32, error) {
	if len(tokens) == 4 {
		adr_r1, err := strconv.ParseUint(tokens[1], 10, 4)
		if err != nil {
			return 0, err
		}
		adr_r2, err := strconv.ParseUint(tokens[2], 10, 4)
		if err != nil {
			return 0, err
		}
		adr_r3, err := strconv.ParseUint(tokens[3], 10, 4)
		if err != nil {
			return 0, err
		}
		code := uint32((0x4 << 28) | (adr_r1 << 24) | (adr_r2 << 20) | (adr_r3 << 16))
		return code, nil
	} else {
		return 0, fmt.Errorf("Unexpected amount of operands for SUB command, expected 4, but got %v", len(tokens))
	}
}

func asbJUMP_LESS(tokens []string) (uint32, error) {
	if len(tokens) == 4 {
		adr_r1, err := strconv.ParseUint(tokens[1], 10, 4)
		if err != nil {
			return 0, err
		}
		adr_r2, err := strconv.ParseUint(tokens[2], 10, 4)
		if err != nil {
			return 0, err
		}
		adr_to_jmp, err := strconv.ParseUint(tokens[3], 10, 10)
		if err != nil {
			return 0, err
		}
		code := uint32((0x5 << 28) | (adr_r1 << 24) | (adr_r2 << 20) | (adr_to_jmp << 10))
		return code, nil
	} else {
		return 0, fmt.Errorf("Unexpected amount of operands for JUMP_LESS command, expected 4, but got %v", len(tokens))
	}
}

func asbMTRK(tokens []string) (uint32, error) {
	if len(tokens) == 3 {
		adr_r1, err := strconv.ParseUint(tokens[1], 10, 4)
		if err != nil {
			return 0, err
		}
		adr_r2, err := strconv.ParseUint(tokens[2], 10, 4)
		if err != nil {
			return 0, err
		}
		code := uint32((0x6 << 28) | (adr_r1 << 24) | (adr_r2 << 20))
		return code, nil
	} else {
		return 0, fmt.Errorf("Unexpected amount of operands for MTRK command, expected 3, but got %v", len(tokens))
	}
}

func asbRTMK(tokens []string) (uint32, error) {
	if len(tokens) == 3 {
		adr_r1, err := strconv.ParseUint(tokens[1], 10, 4)
		if err != nil {
			return 0, err
		}
		adr_r2, err := strconv.ParseUint(tokens[2], 10, 4)
		if err != nil {
			return 0, err
		}
		code := uint32((0x7 << 28) | (adr_r1 << 24) | (adr_r2 << 20))
		return code, nil
	} else {
		return 0, fmt.Errorf("Unexpected amount of operands for RTMK command, expected 3, but got %v", len(tokens))
	}
}

func asbJMP(tokens []string) (uint32, error) {
	if len(tokens) == 2 {
		adr_to_jmp, err := strconv.ParseUint(tokens[1], 10, 10)
		if err != nil {
			return 0, err
		}
		code := uint32((0x8 << 28) | (adr_to_jmp << 18))
		return code, nil
	} else {
		return 0, fmt.Errorf("Unexpected amount of operands for JMP command, expected 2, but got %v", len(tokens))
	}
}

func asbSUM(tokens []string) (uint32, error) {
	if len(tokens) == 4 {
		adr_r1, err := strconv.ParseUint(tokens[1], 10, 4)
		if err != nil {
			return 0, err
		}
		adr_r2, err := strconv.ParseUint(tokens[2], 10, 4)
		if err != nil {
			return 0, err
		}
		adr_r3, err := strconv.ParseUint(tokens[3], 10, 4)
		if err != nil {
			return 0, err
		}
		code := uint32((0x9 << 28) | (adr_r1 << 24) | (adr_r2 << 20) | (adr_r3 << 16))
		return code, nil
	} else {
		return 0, fmt.Errorf("Unexpected amount of operands for SUM command, expected 4, but got %v", len(tokens))
	}
}

func main() {
	fmt.Println(assemble("program.txt"))
}
