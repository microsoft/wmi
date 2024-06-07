// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MDM_Policy_Config01_Printers02 struct
type MDM_Policy_Config01_Printers02 struct { 
	*cim.WmiInstance

	// 
	ApprovedUsbPrintDevices string

	// 
	EnableDeviceControl string

	// 
	InstanceID string

	// 
	ParentID string

	// 
	PointAndPrintRestrictions string

	// 
	PublishPrinters string
}

	func NewMDM_Policy_Config01_Printers02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_Printers02, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MDM_Policy_Config01_Printers02 {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMDM_Policy_Config01_Printers02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MDM_Policy_Config01_Printers02, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MDM_Policy_Config01_Printers02 {
	WmiInstance: tmp,
	}
	return
	}
	

// SetApprovedUsbPrintDevices sets the value of ApprovedUsbPrintDevices for the instance
func (instance *MDM_Policy_Config01_Printers02) SetPropertyApprovedUsbPrintDevices(value string) (err error) { 
    return instance.SetProperty("ApprovedUsbPrintDevices", (value))
}

// GetApprovedUsbPrintDevices gets the value of ApprovedUsbPrintDevices for the instance
func (instance *MDM_Policy_Config01_Printers02) GetPropertyApprovedUsbPrintDevices()(value string, err error) { 
    retValue, err := instance.GetProperty("ApprovedUsbPrintDevices")
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

// SetEnableDeviceControl sets the value of EnableDeviceControl for the instance
func (instance *MDM_Policy_Config01_Printers02) SetPropertyEnableDeviceControl(value string) (err error) { 
    return instance.SetProperty("EnableDeviceControl", (value))
}

// GetEnableDeviceControl gets the value of EnableDeviceControl for the instance
func (instance *MDM_Policy_Config01_Printers02) GetPropertyEnableDeviceControl()(value string, err error) { 
    retValue, err := instance.GetProperty("EnableDeviceControl")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Printers02) SetPropertyInstanceID(value string) (err error) { 
    return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Printers02) GetPropertyInstanceID()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceID")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Printers02) SetPropertyParentID(value string) (err error) { 
    return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Printers02) GetPropertyParentID()(value string, err error) { 
    retValue, err := instance.GetProperty("ParentID")
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

// SetPointAndPrintRestrictions sets the value of PointAndPrintRestrictions for the instance
func (instance *MDM_Policy_Config01_Printers02) SetPropertyPointAndPrintRestrictions(value string) (err error) { 
    return instance.SetProperty("PointAndPrintRestrictions", (value))
}

// GetPointAndPrintRestrictions gets the value of PointAndPrintRestrictions for the instance
func (instance *MDM_Policy_Config01_Printers02) GetPropertyPointAndPrintRestrictions()(value string, err error) { 
    retValue, err := instance.GetProperty("PointAndPrintRestrictions")
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

// SetPublishPrinters sets the value of PublishPrinters for the instance
func (instance *MDM_Policy_Config01_Printers02) SetPropertyPublishPrinters(value string) (err error) { 
    return instance.SetProperty("PublishPrinters", (value))
}

// GetPublishPrinters gets the value of PublishPrinters for the instance
func (instance *MDM_Policy_Config01_Printers02) GetPropertyPublishPrinters()(value string, err error) { 
    retValue, err := instance.GetProperty("PublishPrinters")
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

