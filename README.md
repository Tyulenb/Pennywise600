# Pennywise600 - Simple Processor Emulator

Pennywise600 is a simple processor emulator designed for educational purposes, allowing users to understand basic processor operations and commands.

## Resource Description

| Resource | Size                    | Description                        |
|----------|-------------------------|------------------------------------|
| cmd_mem  | 4 bytes, 1024 locations | Storage for commands               |
| mem      | 2 bytes, 1024 locations | Regular memory                     |
| RF       | 2 bytes, 16 locations   | Register File                      |
| pc       | 2 bytes                 | Program Counter                    |
| cmd_reg  | 4 bytes                 | Current command                    |

## Commands Description

| Command     | Description                              | Format                          | Pseudocode                                         |
|-------------|------------------------------------------|----------------------------------|---------------------------------------------------|
| `NOP`       | Do Nothing                               | [OpCode]                        | void do(){}                                       |
| `LTM`       | Load literal to memory                   | [OpCode][adr_m][literal]       | mem[adr_m] = literal                              |
| `MTR`       | Load memory to register                  | [OpCode][adr_m1][adr_r1]       | RF[adr_r1] = mem[adr_m1]                         |
| `RTR`       | Load register to register                | [OpCode][adr_r1][adr_r2]       | RF[adr_r1] = RF[adr_r2]                           |
| `SUB`       | Subtract                                 | [OpCode][adr_r1][adr_r2][adr_r3]| RF[adr_r3] = RF[adr_r1] - RF[adr_r2]              |
| `JUMP_LESS` | Jump to another operation on condition    | [OpCode][adr_r1][adr_r2][adr_to_jump]| if(RF[adr_r1] >= RF[adr_r2]) { pc = adr_to_jump } |
| `MTRK`      | Load register from memory by address from register | [OpCode][adr_r1][adr_r2] | RF[adr_r1] = mem[RF[adr_r2]]                       |
| `RTMK`      | Load memory from register by address from register | [OpCode][adr_r1][adr_r2] | mem[RF[adr_r1]] = RF[adr_r2]                       |
| `JMP`       | Jump to another operation                | [OpCode][adr_to_jump]          | pc = adr_to_jump                                   |
| `SUM`       | Add                                     | [OpCode][adr_r1][adr_r2][adr_r3]| RF[adr_r3] = RF[adr_r2]+RF[adr_r3]|

## Installation
### Requirements
Go version 1.20 or higher
### Steps
```bash
git clone https://github.com/Tyulenb/Pennywise600.git
cd Pennywise600
go run cmd/cmd.go "path to your program"
```
### Debug Mode
To enable debug mode, you can use the d flag:
```bash
go run cmd/cmd.go "path to your program" d
