// This file was generated by counterfeiter
package fake_bbs

import (
	"sync"
	"time"

	"github.com/cloudfoundry-incubator/runtime-schema/bbs"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
	"github.com/pivotal-golang/lager"
	"github.com/tedsuo/ifrit"
)

type FakeReceptorBBS struct {
	DesireTaskStub        func(lager.Logger, models.Task) error
	desireTaskMutex       sync.RWMutex
	desireTaskArgsForCall []struct {
		arg1 lager.Logger
		arg2 models.Task
	}
	desireTaskReturns struct {
		result1 error
	}
	TasksStub        func(logger lager.Logger) ([]models.Task, error)
	tasksMutex       sync.RWMutex
	tasksArgsForCall []struct {
		logger lager.Logger
	}
	tasksReturns struct {
		result1 []models.Task
		result2 error
	}
	TasksByDomainStub        func(logger lager.Logger, domain string) ([]models.Task, error)
	tasksByDomainMutex       sync.RWMutex
	tasksByDomainArgsForCall []struct {
		logger lager.Logger
		domain string
	}
	tasksByDomainReturns struct {
		result1 []models.Task
		result2 error
	}
	TaskByGuidStub        func(taskGuid string) (models.Task, error)
	taskByGuidMutex       sync.RWMutex
	taskByGuidArgsForCall []struct {
		taskGuid string
	}
	taskByGuidReturns struct {
		result1 models.Task
		result2 error
	}
	ResolvingTaskStub        func(logger lager.Logger, taskGuid string) error
	resolvingTaskMutex       sync.RWMutex
	resolvingTaskArgsForCall []struct {
		logger   lager.Logger
		taskGuid string
	}
	resolvingTaskReturns struct {
		result1 error
	}
	ResolveTaskStub        func(logger lager.Logger, taskGuid string) error
	resolveTaskMutex       sync.RWMutex
	resolveTaskArgsForCall []struct {
		logger   lager.Logger
		taskGuid string
	}
	resolveTaskReturns struct {
		result1 error
	}
	CancelTaskStub        func(logger lager.Logger, taskGuid string) error
	cancelTaskMutex       sync.RWMutex
	cancelTaskArgsForCall []struct {
		logger   lager.Logger
		taskGuid string
	}
	cancelTaskReturns struct {
		result1 error
	}
	DesireLRPStub        func(lager.Logger, models.DesiredLRP) error
	desireLRPMutex       sync.RWMutex
	desireLRPArgsForCall []struct {
		arg1 lager.Logger
		arg2 models.DesiredLRP
	}
	desireLRPReturns struct {
		result1 error
	}
	UpdateDesiredLRPStub        func(logger lager.Logger, processGuid string, update models.DesiredLRPUpdate) error
	updateDesiredLRPMutex       sync.RWMutex
	updateDesiredLRPArgsForCall []struct {
		logger      lager.Logger
		processGuid string
		update      models.DesiredLRPUpdate
	}
	updateDesiredLRPReturns struct {
		result1 error
	}
	RemoveDesiredLRPByProcessGuidStub        func(logger lager.Logger, processGuid string) error
	removeDesiredLRPByProcessGuidMutex       sync.RWMutex
	removeDesiredLRPByProcessGuidArgsForCall []struct {
		logger      lager.Logger
		processGuid string
	}
	removeDesiredLRPByProcessGuidReturns struct {
		result1 error
	}
	DesiredLRPsStub        func() ([]models.DesiredLRP, error)
	desiredLRPsMutex       sync.RWMutex
	desiredLRPsArgsForCall []struct{}
	desiredLRPsReturns     struct {
		result1 []models.DesiredLRP
		result2 error
	}
	DesiredLRPsByDomainStub        func(domain string) ([]models.DesiredLRP, error)
	desiredLRPsByDomainMutex       sync.RWMutex
	desiredLRPsByDomainArgsForCall []struct {
		domain string
	}
	desiredLRPsByDomainReturns struct {
		result1 []models.DesiredLRP
		result2 error
	}
	DesiredLRPByProcessGuidStub        func(processGuid string) (models.DesiredLRP, error)
	desiredLRPByProcessGuidMutex       sync.RWMutex
	desiredLRPByProcessGuidArgsForCall []struct {
		processGuid string
	}
	desiredLRPByProcessGuidReturns struct {
		result1 models.DesiredLRP
		result2 error
	}
	WatchForDesiredLRPChangesStub        func(logger lager.Logger, created func(models.DesiredLRP), changed func(models.DesiredLRPChange), deleted func(models.DesiredLRP)) (stop chan<- bool, errs <-chan error)
	watchForDesiredLRPChangesMutex       sync.RWMutex
	watchForDesiredLRPChangesArgsForCall []struct {
		logger  lager.Logger
		created func(models.DesiredLRP)
		changed func(models.DesiredLRPChange)
		deleted func(models.DesiredLRP)
	}
	watchForDesiredLRPChangesReturns struct {
		result1 chan<- bool
		result2 <-chan error
	}
	ActualLRPsStub        func() ([]models.ActualLRP, error)
	actualLRPsMutex       sync.RWMutex
	actualLRPsArgsForCall []struct{}
	actualLRPsReturns     struct {
		result1 []models.ActualLRP
		result2 error
	}
	ActualLRPsByDomainStub        func(domain string) ([]models.ActualLRP, error)
	actualLRPsByDomainMutex       sync.RWMutex
	actualLRPsByDomainArgsForCall []struct {
		domain string
	}
	actualLRPsByDomainReturns struct {
		result1 []models.ActualLRP
		result2 error
	}
	ActualLRPsByProcessGuidStub        func(string) (models.ActualLRPsByIndex, error)
	actualLRPsByProcessGuidMutex       sync.RWMutex
	actualLRPsByProcessGuidArgsForCall []struct {
		arg1 string
	}
	actualLRPsByProcessGuidReturns struct {
		result1 models.ActualLRPsByIndex
		result2 error
	}
	ActualLRPByProcessGuidAndIndexStub        func(string, int) (models.ActualLRP, error)
	actualLRPByProcessGuidAndIndexMutex       sync.RWMutex
	actualLRPByProcessGuidAndIndexArgsForCall []struct {
		arg1 string
		arg2 int
	}
	actualLRPByProcessGuidAndIndexReturns struct {
		result1 models.ActualLRP
		result2 error
	}
	RequestStopLRPInstanceStub        func(models.ActualLRPKey, models.ActualLRPContainerKey) error
	requestStopLRPInstanceMutex       sync.RWMutex
	requestStopLRPInstanceArgsForCall []struct {
		arg1 models.ActualLRPKey
		arg2 models.ActualLRPContainerKey
	}
	requestStopLRPInstanceReturns struct {
		result1 error
	}
	WatchForActualLRPChangesStub        func(logger lager.Logger, created func(models.ActualLRP), changed func(models.ActualLRPChange), deleted func(models.ActualLRP)) (stop chan<- bool, errs <-chan error)
	watchForActualLRPChangesMutex       sync.RWMutex
	watchForActualLRPChangesArgsForCall []struct {
		logger  lager.Logger
		created func(models.ActualLRP)
		changed func(models.ActualLRPChange)
		deleted func(models.ActualLRP)
	}
	watchForActualLRPChangesReturns struct {
		result1 chan<- bool
		result2 <-chan error
	}
	CellsStub        func() ([]models.CellPresence, error)
	cellsMutex       sync.RWMutex
	cellsArgsForCall []struct{}
	cellsReturns     struct {
		result1 []models.CellPresence
		result2 error
	}
	UpsertDomainStub        func(domain string, ttlInSeconds int) error
	upsertDomainMutex       sync.RWMutex
	upsertDomainArgsForCall []struct {
		domain       string
		ttlInSeconds int
	}
	upsertDomainReturns struct {
		result1 error
	}
	DomainsStub        func() ([]string, error)
	domainsMutex       sync.RWMutex
	domainsArgsForCall []struct{}
	domainsReturns     struct {
		result1 []string
		result2 error
	}
	NewReceptorHeartbeatStub        func(models.ReceptorPresence, time.Duration) ifrit.Runner
	newReceptorHeartbeatMutex       sync.RWMutex
	newReceptorHeartbeatArgsForCall []struct {
		arg1 models.ReceptorPresence
		arg2 time.Duration
	}
	newReceptorHeartbeatReturns struct {
		result1 ifrit.Runner
	}
}

func (fake *FakeReceptorBBS) DesireTask(arg1 lager.Logger, arg2 models.Task) error {
	fake.desireTaskMutex.Lock()
	fake.desireTaskArgsForCall = append(fake.desireTaskArgsForCall, struct {
		arg1 lager.Logger
		arg2 models.Task
	}{arg1, arg2})
	fake.desireTaskMutex.Unlock()
	if fake.DesireTaskStub != nil {
		return fake.DesireTaskStub(arg1, arg2)
	} else {
		return fake.desireTaskReturns.result1
	}
}

func (fake *FakeReceptorBBS) DesireTaskCallCount() int {
	fake.desireTaskMutex.RLock()
	defer fake.desireTaskMutex.RUnlock()
	return len(fake.desireTaskArgsForCall)
}

func (fake *FakeReceptorBBS) DesireTaskArgsForCall(i int) (lager.Logger, models.Task) {
	fake.desireTaskMutex.RLock()
	defer fake.desireTaskMutex.RUnlock()
	return fake.desireTaskArgsForCall[i].arg1, fake.desireTaskArgsForCall[i].arg2
}

func (fake *FakeReceptorBBS) DesireTaskReturns(result1 error) {
	fake.DesireTaskStub = nil
	fake.desireTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReceptorBBS) Tasks(logger lager.Logger) ([]models.Task, error) {
	fake.tasksMutex.Lock()
	fake.tasksArgsForCall = append(fake.tasksArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.tasksMutex.Unlock()
	if fake.TasksStub != nil {
		return fake.TasksStub(logger)
	} else {
		return fake.tasksReturns.result1, fake.tasksReturns.result2
	}
}

func (fake *FakeReceptorBBS) TasksCallCount() int {
	fake.tasksMutex.RLock()
	defer fake.tasksMutex.RUnlock()
	return len(fake.tasksArgsForCall)
}

func (fake *FakeReceptorBBS) TasksArgsForCall(i int) lager.Logger {
	fake.tasksMutex.RLock()
	defer fake.tasksMutex.RUnlock()
	return fake.tasksArgsForCall[i].logger
}

func (fake *FakeReceptorBBS) TasksReturns(result1 []models.Task, result2 error) {
	fake.TasksStub = nil
	fake.tasksReturns = struct {
		result1 []models.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeReceptorBBS) TasksByDomain(logger lager.Logger, domain string) ([]models.Task, error) {
	fake.tasksByDomainMutex.Lock()
	fake.tasksByDomainArgsForCall = append(fake.tasksByDomainArgsForCall, struct {
		logger lager.Logger
		domain string
	}{logger, domain})
	fake.tasksByDomainMutex.Unlock()
	if fake.TasksByDomainStub != nil {
		return fake.TasksByDomainStub(logger, domain)
	} else {
		return fake.tasksByDomainReturns.result1, fake.tasksByDomainReturns.result2
	}
}

func (fake *FakeReceptorBBS) TasksByDomainCallCount() int {
	fake.tasksByDomainMutex.RLock()
	defer fake.tasksByDomainMutex.RUnlock()
	return len(fake.tasksByDomainArgsForCall)
}

func (fake *FakeReceptorBBS) TasksByDomainArgsForCall(i int) (lager.Logger, string) {
	fake.tasksByDomainMutex.RLock()
	defer fake.tasksByDomainMutex.RUnlock()
	return fake.tasksByDomainArgsForCall[i].logger, fake.tasksByDomainArgsForCall[i].domain
}

func (fake *FakeReceptorBBS) TasksByDomainReturns(result1 []models.Task, result2 error) {
	fake.TasksByDomainStub = nil
	fake.tasksByDomainReturns = struct {
		result1 []models.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeReceptorBBS) TaskByGuid(taskGuid string) (models.Task, error) {
	fake.taskByGuidMutex.Lock()
	fake.taskByGuidArgsForCall = append(fake.taskByGuidArgsForCall, struct {
		taskGuid string
	}{taskGuid})
	fake.taskByGuidMutex.Unlock()
	if fake.TaskByGuidStub != nil {
		return fake.TaskByGuidStub(taskGuid)
	} else {
		return fake.taskByGuidReturns.result1, fake.taskByGuidReturns.result2
	}
}

func (fake *FakeReceptorBBS) TaskByGuidCallCount() int {
	fake.taskByGuidMutex.RLock()
	defer fake.taskByGuidMutex.RUnlock()
	return len(fake.taskByGuidArgsForCall)
}

func (fake *FakeReceptorBBS) TaskByGuidArgsForCall(i int) string {
	fake.taskByGuidMutex.RLock()
	defer fake.taskByGuidMutex.RUnlock()
	return fake.taskByGuidArgsForCall[i].taskGuid
}

func (fake *FakeReceptorBBS) TaskByGuidReturns(result1 models.Task, result2 error) {
	fake.TaskByGuidStub = nil
	fake.taskByGuidReturns = struct {
		result1 models.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeReceptorBBS) ResolvingTask(logger lager.Logger, taskGuid string) error {
	fake.resolvingTaskMutex.Lock()
	fake.resolvingTaskArgsForCall = append(fake.resolvingTaskArgsForCall, struct {
		logger   lager.Logger
		taskGuid string
	}{logger, taskGuid})
	fake.resolvingTaskMutex.Unlock()
	if fake.ResolvingTaskStub != nil {
		return fake.ResolvingTaskStub(logger, taskGuid)
	} else {
		return fake.resolvingTaskReturns.result1
	}
}

func (fake *FakeReceptorBBS) ResolvingTaskCallCount() int {
	fake.resolvingTaskMutex.RLock()
	defer fake.resolvingTaskMutex.RUnlock()
	return len(fake.resolvingTaskArgsForCall)
}

func (fake *FakeReceptorBBS) ResolvingTaskArgsForCall(i int) (lager.Logger, string) {
	fake.resolvingTaskMutex.RLock()
	defer fake.resolvingTaskMutex.RUnlock()
	return fake.resolvingTaskArgsForCall[i].logger, fake.resolvingTaskArgsForCall[i].taskGuid
}

func (fake *FakeReceptorBBS) ResolvingTaskReturns(result1 error) {
	fake.ResolvingTaskStub = nil
	fake.resolvingTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReceptorBBS) ResolveTask(logger lager.Logger, taskGuid string) error {
	fake.resolveTaskMutex.Lock()
	fake.resolveTaskArgsForCall = append(fake.resolveTaskArgsForCall, struct {
		logger   lager.Logger
		taskGuid string
	}{logger, taskGuid})
	fake.resolveTaskMutex.Unlock()
	if fake.ResolveTaskStub != nil {
		return fake.ResolveTaskStub(logger, taskGuid)
	} else {
		return fake.resolveTaskReturns.result1
	}
}

func (fake *FakeReceptorBBS) ResolveTaskCallCount() int {
	fake.resolveTaskMutex.RLock()
	defer fake.resolveTaskMutex.RUnlock()
	return len(fake.resolveTaskArgsForCall)
}

func (fake *FakeReceptorBBS) ResolveTaskArgsForCall(i int) (lager.Logger, string) {
	fake.resolveTaskMutex.RLock()
	defer fake.resolveTaskMutex.RUnlock()
	return fake.resolveTaskArgsForCall[i].logger, fake.resolveTaskArgsForCall[i].taskGuid
}

func (fake *FakeReceptorBBS) ResolveTaskReturns(result1 error) {
	fake.ResolveTaskStub = nil
	fake.resolveTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReceptorBBS) CancelTask(logger lager.Logger, taskGuid string) error {
	fake.cancelTaskMutex.Lock()
	fake.cancelTaskArgsForCall = append(fake.cancelTaskArgsForCall, struct {
		logger   lager.Logger
		taskGuid string
	}{logger, taskGuid})
	fake.cancelTaskMutex.Unlock()
	if fake.CancelTaskStub != nil {
		return fake.CancelTaskStub(logger, taskGuid)
	} else {
		return fake.cancelTaskReturns.result1
	}
}

func (fake *FakeReceptorBBS) CancelTaskCallCount() int {
	fake.cancelTaskMutex.RLock()
	defer fake.cancelTaskMutex.RUnlock()
	return len(fake.cancelTaskArgsForCall)
}

func (fake *FakeReceptorBBS) CancelTaskArgsForCall(i int) (lager.Logger, string) {
	fake.cancelTaskMutex.RLock()
	defer fake.cancelTaskMutex.RUnlock()
	return fake.cancelTaskArgsForCall[i].logger, fake.cancelTaskArgsForCall[i].taskGuid
}

func (fake *FakeReceptorBBS) CancelTaskReturns(result1 error) {
	fake.CancelTaskStub = nil
	fake.cancelTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReceptorBBS) DesireLRP(arg1 lager.Logger, arg2 models.DesiredLRP) error {
	fake.desireLRPMutex.Lock()
	fake.desireLRPArgsForCall = append(fake.desireLRPArgsForCall, struct {
		arg1 lager.Logger
		arg2 models.DesiredLRP
	}{arg1, arg2})
	fake.desireLRPMutex.Unlock()
	if fake.DesireLRPStub != nil {
		return fake.DesireLRPStub(arg1, arg2)
	} else {
		return fake.desireLRPReturns.result1
	}
}

func (fake *FakeReceptorBBS) DesireLRPCallCount() int {
	fake.desireLRPMutex.RLock()
	defer fake.desireLRPMutex.RUnlock()
	return len(fake.desireLRPArgsForCall)
}

func (fake *FakeReceptorBBS) DesireLRPArgsForCall(i int) (lager.Logger, models.DesiredLRP) {
	fake.desireLRPMutex.RLock()
	defer fake.desireLRPMutex.RUnlock()
	return fake.desireLRPArgsForCall[i].arg1, fake.desireLRPArgsForCall[i].arg2
}

func (fake *FakeReceptorBBS) DesireLRPReturns(result1 error) {
	fake.DesireLRPStub = nil
	fake.desireLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReceptorBBS) UpdateDesiredLRP(logger lager.Logger, processGuid string, update models.DesiredLRPUpdate) error {
	fake.updateDesiredLRPMutex.Lock()
	fake.updateDesiredLRPArgsForCall = append(fake.updateDesiredLRPArgsForCall, struct {
		logger      lager.Logger
		processGuid string
		update      models.DesiredLRPUpdate
	}{logger, processGuid, update})
	fake.updateDesiredLRPMutex.Unlock()
	if fake.UpdateDesiredLRPStub != nil {
		return fake.UpdateDesiredLRPStub(logger, processGuid, update)
	} else {
		return fake.updateDesiredLRPReturns.result1
	}
}

func (fake *FakeReceptorBBS) UpdateDesiredLRPCallCount() int {
	fake.updateDesiredLRPMutex.RLock()
	defer fake.updateDesiredLRPMutex.RUnlock()
	return len(fake.updateDesiredLRPArgsForCall)
}

func (fake *FakeReceptorBBS) UpdateDesiredLRPArgsForCall(i int) (lager.Logger, string, models.DesiredLRPUpdate) {
	fake.updateDesiredLRPMutex.RLock()
	defer fake.updateDesiredLRPMutex.RUnlock()
	return fake.updateDesiredLRPArgsForCall[i].logger, fake.updateDesiredLRPArgsForCall[i].processGuid, fake.updateDesiredLRPArgsForCall[i].update
}

func (fake *FakeReceptorBBS) UpdateDesiredLRPReturns(result1 error) {
	fake.UpdateDesiredLRPStub = nil
	fake.updateDesiredLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReceptorBBS) RemoveDesiredLRPByProcessGuid(logger lager.Logger, processGuid string) error {
	fake.removeDesiredLRPByProcessGuidMutex.Lock()
	fake.removeDesiredLRPByProcessGuidArgsForCall = append(fake.removeDesiredLRPByProcessGuidArgsForCall, struct {
		logger      lager.Logger
		processGuid string
	}{logger, processGuid})
	fake.removeDesiredLRPByProcessGuidMutex.Unlock()
	if fake.RemoveDesiredLRPByProcessGuidStub != nil {
		return fake.RemoveDesiredLRPByProcessGuidStub(logger, processGuid)
	} else {
		return fake.removeDesiredLRPByProcessGuidReturns.result1
	}
}

func (fake *FakeReceptorBBS) RemoveDesiredLRPByProcessGuidCallCount() int {
	fake.removeDesiredLRPByProcessGuidMutex.RLock()
	defer fake.removeDesiredLRPByProcessGuidMutex.RUnlock()
	return len(fake.removeDesiredLRPByProcessGuidArgsForCall)
}

func (fake *FakeReceptorBBS) RemoveDesiredLRPByProcessGuidArgsForCall(i int) (lager.Logger, string) {
	fake.removeDesiredLRPByProcessGuidMutex.RLock()
	defer fake.removeDesiredLRPByProcessGuidMutex.RUnlock()
	return fake.removeDesiredLRPByProcessGuidArgsForCall[i].logger, fake.removeDesiredLRPByProcessGuidArgsForCall[i].processGuid
}

func (fake *FakeReceptorBBS) RemoveDesiredLRPByProcessGuidReturns(result1 error) {
	fake.RemoveDesiredLRPByProcessGuidStub = nil
	fake.removeDesiredLRPByProcessGuidReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReceptorBBS) DesiredLRPs() ([]models.DesiredLRP, error) {
	fake.desiredLRPsMutex.Lock()
	fake.desiredLRPsArgsForCall = append(fake.desiredLRPsArgsForCall, struct{}{})
	fake.desiredLRPsMutex.Unlock()
	if fake.DesiredLRPsStub != nil {
		return fake.DesiredLRPsStub()
	} else {
		return fake.desiredLRPsReturns.result1, fake.desiredLRPsReturns.result2
	}
}

func (fake *FakeReceptorBBS) DesiredLRPsCallCount() int {
	fake.desiredLRPsMutex.RLock()
	defer fake.desiredLRPsMutex.RUnlock()
	return len(fake.desiredLRPsArgsForCall)
}

func (fake *FakeReceptorBBS) DesiredLRPsReturns(result1 []models.DesiredLRP, result2 error) {
	fake.DesiredLRPsStub = nil
	fake.desiredLRPsReturns = struct {
		result1 []models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeReceptorBBS) DesiredLRPsByDomain(domain string) ([]models.DesiredLRP, error) {
	fake.desiredLRPsByDomainMutex.Lock()
	fake.desiredLRPsByDomainArgsForCall = append(fake.desiredLRPsByDomainArgsForCall, struct {
		domain string
	}{domain})
	fake.desiredLRPsByDomainMutex.Unlock()
	if fake.DesiredLRPsByDomainStub != nil {
		return fake.DesiredLRPsByDomainStub(domain)
	} else {
		return fake.desiredLRPsByDomainReturns.result1, fake.desiredLRPsByDomainReturns.result2
	}
}

func (fake *FakeReceptorBBS) DesiredLRPsByDomainCallCount() int {
	fake.desiredLRPsByDomainMutex.RLock()
	defer fake.desiredLRPsByDomainMutex.RUnlock()
	return len(fake.desiredLRPsByDomainArgsForCall)
}

func (fake *FakeReceptorBBS) DesiredLRPsByDomainArgsForCall(i int) string {
	fake.desiredLRPsByDomainMutex.RLock()
	defer fake.desiredLRPsByDomainMutex.RUnlock()
	return fake.desiredLRPsByDomainArgsForCall[i].domain
}

func (fake *FakeReceptorBBS) DesiredLRPsByDomainReturns(result1 []models.DesiredLRP, result2 error) {
	fake.DesiredLRPsByDomainStub = nil
	fake.desiredLRPsByDomainReturns = struct {
		result1 []models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeReceptorBBS) DesiredLRPByProcessGuid(processGuid string) (models.DesiredLRP, error) {
	fake.desiredLRPByProcessGuidMutex.Lock()
	fake.desiredLRPByProcessGuidArgsForCall = append(fake.desiredLRPByProcessGuidArgsForCall, struct {
		processGuid string
	}{processGuid})
	fake.desiredLRPByProcessGuidMutex.Unlock()
	if fake.DesiredLRPByProcessGuidStub != nil {
		return fake.DesiredLRPByProcessGuidStub(processGuid)
	} else {
		return fake.desiredLRPByProcessGuidReturns.result1, fake.desiredLRPByProcessGuidReturns.result2
	}
}

func (fake *FakeReceptorBBS) DesiredLRPByProcessGuidCallCount() int {
	fake.desiredLRPByProcessGuidMutex.RLock()
	defer fake.desiredLRPByProcessGuidMutex.RUnlock()
	return len(fake.desiredLRPByProcessGuidArgsForCall)
}

func (fake *FakeReceptorBBS) DesiredLRPByProcessGuidArgsForCall(i int) string {
	fake.desiredLRPByProcessGuidMutex.RLock()
	defer fake.desiredLRPByProcessGuidMutex.RUnlock()
	return fake.desiredLRPByProcessGuidArgsForCall[i].processGuid
}

func (fake *FakeReceptorBBS) DesiredLRPByProcessGuidReturns(result1 models.DesiredLRP, result2 error) {
	fake.DesiredLRPByProcessGuidStub = nil
	fake.desiredLRPByProcessGuidReturns = struct {
		result1 models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeReceptorBBS) WatchForDesiredLRPChanges(logger lager.Logger, created func(models.DesiredLRP), changed func(models.DesiredLRPChange), deleted func(models.DesiredLRP)) (stop chan<- bool, errs <-chan error) {
	fake.watchForDesiredLRPChangesMutex.Lock()
	fake.watchForDesiredLRPChangesArgsForCall = append(fake.watchForDesiredLRPChangesArgsForCall, struct {
		logger  lager.Logger
		created func(models.DesiredLRP)
		changed func(models.DesiredLRPChange)
		deleted func(models.DesiredLRP)
	}{logger, created, changed, deleted})
	fake.watchForDesiredLRPChangesMutex.Unlock()
	if fake.WatchForDesiredLRPChangesStub != nil {
		return fake.WatchForDesiredLRPChangesStub(logger, created, changed, deleted)
	} else {
		return fake.watchForDesiredLRPChangesReturns.result1, fake.watchForDesiredLRPChangesReturns.result2
	}
}

func (fake *FakeReceptorBBS) WatchForDesiredLRPChangesCallCount() int {
	fake.watchForDesiredLRPChangesMutex.RLock()
	defer fake.watchForDesiredLRPChangesMutex.RUnlock()
	return len(fake.watchForDesiredLRPChangesArgsForCall)
}

func (fake *FakeReceptorBBS) WatchForDesiredLRPChangesArgsForCall(i int) (lager.Logger, func(models.DesiredLRP), func(models.DesiredLRPChange), func(models.DesiredLRP)) {
	fake.watchForDesiredLRPChangesMutex.RLock()
	defer fake.watchForDesiredLRPChangesMutex.RUnlock()
	return fake.watchForDesiredLRPChangesArgsForCall[i].logger, fake.watchForDesiredLRPChangesArgsForCall[i].created, fake.watchForDesiredLRPChangesArgsForCall[i].changed, fake.watchForDesiredLRPChangesArgsForCall[i].deleted
}

func (fake *FakeReceptorBBS) WatchForDesiredLRPChangesReturns(result1 chan<- bool, result2 <-chan error) {
	fake.WatchForDesiredLRPChangesStub = nil
	fake.watchForDesiredLRPChangesReturns = struct {
		result1 chan<- bool
		result2 <-chan error
	}{result1, result2}
}

func (fake *FakeReceptorBBS) ActualLRPs() ([]models.ActualLRP, error) {
	fake.actualLRPsMutex.Lock()
	fake.actualLRPsArgsForCall = append(fake.actualLRPsArgsForCall, struct{}{})
	fake.actualLRPsMutex.Unlock()
	if fake.ActualLRPsStub != nil {
		return fake.ActualLRPsStub()
	} else {
		return fake.actualLRPsReturns.result1, fake.actualLRPsReturns.result2
	}
}

func (fake *FakeReceptorBBS) ActualLRPsCallCount() int {
	fake.actualLRPsMutex.RLock()
	defer fake.actualLRPsMutex.RUnlock()
	return len(fake.actualLRPsArgsForCall)
}

func (fake *FakeReceptorBBS) ActualLRPsReturns(result1 []models.ActualLRP, result2 error) {
	fake.ActualLRPsStub = nil
	fake.actualLRPsReturns = struct {
		result1 []models.ActualLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeReceptorBBS) ActualLRPsByDomain(domain string) ([]models.ActualLRP, error) {
	fake.actualLRPsByDomainMutex.Lock()
	fake.actualLRPsByDomainArgsForCall = append(fake.actualLRPsByDomainArgsForCall, struct {
		domain string
	}{domain})
	fake.actualLRPsByDomainMutex.Unlock()
	if fake.ActualLRPsByDomainStub != nil {
		return fake.ActualLRPsByDomainStub(domain)
	} else {
		return fake.actualLRPsByDomainReturns.result1, fake.actualLRPsByDomainReturns.result2
	}
}

func (fake *FakeReceptorBBS) ActualLRPsByDomainCallCount() int {
	fake.actualLRPsByDomainMutex.RLock()
	defer fake.actualLRPsByDomainMutex.RUnlock()
	return len(fake.actualLRPsByDomainArgsForCall)
}

func (fake *FakeReceptorBBS) ActualLRPsByDomainArgsForCall(i int) string {
	fake.actualLRPsByDomainMutex.RLock()
	defer fake.actualLRPsByDomainMutex.RUnlock()
	return fake.actualLRPsByDomainArgsForCall[i].domain
}

func (fake *FakeReceptorBBS) ActualLRPsByDomainReturns(result1 []models.ActualLRP, result2 error) {
	fake.ActualLRPsByDomainStub = nil
	fake.actualLRPsByDomainReturns = struct {
		result1 []models.ActualLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeReceptorBBS) ActualLRPsByProcessGuid(arg1 string) (models.ActualLRPsByIndex, error) {
	fake.actualLRPsByProcessGuidMutex.Lock()
	fake.actualLRPsByProcessGuidArgsForCall = append(fake.actualLRPsByProcessGuidArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.actualLRPsByProcessGuidMutex.Unlock()
	if fake.ActualLRPsByProcessGuidStub != nil {
		return fake.ActualLRPsByProcessGuidStub(arg1)
	} else {
		return fake.actualLRPsByProcessGuidReturns.result1, fake.actualLRPsByProcessGuidReturns.result2
	}
}

func (fake *FakeReceptorBBS) ActualLRPsByProcessGuidCallCount() int {
	fake.actualLRPsByProcessGuidMutex.RLock()
	defer fake.actualLRPsByProcessGuidMutex.RUnlock()
	return len(fake.actualLRPsByProcessGuidArgsForCall)
}

func (fake *FakeReceptorBBS) ActualLRPsByProcessGuidArgsForCall(i int) string {
	fake.actualLRPsByProcessGuidMutex.RLock()
	defer fake.actualLRPsByProcessGuidMutex.RUnlock()
	return fake.actualLRPsByProcessGuidArgsForCall[i].arg1
}

func (fake *FakeReceptorBBS) ActualLRPsByProcessGuidReturns(result1 models.ActualLRPsByIndex, result2 error) {
	fake.ActualLRPsByProcessGuidStub = nil
	fake.actualLRPsByProcessGuidReturns = struct {
		result1 models.ActualLRPsByIndex
		result2 error
	}{result1, result2}
}

func (fake *FakeReceptorBBS) ActualLRPByProcessGuidAndIndex(arg1 string, arg2 int) (models.ActualLRP, error) {
	fake.actualLRPByProcessGuidAndIndexMutex.Lock()
	fake.actualLRPByProcessGuidAndIndexArgsForCall = append(fake.actualLRPByProcessGuidAndIndexArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.actualLRPByProcessGuidAndIndexMutex.Unlock()
	if fake.ActualLRPByProcessGuidAndIndexStub != nil {
		return fake.ActualLRPByProcessGuidAndIndexStub(arg1, arg2)
	} else {
		return fake.actualLRPByProcessGuidAndIndexReturns.result1, fake.actualLRPByProcessGuidAndIndexReturns.result2
	}
}

func (fake *FakeReceptorBBS) ActualLRPByProcessGuidAndIndexCallCount() int {
	fake.actualLRPByProcessGuidAndIndexMutex.RLock()
	defer fake.actualLRPByProcessGuidAndIndexMutex.RUnlock()
	return len(fake.actualLRPByProcessGuidAndIndexArgsForCall)
}

func (fake *FakeReceptorBBS) ActualLRPByProcessGuidAndIndexArgsForCall(i int) (string, int) {
	fake.actualLRPByProcessGuidAndIndexMutex.RLock()
	defer fake.actualLRPByProcessGuidAndIndexMutex.RUnlock()
	return fake.actualLRPByProcessGuidAndIndexArgsForCall[i].arg1, fake.actualLRPByProcessGuidAndIndexArgsForCall[i].arg2
}

func (fake *FakeReceptorBBS) ActualLRPByProcessGuidAndIndexReturns(result1 models.ActualLRP, result2 error) {
	fake.ActualLRPByProcessGuidAndIndexStub = nil
	fake.actualLRPByProcessGuidAndIndexReturns = struct {
		result1 models.ActualLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeReceptorBBS) RequestStopLRPInstance(arg1 models.ActualLRPKey, arg2 models.ActualLRPContainerKey) error {
	fake.requestStopLRPInstanceMutex.Lock()
	fake.requestStopLRPInstanceArgsForCall = append(fake.requestStopLRPInstanceArgsForCall, struct {
		arg1 models.ActualLRPKey
		arg2 models.ActualLRPContainerKey
	}{arg1, arg2})
	fake.requestStopLRPInstanceMutex.Unlock()
	if fake.RequestStopLRPInstanceStub != nil {
		return fake.RequestStopLRPInstanceStub(arg1, arg2)
	} else {
		return fake.requestStopLRPInstanceReturns.result1
	}
}

func (fake *FakeReceptorBBS) RequestStopLRPInstanceCallCount() int {
	fake.requestStopLRPInstanceMutex.RLock()
	defer fake.requestStopLRPInstanceMutex.RUnlock()
	return len(fake.requestStopLRPInstanceArgsForCall)
}

func (fake *FakeReceptorBBS) RequestStopLRPInstanceArgsForCall(i int) (models.ActualLRPKey, models.ActualLRPContainerKey) {
	fake.requestStopLRPInstanceMutex.RLock()
	defer fake.requestStopLRPInstanceMutex.RUnlock()
	return fake.requestStopLRPInstanceArgsForCall[i].arg1, fake.requestStopLRPInstanceArgsForCall[i].arg2
}

func (fake *FakeReceptorBBS) RequestStopLRPInstanceReturns(result1 error) {
	fake.RequestStopLRPInstanceStub = nil
	fake.requestStopLRPInstanceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReceptorBBS) WatchForActualLRPChanges(logger lager.Logger, created func(models.ActualLRP), changed func(models.ActualLRPChange), deleted func(models.ActualLRP)) (stop chan<- bool, errs <-chan error) {
	fake.watchForActualLRPChangesMutex.Lock()
	fake.watchForActualLRPChangesArgsForCall = append(fake.watchForActualLRPChangesArgsForCall, struct {
		logger  lager.Logger
		created func(models.ActualLRP)
		changed func(models.ActualLRPChange)
		deleted func(models.ActualLRP)
	}{logger, created, changed, deleted})
	fake.watchForActualLRPChangesMutex.Unlock()
	if fake.WatchForActualLRPChangesStub != nil {
		return fake.WatchForActualLRPChangesStub(logger, created, changed, deleted)
	} else {
		return fake.watchForActualLRPChangesReturns.result1, fake.watchForActualLRPChangesReturns.result2
	}
}

func (fake *FakeReceptorBBS) WatchForActualLRPChangesCallCount() int {
	fake.watchForActualLRPChangesMutex.RLock()
	defer fake.watchForActualLRPChangesMutex.RUnlock()
	return len(fake.watchForActualLRPChangesArgsForCall)
}

func (fake *FakeReceptorBBS) WatchForActualLRPChangesArgsForCall(i int) (lager.Logger, func(models.ActualLRP), func(models.ActualLRPChange), func(models.ActualLRP)) {
	fake.watchForActualLRPChangesMutex.RLock()
	defer fake.watchForActualLRPChangesMutex.RUnlock()
	return fake.watchForActualLRPChangesArgsForCall[i].logger, fake.watchForActualLRPChangesArgsForCall[i].created, fake.watchForActualLRPChangesArgsForCall[i].changed, fake.watchForActualLRPChangesArgsForCall[i].deleted
}

func (fake *FakeReceptorBBS) WatchForActualLRPChangesReturns(result1 chan<- bool, result2 <-chan error) {
	fake.WatchForActualLRPChangesStub = nil
	fake.watchForActualLRPChangesReturns = struct {
		result1 chan<- bool
		result2 <-chan error
	}{result1, result2}
}

func (fake *FakeReceptorBBS) Cells() ([]models.CellPresence, error) {
	fake.cellsMutex.Lock()
	fake.cellsArgsForCall = append(fake.cellsArgsForCall, struct{}{})
	fake.cellsMutex.Unlock()
	if fake.CellsStub != nil {
		return fake.CellsStub()
	} else {
		return fake.cellsReturns.result1, fake.cellsReturns.result2
	}
}

func (fake *FakeReceptorBBS) CellsCallCount() int {
	fake.cellsMutex.RLock()
	defer fake.cellsMutex.RUnlock()
	return len(fake.cellsArgsForCall)
}

func (fake *FakeReceptorBBS) CellsReturns(result1 []models.CellPresence, result2 error) {
	fake.CellsStub = nil
	fake.cellsReturns = struct {
		result1 []models.CellPresence
		result2 error
	}{result1, result2}
}

func (fake *FakeReceptorBBS) UpsertDomain(domain string, ttlInSeconds int) error {
	fake.upsertDomainMutex.Lock()
	fake.upsertDomainArgsForCall = append(fake.upsertDomainArgsForCall, struct {
		domain       string
		ttlInSeconds int
	}{domain, ttlInSeconds})
	fake.upsertDomainMutex.Unlock()
	if fake.UpsertDomainStub != nil {
		return fake.UpsertDomainStub(domain, ttlInSeconds)
	} else {
		return fake.upsertDomainReturns.result1
	}
}

func (fake *FakeReceptorBBS) UpsertDomainCallCount() int {
	fake.upsertDomainMutex.RLock()
	defer fake.upsertDomainMutex.RUnlock()
	return len(fake.upsertDomainArgsForCall)
}

func (fake *FakeReceptorBBS) UpsertDomainArgsForCall(i int) (string, int) {
	fake.upsertDomainMutex.RLock()
	defer fake.upsertDomainMutex.RUnlock()
	return fake.upsertDomainArgsForCall[i].domain, fake.upsertDomainArgsForCall[i].ttlInSeconds
}

func (fake *FakeReceptorBBS) UpsertDomainReturns(result1 error) {
	fake.UpsertDomainStub = nil
	fake.upsertDomainReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReceptorBBS) Domains() ([]string, error) {
	fake.domainsMutex.Lock()
	fake.domainsArgsForCall = append(fake.domainsArgsForCall, struct{}{})
	fake.domainsMutex.Unlock()
	if fake.DomainsStub != nil {
		return fake.DomainsStub()
	} else {
		return fake.domainsReturns.result1, fake.domainsReturns.result2
	}
}

func (fake *FakeReceptorBBS) DomainsCallCount() int {
	fake.domainsMutex.RLock()
	defer fake.domainsMutex.RUnlock()
	return len(fake.domainsArgsForCall)
}

func (fake *FakeReceptorBBS) DomainsReturns(result1 []string, result2 error) {
	fake.DomainsStub = nil
	fake.domainsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeReceptorBBS) NewReceptorHeartbeat(arg1 models.ReceptorPresence, arg2 time.Duration) ifrit.Runner {
	fake.newReceptorHeartbeatMutex.Lock()
	fake.newReceptorHeartbeatArgsForCall = append(fake.newReceptorHeartbeatArgsForCall, struct {
		arg1 models.ReceptorPresence
		arg2 time.Duration
	}{arg1, arg2})
	fake.newReceptorHeartbeatMutex.Unlock()
	if fake.NewReceptorHeartbeatStub != nil {
		return fake.NewReceptorHeartbeatStub(arg1, arg2)
	} else {
		return fake.newReceptorHeartbeatReturns.result1
	}
}

func (fake *FakeReceptorBBS) NewReceptorHeartbeatCallCount() int {
	fake.newReceptorHeartbeatMutex.RLock()
	defer fake.newReceptorHeartbeatMutex.RUnlock()
	return len(fake.newReceptorHeartbeatArgsForCall)
}

func (fake *FakeReceptorBBS) NewReceptorHeartbeatArgsForCall(i int) (models.ReceptorPresence, time.Duration) {
	fake.newReceptorHeartbeatMutex.RLock()
	defer fake.newReceptorHeartbeatMutex.RUnlock()
	return fake.newReceptorHeartbeatArgsForCall[i].arg1, fake.newReceptorHeartbeatArgsForCall[i].arg2
}

func (fake *FakeReceptorBBS) NewReceptorHeartbeatReturns(result1 ifrit.Runner) {
	fake.NewReceptorHeartbeatStub = nil
	fake.newReceptorHeartbeatReturns = struct {
		result1 ifrit.Runner
	}{result1}
}

var _ bbs.ReceptorBBS = new(FakeReceptorBBS)
