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

// MSiSCSIInitiator_SessionClass struct
type MSiSCSIInitiator_SessionClass struct {
	*cim.WmiInstance

	// Information about the connections for this session
	ConnectionInformation []MSiSCSIInitiator_ConnectionInformation

	// Information about the devices exposed by this session
	Devices []MSiSCSIInitiator_DeviceOnSession

	//
	InitiatorName string

	//
	ISID []uint8

	//
	SessionId string

	//
	TargetName string

	//
	TargetNodeName string

	//
	TSID []uint8
}

func NewMSiSCSIInitiator_SessionClassEx1(instance *cim.WmiInstance) (newInstance *MSiSCSIInitiator_SessionClass, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_SessionClass{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSIInitiator_SessionClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSIInitiator_SessionClass, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_SessionClass{
		WmiInstance: tmp,
	}
	return
}

// SetConnectionInformation sets the value of ConnectionInformation for the instance
func (instance *MSiSCSIInitiator_SessionClass) SetPropertyConnectionInformation(value []MSiSCSIInitiator_ConnectionInformation) (err error) {
	return instance.SetProperty("ConnectionInformation", (value))
}

// GetConnectionInformation gets the value of ConnectionInformation for the instance
func (instance *MSiSCSIInitiator_SessionClass) GetPropertyConnectionInformation() (value []MSiSCSIInitiator_ConnectionInformation, err error) {
	retValue, err := instance.GetProperty("ConnectionInformation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSiSCSIInitiator_ConnectionInformation)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSiSCSIInitiator_ConnectionInformation is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSiSCSIInitiator_ConnectionInformation(valuetmp))
	}

	return
}

// SetDevices sets the value of Devices for the instance
func (instance *MSiSCSIInitiator_SessionClass) SetPropertyDevices(value []MSiSCSIInitiator_DeviceOnSession) (err error) {
	return instance.SetProperty("Devices", (value))
}

// GetDevices gets the value of Devices for the instance
func (instance *MSiSCSIInitiator_SessionClass) GetPropertyDevices() (value []MSiSCSIInitiator_DeviceOnSession, err error) {
	retValue, err := instance.GetProperty("Devices")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSiSCSIInitiator_DeviceOnSession)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSiSCSIInitiator_DeviceOnSession is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSiSCSIInitiator_DeviceOnSession(valuetmp))
	}

	return
}

// SetInitiatorName sets the value of InitiatorName for the instance
func (instance *MSiSCSIInitiator_SessionClass) SetPropertyInitiatorName(value string) (err error) {
	return instance.SetProperty("InitiatorName", (value))
}

// GetInitiatorName gets the value of InitiatorName for the instance
func (instance *MSiSCSIInitiator_SessionClass) GetPropertyInitiatorName() (value string, err error) {
	retValue, err := instance.GetProperty("InitiatorName")
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

// SetISID sets the value of ISID for the instance
func (instance *MSiSCSIInitiator_SessionClass) SetPropertyISID(value []uint8) (err error) {
	return instance.SetProperty("ISID", (value))
}

// GetISID gets the value of ISID for the instance
func (instance *MSiSCSIInitiator_SessionClass) GetPropertyISID() (value []uint8, err error) {
	retValue, err := instance.GetProperty("ISID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetSessionId sets the value of SessionId for the instance
func (instance *MSiSCSIInitiator_SessionClass) SetPropertySessionId(value string) (err error) {
	return instance.SetProperty("SessionId", (value))
}

// GetSessionId gets the value of SessionId for the instance
func (instance *MSiSCSIInitiator_SessionClass) GetPropertySessionId() (value string, err error) {
	retValue, err := instance.GetProperty("SessionId")
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

// SetTargetName sets the value of TargetName for the instance
func (instance *MSiSCSIInitiator_SessionClass) SetPropertyTargetName(value string) (err error) {
	return instance.SetProperty("TargetName", (value))
}

// GetTargetName gets the value of TargetName for the instance
func (instance *MSiSCSIInitiator_SessionClass) GetPropertyTargetName() (value string, err error) {
	retValue, err := instance.GetProperty("TargetName")
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

// SetTargetNodeName sets the value of TargetNodeName for the instance
func (instance *MSiSCSIInitiator_SessionClass) SetPropertyTargetNodeName(value string) (err error) {
	return instance.SetProperty("TargetNodeName", (value))
}

// GetTargetNodeName gets the value of TargetNodeName for the instance
func (instance *MSiSCSIInitiator_SessionClass) GetPropertyTargetNodeName() (value string, err error) {
	retValue, err := instance.GetProperty("TargetNodeName")
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

// SetTSID sets the value of TSID for the instance
func (instance *MSiSCSIInitiator_SessionClass) SetPropertyTSID(value []uint8) (err error) {
	return instance.SetProperty("TSID", (value))
}

// GetTSID gets the value of TSID for the instance
func (instance *MSiSCSIInitiator_SessionClass) GetPropertyTSID() (value []uint8, err error) {
	retValue, err := instance.GetProperty("TSID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSiSCSIInitiator_SessionClass) Logout() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Logout")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="EvpdCmddt" type="uint8 "></param>
// <param name="Lun" type="uint64 "></param>
// <param name="PageCode" type="uint8 "></param>

// <param name="ResponseBuffer" type="uint8 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="ScsiStatus" type="uint8 "></param>
// <param name="SenseBuffer" type="uint8 []"></param>
func (instance *MSiSCSIInitiator_SessionClass) SendScsiInquiry( /* IN */ Lun uint64,
	/* IN */ EvpdCmddt uint8,
	/* IN */ PageCode uint8,
	/* OUT */ ScsiStatus uint8,
	/* OUT */ ResponseBuffer []uint8,
	/* OUT */ SenseBuffer []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SendScsiInquiry", Lun, EvpdCmddt, PageCode)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Lun" type="uint64 "></param>

// <param name="ResponseBuffer" type="uint8 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="ScsiStatus" type="uint8 "></param>
// <param name="SenseBuffer" type="uint8 []"></param>
func (instance *MSiSCSIInitiator_SessionClass) SendScsiReadCapacity( /* IN */ Lun uint64,
	/* OUT */ ScsiStatus uint8,
	/* OUT */ ResponseBuffer []uint8,
	/* OUT */ SenseBuffer []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SendScsiReadCapacity", Lun)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ResponseBuffer" type="uint8 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="ScsiStatus" type="uint8 "></param>
// <param name="SenseBuffer" type="uint8 []"></param>
func (instance *MSiSCSIInitiator_SessionClass) SendScsiReportLuns( /* OUT */ ScsiStatus uint8,
	/* OUT */ ResponseBuffer []uint8,
	/* OUT */ SenseBuffer []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SendScsiReportLuns")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="InitiatorPortNumber" type="uint32 "></param>
// <param name="key" type="uint8 []"></param>
// <param name="LoginOptions" type="MSiSCSIInitiator_TargetLoginOptions "></param>
// <param name="SecurityFlags" type="uint64 "></param>
// <param name="TargetPortal" type="MSiSCSIInitiator_Portal "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="UniqueConnectionId" type="string "></param>
func (instance *MSiSCSIInitiator_SessionClass) AddConnection( /* IN */ InitiatorPortNumber uint32,
	/* IN */ TargetPortal MSiSCSIInitiator_Portal,
	/* IN */ SecurityFlags uint64,
	/* IN */ LoginOptions MSiSCSIInitiator_TargetLoginOptions,
	/* IN */ key []uint8,
	/* OUT */ UniqueConnectionId string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddConnection", InitiatorPortNumber, TargetPortal, SecurityFlags, LoginOptions, key)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="UniqueConnectionId" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSiSCSIInitiator_SessionClass) RemoveConnection( /* IN */ UniqueConnectionId string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveConnection", UniqueConnectionId)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
