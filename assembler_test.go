package main

import (
    "testing"
)

func Test_asbNOP(t *testing.T) {
    tokens := []string{"NOP"}
    code, err := asbNOP(tokens)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if code != 0 {
        t.Errorf("Wrong command code\nExpected 0, but got: %d", code)
    }
}

func Test_asbLTM(t *testing.T) {
    tokens := []string{"LTM", "5", "6"}
    code, err := asbLTM(tokens)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    if (code >> 28) != 1 {
        t.Errorf("Wrong command code\nExpected 1, but got: %d", code>>28)
    }
    if (code&0x0FFC0000 >> 18) != 5 {
        t.Errorf("Wrong adr_m\nExpected 5, but got: %d", code&0x0FFC0000 >> 18)
    }
    if (code&0x0003FF00 >> 8) != 6 {
        t.Errorf("Wrong literal\nExpected 6, but got: %d", code&0x0003FF00 >> 8)
    }
}
func Test_asbLTMError(t *testing.T) {
    tokens := []string{"LTM", "5", "6", ""}
    _, err := asbLTM(tokens)
    if err == nil {
        t.Errorf("Expected error: %v", err)
    }
}

func Test_asbMTR(t *testing.T) {
    tokens := []string{"MTR", "2", "3"}
    code, err := asbMTR(tokens)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    if (code >> 28) != 2 {
        t.Errorf("Wrong command code\nExpected 2, but got: %d", code>>28)
    }
    if (code&0x0FFC0000 >> 18) != 2 {
        t.Errorf("Wrong adr_m\nExpected 2, but got: %d", code&0x0FFC0000 >> 18)
    }
    if (code&0x0003C000 >> 14) != 3 {
        t.Errorf("Wrong adr_r\nExpected 3, but got: %d", code&0x0003C000 >> 14)
    }
}

func Test_asbRTR(t *testing.T) {
    tokens := []string{"RTR", "3", "3"}
    code, err := asbRTR(tokens)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    if (code >> 28) != 3 {
        t.Errorf("Wrong command code\nExpected 3, but got: %d", code>>28)
    }
    if (code&0x0F000000 >> 24) != 3 {
        t.Errorf("Wrong adr_r1\nExpected 2, but got: %d", code&0x0F000000 >> 24)
    }
    if (code&0x00F00000 >> 20) != 3 {
        t.Errorf("Wrong adr_r2\nExpected 3, but got: %d", code&0x00F00000 >> 20)
    }
}

func Test_asbSUB(t *testing.T) {
    tokens := []string{"SUB", "3", "1", "2"}
    code, err := asbSUB(tokens)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    if (code >> 28) != 4 {
        t.Errorf("Wrong command code\nExpected 4, but got: %d", code>>28)
    }
    if (code&0x0F000000 >> 24) != 3 {
        t.Errorf("Wrong adr_r1\nExpected 3, but got: %d", code&0x0F000000 >> 24)
    }
    if (code&0x00F00000 >> 20) != 1 {
        t.Errorf("Wrong adr_r2\nExpected 1, but got: %d", code&0x00F00000 >> 20)
    }
    if (code&0x000F0000 >> 16) != 2 {
        t.Errorf("Wrong adr_r3\nExpected 2, but got: %d", code&0x000F0000 >> 16)
    }
}

func Test_asbJUMP_LESS(t *testing.T) {
    tokens := []string{"JUMP_LESS", "4", "2", "3"}
    code, err := asbJUMP_LESS(tokens)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    if (code >> 28) != 5 {
        t.Errorf("Wrong command code\nExpected 5, but got: %d", code>>28)
    }
    if (code&0x0F000000 >> 24) != 4 {
        t.Errorf("Wrong adr_r1\nExpected 3, but got: %d", code&0x0F000000 >> 24)
    }
    if (code&0x00F00000 >> 20) != 2 {
        t.Errorf("Wrong adr_r2\nExpected 1, but got: %d", code&0x00F00000 >> 20)
    }
    if (code&0x000FFC00 >> 10) != 3 {
        t.Errorf("Wrong adr_to_jump\nExpected 3, but got: %d", code&0x000FFC00 >> 10)
    }
}

func Test_asbMTRK(t *testing.T) {
    tokens := []string{"MTRK", "5", "2"}
    code, err := asbMTRK(tokens)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    if (code >> 28) != 6 {
        t.Errorf("Wrong command code\nExpected 6, but got: %d", code>>28)
    }
    if (code&0x0F000000 >> 24) != 5 {
        t.Errorf("Wrong adr_r1\nExpected 5, but got: %d", code&0x0F000000 >> 24)
    }
    if (code&0x00F00000 >> 20) != 2 {
        t.Errorf("Wrong adr_r2\nExpected 2, but got: %d", code&0x00F00000 >> 20)
    }
}

func Test_asbRTMK(t *testing.T) {
    tokens := []string{"RTMK", "2", "5"}
    code, err := asbRTMK(tokens)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    if (code >> 28) != 7 {
        t.Errorf("Wrong command code\nExpected 7, but got: %d", code>>28)
    }
    if (code&0x0F000000 >> 24) != 2 {
        t.Errorf("Wrong adr_r1\nExpected 2, but got: %d", code&0x0F000000 >> 24)
    }
    if (code&0x00F00000 >> 20) != 5 {
        t.Errorf("Wrong adr_r2\nExpected 5, but got: %d", code&0x00F00000 >> 20)
    }
}

func Test_asbJMP(t *testing.T) {
    tokens := []string{"JMP", "10"}
    code, err := asbJMP(tokens)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    if (code >> 28) != 8 {
        t.Errorf("Wrong command code\nExpected 8, but got: %d", code>>28)
    }
    if (code&0x0FFC0000 >> 18) != 10 {
        t.Errorf("Wrong adr_to_jump\nExpected 10, but got: %d", code&0x0FFC0000 >> 18)
    }
}

func Test_asbSUM(t *testing.T) {
    tokens := []string{"SUM", "5", "3", "4"}
    code, err := asbSUM(tokens)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    if (code >> 28) != 9 {
        t.Errorf("Wrong command code\nExpected 9, but got: %d", code>>28)
    }
    if (code&0x0F000000 >> 24) != 5 {
        t.Errorf("Wrong adr_r1\nExpected 5, but got: %d", code&0x0F000000 >> 24)
    }
    if (code&0x00F00000 >> 20) != 3 {
        t.Errorf("Wrong adr_r2\nExpected 3, but got: %d", code&0x00F00000 >> 20)
    }
    if (code&0x000F0000 >> 16) != 4 {
        t.Errorf("Wrong adr_r3\nExpected 4, but got: %d", code&0x000F0000 >> 16)
    }
}
