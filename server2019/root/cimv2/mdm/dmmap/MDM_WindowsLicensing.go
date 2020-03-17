// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_WindowsLicensing struct
type MDM_WindowsLicensing struct {
	cim.WmiInstance

	//
	Edition int32

	//
	InstanceID string

	//
	LicenseKeyType string

	//
	ParentID string

	//
	Status int32
}

// SetEdition sets the value of Edition for the instance
func (instance *MDM_WindowsLicensing) SetPropertyEdition(value int32) (err error) {
	return instance.SetProperty("Edition", value)
}

// GetEdition gets the value of Edition for the instance
func (instance *MDM_WindowsLicensing) GetPropertyEdition() (value int32, err error) {
	retValue, err := instance.GetProperty("Edition")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_WindowsLicensing) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_WindowsLicensing) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLicenseKeyType sets the value of LicenseKeyType for the instance
func (instance *MDM_WindowsLicensing) SetPropertyLicenseKeyType(value string) (err error) {
	return instance.SetProperty("LicenseKeyType", value)
}

// GetLicenseKeyType gets the value of LicenseKeyType for the instance
func (instance *MDM_WindowsLicensing) GetPropertyLicenseKeyType() (value string, err error) {
	retValue, err := instance.GetProperty("LicenseKeyType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_WindowsLicensing) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_WindowsLicensing) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *MDM_WindowsLicensing) SetPropertyStatus(value int32) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MDM_WindowsLicensing) GetPropertyStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="param" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_WindowsLicensing) UpgradeEditionWithProductKeyMethod( /* IN */ param string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("UpgradeEditionWithProductKeyMethod", param)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="param" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_WindowsLicensing) ChangeProductKeyMethod( /* IN */ param string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ChangeProductKeyMethod", param)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="param" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_WindowsLicensing) UpgradeEditionWithLicenseMethod( /* IN */ param string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("UpgradeEditionWithLicenseMethod", param)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="param" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_WindowsLicensing) CheckApplicabilityMethod( /* IN */ param string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CheckApplicabilityMethod", param)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
