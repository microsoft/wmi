// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package controller

type SCSIControllerSettingsCollection []*SCSIControllerSettings

func (scsis *SCSIControllerSettingsCollection) Close() (err error) {
	for _, value := range *scsis {
		value.Close()
	}
	return nil
}

func (scsis *SCSIControllerSettingsCollection) String() string {
	return ""
}
