// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.StandardCimv2
//
// ////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetLldpAgent struct
type MSFT_NetLldpAgent struct {
	*MSFT_NetLldpMsap

	//
	AdminStatus uint32

	//
	InterfaceUp bool

	//
	LastNeighborChangeTime string

	//
	LastRxTime string

	//
	MediaConnected bool

	//
	MsgFastTx uint32

	//
	MsgTxHold uint32

	//
	MsgTxInterval uint32

	//
	Neighbor MSFT_NetLldpMsap

	//
	ReinitDelay uint32

	//
	StatsAgeoutsTotal uint64

	//
	StatsFramesDiscardedTotal uint64

	//
	StatsFramesInErrorsTotal uint64

	//
	StatsFramesInTotal uint64

	//
	StatsFramesOutTotal uint64

	//
	StatsResetTime string

	//
	TooManyNeighbors bool

	//
	TxCreditMax uint32

	//
	TxFastInit uint32
}

func NewMSFT_NetLldpAgentEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetLldpAgent, err error) {
	tmp, err := NewMSFT_NetLldpMsapEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetLldpAgent{
		MSFT_NetLldpMsap: tmp,
	}
	return
}

func NewMSFT_NetLldpAgentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetLldpAgent, err error) {
	tmp, err := NewMSFT_NetLldpMsapEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetLldpAgent{
		MSFT_NetLldpMsap: tmp,
	}
	return
}

// SetAdminStatus sets the value of AdminStatus for the instance
func (instance *MSFT_NetLldpAgent) SetPropertyAdminStatus(value uint32) (err error) {
	return instance.SetProperty("AdminStatus", (value))
}

// GetAdminStatus gets the value of AdminStatus for the instance
func (instance *MSFT_NetLldpAgent) GetPropertyAdminStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("AdminStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetInterfaceUp sets the value of InterfaceUp for the instance
func (instance *MSFT_NetLldpAgent) SetPropertyInterfaceUp(value bool) (err error) {
	return instance.SetProperty("InterfaceUp", (value))
}

// GetInterfaceUp gets the value of InterfaceUp for the instance
func (instance *MSFT_NetLldpAgent) GetPropertyInterfaceUp() (value bool, err error) {
	retValue, err := instance.GetProperty("InterfaceUp")
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

// SetLastNeighborChangeTime sets the value of LastNeighborChangeTime for the instance
func (instance *MSFT_NetLldpAgent) SetPropertyLastNeighborChangeTime(value string) (err error) {
	return instance.SetProperty("LastNeighborChangeTime", (value))
}

// GetLastNeighborChangeTime gets the value of LastNeighborChangeTime for the instance
func (instance *MSFT_NetLldpAgent) GetPropertyLastNeighborChangeTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastNeighborChangeTime")
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

// SetLastRxTime sets the value of LastRxTime for the instance
func (instance *MSFT_NetLldpAgent) SetPropertyLastRxTime(value string) (err error) {
	return instance.SetProperty("LastRxTime", (value))
}

// GetLastRxTime gets the value of LastRxTime for the instance
func (instance *MSFT_NetLldpAgent) GetPropertyLastRxTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastRxTime")
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

// SetMediaConnected sets the value of MediaConnected for the instance
func (instance *MSFT_NetLldpAgent) SetPropertyMediaConnected(value bool) (err error) {
	return instance.SetProperty("MediaConnected", (value))
}

// GetMediaConnected gets the value of MediaConnected for the instance
func (instance *MSFT_NetLldpAgent) GetPropertyMediaConnected() (value bool, err error) {
	retValue, err := instance.GetProperty("MediaConnected")
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

// SetMsgFastTx sets the value of MsgFastTx for the instance
func (instance *MSFT_NetLldpAgent) SetPropertyMsgFastTx(value uint32) (err error) {
	return instance.SetProperty("MsgFastTx", (value))
}

// GetMsgFastTx gets the value of MsgFastTx for the instance
func (instance *MSFT_NetLldpAgent) GetPropertyMsgFastTx() (value uint32, err error) {
	retValue, err := instance.GetProperty("MsgFastTx")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetMsgTxHold sets the value of MsgTxHold for the instance
func (instance *MSFT_NetLldpAgent) SetPropertyMsgTxHold(value uint32) (err error) {
	return instance.SetProperty("MsgTxHold", (value))
}

// GetMsgTxHold gets the value of MsgTxHold for the instance
func (instance *MSFT_NetLldpAgent) GetPropertyMsgTxHold() (value uint32, err error) {
	retValue, err := instance.GetProperty("MsgTxHold")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetMsgTxInterval sets the value of MsgTxInterval for the instance
func (instance *MSFT_NetLldpAgent) SetPropertyMsgTxInterval(value uint32) (err error) {
	return instance.SetProperty("MsgTxInterval", (value))
}

// GetMsgTxInterval gets the value of MsgTxInterval for the instance
func (instance *MSFT_NetLldpAgent) GetPropertyMsgTxInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("MsgTxInterval")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetNeighbor sets the value of Neighbor for the instance
func (instance *MSFT_NetLldpAgent) SetPropertyNeighbor(value MSFT_NetLldpMsap) (err error) {
	return instance.SetProperty("Neighbor", (value))
}

// GetNeighbor gets the value of Neighbor for the instance
func (instance *MSFT_NetLldpAgent) GetPropertyNeighbor() (value MSFT_NetLldpMsap, err error) {
	retValue, err := instance.GetProperty("Neighbor")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_NetLldpMsap)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_NetLldpMsap is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_NetLldpMsap(valuetmp)

	return
}

// SetReinitDelay sets the value of ReinitDelay for the instance
func (instance *MSFT_NetLldpAgent) SetPropertyReinitDelay(value uint32) (err error) {
	return instance.SetProperty("ReinitDelay", (value))
}

// GetReinitDelay gets the value of ReinitDelay for the instance
func (instance *MSFT_NetLldpAgent) GetPropertyReinitDelay() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReinitDelay")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetStatsAgeoutsTotal sets the value of StatsAgeoutsTotal for the instance
func (instance *MSFT_NetLldpAgent) SetPropertyStatsAgeoutsTotal(value uint64) (err error) {
	return instance.SetProperty("StatsAgeoutsTotal", (value))
}

// GetStatsAgeoutsTotal gets the value of StatsAgeoutsTotal for the instance
func (instance *MSFT_NetLldpAgent) GetPropertyStatsAgeoutsTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("StatsAgeoutsTotal")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetStatsFramesDiscardedTotal sets the value of StatsFramesDiscardedTotal for the instance
func (instance *MSFT_NetLldpAgent) SetPropertyStatsFramesDiscardedTotal(value uint64) (err error) {
	return instance.SetProperty("StatsFramesDiscardedTotal", (value))
}

// GetStatsFramesDiscardedTotal gets the value of StatsFramesDiscardedTotal for the instance
func (instance *MSFT_NetLldpAgent) GetPropertyStatsFramesDiscardedTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("StatsFramesDiscardedTotal")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetStatsFramesInErrorsTotal sets the value of StatsFramesInErrorsTotal for the instance
func (instance *MSFT_NetLldpAgent) SetPropertyStatsFramesInErrorsTotal(value uint64) (err error) {
	return instance.SetProperty("StatsFramesInErrorsTotal", (value))
}

// GetStatsFramesInErrorsTotal gets the value of StatsFramesInErrorsTotal for the instance
func (instance *MSFT_NetLldpAgent) GetPropertyStatsFramesInErrorsTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("StatsFramesInErrorsTotal")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetStatsFramesInTotal sets the value of StatsFramesInTotal for the instance
func (instance *MSFT_NetLldpAgent) SetPropertyStatsFramesInTotal(value uint64) (err error) {
	return instance.SetProperty("StatsFramesInTotal", (value))
}

// GetStatsFramesInTotal gets the value of StatsFramesInTotal for the instance
func (instance *MSFT_NetLldpAgent) GetPropertyStatsFramesInTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("StatsFramesInTotal")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetStatsFramesOutTotal sets the value of StatsFramesOutTotal for the instance
func (instance *MSFT_NetLldpAgent) SetPropertyStatsFramesOutTotal(value uint64) (err error) {
	return instance.SetProperty("StatsFramesOutTotal", (value))
}

// GetStatsFramesOutTotal gets the value of StatsFramesOutTotal for the instance
func (instance *MSFT_NetLldpAgent) GetPropertyStatsFramesOutTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("StatsFramesOutTotal")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetStatsResetTime sets the value of StatsResetTime for the instance
func (instance *MSFT_NetLldpAgent) SetPropertyStatsResetTime(value string) (err error) {
	return instance.SetProperty("StatsResetTime", (value))
}

// GetStatsResetTime gets the value of StatsResetTime for the instance
func (instance *MSFT_NetLldpAgent) GetPropertyStatsResetTime() (value string, err error) {
	retValue, err := instance.GetProperty("StatsResetTime")
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

// SetTooManyNeighbors sets the value of TooManyNeighbors for the instance
func (instance *MSFT_NetLldpAgent) SetPropertyTooManyNeighbors(value bool) (err error) {
	return instance.SetProperty("TooManyNeighbors", (value))
}

// GetTooManyNeighbors gets the value of TooManyNeighbors for the instance
func (instance *MSFT_NetLldpAgent) GetPropertyTooManyNeighbors() (value bool, err error) {
	retValue, err := instance.GetProperty("TooManyNeighbors")
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

// SetTxCreditMax sets the value of TxCreditMax for the instance
func (instance *MSFT_NetLldpAgent) SetPropertyTxCreditMax(value uint32) (err error) {
	return instance.SetProperty("TxCreditMax", (value))
}

// GetTxCreditMax gets the value of TxCreditMax for the instance
func (instance *MSFT_NetLldpAgent) GetPropertyTxCreditMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("TxCreditMax")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetTxFastInit sets the value of TxFastInit for the instance
func (instance *MSFT_NetLldpAgent) SetPropertyTxFastInit(value uint32) (err error) {
	return instance.SetProperty("TxFastInit", (value))
}

// GetTxFastInit gets the value of TxFastInit for the instance
func (instance *MSFT_NetLldpAgent) GetPropertyTxFastInit() (value uint32, err error) {
	retValue, err := instance.GetProperty("TxFastInit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetLldpAgent) Enable() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetLldpAgent) Disable() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disable")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
