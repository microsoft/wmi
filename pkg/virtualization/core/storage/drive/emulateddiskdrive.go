// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package drive

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type EmulatedDiskDrive struct {
	*VirtualDrive
}

// NewEmulatedDiskDrive
func NewEmulatedDiskDrive(instance *wmi.WmiInstance) (*EmulatedDiskDrive, error) {
	wmivm, err := NewVirtualDrive(instance)
	if err != nil {
		return nil, err
	}
	return &EmulatedDiskDrive{wmivm}, nil
}
