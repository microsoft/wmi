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

// ListenerAdaptersSection struct
type ListenerAdaptersSection struct { 
	*ConfigurationSectionWithCollection

	// 
	ListenerAdapters []ListenerAdapterElement
}

	func NewListenerAdaptersSectionEx1(instance *cim.WmiInstance) (newInstance *ListenerAdaptersSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &ListenerAdaptersSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewListenerAdaptersSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ListenerAdaptersSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ListenerAdaptersSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetListenerAdapters sets the value of ListenerAdapters for the instance
func (instance *ListenerAdaptersSection) SetPropertyListenerAdapters(value []ListenerAdapterElement) (err error) { 
    return instance.SetProperty("ListenerAdapters", (value))
}

// GetListenerAdapters gets the value of ListenerAdapters for the instance
func (instance *ListenerAdaptersSection) GetPropertyListenerAdapters()(value []ListenerAdapterElement, err error) { 
    retValue, err := instance.GetProperty("ListenerAdapters")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(ListenerAdapterElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " ListenerAdapterElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, ListenerAdapterElement(valuetmp))
    }

    return
}

