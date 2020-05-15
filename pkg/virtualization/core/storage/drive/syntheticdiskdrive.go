// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package drive

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type SyntheticDiskDrive struct {
	*VirtualDrive
}

// NewSyntheticDiskDrive
func NewSyntheticDiskDrive(instance *wmi.WmiInstance) (*SyntheticDiskDrive, error) {
	wmivm, err := NewVirtualDrive(instance)
	if err != nil {
		return nil, err
	}
	return &SyntheticDiskDrive{wmivm}, nil
}
