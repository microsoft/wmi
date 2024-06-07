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
)

// ConfigurationSectionWithCollection struct
type ConfigurationSectionWithCollection struct { 
	*ConfigurationSection
}

	func NewConfigurationSectionWithCollectionEx1(instance *cim.WmiInstance) (newInstance *ConfigurationSectionWithCollection, err error) {tmp, err := NewConfigurationSectionEx1(instance)
		
	if err != nil { return }
	newInstance = &ConfigurationSectionWithCollection {
	ConfigurationSection: tmp,
	}
	return
	}
	

	func NewConfigurationSectionWithCollectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ConfigurationSectionWithCollection, err error) {tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ConfigurationSectionWithCollection {
	ConfigurationSection: tmp,
	}
	return
	}
	

// 

// <param name="CollectionName" type="string "></param>
// <param name="ElementIdentifier" type="string "></param>

// <param name="Element" type="CollectionElement "></param>
func (instance *ConfigurationSectionWithCollection) Get( /* IN */ CollectionName string,
 /* IN */ ElementIdentifier string,
 /* OUT */ Element CollectionElement) (err error) {_, err = instance.InvokeMethod("Get" , CollectionName, ElementIdentifier)
	if err != nil { return }
	return
	
}


// 

// <param name="CollectionName" type="string "></param>
// <param name="Element" type="CollectionElement "></param>
func (instance *ConfigurationSectionWithCollection) Add( /* IN */ CollectionName string,
 /* IN */ Element CollectionElement) (err error) {_, err = instance.InvokeMethodWithReturn("Add" , CollectionName, Element);
	if err != nil { return }
	return
	
}


// 

// <param name="CollectionName" type="string "></param>
// <param name="Element" type="CollectionElement "></param>
func (instance *ConfigurationSectionWithCollection) Remove( /* IN */ CollectionName string,
 /* IN */ Element CollectionElement) (err error) {_, err = instance.InvokeMethodWithReturn("Remove" , CollectionName, Element);
	if err != nil { return }
	return
	
}


// 

// <param name="CollectionName" type="string "></param>
func (instance *ConfigurationSectionWithCollection) Clear( /* IN */ CollectionName string) (err error) {_, err = instance.InvokeMethodWithReturn("Clear" , CollectionName);
	if err != nil { return }
	return
	
}


