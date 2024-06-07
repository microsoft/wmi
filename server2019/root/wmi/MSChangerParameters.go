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

// MSChangerParameters struct
type MSChangerParameters struct { 
	*MSChangerDriver

	// 
	Active bool

	// 
	InstanceName string

	// 
	MagazineSize uint32

	// 
	NumberOfCleanerSlots uint32

	// 
	NumberOfDoors uint32

	// 
	NumberOfDrives uint32

	// 
	NumberOfIEPorts uint32

	// 
	NumberOfSlots uint32

	// 
	NumberOfTransports uint32
}

	func NewMSChangerParametersEx1(instance *cim.WmiInstance) (newInstance *MSChangerParameters, err error) {tmp, err := NewMSChangerDriverEx1(instance)
		
	if err != nil { return }
	newInstance = &MSChangerParameters {
	MSChangerDriver: tmp,
	}
	return
	}
	

	func NewMSChangerParametersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSChangerParameters, err error) {tmp, err := NewMSChangerDriverEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSChangerParameters {
	MSChangerDriver: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSChangerParameters) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSChangerParameters) GetPropertyActive()(value bool, err error) { 
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
func (instance *MSChangerParameters) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSChangerParameters) GetPropertyInstanceName()(value string, err error) { 
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

// SetMagazineSize sets the value of MagazineSize for the instance
func (instance *MSChangerParameters) SetPropertyMagazineSize(value uint32) (err error) { 
    return instance.SetProperty("MagazineSize", (value))
}

// GetMagazineSize gets the value of MagazineSize for the instance
func (instance *MSChangerParameters) GetPropertyMagazineSize()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MagazineSize")
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

// SetNumberOfCleanerSlots sets the value of NumberOfCleanerSlots for the instance
func (instance *MSChangerParameters) SetPropertyNumberOfCleanerSlots(value uint32) (err error) { 
    return instance.SetProperty("NumberOfCleanerSlots", (value))
}

// GetNumberOfCleanerSlots gets the value of NumberOfCleanerSlots for the instance
func (instance *MSChangerParameters) GetPropertyNumberOfCleanerSlots()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NumberOfCleanerSlots")
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

// SetNumberOfDoors sets the value of NumberOfDoors for the instance
func (instance *MSChangerParameters) SetPropertyNumberOfDoors(value uint32) (err error) { 
    return instance.SetProperty("NumberOfDoors", (value))
}

// GetNumberOfDoors gets the value of NumberOfDoors for the instance
func (instance *MSChangerParameters) GetPropertyNumberOfDoors()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NumberOfDoors")
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

// SetNumberOfDrives sets the value of NumberOfDrives for the instance
func (instance *MSChangerParameters) SetPropertyNumberOfDrives(value uint32) (err error) { 
    return instance.SetProperty("NumberOfDrives", (value))
}

// GetNumberOfDrives gets the value of NumberOfDrives for the instance
func (instance *MSChangerParameters) GetPropertyNumberOfDrives()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NumberOfDrives")
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

// SetNumberOfIEPorts sets the value of NumberOfIEPorts for the instance
func (instance *MSChangerParameters) SetPropertyNumberOfIEPorts(value uint32) (err error) { 
    return instance.SetProperty("NumberOfIEPorts", (value))
}

// GetNumberOfIEPorts gets the value of NumberOfIEPorts for the instance
func (instance *MSChangerParameters) GetPropertyNumberOfIEPorts()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NumberOfIEPorts")
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

// SetNumberOfSlots sets the value of NumberOfSlots for the instance
func (instance *MSChangerParameters) SetPropertyNumberOfSlots(value uint32) (err error) { 
    return instance.SetProperty("NumberOfSlots", (value))
}

// GetNumberOfSlots gets the value of NumberOfSlots for the instance
func (instance *MSChangerParameters) GetPropertyNumberOfSlots()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NumberOfSlots")
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

// SetNumberOfTransports sets the value of NumberOfTransports for the instance
func (instance *MSChangerParameters) SetPropertyNumberOfTransports(value uint32) (err error) { 
    return instance.SetProperty("NumberOfTransports", (value))
}

// GetNumberOfTransports gets the value of NumberOfTransports for the instance
func (instance *MSChangerParameters) GetPropertyNumberOfTransports()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NumberOfTransports")
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

