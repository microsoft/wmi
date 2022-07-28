// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package pcie

import (
	"log"

	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type PcieDeviceCollection []*PcieDevice

func NewPcieDeviceCollection(instances []*wmi.WmiInstance) (col PcieDeviceCollection, err error) {
	log.Printf("[DHAVAL POPAT] Length of instances [%d]", len(instances))
	for _, inst := range instances {
		na, err1 := NewPcieDeviceEx1(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	log.Printf("[DHAVAL POPAT] Length of collection [%d]", len(col))
	return
}

func (vms *PcieDeviceCollection) Close() (err error) {
	for _, value := range *vms {
		value.Close()
	}
	return nil
}
