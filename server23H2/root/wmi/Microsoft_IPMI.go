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

// Microsoft_IPMI struct
type Microsoft_IPMI struct {
	*cim.WmiInstance

	//
	Active bool

	//
	BMCAddress uint8

	//
	InstanceName string
}

func NewMicrosoft_IPMIEx1(instance *cim.WmiInstance) (newInstance *Microsoft_IPMI, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Microsoft_IPMI{
		WmiInstance: tmp,
	}
	return
}

func NewMicrosoft_IPMIEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Microsoft_IPMI, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Microsoft_IPMI{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *Microsoft_IPMI) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *Microsoft_IPMI) GetPropertyActive() (value bool, err error) {
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

// SetBMCAddress sets the value of BMCAddress for the instance
func (instance *Microsoft_IPMI) SetPropertyBMCAddress(value uint8) (err error) {
	return instance.SetProperty("BMCAddress", (value))
}

// GetBMCAddress gets the value of BMCAddress for the instance
func (instance *Microsoft_IPMI) GetPropertyBMCAddress() (value uint8, err error) {
	retValue, err := instance.GetProperty("BMCAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *Microsoft_IPMI) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *Microsoft_IPMI) GetPropertyInstanceName() (value string, err error) {
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

// <param name="Command" type="uint8 "></param>
// <param name="Lun" type="uint8 "></param>
// <param name="NetworkFunction" type="uint8 "></param>
// <param name="RequestData" type="uint8 []"></param>
// <param name="RequestDataSize" type="uint32 "></param>
// <param name="ResponderAddress" type="uint8 "></param>

// <param name="CompletionCode" type="uint8 "></param>
// <param name="ResponseData" type="uint8 []"></param>
// <param name="ResponseDataSize" type="uint32 "></param>
func (instance *Microsoft_IPMI) RequestResponse( /* IN */ NetworkFunction uint8,
	/* IN */ Lun uint8,
	/* IN */ ResponderAddress uint8,
	/* IN */ Command uint8,
	/* IN */ RequestDataSize uint32,
	/* IN */ RequestData []uint8,
	/* OUT */ CompletionCode uint8,
	/* OUT */ ResponseDataSize uint32,
	/* OUT */ ResponseData []uint8) (err error) {
	_, err = instance.InvokeMethod("RequestResponse", NetworkFunction, Lun, ResponderAddress, Command, RequestDataSize, RequestData)
	if err != nil {
		return
	}
	return

}

//

// <param name="AttentionSet" type="bool "></param>
// <param name="StatusRegisterValue" type="uint8 "></param>
func (instance *Microsoft_IPMI) SMS_Attention( /* OUT */ AttentionSet bool,
	/* OUT */ StatusRegisterValue uint8) (err error) {
	_, err = instance.InvokeMethod("SMS_Attention")
	if err != nil {
		return
	}
	return

}

//

// <param name="Command" type="uint8 "></param>
// <param name="Data" type="uint8 []"></param>
// <param name="DataSize" type="uint32 "></param>
// <param name="Lun" type="uint8 "></param>
// <param name="NetworkFunction" type="uint8 "></param>
// <param name="RequestDataSize" type="uint32 "></param>
// <param name="ResponderAddress" type="uint8 "></param>

// <param name="CompletionCode" type="uint8 "></param>
// <param name="ResponseData" type="uint8 []"></param>
// <param name="ResponseDataSize" type="uint32 "></param>
func (instance *Microsoft_IPMI) RequestResponseEx( /* IN */ NetworkFunction uint8,
	/* IN */ Lun uint8,
	/* IN */ ResponderAddress uint8,
	/* IN */ Command uint8,
	/* IN */ RequestDataSize uint32,
	/* IN */ DataSize uint32,
	/* IN */ Data []uint8,
	/* OUT */ CompletionCode uint8,
	/* OUT */ ResponseDataSize uint32,
	/* OUT */ ResponseData []uint8) (err error) {
	_, err = instance.InvokeMethod("RequestResponseEx", NetworkFunction, Lun, ResponderAddress, Command, RequestDataSize, DataSize, Data)
	if err != nil {
		return
	}
	return

}

//

// <param name="Does0x80Dump" type="bool "></param>
// <param name="Does0x9cDump" type="bool "></param>
func (instance *Microsoft_IPMI) GetEnableMsrDumpSettings( /* OUT */ Does0x80Dump bool,
	/* OUT */ Does0x9cDump bool) (err error) {
	_, err = instance.InvokeMethod("GetEnableMsrDumpSettings")
	if err != nil {
		return
	}
	return

}

//

// <param name="Does0x80Dump" type="bool "></param>
// <param name="Does0x9cDump" type="bool "></param>
func (instance *Microsoft_IPMI) SetEnableMsrDumpSettings( /* IN */ Does0x80Dump bool,
	/* IN */ Does0x9cDump bool) (err error) {
	_, err = instance.InvokeMethodWithReturn("SetEnableMsrDumpSettings", Does0x80Dump, Does0x9cDump)
	if err != nil {
		return
	}
	return

}

//

// <param name="ClearMsrDump" type="bool "></param>
func (instance *Microsoft_IPMI) GetClearMsrDumpSettings( /* OUT */ ClearMsrDump bool) (err error) {
	_, err = instance.InvokeMethod("GetClearMsrDumpSettings")
	if err != nil {
		return
	}
	return

}

//

// <param name="ClearMsrDump" type="bool "></param>
func (instance *Microsoft_IPMI) SetClearMsrDumpSettings( /* IN */ ClearMsrDump bool) (err error) {
	_, err = instance.InvokeMethodWithReturn("SetClearMsrDumpSettings", ClearMsrDump)
	if err != nil {
		return
	}
	return

}

//
func (instance *Microsoft_IPMI) EnableSelCachingSupport() (err error) {
	_, err = instance.InvokeMethodWithReturn("EnableSelCachingSupport")
	if err != nil {
		return
	}
	return

}

//
func (instance *Microsoft_IPMI) DisableSelCachingSupport() (err error) {
	_, err = instance.InvokeMethodWithReturn("DisableSelCachingSupport")
	if err != nil {
		return
	}
	return

}

//

// <param name="SelCachePollingIntervalMs" type="uint32 "></param>
func (instance *Microsoft_IPMI) SetSelCachePollingInterval( /* IN */ SelCachePollingIntervalMs uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("SetSelCachePollingInterval", SelCachePollingIntervalMs)
	if err != nil {
		return
	}
	return

}

//

// <param name="SelCacheReadRecordLimit" type="uint32 "></param>
func (instance *Microsoft_IPMI) SetSelCacheReadRecordLimit( /* IN */ SelCacheReadRecordLimit uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("SetSelCacheReadRecordLimit", SelCacheReadRecordLimit)
	if err != nil {
		return
	}
	return

}

//
func (instance *Microsoft_IPMI) EnableMsrDumpCommands() (err error) {
	_, err = instance.InvokeMethodWithReturn("EnableMsrDumpCommands")
	if err != nil {
		return
	}
	return

}

//
func (instance *Microsoft_IPMI) DisableMsrDumpCommands() (err error) {
	_, err = instance.InvokeMethodWithReturn("DisableMsrDumpCommands")
	if err != nil {
		return
	}
	return

}
