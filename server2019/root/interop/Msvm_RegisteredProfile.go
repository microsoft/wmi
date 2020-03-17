// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Interop
//////////////////////////////////////////////
package interop

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_RegisteredProfile struct
type Msvm_RegisteredProfile struct {
	CIM_RegisteredProfile
}

func (instance *Msvm_RegisteredProfile) GetRelatedRegisteredProfile() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_RegisteredProfile")
}
