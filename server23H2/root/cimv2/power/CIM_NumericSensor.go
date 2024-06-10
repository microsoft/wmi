// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2.power
//
// ////////////////////////////////////////////
package power

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_NumericSensor struct
type CIM_NumericSensor struct {
	*CIM_Sensor

	//
	Accuracy int32

	//
	BaseUnits uint16

	//
	CurrentReading int32

	//
	EnabledThresholds []uint16

	//
	Hysteresis uint32

	//
	IsLinear bool

	//
	LowerThresholdCritical int32

	//
	LowerThresholdFatal int32

	//
	LowerThresholdNonCritical int32

	//
	MaxReadable int32

	//
	MinReadable int32

	//
	NominalReading int32

	//
	NormalMax int32

	//
	NormalMin int32

	//
	RateUnits uint16

	//
	Resolution uint32

	//
	SettableThresholds []uint16

	//
	SupportedThresholds []uint16

	//
	Tolerance int32

	//
	UnitModifier int32

	//
	UpperThresholdCritical int32

	//
	UpperThresholdFatal int32

	//
	UpperThresholdNonCritical int32
}

func NewCIM_NumericSensorEx1(instance *cim.WmiInstance) (newInstance *CIM_NumericSensor, err error) {
	tmp, err := NewCIM_SensorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_NumericSensor{
		CIM_Sensor: tmp,
	}
	return
}

func NewCIM_NumericSensorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_NumericSensor, err error) {
	tmp, err := NewCIM_SensorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_NumericSensor{
		CIM_Sensor: tmp,
	}
	return
}

// SetAccuracy sets the value of Accuracy for the instance
func (instance *CIM_NumericSensor) SetPropertyAccuracy(value int32) (err error) {
	return instance.SetProperty("Accuracy", (value))
}

// GetAccuracy gets the value of Accuracy for the instance
func (instance *CIM_NumericSensor) GetPropertyAccuracy() (value int32, err error) {
	retValue, err := instance.GetProperty("Accuracy")
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

	value = int32(valuetmp)

	return
}

// SetBaseUnits sets the value of BaseUnits for the instance
func (instance *CIM_NumericSensor) SetPropertyBaseUnits(value uint16) (err error) {
	return instance.SetProperty("BaseUnits", (value))
}

// GetBaseUnits gets the value of BaseUnits for the instance
func (instance *CIM_NumericSensor) GetPropertyBaseUnits() (value uint16, err error) {
	retValue, err := instance.GetProperty("BaseUnits")
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

// SetCurrentReading sets the value of CurrentReading for the instance
func (instance *CIM_NumericSensor) SetPropertyCurrentReading(value int32) (err error) {
	return instance.SetProperty("CurrentReading", (value))
}

// GetCurrentReading gets the value of CurrentReading for the instance
func (instance *CIM_NumericSensor) GetPropertyCurrentReading() (value int32, err error) {
	retValue, err := instance.GetProperty("CurrentReading")
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

	value = int32(valuetmp)

	return
}

// SetEnabledThresholds sets the value of EnabledThresholds for the instance
func (instance *CIM_NumericSensor) SetPropertyEnabledThresholds(value []uint16) (err error) {
	return instance.SetProperty("EnabledThresholds", (value))
}

// GetEnabledThresholds gets the value of EnabledThresholds for the instance
func (instance *CIM_NumericSensor) GetPropertyEnabledThresholds() (value []uint16, err error) {
	retValue, err := instance.GetProperty("EnabledThresholds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetHysteresis sets the value of Hysteresis for the instance
func (instance *CIM_NumericSensor) SetPropertyHysteresis(value uint32) (err error) {
	return instance.SetProperty("Hysteresis", (value))
}

// GetHysteresis gets the value of Hysteresis for the instance
func (instance *CIM_NumericSensor) GetPropertyHysteresis() (value uint32, err error) {
	retValue, err := instance.GetProperty("Hysteresis")
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

// SetIsLinear sets the value of IsLinear for the instance
func (instance *CIM_NumericSensor) SetPropertyIsLinear(value bool) (err error) {
	return instance.SetProperty("IsLinear", (value))
}

// GetIsLinear gets the value of IsLinear for the instance
func (instance *CIM_NumericSensor) GetPropertyIsLinear() (value bool, err error) {
	retValue, err := instance.GetProperty("IsLinear")
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

// SetLowerThresholdCritical sets the value of LowerThresholdCritical for the instance
func (instance *CIM_NumericSensor) SetPropertyLowerThresholdCritical(value int32) (err error) {
	return instance.SetProperty("LowerThresholdCritical", (value))
}

// GetLowerThresholdCritical gets the value of LowerThresholdCritical for the instance
func (instance *CIM_NumericSensor) GetPropertyLowerThresholdCritical() (value int32, err error) {
	retValue, err := instance.GetProperty("LowerThresholdCritical")
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

	value = int32(valuetmp)

	return
}

// SetLowerThresholdFatal sets the value of LowerThresholdFatal for the instance
func (instance *CIM_NumericSensor) SetPropertyLowerThresholdFatal(value int32) (err error) {
	return instance.SetProperty("LowerThresholdFatal", (value))
}

// GetLowerThresholdFatal gets the value of LowerThresholdFatal for the instance
func (instance *CIM_NumericSensor) GetPropertyLowerThresholdFatal() (value int32, err error) {
	retValue, err := instance.GetProperty("LowerThresholdFatal")
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

	value = int32(valuetmp)

	return
}

// SetLowerThresholdNonCritical sets the value of LowerThresholdNonCritical for the instance
func (instance *CIM_NumericSensor) SetPropertyLowerThresholdNonCritical(value int32) (err error) {
	return instance.SetProperty("LowerThresholdNonCritical", (value))
}

// GetLowerThresholdNonCritical gets the value of LowerThresholdNonCritical for the instance
func (instance *CIM_NumericSensor) GetPropertyLowerThresholdNonCritical() (value int32, err error) {
	retValue, err := instance.GetProperty("LowerThresholdNonCritical")
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

	value = int32(valuetmp)

	return
}

// SetMaxReadable sets the value of MaxReadable for the instance
func (instance *CIM_NumericSensor) SetPropertyMaxReadable(value int32) (err error) {
	return instance.SetProperty("MaxReadable", (value))
}

// GetMaxReadable gets the value of MaxReadable for the instance
func (instance *CIM_NumericSensor) GetPropertyMaxReadable() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxReadable")
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

	value = int32(valuetmp)

	return
}

// SetMinReadable sets the value of MinReadable for the instance
func (instance *CIM_NumericSensor) SetPropertyMinReadable(value int32) (err error) {
	return instance.SetProperty("MinReadable", (value))
}

// GetMinReadable gets the value of MinReadable for the instance
func (instance *CIM_NumericSensor) GetPropertyMinReadable() (value int32, err error) {
	retValue, err := instance.GetProperty("MinReadable")
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

	value = int32(valuetmp)

	return
}

// SetNominalReading sets the value of NominalReading for the instance
func (instance *CIM_NumericSensor) SetPropertyNominalReading(value int32) (err error) {
	return instance.SetProperty("NominalReading", (value))
}

// GetNominalReading gets the value of NominalReading for the instance
func (instance *CIM_NumericSensor) GetPropertyNominalReading() (value int32, err error) {
	retValue, err := instance.GetProperty("NominalReading")
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

	value = int32(valuetmp)

	return
}

// SetNormalMax sets the value of NormalMax for the instance
func (instance *CIM_NumericSensor) SetPropertyNormalMax(value int32) (err error) {
	return instance.SetProperty("NormalMax", (value))
}

// GetNormalMax gets the value of NormalMax for the instance
func (instance *CIM_NumericSensor) GetPropertyNormalMax() (value int32, err error) {
	retValue, err := instance.GetProperty("NormalMax")
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

	value = int32(valuetmp)

	return
}

// SetNormalMin sets the value of NormalMin for the instance
func (instance *CIM_NumericSensor) SetPropertyNormalMin(value int32) (err error) {
	return instance.SetProperty("NormalMin", (value))
}

// GetNormalMin gets the value of NormalMin for the instance
func (instance *CIM_NumericSensor) GetPropertyNormalMin() (value int32, err error) {
	retValue, err := instance.GetProperty("NormalMin")
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

	value = int32(valuetmp)

	return
}

// SetRateUnits sets the value of RateUnits for the instance
func (instance *CIM_NumericSensor) SetPropertyRateUnits(value uint16) (err error) {
	return instance.SetProperty("RateUnits", (value))
}

// GetRateUnits gets the value of RateUnits for the instance
func (instance *CIM_NumericSensor) GetPropertyRateUnits() (value uint16, err error) {
	retValue, err := instance.GetProperty("RateUnits")
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

// SetResolution sets the value of Resolution for the instance
func (instance *CIM_NumericSensor) SetPropertyResolution(value uint32) (err error) {
	return instance.SetProperty("Resolution", (value))
}

// GetResolution gets the value of Resolution for the instance
func (instance *CIM_NumericSensor) GetPropertyResolution() (value uint32, err error) {
	retValue, err := instance.GetProperty("Resolution")
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

// SetSettableThresholds sets the value of SettableThresholds for the instance
func (instance *CIM_NumericSensor) SetPropertySettableThresholds(value []uint16) (err error) {
	return instance.SetProperty("SettableThresholds", (value))
}

// GetSettableThresholds gets the value of SettableThresholds for the instance
func (instance *CIM_NumericSensor) GetPropertySettableThresholds() (value []uint16, err error) {
	retValue, err := instance.GetProperty("SettableThresholds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetSupportedThresholds sets the value of SupportedThresholds for the instance
func (instance *CIM_NumericSensor) SetPropertySupportedThresholds(value []uint16) (err error) {
	return instance.SetProperty("SupportedThresholds", (value))
}

// GetSupportedThresholds gets the value of SupportedThresholds for the instance
func (instance *CIM_NumericSensor) GetPropertySupportedThresholds() (value []uint16, err error) {
	retValue, err := instance.GetProperty("SupportedThresholds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetTolerance sets the value of Tolerance for the instance
func (instance *CIM_NumericSensor) SetPropertyTolerance(value int32) (err error) {
	return instance.SetProperty("Tolerance", (value))
}

// GetTolerance gets the value of Tolerance for the instance
func (instance *CIM_NumericSensor) GetPropertyTolerance() (value int32, err error) {
	retValue, err := instance.GetProperty("Tolerance")
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

	value = int32(valuetmp)

	return
}

// SetUnitModifier sets the value of UnitModifier for the instance
func (instance *CIM_NumericSensor) SetPropertyUnitModifier(value int32) (err error) {
	return instance.SetProperty("UnitModifier", (value))
}

// GetUnitModifier gets the value of UnitModifier for the instance
func (instance *CIM_NumericSensor) GetPropertyUnitModifier() (value int32, err error) {
	retValue, err := instance.GetProperty("UnitModifier")
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

	value = int32(valuetmp)

	return
}

// SetUpperThresholdCritical sets the value of UpperThresholdCritical for the instance
func (instance *CIM_NumericSensor) SetPropertyUpperThresholdCritical(value int32) (err error) {
	return instance.SetProperty("UpperThresholdCritical", (value))
}

// GetUpperThresholdCritical gets the value of UpperThresholdCritical for the instance
func (instance *CIM_NumericSensor) GetPropertyUpperThresholdCritical() (value int32, err error) {
	retValue, err := instance.GetProperty("UpperThresholdCritical")
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

	value = int32(valuetmp)

	return
}

// SetUpperThresholdFatal sets the value of UpperThresholdFatal for the instance
func (instance *CIM_NumericSensor) SetPropertyUpperThresholdFatal(value int32) (err error) {
	return instance.SetProperty("UpperThresholdFatal", (value))
}

// GetUpperThresholdFatal gets the value of UpperThresholdFatal for the instance
func (instance *CIM_NumericSensor) GetPropertyUpperThresholdFatal() (value int32, err error) {
	retValue, err := instance.GetProperty("UpperThresholdFatal")
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

	value = int32(valuetmp)

	return
}

// SetUpperThresholdNonCritical sets the value of UpperThresholdNonCritical for the instance
func (instance *CIM_NumericSensor) SetPropertyUpperThresholdNonCritical(value int32) (err error) {
	return instance.SetProperty("UpperThresholdNonCritical", (value))
}

// GetUpperThresholdNonCritical gets the value of UpperThresholdNonCritical for the instance
func (instance *CIM_NumericSensor) GetPropertyUpperThresholdNonCritical() (value int32, err error) {
	retValue, err := instance.GetProperty("UpperThresholdNonCritical")
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

	value = int32(valuetmp)

	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_NumericSensor) RestoreDefaultThresholds() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RestoreDefaultThresholds")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="SensorReading" type="int32 "></param>

// <param name="Accuracy" type="int32 "></param>
// <param name="Hysteresis" type="uint32 "></param>
// <param name="Resolution" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="Tolerance" type="int32 "></param>
func (instance *CIM_NumericSensor) GetNonLinearFactors( /* IN */ SensorReading int32,
	/* OUT */ Accuracy int32,
	/* OUT */ Resolution uint32,
	/* OUT */ Tolerance int32,
	/* OUT */ Hysteresis uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetNonLinearFactors", SensorReading)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
