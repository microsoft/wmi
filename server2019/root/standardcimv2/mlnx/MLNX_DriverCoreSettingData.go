// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MLNX_DriverCoreSettingData struct
type MLNX_DriverCoreSettingData struct {
	*MLNX_DriverSettingData

	//
	LogMttsPerSeg uint32

	//
	LogNumCq uint32

	//
	LogNumMac uint32

	//
	LogNumMcg uint32

	//
	LogNumMpt uint32

	//
	LogNumMtt uint32

	//
	LogNumQp uint32

	//
	LogNumRdmaRc uint32

	//
	LogNumSrq uint32

	//
	LogNumVlan uint32

	//
	MaximumWorkingThreads uint32

	//
	RoceMode string

	//
	Set4kMtu bool

	//
	SriovPort1NumVFs uint32

	//
	SriovPort2NumVFs uint32

	//
	SriovPortMode uint32
}

func NewMLNX_DriverCoreSettingDataEx1(instance *cim.WmiInstance) (newInstance *MLNX_DriverCoreSettingData, err error) {
	tmp, err := NewMLNX_DriverSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverCoreSettingData{
		MLNX_DriverSettingData: tmp,
	}
	return
}

func NewMLNX_DriverCoreSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_DriverCoreSettingData, err error) {
	tmp, err := NewMLNX_DriverSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverCoreSettingData{
		MLNX_DriverSettingData: tmp,
	}
	return
}

// SetLogMttsPerSeg sets the value of LogMttsPerSeg for the instance
func (instance *MLNX_DriverCoreSettingData) SetPropertyLogMttsPerSeg(value uint32) (err error) {
	return instance.SetProperty("LogMttsPerSeg", value)
}

// GetLogMttsPerSeg gets the value of LogMttsPerSeg for the instance
func (instance *MLNX_DriverCoreSettingData) GetPropertyLogMttsPerSeg() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogMttsPerSeg")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumCq sets the value of LogNumCq for the instance
func (instance *MLNX_DriverCoreSettingData) SetPropertyLogNumCq(value uint32) (err error) {
	return instance.SetProperty("LogNumCq", value)
}

// GetLogNumCq gets the value of LogNumCq for the instance
func (instance *MLNX_DriverCoreSettingData) GetPropertyLogNumCq() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumCq")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumMac sets the value of LogNumMac for the instance
func (instance *MLNX_DriverCoreSettingData) SetPropertyLogNumMac(value uint32) (err error) {
	return instance.SetProperty("LogNumMac", value)
}

// GetLogNumMac gets the value of LogNumMac for the instance
func (instance *MLNX_DriverCoreSettingData) GetPropertyLogNumMac() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumMac")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumMcg sets the value of LogNumMcg for the instance
func (instance *MLNX_DriverCoreSettingData) SetPropertyLogNumMcg(value uint32) (err error) {
	return instance.SetProperty("LogNumMcg", value)
}

// GetLogNumMcg gets the value of LogNumMcg for the instance
func (instance *MLNX_DriverCoreSettingData) GetPropertyLogNumMcg() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumMcg")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumMpt sets the value of LogNumMpt for the instance
func (instance *MLNX_DriverCoreSettingData) SetPropertyLogNumMpt(value uint32) (err error) {
	return instance.SetProperty("LogNumMpt", value)
}

// GetLogNumMpt gets the value of LogNumMpt for the instance
func (instance *MLNX_DriverCoreSettingData) GetPropertyLogNumMpt() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumMpt")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumMtt sets the value of LogNumMtt for the instance
func (instance *MLNX_DriverCoreSettingData) SetPropertyLogNumMtt(value uint32) (err error) {
	return instance.SetProperty("LogNumMtt", value)
}

// GetLogNumMtt gets the value of LogNumMtt for the instance
func (instance *MLNX_DriverCoreSettingData) GetPropertyLogNumMtt() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumMtt")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumQp sets the value of LogNumQp for the instance
func (instance *MLNX_DriverCoreSettingData) SetPropertyLogNumQp(value uint32) (err error) {
	return instance.SetProperty("LogNumQp", value)
}

// GetLogNumQp gets the value of LogNumQp for the instance
func (instance *MLNX_DriverCoreSettingData) GetPropertyLogNumQp() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumQp")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumRdmaRc sets the value of LogNumRdmaRc for the instance
func (instance *MLNX_DriverCoreSettingData) SetPropertyLogNumRdmaRc(value uint32) (err error) {
	return instance.SetProperty("LogNumRdmaRc", value)
}

// GetLogNumRdmaRc gets the value of LogNumRdmaRc for the instance
func (instance *MLNX_DriverCoreSettingData) GetPropertyLogNumRdmaRc() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumRdmaRc")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumSrq sets the value of LogNumSrq for the instance
func (instance *MLNX_DriverCoreSettingData) SetPropertyLogNumSrq(value uint32) (err error) {
	return instance.SetProperty("LogNumSrq", value)
}

// GetLogNumSrq gets the value of LogNumSrq for the instance
func (instance *MLNX_DriverCoreSettingData) GetPropertyLogNumSrq() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumSrq")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumVlan sets the value of LogNumVlan for the instance
func (instance *MLNX_DriverCoreSettingData) SetPropertyLogNumVlan(value uint32) (err error) {
	return instance.SetProperty("LogNumVlan", value)
}

// GetLogNumVlan gets the value of LogNumVlan for the instance
func (instance *MLNX_DriverCoreSettingData) GetPropertyLogNumVlan() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumVlan")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumWorkingThreads sets the value of MaximumWorkingThreads for the instance
func (instance *MLNX_DriverCoreSettingData) SetPropertyMaximumWorkingThreads(value uint32) (err error) {
	return instance.SetProperty("MaximumWorkingThreads", value)
}

// GetMaximumWorkingThreads gets the value of MaximumWorkingThreads for the instance
func (instance *MLNX_DriverCoreSettingData) GetPropertyMaximumWorkingThreads() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumWorkingThreads")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoceMode sets the value of RoceMode for the instance
func (instance *MLNX_DriverCoreSettingData) SetPropertyRoceMode(value string) (err error) {
	return instance.SetProperty("RoceMode", value)
}

// GetRoceMode gets the value of RoceMode for the instance
func (instance *MLNX_DriverCoreSettingData) GetPropertyRoceMode() (value string, err error) {
	retValue, err := instance.GetProperty("RoceMode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSet4kMtu sets the value of Set4kMtu for the instance
func (instance *MLNX_DriverCoreSettingData) SetPropertySet4kMtu(value bool) (err error) {
	return instance.SetProperty("Set4kMtu", value)
}

// GetSet4kMtu gets the value of Set4kMtu for the instance
func (instance *MLNX_DriverCoreSettingData) GetPropertySet4kMtu() (value bool, err error) {
	retValue, err := instance.GetProperty("Set4kMtu")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSriovPort1NumVFs sets the value of SriovPort1NumVFs for the instance
func (instance *MLNX_DriverCoreSettingData) SetPropertySriovPort1NumVFs(value uint32) (err error) {
	return instance.SetProperty("SriovPort1NumVFs", value)
}

// GetSriovPort1NumVFs gets the value of SriovPort1NumVFs for the instance
func (instance *MLNX_DriverCoreSettingData) GetPropertySriovPort1NumVFs() (value uint32, err error) {
	retValue, err := instance.GetProperty("SriovPort1NumVFs")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSriovPort2NumVFs sets the value of SriovPort2NumVFs for the instance
func (instance *MLNX_DriverCoreSettingData) SetPropertySriovPort2NumVFs(value uint32) (err error) {
	return instance.SetProperty("SriovPort2NumVFs", value)
}

// GetSriovPort2NumVFs gets the value of SriovPort2NumVFs for the instance
func (instance *MLNX_DriverCoreSettingData) GetPropertySriovPort2NumVFs() (value uint32, err error) {
	retValue, err := instance.GetProperty("SriovPort2NumVFs")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSriovPortMode sets the value of SriovPortMode for the instance
func (instance *MLNX_DriverCoreSettingData) SetPropertySriovPortMode(value uint32) (err error) {
	return instance.SetProperty("SriovPortMode", value)
}

// GetSriovPortMode gets the value of SriovPortMode for the instance
func (instance *MLNX_DriverCoreSettingData) GetPropertySriovPortMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("SriovPortMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="DebugFlags" type="uint32 "></param>
// <param name="DebugLevel" type="uint32 "></param>
// <param name="LogMttsPerSeg" type="uint32 "></param>
// <param name="LogNumCq" type="uint32 "></param>
// <param name="LogNumMac" type="uint32 "></param>
// <param name="LogNumMcg" type="uint32 "></param>
// <param name="LogNumMpt" type="uint32 "></param>
// <param name="LogNumMtt" type="uint32 "></param>
// <param name="LogNumRdmaRc" type="uint32 "></param>
// <param name="LogNumSrq" type="uint32 "></param>
// <param name="LogNumVlan" type="uint32 "></param>
// <param name="MaxContQuant" type="uint32 "></param>
// <param name="MaximumWorkingThreads" type="uint32 "></param>
// <param name="ModeFlags" type="uint32 "></param>
// <param name="MultiEqNum" type="uint32 "></param>
// <param name="MultiMsixNum" type="uint32 "></param>
// <param name="NumFcExch" type="uint32 "></param>
// <param name="RoceMode" type="string "></param>
// <param name="Set4kMtu" type="bool "></param>
// <param name="SingleEqNum" type="uint32 "></param>
// <param name="SingleMsixNum" type="uint32 "></param>
// <param name="SriovPort1NumVFs" type="uint32 "></param>
// <param name="SriovPort2NumVFs" type="uint32 "></param>
// <param name="SriovPortMode" type="uint32 "></param>
// <param name="StatFlags" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_DriverCoreSettingData) SetValue( /* IN */ LogNumRdmaRc uint32,
	/* IN */ LogNumSrq uint32,
	/* IN */ LogNumCq uint32,
	/* IN */ LogNumMcg uint32,
	/* IN */ LogNumMpt uint32,
	/* IN */ LogNumMtt uint32,
	/* IN */ LogNumMac uint32,
	/* IN */ LogNumVlan uint32,
	/* IN */ NumFcExch uint32,
	/* IN */ LogMttsPerSeg uint32,
	/* IN */ ModeFlags uint32,
	/* IN */ StatFlags uint32,
	/* IN */ SingleMsixNum uint32,
	/* IN */ MultiMsixNum uint32,
	/* IN */ SingleEqNum uint32,
	/* IN */ MultiEqNum uint32,
	/* IN */ MaxContQuant uint32,
	/* IN */ DebugFlags uint32,
	/* IN */ DebugLevel uint32,
	/* IN */ MaximumWorkingThreads uint32,
	/* IN */ Set4kMtu bool,
	/* IN */ SriovPortMode uint32,
	/* IN */ SriovPort1NumVFs uint32,
	/* IN */ SriovPort2NumVFs uint32,
	/* IN */ RoceMode string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetValue", LogNumRdmaRc, LogNumSrq, LogNumCq, LogNumMcg, LogNumMpt, LogNumMtt, LogNumMac, LogNumVlan, NumFcExch, LogMttsPerSeg, ModeFlags, StatFlags, SingleMsixNum, MultiMsixNum, SingleEqNum, MultiEqNum, MaxContQuant, DebugFlags, DebugLevel, MaximumWorkingThreads, Set4kMtu, SriovPortMode, SriovPort1NumVFs, SriovPort2NumVFs, RoceMode)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_DriverCoreSettingData) SetDefault() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetDefault")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="RoceMode" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_DriverCoreSettingData) SetRoceMode( /* IN */ RoceMode string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetRoceMode", RoceMode)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
