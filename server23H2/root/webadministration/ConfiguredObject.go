// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ConfiguredObject struct
type ConfiguredObject struct {
	*Object
}

func NewConfiguredObjectEx1(instance *cim.WmiInstance) (newInstance *ConfiguredObject, err error) {
	tmp, err := NewObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ConfiguredObject{
		Object: tmp,
	}
	return
}

func NewConfiguredObjectEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ConfiguredObject, err error) {
	tmp, err := NewObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ConfiguredObject{
		Object: tmp,
	}
	return
}

//

// <param name="Sections" type="ConfigurationSection []"></param>
func (instance *ConfiguredObject) GetAllSections( /* OUT */ Sections []ConfigurationSection) (err error) {
	_, err = instance.InvokeMethod("GetAllSections")
	if err != nil {
		return
	}
	return

}

//

// <param name="SectionName" type="string "></param>

// <param name="Section" type="ConfigurationSection "></param>
func (instance *ConfiguredObject) GetSection( /* IN */ SectionName string,
	/* OUT */ Section ConfigurationSection) (err error) {
	_, err = instance.InvokeMethod("GetSection", SectionName)
	if err != nil {
		return
	}
	return

}

//

// <param name="AttributeName" type="string "></param>
// <param name="Location" type="string "></param>
// <param name="RelativePath" type="string "></param>
// <param name="SectionPath" type="string "></param>

// <param name="ReturnValue" type="string "></param>
func (instance *ConfiguredObject) GetAttribute( /* IN */ SectionPath string,
	/* IN */ AttributeName string,
	/* OPTIONAL IN */ RelativePath string,
	/* OPTIONAL IN */ Location string) (result string, err error) {
	retVal, err := instance.InvokeMethodWithReturn("GetAttribute", SectionPath, AttributeName, RelativePath, Location)
	if err != nil {
		return
	}
	result = string(retVal)
	return

}

//

// <param name="AttributeName" type="string "></param>
// <param name="Location" type="string "></param>
// <param name="RelativePath" type="string "></param>
// <param name="SectionPath" type="string "></param>
// <param name="Value" type="string "></param>
func (instance *ConfiguredObject) SetAttribute( /* IN */ SectionPath string,
	/* IN */ AttributeName string,
	/* OPTIONAL IN */ RelativePath string,
	/* OPTIONAL IN */ Location string,
	/* IN */ Value string) (err error) {
	_, err = instance.InvokeMethodWithReturn("SetAttribute", SectionPath, AttributeName, RelativePath, Location, Value)
	if err != nil {
		return
	}
	return

}

//

// <param name="AttributeName" type="string "></param>
// <param name="Location" type="string "></param>
// <param name="RelativePath" type="string "></param>
// <param name="SectionPath" type="string "></param>
func (instance *ConfiguredObject) ClearAttribute( /* IN */ SectionPath string,
	/* IN */ AttributeName string,
	/* OPTIONAL IN */ RelativePath string,
	/* OPTIONAL IN */ Location string) (err error) {
	_, err = instance.InvokeMethodWithReturn("ClearAttribute", SectionPath, AttributeName, RelativePath, Location)
	if err != nil {
		return
	}
	return

}

//

// <param name="AddElementName" type="string "></param>
// <param name="CollectionElementName" type="string "></param>
// <param name="ElementParameters" type="NameValueConfigurationElement []"></param>
// <param name="Location" type="string "></param>
// <param name="RelativePath" type="string "></param>
// <param name="SectionPath" type="string "></param>
func (instance *ConfiguredObject) AddCollectionElement( /* IN */ SectionPath string,
	/* IN */ CollectionElementName string,
	/* IN */ ElementParameters []NameValueConfigurationElement,
	/* OPTIONAL IN */ AddElementName string,
	/* OPTIONAL IN */ RelativePath string,
	/* OPTIONAL IN */ Location string) (err error) {
	_, err = instance.InvokeMethodWithReturn("AddCollectionElement", SectionPath, CollectionElementName, ElementParameters, AddElementName, RelativePath, Location)
	if err != nil {
		return
	}
	return

}

//

// <param name="CollectionElementName" type="string "></param>
// <param name="ElementParameters" type="NameValueConfigurationElement []"></param>
// <param name="Location" type="string "></param>
// <param name="RelativePath" type="string "></param>
// <param name="SectionPath" type="string "></param>
func (instance *ConfiguredObject) RemoveCollectionElement( /* IN */ SectionPath string,
	/* IN */ CollectionElementName string,
	/* IN */ ElementParameters []NameValueConfigurationElement,
	/* OPTIONAL IN */ RelativePath string,
	/* OPTIONAL IN */ Location string) (err error) {
	_, err = instance.InvokeMethodWithReturn("RemoveCollectionElement", SectionPath, CollectionElementName, ElementParameters, RelativePath, Location)
	if err != nil {
		return
	}
	return

}

//

// <param name="Location" type="string "></param>
// <param name="MetadataName" type="string "></param>
// <param name="RelativePath" type="string "></param>
// <param name="SectionPath" type="string "></param>
// <param name="TargetName" type="string "></param>

// <param name="ReturnValue" type="string "></param>
func (instance *ConfiguredObject) GetRawMetadata( /* IN */ SectionPath string,
	/* IN */ TargetName string,
	/* IN */ MetadataName string,
	/* OPTIONAL IN */ RelativePath string,
	/* OPTIONAL IN */ Location string) (result string, err error) {
	retVal, err := instance.InvokeMethodWithReturn("GetRawMetadata", SectionPath, TargetName, MetadataName, RelativePath, Location)
	if err != nil {
		return
	}
	result = string(retVal)
	return

}

//

// <param name="Location" type="string "></param>
// <param name="MetadataName" type="string "></param>
// <param name="MetadataValue" type="string "></param>
// <param name="RelativePath" type="string "></param>
// <param name="SectionPath" type="string "></param>
// <param name="TargetName" type="string "></param>

// <param name="ReturnValue" type="string "></param>
func (instance *ConfiguredObject) SetRawMetadata( /* IN */ SectionPath string,
	/* IN */ TargetName string,
	/* IN */ MetadataName string,
	/* IN */ MetadataValue string,
	/* OPTIONAL IN */ RelativePath string,
	/* OPTIONAL IN */ Location string) (result string, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetRawMetadata", SectionPath, TargetName, MetadataName, MetadataValue, RelativePath, Location)
	if err != nil {
		return
	}
	result = string(retVal)
	return

}
