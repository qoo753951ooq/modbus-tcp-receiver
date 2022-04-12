package modbus

import (
	"fmt"
	"log"
	"modbus-tcp-receiver/util"
	"os"
	"time"

	"github.com/goburrow/modbus"
)

const function_code_01 = "01"
const function_code_02 = "02"
const function_code_03 = "03"

const read_coils = "ReadCoils"
const read_discrete_inputs = "ReadDiscreteInputs"
const read_holding_registers = "ReadHoldingRegisters"

const Show_binary_log = 1
const Show_decimal_log = 2

const timeout = 3 * time.Second
const slaveId = 1

var client modbus.Client

var handler *modbus.TCPClientHandler

func NewModbusTcp(ip string, showLog bool) error {

	handler = modbus.NewTCPClientHandler(ip)
	handler.Timeout = timeout
	handler.SlaveId = slaveId

	if showLog {
		handler.Logger = log.New(os.Stdout, "tcp:", log.LstdFlags)
	}

	connErr := handler.Connect()

	if connErr != nil {
		fmt.Printf("%v \n", connErr)
		return connErr
	}

	client = modbus.NewClient(handler)
	return nil
}

func GetReadCoils(address, quantity uint16) []int {
	var binary []int

	fmt.Printf("func code: %s method: %s hex: %d \n", function_code_01, read_coils, address+1)
	result, _ := client.ReadCoils(address, quantity)
	fmt.Printf("Coil bytes:  %v \n", result)

	binary = util.Bytes2Bits(result)
	fmt.Printf("Binary: %v \n", binary)

	return binary
}

func GetReadDiscreteInputs(address, quantity uint16) []int {

	var binary []int

	fmt.Printf("func code: %s method: %s hex: %d \n", function_code_02, read_discrete_inputs, address+1)
	result, _ := client.ReadDiscreteInputs(address, quantity)
	fmt.Printf("Discrete Input bytes: %v \n", result)

	binary = util.Bytes2Bits(result)
	fmt.Printf("Binary: %v \n", binary)

	return binary
}

func GetReadHoldingRegisters(address, quantity uint16) ([]int, []uint16) {

	var binary []int
	var decimal []uint16

	fmt.Printf("func code: %s method: %s hex: %d \n", function_code_03, read_holding_registers, address+1)
	result, _ := client.ReadHoldingRegisters(address, quantity)
	fmt.Printf("Holding Register bytes: %v \n", result)

	binary = util.Bytes2Bits(result)
	decimal = util.Bytes2Decimal(result)

	return binary, decimal
}

func CloseModBusTcp() {
	handler.Close()
}

func ShowModBusLog(address uint16, showlog int, logData interface{}) {

	switch showlog {
	case Show_binary_log:
		//fmt.Printf("Binary: %v \n", logData)
		util.Writelog(fmt.Sprintf("Address: %d Binary: %v", address, logData))
	case Show_decimal_log:
		//fmt.Printf("Decimal: %v \n", logData)
		util.Writelog(fmt.Sprintf("Address: %d Decimal: %v", address, logData))
	}
}
