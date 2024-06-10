// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSiSCSIInitiator_MethodClass struct
type MSiSCSIInitiator_MethodClass struct {
	*cim.WmiInstance

	//
	iSCSINodeName string
}

func NewMSiSCSIInitiator_MethodClassEx1(instance *cim.WmiInstance) (newInstance *MSiSCSIInitiator_MethodClass, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_MethodClass{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSIInitiator_MethodClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSIInitiator_MethodClass, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_MethodClass{
		WmiInstance: tmp,
	}
	return
}

// SetiSCSINodeName sets the value of iSCSINodeName for the instance
func (instance *MSiSCSIInitiator_MethodClass) SetPropertyiSCSINodeName(value string) (err error) {
	return instance.SetProperty("iSCSINodeName", (value))
}

// GetiSCSINodeName gets the value of iSCSINodeName for the instance
func (instance *MSiSCSIInitiator_MethodClass) GetPropertyiSCSINodeName() (value string, err error) {
	retValue, err := instance.GetProperty("iSCSINodeName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSiSCSIInitiator_MethodClass) RefreshTargetList() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RefreshTargetList")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="DestinationAddress" type="string "></param>
// <param name="InitiatorName" type="string "></param>
// <param name="InitiatorPortNumber" type="uint32 "></param>
// <param name="OuterModeAddress" type="string "></param>
// <param name="Persist" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSiSCSIInitiator_MethodClass) SetIScsiTunnelModeOuterAddress( /* IN */ InitiatorName string,
	/* IN */ InitiatorPortNumber uint32,
	/* IN */ DestinationAddress string,
	/* IN */ OuterModeAddress string,
	/* IN */ Persist bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetIScsiTunnelModeOuterAddress", InitiatorName, InitiatorPortNumber, DestinationAddress, OuterModeAddress, Persist)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="AuthInfo" type="MSiSCSIInitiator_IKEPresharedKeyAuthenticationInfo "></param>
// <param name="InitiatorName" type="string "></param>
// <param name="InitiatorPortNumber" type="uint32 "></param>
// <param name="Persist" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSiSCSIInitiator_MethodClass) SetIScsiIKEInfo( /* IN */ InitiatorName string,
	/* IN */ InitiatorPortNumber uint32,
	/* IN */ AuthInfo MSiSCSIInitiator_IKEPresharedKeyAuthenticationInfo,
	/* IN */ Persist bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetIScsiIKEInfo", InitiatorName, InitiatorPortNumber, AuthInfo, Persist)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="key" type="uint8 []"></param>
// <param name="Persist" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSiSCSIInitiator_MethodClass) SetIScsiGroupPresharedKey( /* IN */ key []uint8,
	/* IN */ Persist bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetIScsiGroupPresharedKey", key, Persist)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="SharedSecret" type="uint8 []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSiSCSIInitiator_MethodClass) SetIScsiInitiatorCHAPSharedSecret( /* IN */ SharedSecret []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetIScsiInitiatorCHAPSharedSecret", SharedSecret)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="SharedSecret" type="uint8 []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSiSCSIInitiator_MethodClass) SetIScsiInitiatorRADIUSSharedSecret( /* IN */ SharedSecret []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetIScsiInitiatorRADIUSSharedSecret", SharedSecret)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="InitiatorNodeName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSiSCSIInitiator_MethodClass) SetIScsiInitiatorNodeName( /* IN */ InitiatorNodeName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetIScsiInitiatorNodeName", InitiatorNodeName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSiSCSIInitiator_MethodClass) SetupPersistentIScsiVolumes() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetupPersistentIScsiVolumes")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSiSCSIInitiator_MethodClass) ClearPersistentIScsiVolumes() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ClearPersistentIScsiVolumes")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
