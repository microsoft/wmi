// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package drive

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type DvdDriveCollection []*DvdDrive

func NewDvdDriveCollection(instances []*wmi.WmiInstance) (col DvdDriveCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewDvdDrive(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (dvdcol *DvdDriveCollection) Close() (err error) {
	for _, value := range *dvdcol {
		value.Close()
	}
	return nil
}
