// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package snapshot

import (
	"log"
	"sync"
	"time"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/service"
	"github.com/microsoft/wmi/pkg/virtualization/core/virtualsystem"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

var (
	serviceStoreMap map[string]*VirtualSystemSnapshotService
	mux             sync.Mutex
)

func init() {
	serviceStoreMap = map[string]*VirtualSystemSnapshotService{}
}

type VirtualSystemSnapshotService struct {
	*v2.Msvm_VirtualSystemSnapshotService
}

func GetVirtualSystemSnapshotService(whost *host.WmiHost) (snapshotSvc *VirtualSystemSnapshotService, err error) {
	if val, ok := serviceStoreMap[whost.HostName]; ok {
		snapshotSvc = val
		return
	}

	snapshotSvc, err = getService(whost)
	if err != nil {
		return
	}

	mux.Lock()
	defer mux.Unlock()
	if val, ok := serviceStoreMap[whost.HostName]; ok {
		snapshotSvc.Close()
		snapshotSvc = val
		return
	}

	serviceStoreMap[whost.HostName] = snapshotSvc
	return
}

func getService(whost *host.WmiHost) (snapshotSvc *VirtualSystemSnapshotService, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("Msvm_VirtualSystemSnapshotService")
	snapshotSvcWmi, err := v2.NewMsvm_VirtualSystemSnapshotServiceEx6(whost.HostName, string(constant.Virtualization), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}

	snapshotSvc = &VirtualSystemSnapshotService{snapshotSvcWmi}
	return
}

func (snapshotSvc *VirtualSystemSnapshotService) CreateSnapshot(vm *virtualsystem.VirtualMachine, name string, timeoutSeconds int16) (err error) {
	method, err := snapshotSvc.GetWmiMethod("CreateSnapshot")
	if err != nil {
		return
	}
	defer method.Close()

	inparams := wmi.WmiMethodParamCollection{}
	inparams = append(inparams, wmi.NewWmiMethodParam("AffectedSystem", vm.InstancePath()))
	// Full snapshot
	inparams = append(inparams, wmi.NewWmiMethodParam("SnapshotType", 2))
	outparams := wmi.WmiMethodParamCollection{wmi.NewWmiMethodParam("Job", nil)}
	var wmijob *wmi.WmiJob

	for {
		result, err1 := method.Execute(inparams, outparams)
		if err1 != nil {
			err = err1
			return
		}

		returnVal := result.ReturnValue
		if returnVal != 0 && returnVal != 4096 {
			// Virtual System is in Invalid State, try to retry
			if returnVal == 32775 {
				log.Printf("[WMI] Method [%s] failed with [%d]. Retrying ...", method.Name, returnVal)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			err = errors.Wrapf(errors.Failed, "Method failed with [%d]", result.ReturnValue)
			return
		}

		if result.ReturnValue == 0 {
			return
		}

		job, ok := result.OutMethodParams["Job"]
		if !ok || job.Value == nil {
			err = errors.Wrapf(errors.NotFound, "Job")
			return
		}

		wmijob, err = instance.GetWmiJob(snapshotSvc.GetWmiHost(), string(constant.Virtualization), job.Value.(string))
		if err != nil {
			return
		}

		defer job.Close()

		err = wmijob.WaitForJobCompletion(result.ReturnValue, timeoutSeconds)
		if err != nil {
			return
		}

		break
	}

	/* TODO: rename fail, remove the snapshot
	defer func() {
		if err != nil {
		}
	}()
	*/

	if name != "" {
		// Rename snapshot
		inst, err1 := wmijob.GetRelated("Msvm_VirtualSystemSettingData")
		if err1 != nil {
			err = err1
			return
		}

		virtualSystemSettingData, err1 := virtualsystem.NewVirtualSystemSettingData(inst)
		if err1 != nil {
			err = err1
			return
		}

		virtualSystemSettingData.SetPropertyElementName(name)
		vmms, err1 := service.GetVirtualSystemManagementService(snapshotSvc.GetWmiHost())
		if err != nil {
			err = err1
			return
		}

		err = vmms.ModifyVirtualSystemSettings(virtualSystemSettingData, -1)
		if err != nil {
			log.Printf("[WMI] Method [%s] failed to rename snapshot to %s", method.Name, name)
			return
		}
	}

	return
}
