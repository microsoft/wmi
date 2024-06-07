// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// FtpHiddenSegmentSettings struct
type FtpHiddenSegmentSettings struct { 
	*EmbeddedObject

	// 
	HiddenSegments []SegmentElement
}

	func NewFtpHiddenSegmentSettingsEx1(instance *cim.WmiInstance) (newInstance *FtpHiddenSegmentSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &FtpHiddenSegmentSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewFtpHiddenSegmentSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *FtpHiddenSegmentSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &FtpHiddenSegmentSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetHiddenSegments sets the value of HiddenSegments for the instance
func (instance *FtpHiddenSegmentSettings) SetPropertyHiddenSegments(value []SegmentElement) (err error) { 
    return instance.SetProperty("HiddenSegments", (value))
}

// GetHiddenSegments gets the value of HiddenSegments for the instance
func (instance *FtpHiddenSegmentSettings) GetPropertyHiddenSegments()(value []SegmentElement, err error) { 
    retValue, err := instance.GetProperty("HiddenSegments")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(SegmentElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " SegmentElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, SegmentElement(valuetmp))
    }

    return
}

