// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package disk

import (
	"encoding/xml"
	"log"
	"strconv"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type VirtualHardDiskType uint16

const (
	VirtualHardDiskType_NONE   = 0
	VirtualHardDiskType_LEGACY = 1
	VirtualHardDiskType_FLAT   = 2
	VirtualHardDiskType_SPARSE = 3
)

type VirtualHardDiskFormat uint16

const (
	VirtualHardDiskFormat_NONE = 0
	VirtualHardDiskFormat_ISO  = 1
	VirtualHardDiskFormat_1    = 2
	VirtualHardDiskFormat_2    = 3
)

type INSTANCE struct {
	XMLName   xml.Name `xml:"INSTANCE"`
	Text      string   `xml:",chardata"`
	CLASSNAME string   `xml:"CLASSNAME,attr"`
	PROPERTY  []struct {
		Text       string `xml:",chardata"`
		NAME       string `xml:"NAME,attr"`
		TYPE       string `xml:"TYPE,attr"`
		PROPAGATED string `xml:"PROPAGATED,attr"`
		VALUE      string `xml:"VALUE"`
	} `xml:"PROPERTY"`
}

type VirtualHardDiskSettingData struct {
	*v2.Msvm_VirtualHardDiskSettingData
}

// NewVirtualHardDiskSettingData
func NewVirtualHardDiskSettingData(instance *wmi.WmiInstance) (*VirtualHardDiskSettingData, error) {
	wmivm, err := v2.NewMsvm_VirtualHardDiskSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &VirtualHardDiskSettingData{wmivm}, nil
}

func GetDefaultVirtualHardDiskSettingData(whost *host.WmiHost) (*VirtualHardDiskSettingData, error) {
	inst, err := instance.CreateWmiInstance(whost, string(constant.Virtualization), "Msvm_VirtualHardDiskSettingData")
	if err != nil {
		return nil, err
	}
	return NewVirtualHardDiskSettingData(inst)
}

func GetVirtualHardDiskSettingData(whost *host.WmiHost, path string,
	logicalSectorSize, physicalSectorSize, blockSize uint32,
	diskSize uint64, dynamic bool, diskFileFormat VirtualHardDiskFormat) (vhdsetting *VirtualHardDiskSettingData, err error) {

	vhdsetting, err = GetDefaultVirtualHardDiskSettingData(whost)
	if err != nil {
		return nil, err
	}

	vhdsetting.SetPropertyPath(path)
	vhdsetting.SetPropertyFormat(uint16(diskFileFormat))
	vhdsetting.SetPropertyBlockSize(blockSize)
	vhdsetting.SetPropertyLogicalSectorSize(logicalSectorSize)
	vhdsetting.SetPropertyPhysicalSectorSize(physicalSectorSize)
	vhdsetting.SetPropertyMaxInternalSize(diskSize)
	if dynamic {
		vhdsetting.SetPropertyType(uint16(VirtualHardDiskType_SPARSE))
	} else {
		// Fixed
		vhdsetting.SetPropertyType(uint16(VirtualHardDiskType_FLAT))
	}

	return
}

func GetVirtualHardDiskSettingDataFromXml(whost *host.WmiHost, xmlInstance string) (size uint64,
	blockSize uint32,
	lSectorSize uint32,
	pSectorSize uint32,
	format uint16,
	err error) {
	/*vhdTemp, err := GetDefaultVirtualHardDiskSettingData(whost)
	if err != nil {
		return nil, err
	}
	defer vhdTemp.Close()

	// clone vhdsetting and populated the cloned instance
	vhdWmi, err := NewVirtualHardDiskSettingData(vhdTemp.WmiInstance)
	if err != nil {
		return nil, err
	}

	vhdWmiClone, err := vhdWmi.Clone()
	if err != nil {
		return nil, err
	}

	vhdsetting, err = NewVirtualHardDiskSettingData(vhdWmiClone)
	if err != nil {
		return nil, err
	}*/

	log.Printf("Decoding WMI response [%s]\n", xmlInstance)
	var tempvar uint64
	var diskData INSTANCE
	err = xml.Unmarshal([]byte(xmlInstance), &diskData)
	if err != nil {
		return
	}

	for _, property := range diskData.PROPERTY {
		if property.NAME == "MaxInternalSize" {
			size, err = strconv.ParseUint(property.VALUE, 10, 64)
			if err != nil {
				return
			}
		} else if property.NAME == "BlockSize" {
			tempvar, err = strconv.ParseUint(property.VALUE, 10, 32)
			if err != nil {
				return
			}
			blockSize = uint32(tempvar)
		} else if property.NAME == "LogicalSectorSize" {
			tempvar, err = strconv.ParseUint(property.VALUE, 10, 32)
			if err != nil {
				return
			}
			lSectorSize = uint32(tempvar)
		} else if property.NAME == "PhysicalSectorSize" {
			tempvar, err = strconv.ParseUint(property.VALUE, 10, 32)
			if err != nil {
				return
			}
			pSectorSize = uint32(tempvar)
		} else if property.NAME == "Format" {
			tempvar, err = strconv.ParseUint(property.VALUE, 10, 16)
			if err != nil {
				return
			}
			format = uint16(tempvar)
		}
	}

	return
}
