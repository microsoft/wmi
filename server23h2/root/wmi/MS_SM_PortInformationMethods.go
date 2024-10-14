// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
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

// MS_SM_PortInformationMethods struct
type MS_SM_PortInformationMethods struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string
}

func NewMS_SM_PortInformationMethodsEx1(instance *cim.WmiInstance) (newInstance *MS_SM_PortInformationMethods, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MS_SM_PortInformationMethods{
		WmiInstance: tmp,
	}
	return
}

func NewMS_SM_PortInformationMethodsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MS_SM_PortInformationMethods, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MS_SM_PortInformationMethods{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MS_SM_PortInformationMethods) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MS_SM_PortInformationMethods) GetPropertyActive() (value bool, err error) {
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
func (instance *MS_SM_PortInformationMethods) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MS_SM_PortInformationMethods) GetPropertyInstanceName() (value string, err error) {
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

// <param name="PortIndex" type="uint32 "></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="PortType" type="uint32 "></param>
func (instance *MS_SM_PortInformationMethods) SM_GetPortType( /* IN */ PortIndex uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ PortType uint32) (err error) {
	_, err = instance.InvokeMethod("SM_GetPortType", PortIndex)
	if err != nil {
		return
	}
	return

}

//

// <param name="PortIndex" type="uint32 "></param>
// <param name="PortSpecificAttributesMaxSize" type="uint32 "></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="PortAttributes" type="MS_SMHBA_PORTATTRIBUTES "></param>
func (instance *MS_SM_PortInformationMethods) SM_GetAdapterPortAttributes( /* IN */ PortIndex uint32,
	/* IN */ PortSpecificAttributesMaxSize uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ PortAttributes MS_SMHBA_PORTATTRIBUTES) (err error) {
	_, err = instance.InvokeMethod("SM_GetAdapterPortAttributes", PortIndex, PortSpecificAttributesMaxSize)
	if err != nil {
		return
	}
	return

}

//

// <param name="DiscoveredPortIndex" type="uint32 "></param>
// <param name="PortIndex" type="uint32 "></param>
// <param name="PortSpecificAttributesMaxSize" type="uint32 "></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="PortAttributes" type="MS_SMHBA_PORTATTRIBUTES "></param>
func (instance *MS_SM_PortInformationMethods) SM_GetDiscoveredPortAttributes( /* IN */ PortIndex uint32,
	/* IN */ DiscoveredPortIndex uint32,
	/* IN */ PortSpecificAttributesMaxSize uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ PortAttributes MS_SMHBA_PORTATTRIBUTES) (err error) {
	_, err = instance.InvokeMethod("SM_GetDiscoveredPortAttributes", PortIndex, DiscoveredPortIndex, PortSpecificAttributesMaxSize)
	if err != nil {
		return
	}
	return

}

//

// <param name="DomainPortWWN" type="uint8 []"></param>
// <param name="PortSpecificAttributesMaxSize" type="uint32 "></param>
// <param name="PortWWN" type="uint8 []"></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="PortAttributes" type="MS_SMHBA_PORTATTRIBUTES "></param>
func (instance *MS_SM_PortInformationMethods) SM_GetPortAttributesByWWN( /* IN */ PortWWN []uint8,
	/* IN */ DomainPortWWN []uint8,
	/* IN */ PortSpecificAttributesMaxSize uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ PortAttributes MS_SMHBA_PORTATTRIBUTES) (err error) {
	_, err = instance.InvokeMethod("SM_GetPortAttributesByWWN", PortWWN, DomainPortWWN, PortSpecificAttributesMaxSize)
	if err != nil {
		return
	}
	return

}

//

// <param name="PortIndex" type="uint32 "></param>
// <param name="ProtocolType" type="uint32 "></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="ProtocolStatistics" type="MS_SMHBA_PROTOCOLSTATISTICS "></param>
func (instance *MS_SM_PortInformationMethods) SM_GetProtocolStatistics( /* IN */ PortIndex uint32,
	/* IN */ ProtocolType uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ ProtocolStatistics MS_SMHBA_PROTOCOLSTATISTICS) (err error) {
	_, err = instance.InvokeMethod("SM_GetProtocolStatistics", PortIndex, ProtocolType)
	if err != nil {
		return
	}
	return

}

//

// <param name="InNumOfPhyCounters" type="uint32 "></param>
// <param name="PhyIndex" type="uint32 "></param>
// <param name="PortIndex" type="uint32 "></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="OutNumOfPhyCounters" type="uint32 "></param>
// <param name="PhyCounter" type="int64 []"></param>
// <param name="TotalNumOfPhyCounters" type="uint32 "></param>
func (instance *MS_SM_PortInformationMethods) SM_GetPhyStatistics( /* IN */ PortIndex uint32,
	/* IN */ PhyIndex uint32,
	/* IN */ InNumOfPhyCounters uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ TotalNumOfPhyCounters uint32,
	/* OUT */ OutNumOfPhyCounters uint32,
	/* OUT */ PhyCounter []int64) (err error) {
	_, err = instance.InvokeMethod("SM_GetPhyStatistics", PortIndex, PhyIndex, InNumOfPhyCounters)
	if err != nil {
		return
	}
	return

}

//

// <param name="PhyIndex" type="uint32 "></param>
// <param name="PortIndex" type="uint32 "></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="PhyType" type="MS_SMHBA_FC_PHY "></param>
func (instance *MS_SM_PortInformationMethods) SM_GetFCPhyAttributes( /* IN */ PortIndex uint32,
	/* IN */ PhyIndex uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ PhyType MS_SMHBA_FC_PHY) (err error) {
	_, err = instance.InvokeMethod("SM_GetFCPhyAttributes", PortIndex, PhyIndex)
	if err != nil {
		return
	}
	return

}

//

// <param name="PhyIndex" type="uint32 "></param>
// <param name="PortIndex" type="uint32 "></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="PhyType" type="MS_SMHBA_SAS_PHY "></param>
func (instance *MS_SM_PortInformationMethods) SM_GetSASPhyAttributes( /* IN */ PortIndex uint32,
	/* IN */ PhyIndex uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ PhyType MS_SMHBA_SAS_PHY) (err error) {
	_, err = instance.InvokeMethod("SM_GetSASPhyAttributes", PortIndex, PhyIndex)
	if err != nil {
		return
	}
	return

}

//
func (instance *MS_SM_PortInformationMethods) SM_RefreshInformation() (err error) {
	_, err = instance.InvokeMethodWithReturn("SM_RefreshInformation")
	if err != nil {
		return
	}
	return

}
