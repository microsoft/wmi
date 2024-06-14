// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.SDDC.Management
//////////////////////////////////////////////
package management

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SDDC_Vhd struct
type SDDC_Vhd struct {
	*cim.WmiInstance

	//
	Alerts []SDDC_Alert

	//
	AverageLatency float64

	//
	FilePath string

	//
	Id string

	//
	Size uint64

	//
	SizeUsed uint64

	//
	TotalIops float64

	//
	TotalThroughput float64

	//
	VhdFormat uint16

	//
	VhdType uint16

	//
	VolumeId string
}

func NewSDDC_VhdEx1(instance *cim.WmiInstance) (newInstance *SDDC_Vhd, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &SDDC_Vhd{
		WmiInstance: tmp,
	}
	return
}

func NewSDDC_VhdEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SDDC_Vhd, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SDDC_Vhd{
		WmiInstance: tmp,
	}
	return
}

// SetAlerts sets the value of Alerts for the instance
func (instance *SDDC_Vhd) SetPropertyAlerts(value []SDDC_Alert) (err error) {
	return instance.SetProperty("Alerts", (value))
}

// GetAlerts gets the value of Alerts for the instance
func (instance *SDDC_Vhd) GetPropertyAlerts() (value []SDDC_Alert, err error) {
	retValue, err := instance.GetProperty("Alerts")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(SDDC_Alert)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " SDDC_Alert is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, SDDC_Alert(valuetmp))
	}

	return
}

// SetAverageLatency sets the value of AverageLatency for the instance
func (instance *SDDC_Vhd) SetPropertyAverageLatency(value float64) (err error) {
	return instance.SetProperty("AverageLatency", (value))
}

// GetAverageLatency gets the value of AverageLatency for the instance
func (instance *SDDC_Vhd) GetPropertyAverageLatency() (value float64, err error) {
	retValue, err := instance.GetProperty("AverageLatency")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(float64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " float64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = float64(valuetmp)

	return
}

// SetFilePath sets the value of FilePath for the instance
func (instance *SDDC_Vhd) SetPropertyFilePath(value string) (err error) {
	return instance.SetProperty("FilePath", (value))
}

// GetFilePath gets the value of FilePath for the instance
func (instance *SDDC_Vhd) GetPropertyFilePath() (value string, err error) {
	retValue, err := instance.GetProperty("FilePath")
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

// SetId sets the value of Id for the instance
func (instance *SDDC_Vhd) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *SDDC_Vhd) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetSize sets the value of Size for the instance
func (instance *SDDC_Vhd) SetPropertySize(value uint64) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *SDDC_Vhd) GetPropertySize() (value uint64, err error) {
	retValue, err := instance.GetProperty("Size")
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

// SetSizeUsed sets the value of SizeUsed for the instance
func (instance *SDDC_Vhd) SetPropertySizeUsed(value uint64) (err error) {
	return instance.SetProperty("SizeUsed", (value))
}

// GetSizeUsed gets the value of SizeUsed for the instance
func (instance *SDDC_Vhd) GetPropertySizeUsed() (value uint64, err error) {
	retValue, err := instance.GetProperty("SizeUsed")
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

// SetTotalIops sets the value of TotalIops for the instance
func (instance *SDDC_Vhd) SetPropertyTotalIops(value float64) (err error) {
	return instance.SetProperty("TotalIops", (value))
}

// GetTotalIops gets the value of TotalIops for the instance
func (instance *SDDC_Vhd) GetPropertyTotalIops() (value float64, err error) {
	retValue, err := instance.GetProperty("TotalIops")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(float64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " float64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = float64(valuetmp)

	return
}

// SetTotalThroughput sets the value of TotalThroughput for the instance
func (instance *SDDC_Vhd) SetPropertyTotalThroughput(value float64) (err error) {
	return instance.SetProperty("TotalThroughput", (value))
}

// GetTotalThroughput gets the value of TotalThroughput for the instance
func (instance *SDDC_Vhd) GetPropertyTotalThroughput() (value float64, err error) {
	retValue, err := instance.GetProperty("TotalThroughput")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(float64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " float64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = float64(valuetmp)

	return
}

// SetVhdFormat sets the value of VhdFormat for the instance
func (instance *SDDC_Vhd) SetPropertyVhdFormat(value uint16) (err error) {
	return instance.SetProperty("VhdFormat", (value))
}

// GetVhdFormat gets the value of VhdFormat for the instance
func (instance *SDDC_Vhd) GetPropertyVhdFormat() (value uint16, err error) {
	retValue, err := instance.GetProperty("VhdFormat")
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

// SetVhdType sets the value of VhdType for the instance
func (instance *SDDC_Vhd) SetPropertyVhdType(value uint16) (err error) {
	return instance.SetProperty("VhdType", (value))
}

// GetVhdType gets the value of VhdType for the instance
func (instance *SDDC_Vhd) GetPropertyVhdType() (value uint16, err error) {
	retValue, err := instance.GetProperty("VhdType")
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

// SetVolumeId sets the value of VolumeId for the instance
func (instance *SDDC_Vhd) SetPropertyVolumeId(value string) (err error) {
	return instance.SetProperty("VolumeId", (value))
}

// GetVolumeId gets the value of VolumeId for the instance
func (instance *SDDC_Vhd) GetPropertyVolumeId() (value string, err error) {
	retValue, err := instance.GetProperty("VolumeId")
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

// <param name="FilePath" type="string "></param>
// <param name="SeriesName" type="string "></param>
// <param name="TimeFrame" type="uint16 "></param>

// <param name="Metric" type="SDDC_Metric "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Vhd) GetMetrics( /* IN */ FilePath string,
	/* IN */ SeriesName string,
	/* IN */ TimeFrame uint16,
	/* OUT */ Metric SDDC_Metric) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetMetrics", FilePath, SeriesName, TimeFrame)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
