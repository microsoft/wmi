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

// MS_SM_TargetInformationMethods struct
type MS_SM_TargetInformationMethods struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string
}

func NewMS_SM_TargetInformationMethodsEx1(instance *cim.WmiInstance) (newInstance *MS_SM_TargetInformationMethods, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MS_SM_TargetInformationMethods{
		WmiInstance: tmp,
	}
	return
}

func NewMS_SM_TargetInformationMethodsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MS_SM_TargetInformationMethods, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MS_SM_TargetInformationMethods{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MS_SM_TargetInformationMethods) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MS_SM_TargetInformationMethods) GetPropertyActive() (value bool, err error) {
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
func (instance *MS_SM_TargetInformationMethods) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MS_SM_TargetInformationMethods) GetPropertyInstanceName() (value string, err error) {
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

// <param name="DomainPortWWN" type="uint8 []"></param>
// <param name="HbaPortWWN" type="uint8 []"></param>
// <param name="InEntryCount" type="uint32 "></param>

// <param name="Entry" type="MS_SMHBA_SCSIENTRY []"></param>
// <param name="HBAStatus" type="uint32 "></param>
// <param name="OutEntryCount" type="uint32 "></param>
// <param name="TotalEntryCount" type="uint32 "></param>
func (instance *MS_SM_TargetInformationMethods) SM_GetTargetMapping( /* IN */ HbaPortWWN []uint8,
	/* IN */ DomainPortWWN []uint8,
	/* IN */ InEntryCount uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ TotalEntryCount uint32,
	/* OUT */ OutEntryCount uint32,
	/* OUT */ Entry []MS_SMHBA_SCSIENTRY) (err error) {
	_, err = instance.InvokeMethod("SM_GetTargetMapping", HbaPortWWN, DomainPortWWN, InEntryCount)
	if err != nil {
		return
	}
	return

}

//

// <param name="DomainPortWWN" type="uint8 []"></param>
// <param name="HbaPortWWN" type="uint8 []"></param>

// <param name="Flags" type="uint32 "></param>
// <param name="HBAStatus" type="uint32 "></param>
func (instance *MS_SM_TargetInformationMethods) SM_GetBindingCapability( /* IN */ HbaPortWWN []uint8,
	/* IN */ DomainPortWWN []uint8,
	/* OUT */ HBAStatus uint32,
	/* OUT */ Flags uint32) (err error) {
	_, err = instance.InvokeMethod("SM_GetBindingCapability", HbaPortWWN, DomainPortWWN)
	if err != nil {
		return
	}
	return

}

//

// <param name="DomainPortWWN" type="uint8 []"></param>
// <param name="HbaPortWWN" type="uint8 []"></param>

// <param name="Flags" type="uint32 "></param>
// <param name="HBAStatus" type="uint32 "></param>
func (instance *MS_SM_TargetInformationMethods) SM_GetBindingSupport( /* IN */ HbaPortWWN []uint8,
	/* IN */ DomainPortWWN []uint8,
	/* OUT */ HBAStatus uint32,
	/* OUT */ Flags uint32) (err error) {
	_, err = instance.InvokeMethod("SM_GetBindingSupport", HbaPortWWN, DomainPortWWN)
	if err != nil {
		return
	}
	return

}

//

// <param name="DomainPortWWN" type="uint8 []"></param>
// <param name="Flags" type="uint32 "></param>
// <param name="HbaPortWWN" type="uint8 []"></param>

// <param name="HBAStatus" type="uint32 "></param>
func (instance *MS_SM_TargetInformationMethods) SM_SetBindingSupport( /* IN */ HbaPortWWN []uint8,
	/* IN */ DomainPortWWN []uint8,
	/* IN */ Flags uint32,
	/* OUT */ HBAStatus uint32) (err error) {
	_, err = instance.InvokeMethod("SM_SetBindingSupport", HbaPortWWN, DomainPortWWN, Flags)
	if err != nil {
		return
	}
	return

}

//

// <param name="DomainPortWWN" type="uint8 []"></param>
// <param name="HbaPortWWN" type="uint8 []"></param>
// <param name="InEntryCount" type="uint32 "></param>

// <param name="Entry" type="MS_SMHBA_BINDINGENTRY []"></param>
// <param name="HBAStatus" type="uint32 "></param>
// <param name="OutEntryCount" type="uint32 "></param>
// <param name="TotalEntryCount" type="uint32 "></param>
func (instance *MS_SM_TargetInformationMethods) SM_GetPersistentBinding( /* IN */ HbaPortWWN []uint8,
	/* IN */ DomainPortWWN []uint8,
	/* IN */ InEntryCount uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ TotalEntryCount uint32,
	/* OUT */ OutEntryCount uint32,
	/* OUT */ Entry []MS_SMHBA_BINDINGENTRY) (err error) {
	_, err = instance.InvokeMethod("SM_GetPersistentBinding", HbaPortWWN, DomainPortWWN, InEntryCount)
	if err != nil {
		return
	}
	return

}

//

// <param name="DomainPortWWN" type="uint8 []"></param>
// <param name="Entry" type="MS_SMHBA_BINDINGENTRY []"></param>
// <param name="HbaPortWWN" type="uint8 []"></param>
// <param name="InEntryCount" type="uint32 "></param>

// <param name="EntryStatus" type="uint32 []"></param>
// <param name="HBAStatus" type="uint32 "></param>
// <param name="OutStatusCount" type="uint32 "></param>
func (instance *MS_SM_TargetInformationMethods) SM_SetPersistentBinding( /* IN */ HbaPortWWN []uint8,
	/* IN */ DomainPortWWN []uint8,
	/* IN */ InEntryCount uint32,
	/* IN */ Entry []MS_SMHBA_BINDINGENTRY,
	/* OUT */ HBAStatus uint32,
	/* OUT */ OutStatusCount uint32,
	/* OUT */ EntryStatus []uint32) (err error) {
	_, err = instance.InvokeMethod("SM_SetPersistentBinding", HbaPortWWN, DomainPortWWN, InEntryCount, Entry)
	if err != nil {
		return
	}
	return

}

//

// <param name="DomainPortWWN" type="uint8 []"></param>
// <param name="Entry" type="MS_SMHBA_BINDINGENTRY []"></param>
// <param name="EntryCount" type="uint32 "></param>
// <param name="HbaPortWWN" type="uint8 []"></param>

// <param name="HBAStatus" type="uint32 "></param>
func (instance *MS_SM_TargetInformationMethods) SM_RemovePersistentBinding( /* IN */ HbaPortWWN []uint8,
	/* IN */ DomainPortWWN []uint8,
	/* IN */ EntryCount uint32,
	/* IN */ Entry []MS_SMHBA_BINDINGENTRY,
	/* OUT */ HBAStatus uint32) (err error) {
	_, err = instance.InvokeMethod("SM_RemovePersistentBinding", HbaPortWWN, DomainPortWWN, EntryCount, Entry)
	if err != nil {
		return
	}
	return

}

//

// <param name="Lunit" type="HBAScsiID "></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="ProtocolStatistics" type="MS_SMHBA_PROTOCOLSTATISTICS "></param>
func (instance *MS_SM_TargetInformationMethods) SM_GetLUNStatistics( /* IN */ Lunit HBAScsiID,
	/* OUT */ HBAStatus uint32,
	/* OUT */ ProtocolStatistics MS_SMHBA_PROTOCOLSTATISTICS) (err error) {
	_, err = instance.InvokeMethod("SM_GetLUNStatistics", Lunit)
	if err != nil {
		return
	}
	return

}
