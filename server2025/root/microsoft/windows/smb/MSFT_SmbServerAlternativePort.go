// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_SmbServerAlternativePort struct
type MSFT_SmbServerAlternativePort struct {
	*cim.WmiInstance

	//
	Instances SmbServerAlternativePort_Instances

	//
	Port uint16

	//
	TransportType SmbServerAlternativePort_TransportType
}

func NewMSFT_SmbServerAlternativePortEx1(instance *cim.WmiInstance) (newInstance *MSFT_SmbServerAlternativePort, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbServerAlternativePort{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_SmbServerAlternativePortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_SmbServerAlternativePort, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbServerAlternativePort{
		WmiInstance: tmp,
	}
	return
}

// SetInstances sets the value of Instances for the instance
func (instance *MSFT_SmbServerAlternativePort) SetPropertyInstances(value SmbServerAlternativePort_Instances) (err error) {
	return instance.SetProperty("Instances", (value))
}

// GetInstances gets the value of Instances for the instance
func (instance *MSFT_SmbServerAlternativePort) GetPropertyInstances() (value SmbServerAlternativePort_Instances, err error) {
	retValue, err := instance.GetProperty("Instances")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = SmbServerAlternativePort_Instances(valuetmp)

	return
}

// SetPort sets the value of Port for the instance
func (instance *MSFT_SmbServerAlternativePort) SetPropertyPort(value uint16) (err error) {
	return instance.SetProperty("Port", (value))
}

// GetPort gets the value of Port for the instance
func (instance *MSFT_SmbServerAlternativePort) GetPropertyPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("Port")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetTransportType sets the value of TransportType for the instance
func (instance *MSFT_SmbServerAlternativePort) SetPropertyTransportType(value SmbServerAlternativePort_TransportType) (err error) {
	return instance.SetProperty("TransportType", (value))
}

// GetTransportType gets the value of TransportType for the instance
func (instance *MSFT_SmbServerAlternativePort) GetPropertyTransportType() (value SmbServerAlternativePort_TransportType, err error) {
	retValue, err := instance.GetProperty("TransportType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = SmbServerAlternativePort_TransportType(valuetmp)

	return
}

//

// <param name="EnableInstances" type="uint32 "></param>
// <param name="Port" type="uint16 "></param>
// <param name="TransportType" type="uint32 "></param>

// <param name="CreatedAlternativePort" type="MSFT_SmbServerAlternativePort "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbServerAlternativePort) NewServerAlternativePort( /* IN */ TransportType uint32,
	/* IN */ Port uint16,
	/* IN */ EnableInstances uint32,
	/* OUT */ CreatedAlternativePort MSFT_SmbServerAlternativePort) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewServerAlternativePort", TransportType, Port, EnableInstances)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DisableInstances" type="uint32 "></param>
// <param name="EnableInstances" type="uint32 "></param>
// <param name="Port" type="uint16 "></param>
// <param name="TransportType" type="uint32 "></param>

// <param name="CreatedAlternativePort" type="MSFT_SmbServerAlternativePort "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbServerAlternativePort) SetServerAlternativePort( /* IN */ TransportType uint32,
	/* IN */ Port uint16,
	/* IN */ EnableInstances uint32,
	/* IN */ DisableInstances uint32,
	/* OUT */ CreatedAlternativePort MSFT_SmbServerAlternativePort) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetServerAlternativePort", TransportType, Port, EnableInstances, DisableInstances)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
