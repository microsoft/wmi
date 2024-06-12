// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MS_SM_ScsiInformationMethods struct
type MS_SM_ScsiInformationMethods struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string
}

func NewMS_SM_ScsiInformationMethodsEx1(instance *cim.WmiInstance) (newInstance *MS_SM_ScsiInformationMethods, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MS_SM_ScsiInformationMethods{
		WmiInstance: tmp,
	}
	return
}

func NewMS_SM_ScsiInformationMethodsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MS_SM_ScsiInformationMethods, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MS_SM_ScsiInformationMethods{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MS_SM_ScsiInformationMethods) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MS_SM_ScsiInformationMethods) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MS_SM_ScsiInformationMethods) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MS_SM_ScsiInformationMethods) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// <param name="Cdb" type="uint8 []"></param>
// <param name="DiscoveredPortWWN" type="uint8 []"></param>
// <param name="DomainPortWWN" type="uint8 []"></param>
// <param name="HbaPortWWN" type="uint8 []"></param>
// <param name="InRespBufferMaxSize" type="uint32 "></param>
// <param name="InSenseBufferMaxSize" type="uint32 "></param>
// <param name="SmhbaLUN" type="uint64 "></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="OutRespBufferSize" type="uint32 "></param>
// <param name="OutSenseBufferSize" type="uint32 "></param>
// <param name="RespBuffer" type="uint8 []"></param>
// <param name="ScsiStatus" type="uint8 "></param>
// <param name="SenseBuffer" type="uint8 []"></param>
func (instance *MS_SM_ScsiInformationMethods) SM_ScsiInquiry( /* IN */ HbaPortWWN []uint8,
	/* IN */ DiscoveredPortWWN []uint8,
	/* IN */ DomainPortWWN []uint8,
	/* IN */ SmhbaLUN uint64,
	/* IN */ Cdb []uint8,
	/* IN */ InRespBufferMaxSize uint32,
	/* IN */ InSenseBufferMaxSize uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ ScsiStatus uint8,
	/* OUT */ OutRespBufferSize uint32,
	/* OUT */ OutSenseBufferSize uint32,
	/* OUT */ RespBuffer []uint8,
	/* OUT */ SenseBuffer []uint8) (err error) {
	_, err = instance.InvokeMethod("SM_ScsiInquiry", HbaPortWWN, DiscoveredPortWWN, DomainPortWWN, SmhbaLUN, Cdb, InRespBufferMaxSize, InSenseBufferMaxSize)
	if err != nil {
		return
	}
	return

}

//

// <param name="Cdb" type="uint8 []"></param>
// <param name="DiscoveredPortWWN" type="uint8 []"></param>
// <param name="DomainPortWWN" type="uint8 []"></param>
// <param name="HbaPortWWN" type="uint8 []"></param>
// <param name="InRespBufferMaxSize" type="uint32 "></param>
// <param name="InSenseBufferMaxSize" type="uint32 "></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="OutRespBufferSize" type="uint32 "></param>
// <param name="OutSenseBufferSize" type="uint32 "></param>
// <param name="RespBuffer" type="uint8 []"></param>
// <param name="ScsiStatus" type="uint8 "></param>
// <param name="SenseBuffer" type="uint8 []"></param>
// <param name="TotalRespBufferSize" type="uint32 "></param>
func (instance *MS_SM_ScsiInformationMethods) SM_ScsiReportLuns( /* IN */ HbaPortWWN []uint8,
	/* IN */ DiscoveredPortWWN []uint8,
	/* IN */ DomainPortWWN []uint8,
	/* IN */ Cdb []uint8,
	/* IN */ InRespBufferMaxSize uint32,
	/* IN */ InSenseBufferMaxSize uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ ScsiStatus uint8,
	/* OUT */ TotalRespBufferSize uint32,
	/* OUT */ OutRespBufferSize uint32,
	/* OUT */ OutSenseBufferSize uint32,
	/* OUT */ RespBuffer []uint8,
	/* OUT */ SenseBuffer []uint8) (err error) {
	_, err = instance.InvokeMethod("SM_ScsiReportLuns", HbaPortWWN, DiscoveredPortWWN, DomainPortWWN, Cdb, InRespBufferMaxSize, InSenseBufferMaxSize)
	if err != nil {
		return
	}
	return

}

//

// <param name="Cdb" type="uint8 []"></param>
// <param name="DiscoveredPortWWN" type="uint8 []"></param>
// <param name="DomainPortWWN" type="uint8 []"></param>
// <param name="HbaPortWWN" type="uint8 []"></param>
// <param name="InRespBufferMaxSize" type="uint32 "></param>
// <param name="InSenseBufferMaxSize" type="uint32 "></param>
// <param name="SmhbaLUN" type="uint64 "></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="OutRespBufferSize" type="uint32 "></param>
// <param name="OutSenseBufferSize" type="uint32 "></param>
// <param name="RespBuffer" type="uint8 []"></param>
// <param name="ScsiStatus" type="uint8 "></param>
// <param name="SenseBuffer" type="uint8 []"></param>
func (instance *MS_SM_ScsiInformationMethods) SM_ScsiReadCapacity( /* IN */ HbaPortWWN []uint8,
	/* IN */ DiscoveredPortWWN []uint8,
	/* IN */ DomainPortWWN []uint8,
	/* IN */ SmhbaLUN uint64,
	/* IN */ Cdb []uint8,
	/* IN */ InRespBufferMaxSize uint32,
	/* IN */ InSenseBufferMaxSize uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ ScsiStatus uint8,
	/* OUT */ OutRespBufferSize uint32,
	/* OUT */ OutSenseBufferSize uint32,
	/* OUT */ RespBuffer []uint8,
	/* OUT */ SenseBuffer []uint8) (err error) {
	_, err = instance.InvokeMethod("SM_ScsiReadCapacity", HbaPortWWN, DiscoveredPortWWN, DomainPortWWN, SmhbaLUN, Cdb, InRespBufferMaxSize, InSenseBufferMaxSize)
	if err != nil {
		return
	}
	return

}
