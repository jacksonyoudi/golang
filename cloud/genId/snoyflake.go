package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sony/sonyflake"
)


// 通过配置写到服务id
func readMachineIDFromLocalFile()  uint16 {
	return uint16(1)
}

func generateMachineID() (uint16, error) {
	return uint16(2), nil
}




func getMachineID() (uint16, error) {
	var machineID uint16
	var err error
	machineID = readMachineIDFromLocalFile()
	if machineID == 0 {
		machineID, err = generateMachineID()
		if err != nil {
			return 0, err
		}
	}

	return machineID, nil
}

func saddMachineIDToRedisSet() (uint16, error) {
	return 3, nil
}

func saveMachineIDToLocalFile(i uint16) error  {
	return nil
}


func checkMachineID(machineID uint16) bool {
	saddResult, err := saddMachineIDToRedisSet()
	if err != nil || saddResult == 0 {
		return true
	}

	err = saveMachineIDToLocalFile(machineID)
	if err != nil {
		return true
	}

	return false
}

func main() {
	t, _ := time.Parse("2006-01-02", "2018-01-01")
	settings := sonyflake.Settings{
		StartTime:      t,
		MachineID:      getMachineID,
		CheckMachineID: checkMachineID,
	}

	sf := sonyflake.NewSonyflake(settings)
	id, err := sf.NextID()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(id)
}