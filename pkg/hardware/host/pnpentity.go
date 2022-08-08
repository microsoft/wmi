// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package host

import (
	"log"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/virtualization/core/pnp"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"

	"github.com/microsoft/wmi/server2019/root/cimv2"
)

type PnpEntity struct {
	*cimv2.Win32_PnPEntity
}

// NewPnpEntity
func NewPnpEntity(instance *wmi.WmiInstance) (*PnpEntity, error) {
	wmivm, err := cimv2.NewWin32_PnPEntityEx1(instance)
	if err != nil {
		return nil, err
	}
	return &PnpEntity{wmivm}, nil
}

// GetPnpEntityCollectionByName
func GetPnpEntityCollection(whost *host.WmiHost) (entities pnp.PnpEntityCollection, err error) {
	query := query.NewWmiQuery("Win32_PnPEntity")
	rasdcollection, err := instance.GetWmiInstancesFromHost(whost, string(constant.CimV2), query)
	if err != nil {
		return
	}
	defer rasdcollection.Close()

	entities, err = pnp.NewPnpEntityCollection(rasdcollection)
	return
}

func GetPnpEntityCollectionByName(whost *host.WmiHost, pnpEntityName string) (entities pnp.PnpEntityCollection, err error) {
	pnpEntityCollection, err := GetPnpEntityCollection(whost)
	if err != nil {
		return
	}
	defer pnpEntityCollection.Close()

	for _, pnpEntity := range pnpEntityCollection {
		entityName, err := pnpEntity.GetPropertyName()
		if err != nil {
			log.Printf("[WMI] Error getting property name for pnp entity [%s] - Error details [%+v]\n", pnpEntity, err)
			continue
		}
		if entityName == pnpEntityName {
			// Found the match
			log.Printf("[WMI] Found pnp entity [%s] for name [%s]", pnpEntity, pnpEntityName)
			clonedEntity, err := pnpEntity.CloneEx1()
			if err != nil {
				log.Printf("[WMI] Error cloning the pnp entity [%s] - Error details [%+v]\n", pnpEntity, err)
				continue
			}
			entities = append(entities, clonedEntity)
		}
	}
	return
}
