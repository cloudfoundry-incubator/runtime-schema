// This file was generated by counterfeiter
package fake_bbs

import (
	"sync"
	"time"

	"github.com/cloudfoundry-incubator/runtime-schema/bbs"
	"github.com/cloudfoundry-incubator/runtime-schema/bbs/shared"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
	"github.com/pivotal-golang/lager"
	"github.com/tedsuo/ifrit"
)

type FakeRepBBS struct {
	NewCellPresenceStub        func(cellPresence models.CellPresence, retryInterval time.Duration) ifrit.Runner
	newCellPresenceMutex       sync.RWMutex
	newCellPresenceArgsForCall []struct {
		cellPresence  models.CellPresence
		retryInterval time.Duration
	}
	newCellPresenceReturns struct {
		result1 ifrit.Runner
	}
	StartTaskStub        func(logger lager.Logger, taskGuid string, cellID string) (bool, error)
	startTaskMutex       sync.RWMutex
	startTaskArgsForCall []struct {
		logger   lager.Logger
		taskGuid string
		cellID   string
	}
	startTaskReturns struct {
		result1 bool
		result2 error
	}
	TaskByGuidStub        func(logger lager.Logger, taskGuid string) (models.Task, error)
	taskByGuidMutex       sync.RWMutex
	taskByGuidArgsForCall []struct {
		logger   lager.Logger
		taskGuid string
	}
	taskByGuidReturns struct {
		result1 models.Task
		result2 error
	}
	TasksByCellIDStub        func(logger lager.Logger, cellID string) ([]models.Task, error)
	tasksByCellIDMutex       sync.RWMutex
	tasksByCellIDArgsForCall []struct {
		logger lager.Logger
		cellID string
	}
	tasksByCellIDReturns struct {
		result1 []models.Task
		result2 error
	}
	FailTaskStub        func(logger lager.Logger, taskGuid string, failureReason string) error
	failTaskMutex       sync.RWMutex
	failTaskArgsForCall []struct {
		logger        lager.Logger
		taskGuid      string
		failureReason string
	}
	failTaskReturns struct {
		result1 error
	}
	CompleteTaskStub        func(logger lager.Logger, taskGuid string, cellID string, failed bool, failureReason string, result string) error
	completeTaskMutex       sync.RWMutex
	completeTaskArgsForCall []struct {
		logger        lager.Logger
		taskGuid      string
		cellID        string
		failed        bool
		failureReason string
		result        string
	}
	completeTaskReturns struct {
		result1 error
	}
	StartActualLRPStub        func(lager.Logger, models.ActualLRPKey, models.ActualLRPInstanceKey, models.ActualLRPNetInfo) error
	startActualLRPMutex       sync.RWMutex
	startActualLRPArgsForCall []struct {
		arg1 lager.Logger
		arg2 models.ActualLRPKey
		arg3 models.ActualLRPInstanceKey
		arg4 models.ActualLRPNetInfo
	}
	startActualLRPReturns struct {
		result1 error
	}
	CrashActualLRPStub        func(lager.Logger, models.ActualLRPKey, models.ActualLRPInstanceKey, string) error
	crashActualLRPMutex       sync.RWMutex
	crashActualLRPArgsForCall []struct {
		arg1 lager.Logger
		arg2 models.ActualLRPKey
		arg3 models.ActualLRPInstanceKey
		arg4 string
	}
	crashActualLRPReturns struct {
		result1 error
	}
	RemoveActualLRPStub        func(lager.Logger, models.ActualLRPKey, models.ActualLRPInstanceKey) error
	removeActualLRPMutex       sync.RWMutex
	removeActualLRPArgsForCall []struct {
		arg1 lager.Logger
		arg2 models.ActualLRPKey
		arg3 models.ActualLRPInstanceKey
	}
	removeActualLRPReturns struct {
		result1 error
	}
	EvacuateClaimedActualLRPStub        func(lager.Logger, models.ActualLRPKey, models.ActualLRPInstanceKey) (shared.ContainerRetainment, error)
	evacuateClaimedActualLRPMutex       sync.RWMutex
	evacuateClaimedActualLRPArgsForCall []struct {
		arg1 lager.Logger
		arg2 models.ActualLRPKey
		arg3 models.ActualLRPInstanceKey
	}
	evacuateClaimedActualLRPReturns struct {
		result1 shared.ContainerRetainment
		result2 error
	}
	EvacuateRunningActualLRPStub        func(lager.Logger, models.ActualLRPKey, models.ActualLRPInstanceKey, models.ActualLRPNetInfo, uint64) (shared.ContainerRetainment, error)
	evacuateRunningActualLRPMutex       sync.RWMutex
	evacuateRunningActualLRPArgsForCall []struct {
		arg1 lager.Logger
		arg2 models.ActualLRPKey
		arg3 models.ActualLRPInstanceKey
		arg4 models.ActualLRPNetInfo
		arg5 uint64
	}
	evacuateRunningActualLRPReturns struct {
		result1 shared.ContainerRetainment
		result2 error
	}
	EvacuateStoppedActualLRPStub        func(lager.Logger, models.ActualLRPKey, models.ActualLRPInstanceKey) (shared.ContainerRetainment, error)
	evacuateStoppedActualLRPMutex       sync.RWMutex
	evacuateStoppedActualLRPArgsForCall []struct {
		arg1 lager.Logger
		arg2 models.ActualLRPKey
		arg3 models.ActualLRPInstanceKey
	}
	evacuateStoppedActualLRPReturns struct {
		result1 shared.ContainerRetainment
		result2 error
	}
	EvacuateCrashedActualLRPStub        func(lager.Logger, models.ActualLRPKey, models.ActualLRPInstanceKey, string) (shared.ContainerRetainment, error)
	evacuateCrashedActualLRPMutex       sync.RWMutex
	evacuateCrashedActualLRPArgsForCall []struct {
		arg1 lager.Logger
		arg2 models.ActualLRPKey
		arg3 models.ActualLRPInstanceKey
		arg4 string
	}
	evacuateCrashedActualLRPReturns struct {
		result1 shared.ContainerRetainment
		result2 error
	}
	RemoveEvacuatingActualLRPStub        func(lager.Logger, models.ActualLRPKey, models.ActualLRPInstanceKey) error
	removeEvacuatingActualLRPMutex       sync.RWMutex
	removeEvacuatingActualLRPArgsForCall []struct {
		arg1 lager.Logger
		arg2 models.ActualLRPKey
		arg3 models.ActualLRPInstanceKey
	}
	removeEvacuatingActualLRPReturns struct {
		result1 error
	}
}

func (fake *FakeRepBBS) NewCellPresence(cellPresence models.CellPresence, retryInterval time.Duration) ifrit.Runner {
	fake.newCellPresenceMutex.Lock()
	fake.newCellPresenceArgsForCall = append(fake.newCellPresenceArgsForCall, struct {
		cellPresence  models.CellPresence
		retryInterval time.Duration
	}{cellPresence, retryInterval})
	fake.newCellPresenceMutex.Unlock()
	if fake.NewCellPresenceStub != nil {
		return fake.NewCellPresenceStub(cellPresence, retryInterval)
	} else {
		return fake.newCellPresenceReturns.result1
	}
}

func (fake *FakeRepBBS) NewCellPresenceCallCount() int {
	fake.newCellPresenceMutex.RLock()
	defer fake.newCellPresenceMutex.RUnlock()
	return len(fake.newCellPresenceArgsForCall)
}

func (fake *FakeRepBBS) NewCellPresenceArgsForCall(i int) (models.CellPresence, time.Duration) {
	fake.newCellPresenceMutex.RLock()
	defer fake.newCellPresenceMutex.RUnlock()
	return fake.newCellPresenceArgsForCall[i].cellPresence, fake.newCellPresenceArgsForCall[i].retryInterval
}

func (fake *FakeRepBBS) NewCellPresenceReturns(result1 ifrit.Runner) {
	fake.NewCellPresenceStub = nil
	fake.newCellPresenceReturns = struct {
		result1 ifrit.Runner
	}{result1}
}

func (fake *FakeRepBBS) StartTask(logger lager.Logger, taskGuid string, cellID string) (bool, error) {
	fake.startTaskMutex.Lock()
	fake.startTaskArgsForCall = append(fake.startTaskArgsForCall, struct {
		logger   lager.Logger
		taskGuid string
		cellID   string
	}{logger, taskGuid, cellID})
	fake.startTaskMutex.Unlock()
	if fake.StartTaskStub != nil {
		return fake.StartTaskStub(logger, taskGuid, cellID)
	} else {
		return fake.startTaskReturns.result1, fake.startTaskReturns.result2
	}
}

func (fake *FakeRepBBS) StartTaskCallCount() int {
	fake.startTaskMutex.RLock()
	defer fake.startTaskMutex.RUnlock()
	return len(fake.startTaskArgsForCall)
}

func (fake *FakeRepBBS) StartTaskArgsForCall(i int) (lager.Logger, string, string) {
	fake.startTaskMutex.RLock()
	defer fake.startTaskMutex.RUnlock()
	return fake.startTaskArgsForCall[i].logger, fake.startTaskArgsForCall[i].taskGuid, fake.startTaskArgsForCall[i].cellID
}

func (fake *FakeRepBBS) StartTaskReturns(result1 bool, result2 error) {
	fake.StartTaskStub = nil
	fake.startTaskReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeRepBBS) TaskByGuid(logger lager.Logger, taskGuid string) (models.Task, error) {
	fake.taskByGuidMutex.Lock()
	fake.taskByGuidArgsForCall = append(fake.taskByGuidArgsForCall, struct {
		logger   lager.Logger
		taskGuid string
	}{logger, taskGuid})
	fake.taskByGuidMutex.Unlock()
	if fake.TaskByGuidStub != nil {
		return fake.TaskByGuidStub(logger, taskGuid)
	} else {
		return fake.taskByGuidReturns.result1, fake.taskByGuidReturns.result2
	}
}

func (fake *FakeRepBBS) TaskByGuidCallCount() int {
	fake.taskByGuidMutex.RLock()
	defer fake.taskByGuidMutex.RUnlock()
	return len(fake.taskByGuidArgsForCall)
}

func (fake *FakeRepBBS) TaskByGuidArgsForCall(i int) (lager.Logger, string) {
	fake.taskByGuidMutex.RLock()
	defer fake.taskByGuidMutex.RUnlock()
	return fake.taskByGuidArgsForCall[i].logger, fake.taskByGuidArgsForCall[i].taskGuid
}

func (fake *FakeRepBBS) TaskByGuidReturns(result1 models.Task, result2 error) {
	fake.TaskByGuidStub = nil
	fake.taskByGuidReturns = struct {
		result1 models.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeRepBBS) TasksByCellID(logger lager.Logger, cellID string) ([]models.Task, error) {
	fake.tasksByCellIDMutex.Lock()
	fake.tasksByCellIDArgsForCall = append(fake.tasksByCellIDArgsForCall, struct {
		logger lager.Logger
		cellID string
	}{logger, cellID})
	fake.tasksByCellIDMutex.Unlock()
	if fake.TasksByCellIDStub != nil {
		return fake.TasksByCellIDStub(logger, cellID)
	} else {
		return fake.tasksByCellIDReturns.result1, fake.tasksByCellIDReturns.result2
	}
}

func (fake *FakeRepBBS) TasksByCellIDCallCount() int {
	fake.tasksByCellIDMutex.RLock()
	defer fake.tasksByCellIDMutex.RUnlock()
	return len(fake.tasksByCellIDArgsForCall)
}

func (fake *FakeRepBBS) TasksByCellIDArgsForCall(i int) (lager.Logger, string) {
	fake.tasksByCellIDMutex.RLock()
	defer fake.tasksByCellIDMutex.RUnlock()
	return fake.tasksByCellIDArgsForCall[i].logger, fake.tasksByCellIDArgsForCall[i].cellID
}

func (fake *FakeRepBBS) TasksByCellIDReturns(result1 []models.Task, result2 error) {
	fake.TasksByCellIDStub = nil
	fake.tasksByCellIDReturns = struct {
		result1 []models.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeRepBBS) FailTask(logger lager.Logger, taskGuid string, failureReason string) error {
	fake.failTaskMutex.Lock()
	fake.failTaskArgsForCall = append(fake.failTaskArgsForCall, struct {
		logger        lager.Logger
		taskGuid      string
		failureReason string
	}{logger, taskGuid, failureReason})
	fake.failTaskMutex.Unlock()
	if fake.FailTaskStub != nil {
		return fake.FailTaskStub(logger, taskGuid, failureReason)
	} else {
		return fake.failTaskReturns.result1
	}
}

func (fake *FakeRepBBS) FailTaskCallCount() int {
	fake.failTaskMutex.RLock()
	defer fake.failTaskMutex.RUnlock()
	return len(fake.failTaskArgsForCall)
}

func (fake *FakeRepBBS) FailTaskArgsForCall(i int) (lager.Logger, string, string) {
	fake.failTaskMutex.RLock()
	defer fake.failTaskMutex.RUnlock()
	return fake.failTaskArgsForCall[i].logger, fake.failTaskArgsForCall[i].taskGuid, fake.failTaskArgsForCall[i].failureReason
}

func (fake *FakeRepBBS) FailTaskReturns(result1 error) {
	fake.FailTaskStub = nil
	fake.failTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepBBS) CompleteTask(logger lager.Logger, taskGuid string, cellID string, failed bool, failureReason string, result string) error {
	fake.completeTaskMutex.Lock()
	fake.completeTaskArgsForCall = append(fake.completeTaskArgsForCall, struct {
		logger        lager.Logger
		taskGuid      string
		cellID        string
		failed        bool
		failureReason string
		result        string
	}{logger, taskGuid, cellID, failed, failureReason, result})
	fake.completeTaskMutex.Unlock()
	if fake.CompleteTaskStub != nil {
		return fake.CompleteTaskStub(logger, taskGuid, cellID, failed, failureReason, result)
	} else {
		return fake.completeTaskReturns.result1
	}
}

func (fake *FakeRepBBS) CompleteTaskCallCount() int {
	fake.completeTaskMutex.RLock()
	defer fake.completeTaskMutex.RUnlock()
	return len(fake.completeTaskArgsForCall)
}

func (fake *FakeRepBBS) CompleteTaskArgsForCall(i int) (lager.Logger, string, string, bool, string, string) {
	fake.completeTaskMutex.RLock()
	defer fake.completeTaskMutex.RUnlock()
	return fake.completeTaskArgsForCall[i].logger, fake.completeTaskArgsForCall[i].taskGuid, fake.completeTaskArgsForCall[i].cellID, fake.completeTaskArgsForCall[i].failed, fake.completeTaskArgsForCall[i].failureReason, fake.completeTaskArgsForCall[i].result
}

func (fake *FakeRepBBS) CompleteTaskReturns(result1 error) {
	fake.CompleteTaskStub = nil
	fake.completeTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepBBS) StartActualLRP(arg1 lager.Logger, arg2 models.ActualLRPKey, arg3 models.ActualLRPInstanceKey, arg4 models.ActualLRPNetInfo) error {
	fake.startActualLRPMutex.Lock()
	fake.startActualLRPArgsForCall = append(fake.startActualLRPArgsForCall, struct {
		arg1 lager.Logger
		arg2 models.ActualLRPKey
		arg3 models.ActualLRPInstanceKey
		arg4 models.ActualLRPNetInfo
	}{arg1, arg2, arg3, arg4})
	fake.startActualLRPMutex.Unlock()
	if fake.StartActualLRPStub != nil {
		return fake.StartActualLRPStub(arg1, arg2, arg3, arg4)
	} else {
		return fake.startActualLRPReturns.result1
	}
}

func (fake *FakeRepBBS) StartActualLRPCallCount() int {
	fake.startActualLRPMutex.RLock()
	defer fake.startActualLRPMutex.RUnlock()
	return len(fake.startActualLRPArgsForCall)
}

func (fake *FakeRepBBS) StartActualLRPArgsForCall(i int) (lager.Logger, models.ActualLRPKey, models.ActualLRPInstanceKey, models.ActualLRPNetInfo) {
	fake.startActualLRPMutex.RLock()
	defer fake.startActualLRPMutex.RUnlock()
	return fake.startActualLRPArgsForCall[i].arg1, fake.startActualLRPArgsForCall[i].arg2, fake.startActualLRPArgsForCall[i].arg3, fake.startActualLRPArgsForCall[i].arg4
}

func (fake *FakeRepBBS) StartActualLRPReturns(result1 error) {
	fake.StartActualLRPStub = nil
	fake.startActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepBBS) CrashActualLRP(arg1 lager.Logger, arg2 models.ActualLRPKey, arg3 models.ActualLRPInstanceKey, arg4 string) error {
	fake.crashActualLRPMutex.Lock()
	fake.crashActualLRPArgsForCall = append(fake.crashActualLRPArgsForCall, struct {
		arg1 lager.Logger
		arg2 models.ActualLRPKey
		arg3 models.ActualLRPInstanceKey
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.crashActualLRPMutex.Unlock()
	if fake.CrashActualLRPStub != nil {
		return fake.CrashActualLRPStub(arg1, arg2, arg3, arg4)
	} else {
		return fake.crashActualLRPReturns.result1
	}
}

func (fake *FakeRepBBS) CrashActualLRPCallCount() int {
	fake.crashActualLRPMutex.RLock()
	defer fake.crashActualLRPMutex.RUnlock()
	return len(fake.crashActualLRPArgsForCall)
}

func (fake *FakeRepBBS) CrashActualLRPArgsForCall(i int) (lager.Logger, models.ActualLRPKey, models.ActualLRPInstanceKey, string) {
	fake.crashActualLRPMutex.RLock()
	defer fake.crashActualLRPMutex.RUnlock()
	return fake.crashActualLRPArgsForCall[i].arg1, fake.crashActualLRPArgsForCall[i].arg2, fake.crashActualLRPArgsForCall[i].arg3, fake.crashActualLRPArgsForCall[i].arg4
}

func (fake *FakeRepBBS) CrashActualLRPReturns(result1 error) {
	fake.CrashActualLRPStub = nil
	fake.crashActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepBBS) RemoveActualLRP(arg1 lager.Logger, arg2 models.ActualLRPKey, arg3 models.ActualLRPInstanceKey) error {
	fake.removeActualLRPMutex.Lock()
	fake.removeActualLRPArgsForCall = append(fake.removeActualLRPArgsForCall, struct {
		arg1 lager.Logger
		arg2 models.ActualLRPKey
		arg3 models.ActualLRPInstanceKey
	}{arg1, arg2, arg3})
	fake.removeActualLRPMutex.Unlock()
	if fake.RemoveActualLRPStub != nil {
		return fake.RemoveActualLRPStub(arg1, arg2, arg3)
	} else {
		return fake.removeActualLRPReturns.result1
	}
}

func (fake *FakeRepBBS) RemoveActualLRPCallCount() int {
	fake.removeActualLRPMutex.RLock()
	defer fake.removeActualLRPMutex.RUnlock()
	return len(fake.removeActualLRPArgsForCall)
}

func (fake *FakeRepBBS) RemoveActualLRPArgsForCall(i int) (lager.Logger, models.ActualLRPKey, models.ActualLRPInstanceKey) {
	fake.removeActualLRPMutex.RLock()
	defer fake.removeActualLRPMutex.RUnlock()
	return fake.removeActualLRPArgsForCall[i].arg1, fake.removeActualLRPArgsForCall[i].arg2, fake.removeActualLRPArgsForCall[i].arg3
}

func (fake *FakeRepBBS) RemoveActualLRPReturns(result1 error) {
	fake.RemoveActualLRPStub = nil
	fake.removeActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepBBS) EvacuateClaimedActualLRP(arg1 lager.Logger, arg2 models.ActualLRPKey, arg3 models.ActualLRPInstanceKey) (shared.ContainerRetainment, error) {
	fake.evacuateClaimedActualLRPMutex.Lock()
	fake.evacuateClaimedActualLRPArgsForCall = append(fake.evacuateClaimedActualLRPArgsForCall, struct {
		arg1 lager.Logger
		arg2 models.ActualLRPKey
		arg3 models.ActualLRPInstanceKey
	}{arg1, arg2, arg3})
	fake.evacuateClaimedActualLRPMutex.Unlock()
	if fake.EvacuateClaimedActualLRPStub != nil {
		return fake.EvacuateClaimedActualLRPStub(arg1, arg2, arg3)
	} else {
		return fake.evacuateClaimedActualLRPReturns.result1, fake.evacuateClaimedActualLRPReturns.result2
	}
}

func (fake *FakeRepBBS) EvacuateClaimedActualLRPCallCount() int {
	fake.evacuateClaimedActualLRPMutex.RLock()
	defer fake.evacuateClaimedActualLRPMutex.RUnlock()
	return len(fake.evacuateClaimedActualLRPArgsForCall)
}

func (fake *FakeRepBBS) EvacuateClaimedActualLRPArgsForCall(i int) (lager.Logger, models.ActualLRPKey, models.ActualLRPInstanceKey) {
	fake.evacuateClaimedActualLRPMutex.RLock()
	defer fake.evacuateClaimedActualLRPMutex.RUnlock()
	return fake.evacuateClaimedActualLRPArgsForCall[i].arg1, fake.evacuateClaimedActualLRPArgsForCall[i].arg2, fake.evacuateClaimedActualLRPArgsForCall[i].arg3
}

func (fake *FakeRepBBS) EvacuateClaimedActualLRPReturns(result1 shared.ContainerRetainment, result2 error) {
	fake.EvacuateClaimedActualLRPStub = nil
	fake.evacuateClaimedActualLRPReturns = struct {
		result1 shared.ContainerRetainment
		result2 error
	}{result1, result2}
}

func (fake *FakeRepBBS) EvacuateRunningActualLRP(arg1 lager.Logger, arg2 models.ActualLRPKey, arg3 models.ActualLRPInstanceKey, arg4 models.ActualLRPNetInfo, arg5 uint64) (shared.ContainerRetainment, error) {
	fake.evacuateRunningActualLRPMutex.Lock()
	fake.evacuateRunningActualLRPArgsForCall = append(fake.evacuateRunningActualLRPArgsForCall, struct {
		arg1 lager.Logger
		arg2 models.ActualLRPKey
		arg3 models.ActualLRPInstanceKey
		arg4 models.ActualLRPNetInfo
		arg5 uint64
	}{arg1, arg2, arg3, arg4, arg5})
	fake.evacuateRunningActualLRPMutex.Unlock()
	if fake.EvacuateRunningActualLRPStub != nil {
		return fake.EvacuateRunningActualLRPStub(arg1, arg2, arg3, arg4, arg5)
	} else {
		return fake.evacuateRunningActualLRPReturns.result1, fake.evacuateRunningActualLRPReturns.result2
	}
}

func (fake *FakeRepBBS) EvacuateRunningActualLRPCallCount() int {
	fake.evacuateRunningActualLRPMutex.RLock()
	defer fake.evacuateRunningActualLRPMutex.RUnlock()
	return len(fake.evacuateRunningActualLRPArgsForCall)
}

func (fake *FakeRepBBS) EvacuateRunningActualLRPArgsForCall(i int) (lager.Logger, models.ActualLRPKey, models.ActualLRPInstanceKey, models.ActualLRPNetInfo, uint64) {
	fake.evacuateRunningActualLRPMutex.RLock()
	defer fake.evacuateRunningActualLRPMutex.RUnlock()
	return fake.evacuateRunningActualLRPArgsForCall[i].arg1, fake.evacuateRunningActualLRPArgsForCall[i].arg2, fake.evacuateRunningActualLRPArgsForCall[i].arg3, fake.evacuateRunningActualLRPArgsForCall[i].arg4, fake.evacuateRunningActualLRPArgsForCall[i].arg5
}

func (fake *FakeRepBBS) EvacuateRunningActualLRPReturns(result1 shared.ContainerRetainment, result2 error) {
	fake.EvacuateRunningActualLRPStub = nil
	fake.evacuateRunningActualLRPReturns = struct {
		result1 shared.ContainerRetainment
		result2 error
	}{result1, result2}
}

func (fake *FakeRepBBS) EvacuateStoppedActualLRP(arg1 lager.Logger, arg2 models.ActualLRPKey, arg3 models.ActualLRPInstanceKey) (shared.ContainerRetainment, error) {
	fake.evacuateStoppedActualLRPMutex.Lock()
	fake.evacuateStoppedActualLRPArgsForCall = append(fake.evacuateStoppedActualLRPArgsForCall, struct {
		arg1 lager.Logger
		arg2 models.ActualLRPKey
		arg3 models.ActualLRPInstanceKey
	}{arg1, arg2, arg3})
	fake.evacuateStoppedActualLRPMutex.Unlock()
	if fake.EvacuateStoppedActualLRPStub != nil {
		return fake.EvacuateStoppedActualLRPStub(arg1, arg2, arg3)
	} else {
		return fake.evacuateStoppedActualLRPReturns.result1, fake.evacuateStoppedActualLRPReturns.result2
	}
}

func (fake *FakeRepBBS) EvacuateStoppedActualLRPCallCount() int {
	fake.evacuateStoppedActualLRPMutex.RLock()
	defer fake.evacuateStoppedActualLRPMutex.RUnlock()
	return len(fake.evacuateStoppedActualLRPArgsForCall)
}

func (fake *FakeRepBBS) EvacuateStoppedActualLRPArgsForCall(i int) (lager.Logger, models.ActualLRPKey, models.ActualLRPInstanceKey) {
	fake.evacuateStoppedActualLRPMutex.RLock()
	defer fake.evacuateStoppedActualLRPMutex.RUnlock()
	return fake.evacuateStoppedActualLRPArgsForCall[i].arg1, fake.evacuateStoppedActualLRPArgsForCall[i].arg2, fake.evacuateStoppedActualLRPArgsForCall[i].arg3
}

func (fake *FakeRepBBS) EvacuateStoppedActualLRPReturns(result1 shared.ContainerRetainment, result2 error) {
	fake.EvacuateStoppedActualLRPStub = nil
	fake.evacuateStoppedActualLRPReturns = struct {
		result1 shared.ContainerRetainment
		result2 error
	}{result1, result2}
}

func (fake *FakeRepBBS) EvacuateCrashedActualLRP(arg1 lager.Logger, arg2 models.ActualLRPKey, arg3 models.ActualLRPInstanceKey, arg4 string) (shared.ContainerRetainment, error) {
	fake.evacuateCrashedActualLRPMutex.Lock()
	fake.evacuateCrashedActualLRPArgsForCall = append(fake.evacuateCrashedActualLRPArgsForCall, struct {
		arg1 lager.Logger
		arg2 models.ActualLRPKey
		arg3 models.ActualLRPInstanceKey
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.evacuateCrashedActualLRPMutex.Unlock()
	if fake.EvacuateCrashedActualLRPStub != nil {
		return fake.EvacuateCrashedActualLRPStub(arg1, arg2, arg3, arg4)
	} else {
		return fake.evacuateCrashedActualLRPReturns.result1, fake.evacuateCrashedActualLRPReturns.result2
	}
}

func (fake *FakeRepBBS) EvacuateCrashedActualLRPCallCount() int {
	fake.evacuateCrashedActualLRPMutex.RLock()
	defer fake.evacuateCrashedActualLRPMutex.RUnlock()
	return len(fake.evacuateCrashedActualLRPArgsForCall)
}

func (fake *FakeRepBBS) EvacuateCrashedActualLRPArgsForCall(i int) (lager.Logger, models.ActualLRPKey, models.ActualLRPInstanceKey, string) {
	fake.evacuateCrashedActualLRPMutex.RLock()
	defer fake.evacuateCrashedActualLRPMutex.RUnlock()
	return fake.evacuateCrashedActualLRPArgsForCall[i].arg1, fake.evacuateCrashedActualLRPArgsForCall[i].arg2, fake.evacuateCrashedActualLRPArgsForCall[i].arg3, fake.evacuateCrashedActualLRPArgsForCall[i].arg4
}

func (fake *FakeRepBBS) EvacuateCrashedActualLRPReturns(result1 shared.ContainerRetainment, result2 error) {
	fake.EvacuateCrashedActualLRPStub = nil
	fake.evacuateCrashedActualLRPReturns = struct {
		result1 shared.ContainerRetainment
		result2 error
	}{result1, result2}
}

func (fake *FakeRepBBS) RemoveEvacuatingActualLRP(arg1 lager.Logger, arg2 models.ActualLRPKey, arg3 models.ActualLRPInstanceKey) error {
	fake.removeEvacuatingActualLRPMutex.Lock()
	fake.removeEvacuatingActualLRPArgsForCall = append(fake.removeEvacuatingActualLRPArgsForCall, struct {
		arg1 lager.Logger
		arg2 models.ActualLRPKey
		arg3 models.ActualLRPInstanceKey
	}{arg1, arg2, arg3})
	fake.removeEvacuatingActualLRPMutex.Unlock()
	if fake.RemoveEvacuatingActualLRPStub != nil {
		return fake.RemoveEvacuatingActualLRPStub(arg1, arg2, arg3)
	} else {
		return fake.removeEvacuatingActualLRPReturns.result1
	}
}

func (fake *FakeRepBBS) RemoveEvacuatingActualLRPCallCount() int {
	fake.removeEvacuatingActualLRPMutex.RLock()
	defer fake.removeEvacuatingActualLRPMutex.RUnlock()
	return len(fake.removeEvacuatingActualLRPArgsForCall)
}

func (fake *FakeRepBBS) RemoveEvacuatingActualLRPArgsForCall(i int) (lager.Logger, models.ActualLRPKey, models.ActualLRPInstanceKey) {
	fake.removeEvacuatingActualLRPMutex.RLock()
	defer fake.removeEvacuatingActualLRPMutex.RUnlock()
	return fake.removeEvacuatingActualLRPArgsForCall[i].arg1, fake.removeEvacuatingActualLRPArgsForCall[i].arg2, fake.removeEvacuatingActualLRPArgsForCall[i].arg3
}

func (fake *FakeRepBBS) RemoveEvacuatingActualLRPReturns(result1 error) {
	fake.RemoveEvacuatingActualLRPStub = nil
	fake.removeEvacuatingActualLRPReturns = struct {
		result1 error
	}{result1}
}

var _ bbs.RepBBS = new(FakeRepBBS)
