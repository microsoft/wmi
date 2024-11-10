// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package internal

import (
	"encoding/xml"
	"github.com/microsoft/wmi/pkg/errors"
)

type Instance struct {
	XMLName          xml.Name         `xml:"INSTANCE"`
	ClassName        string           `xml:"CLASSNAME,attr"`
	Properties       []Property       `xml:"PROPERTY"`
	PropertiesObject []PropertyObject `xml:"PROPERTY.OBJECT"`
	PropertiesArray  []PropertyArray  `xml:"PROPERTY.ARRAY"`
}

type Property struct {
	XMLName     xml.Name `xml:"PROPERTY"`
	Name        string   `xml:"NAME,attr"`
	ClassOrigin string   `xml:"CLASSORIGIN,attr"`
	Type        string   `xml:"TYPE,attr"`
	Value       string   `xml:"VALUE"`
}

type PropertyObject struct {
	XMLName        xml.Name    `xml:"PROPERTY.OBJECT"`
	Name           string      `xml:"NAME,attr"`
	ClassOrigin    string      `xml:"CLASSORIGIN,attr"`
	ReferenceClass string      `xml:"REFERENCECLASS,attr"`
	Value          ValueObject `xml:"VALUE.OBJECT"`
}

type PropertyArray struct {
	XMLName        xml.Name `xml:"PROPERTY.ARRAY"`
	Name           string   `xml:"NAME,attr"`
	ClassOrigin    string   `xml:"CLASSORIGIN,attr"`
	ReferenceClass string   `xml:"REFERENCECLASS,attr"`
	Value          Values   `xml:"VALUE.ARRAY"`
}

type Values struct {
	Value []string `xml:"VALUE"`
}

type ValueObject struct {
	XMLName  xml.Name `xml:"VALUE.OBJECT"`
	Name     string   `xml:"NAME,attr"`
	Instance *Instance
}

// GetVMID This function is a hack to obtain VM ID.
// There could be a better pure WMI way to obtain the ID which requires more investigation
// We are currently using the XML from the embedded instance to obtain the VM ID
func GetVMID(xmlContent string) (vmID string, err error) {
	var clusterResource Instance
	err = xml.Unmarshal([]byte(xmlContent), &clusterResource)
	if err != nil {
		return "", err
	}

	for _, propertyObject := range clusterResource.PropertiesObject {
		if propertyObject.Name == "PrivateProperties" {
			for _, property := range propertyObject.Value.Instance.Properties {
				if property.Name == "VmID" {
					return property.Value, nil
				}
			}
		}
	}
	return "", errors.Wrapf(errors.NotFound, "Missing VMID")
}
