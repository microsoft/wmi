// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSTapeProblemIoError struct
type MSTapeProblemIoError struct { 
	*MSTapeDriver

	// 
	Active bool

	// 
	InstanceName string

	// 
	NonMediumErrors uint32

	// 
	ReadCorrectedWithDelay uint32

	// 
	ReadCorrectedWithoutDelay uint32

	// 
	ReadCorrectionAlgorithmProcessed uint32

	// 
	ReadTotalCorrectedErrors uint32

	// 
	ReadTotalErrors uint32

	// 
	ReadTotalUncorrectedErrors uint32

	// 
	WriteCorrectedWithDelay uint32

	// 
	WriteCorrectedWithoutDelay uint32

	// 
	WriteCorrectionAlgorithmProcessed uint32

	// 
	WriteTotalCorrectedErrors uint32

	// 
	WriteTotalErrors uint32

	// 
	WriteTotalUncorrectedErrors uint32
}

	func NewMSTapeProblemIoErrorEx1(instance *cim.WmiInstance) (newInstance *MSTapeProblemIoError, err error) {tmp, err := NewMSTapeDriverEx1(instance)
		
	if err != nil { return }
	newInstance = &MSTapeProblemIoError {
	MSTapeDriver: tmp,
	}
	return
	}
	

	func NewMSTapeProblemIoErrorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSTapeProblemIoError, err error) {tmp, err := NewMSTapeDriverEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSTapeProblemIoError {
	MSTapeDriver: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSTapeProblemIoError) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSTapeProblemIoError) GetPropertyActive()(value bool, err error) { 
    retValue, err := instance.GetProperty("Active")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSTapeProblemIoError) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSTapeProblemIoError) GetPropertyInstanceName()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceName")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetNonMediumErrors sets the value of NonMediumErrors for the instance
func (instance *MSTapeProblemIoError) SetPropertyNonMediumErrors(value uint32) (err error) { 
    return instance.SetProperty("NonMediumErrors", (value))
}

// GetNonMediumErrors gets the value of NonMediumErrors for the instance
func (instance *MSTapeProblemIoError) GetPropertyNonMediumErrors()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NonMediumErrors")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetReadCorrectedWithDelay sets the value of ReadCorrectedWithDelay for the instance
func (instance *MSTapeProblemIoError) SetPropertyReadCorrectedWithDelay(value uint32) (err error) { 
    return instance.SetProperty("ReadCorrectedWithDelay", (value))
}

// GetReadCorrectedWithDelay gets the value of ReadCorrectedWithDelay for the instance
func (instance *MSTapeProblemIoError) GetPropertyReadCorrectedWithDelay()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ReadCorrectedWithDelay")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetReadCorrectedWithoutDelay sets the value of ReadCorrectedWithoutDelay for the instance
func (instance *MSTapeProblemIoError) SetPropertyReadCorrectedWithoutDelay(value uint32) (err error) { 
    return instance.SetProperty("ReadCorrectedWithoutDelay", (value))
}

// GetReadCorrectedWithoutDelay gets the value of ReadCorrectedWithoutDelay for the instance
func (instance *MSTapeProblemIoError) GetPropertyReadCorrectedWithoutDelay()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ReadCorrectedWithoutDelay")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetReadCorrectionAlgorithmProcessed sets the value of ReadCorrectionAlgorithmProcessed for the instance
func (instance *MSTapeProblemIoError) SetPropertyReadCorrectionAlgorithmProcessed(value uint32) (err error) { 
    return instance.SetProperty("ReadCorrectionAlgorithmProcessed", (value))
}

// GetReadCorrectionAlgorithmProcessed gets the value of ReadCorrectionAlgorithmProcessed for the instance
func (instance *MSTapeProblemIoError) GetPropertyReadCorrectionAlgorithmProcessed()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ReadCorrectionAlgorithmProcessed")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetReadTotalCorrectedErrors sets the value of ReadTotalCorrectedErrors for the instance
func (instance *MSTapeProblemIoError) SetPropertyReadTotalCorrectedErrors(value uint32) (err error) { 
    return instance.SetProperty("ReadTotalCorrectedErrors", (value))
}

// GetReadTotalCorrectedErrors gets the value of ReadTotalCorrectedErrors for the instance
func (instance *MSTapeProblemIoError) GetPropertyReadTotalCorrectedErrors()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ReadTotalCorrectedErrors")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetReadTotalErrors sets the value of ReadTotalErrors for the instance
func (instance *MSTapeProblemIoError) SetPropertyReadTotalErrors(value uint32) (err error) { 
    return instance.SetProperty("ReadTotalErrors", (value))
}

// GetReadTotalErrors gets the value of ReadTotalErrors for the instance
func (instance *MSTapeProblemIoError) GetPropertyReadTotalErrors()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ReadTotalErrors")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetReadTotalUncorrectedErrors sets the value of ReadTotalUncorrectedErrors for the instance
func (instance *MSTapeProblemIoError) SetPropertyReadTotalUncorrectedErrors(value uint32) (err error) { 
    return instance.SetProperty("ReadTotalUncorrectedErrors", (value))
}

// GetReadTotalUncorrectedErrors gets the value of ReadTotalUncorrectedErrors for the instance
func (instance *MSTapeProblemIoError) GetPropertyReadTotalUncorrectedErrors()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ReadTotalUncorrectedErrors")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetWriteCorrectedWithDelay sets the value of WriteCorrectedWithDelay for the instance
func (instance *MSTapeProblemIoError) SetPropertyWriteCorrectedWithDelay(value uint32) (err error) { 
    return instance.SetProperty("WriteCorrectedWithDelay", (value))
}

// GetWriteCorrectedWithDelay gets the value of WriteCorrectedWithDelay for the instance
func (instance *MSTapeProblemIoError) GetPropertyWriteCorrectedWithDelay()(value uint32, err error) { 
    retValue, err := instance.GetProperty("WriteCorrectedWithDelay")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetWriteCorrectedWithoutDelay sets the value of WriteCorrectedWithoutDelay for the instance
func (instance *MSTapeProblemIoError) SetPropertyWriteCorrectedWithoutDelay(value uint32) (err error) { 
    return instance.SetProperty("WriteCorrectedWithoutDelay", (value))
}

// GetWriteCorrectedWithoutDelay gets the value of WriteCorrectedWithoutDelay for the instance
func (instance *MSTapeProblemIoError) GetPropertyWriteCorrectedWithoutDelay()(value uint32, err error) { 
    retValue, err := instance.GetProperty("WriteCorrectedWithoutDelay")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetWriteCorrectionAlgorithmProcessed sets the value of WriteCorrectionAlgorithmProcessed for the instance
func (instance *MSTapeProblemIoError) SetPropertyWriteCorrectionAlgorithmProcessed(value uint32) (err error) { 
    return instance.SetProperty("WriteCorrectionAlgorithmProcessed", (value))
}

// GetWriteCorrectionAlgorithmProcessed gets the value of WriteCorrectionAlgorithmProcessed for the instance
func (instance *MSTapeProblemIoError) GetPropertyWriteCorrectionAlgorithmProcessed()(value uint32, err error) { 
    retValue, err := instance.GetProperty("WriteCorrectionAlgorithmProcessed")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetWriteTotalCorrectedErrors sets the value of WriteTotalCorrectedErrors for the instance
func (instance *MSTapeProblemIoError) SetPropertyWriteTotalCorrectedErrors(value uint32) (err error) { 
    return instance.SetProperty("WriteTotalCorrectedErrors", (value))
}

// GetWriteTotalCorrectedErrors gets the value of WriteTotalCorrectedErrors for the instance
func (instance *MSTapeProblemIoError) GetPropertyWriteTotalCorrectedErrors()(value uint32, err error) { 
    retValue, err := instance.GetProperty("WriteTotalCorrectedErrors")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetWriteTotalErrors sets the value of WriteTotalErrors for the instance
func (instance *MSTapeProblemIoError) SetPropertyWriteTotalErrors(value uint32) (err error) { 
    return instance.SetProperty("WriteTotalErrors", (value))
}

// GetWriteTotalErrors gets the value of WriteTotalErrors for the instance
func (instance *MSTapeProblemIoError) GetPropertyWriteTotalErrors()(value uint32, err error) { 
    retValue, err := instance.GetProperty("WriteTotalErrors")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetWriteTotalUncorrectedErrors sets the value of WriteTotalUncorrectedErrors for the instance
func (instance *MSTapeProblemIoError) SetPropertyWriteTotalUncorrectedErrors(value uint32) (err error) { 
    return instance.SetProperty("WriteTotalUncorrectedErrors", (value))
}

// GetWriteTotalUncorrectedErrors gets the value of WriteTotalUncorrectedErrors for the instance
func (instance *MSTapeProblemIoError) GetPropertyWriteTotalUncorrectedErrors()(value uint32, err error) { 
    retValue, err := instance.GetProperty("WriteTotalUncorrectedErrors")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

