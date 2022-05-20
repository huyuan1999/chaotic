// @Time    : 2022/5/20 4:14 下午
// @Author  : HuYuan
// @File    : types.go
// @Email   : huyuan@virtaitech.com

package scheme

import (
	"github.com/huyuan1999/chaotic/apimachinery/apis/meta"
	"time"
)

type OwnerReference struct {
	Name string
	UID  string
}

type ObjectMeta struct {
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	UID               string            `json:"uid"`
	CreationTimestamp string            `json:"creationTimestamp"`
	ResourceVersion   string            `json:"resourceVersion"`
	DeletionTimestamp time.Time         `json:"deletionTimestamp"`
	Annotations       map[string]string `json:"annotations"`
	OwnerReferences   []OwnerReference  `json:"ownerReferences"`
	Finalizers        []string          `json:"finalizers"`
}

type ImageSpec struct {
	Name            string `json:"name"`
	Path            string `json:"path"`
	Checksum        string `json:"checksum"`
	UpdateTimestamp string `json:"UpdateTimestamp"`
	CreateTimestamp string `json:"creationTimestamp"`
	DiskFormat      string `json:"diskFormat"`
	VirtualSize     uint64 `json:"virtualSize"`
	ActualSize      uint64 `json:"actualSize"`
}

type Image struct {
	Kind     meta.KindType `json:"kind"`
	Metadata ObjectMeta    `json:"metadata"`
	Spec     ImageSpec     `json:"spec"`
}

type FlavorSpec struct {
	Name   string `json:"name"`
	VCpus  uint   `json:"vcpus"`
	Memory uint64 `json:"memory"`
}

type Flavor struct {
	Kind     meta.KindType `json:"kind"`
	Metadata ObjectMeta    `json:"metadata"`
	Spec     FlavorSpec    `json:"spec"`
}

type VirtualMachineSpec struct {
	Name     string `json:"name"`
	UUID     string `json:"uuid"`
	NodeName string `json:"nodeName"`
	Image    Image  `json:"image"`
	Flavor   Flavor `json:"flavor"`
}

type VirtualMachineStatus struct {
	Phase             meta.VirtualMachineLifeCycleState `json:"phase"`
	VirtualMachineIPs []string                          `json:"virtualMachineIPs"`
	StartTime         time.Time                         `json:"startTime"`
	QOSClass          string                            `json:"qosClass"`
}

type VirtualMachine struct {
	Kind     meta.KindType        `json:"kind"`
	Metadata ObjectMeta           `json:"metadata"`
	Spec     VirtualMachineSpec   `json:"spec"`
	Status   VirtualMachineStatus `json:"status"`
}
