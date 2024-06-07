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
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MS_SMHBA_SCSIENTRY struct
type MS_SMHBA_SCSIENTRY struct { 
	*cim.WmiInstance

	// 
	LUID []uint8

	// 
	PortLun MS_SMHBA_PORTLUN

	// 
	ScsiId HBAScsiID
}

	func NewMS_SMHBA_SCSIENTRYEx1(instance *cim.WmiInstance) (newInstance *MS_SMHBA_SCSIENTRY, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MS_SMHBA_SCSIENTRY {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMS_SMHBA_SCSIENTRYEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MS_SMHBA_SCSIENTRY, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MS_SMHBA_SCSIENTRY {
	WmiInstance: tmp,
	}
	return
	}
	

// SetLUID sets the value of LUID for the instance
func (instance *MS_SMHBA_SCSIENTRY) SetPropertyLUID(value []uint8) (err error) { 
    return instance.SetProperty("LUID", (value))
}

// GetLUID gets the value of LUID for the instance
func (instance *MS_SMHBA_SCSIENTRY) GetPropertyLUID()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("LUID")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint8); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint8(valuetmp))
    }

    return
}

// SetPortLun sets the value of PortLun for the instance
func (instance *MS_SMHBA_SCSIENTRY) SetPropertyPortLun(value MS_SMHBA_PORTLUN) (err error) { 
    return instance.SetProperty("PortLun", (value))
}

// GetPortLun gets the value of PortLun for the instance
func (instance *MS_SMHBA_SCSIENTRY) GetPropertyPortLun()(value MS_SMHBA_PORTLUN, err error) { 
    retValue, err := instance.GetProperty("PortLun")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MS_SMHBA_PORTLUN); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MS_SMHBA_PORTLUN is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MS_SMHBA_PORTLUN(valuetmp)

    return
}

// SetScsiId sets the value of ScsiId for the instance
func (instance *MS_SMHBA_SCSIENTRY) SetPropertyScsiId(value HBAScsiID) (err error) { 
    return instance.SetProperty("ScsiId", (value))
}

// GetScsiId gets the value of ScsiId for the instance
func (instance *MS_SMHBA_SCSIENTRY) GetPropertyScsiId()(value HBAScsiID, err error) { 
    retValue, err := instance.GetProperty("ScsiId")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(HBAScsiID); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " HBAScsiID is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = HBAScsiID(valuetmp)

    return
}

