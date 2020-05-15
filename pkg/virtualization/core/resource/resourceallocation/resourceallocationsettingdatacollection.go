// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package resourceallocation

type ResourceAllocationSettingDataCollection []*ResourceAllocationSettingData

func (vms *ResourceAllocationSettingDataCollection) Close() (err error) {
	for _, value := range *vms {
		value.Close()
	}
	return nil
}

func (vms *ResourceAllocationSettingDataCollection) String() string {
	return ""
}
func (vms *ResourceAllocationSettingDataCollection) EmbeddedXMLInstances() (xmls []string, err error) {
	for _, value := range *vms {
		xmlvalue, err := value.EmbeddedXMLInstance()
		if err != nil {
		}
		xmls = append(xmls, xmlvalue)
	}

	return
}
