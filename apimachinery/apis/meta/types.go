// @Time    : 2022/5/20 4:09 下午
// @Author  : HuYuan
// @File    : image.go
// @Email   : huyuan@virtaitech.com

package meta

type VirtualMachineLifeCycleState string

const (
	Pending       VirtualMachineLifeCycleState = "Pending"
	Scheduler     VirtualMachineLifeCycleState = "Scheduler"
	Create        VirtualMachineLifeCycleState = "Create"
	Running       VirtualMachineLifeCycleState = "Running"
	Crash         VirtualMachineLifeCycleState = "Crash"
	Restart       VirtualMachineLifeCycleState = "Restart"
	Shutdown      VirtualMachineLifeCycleState = "Shutdown"
	Suspend       VirtualMachineLifeCycleState = "Suspend"
	Unschedulable VirtualMachineLifeCycleState = "Unschedulable"
	StaticMigrate VirtualMachineLifeCycleState = "StaticMigrate"
	LiveMigrate   VirtualMachineLifeCycleState = "LiveMigrate"
	Evicted       VirtualMachineLifeCycleState = "Evicted"
	Unknown       VirtualMachineLifeCycleState = "Unknown"
)

type KindType string

const (
	Instance KindType = "instance"
	Image    KindType = "image"
	Flavor   KindType = "flavor"
)
