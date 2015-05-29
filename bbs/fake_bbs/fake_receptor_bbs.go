// This file was generated by counterfeiter
package fake_bbs

import (
	"sync"

	"github.com/cloudfoundry-incubator/runtime-schema/bbs"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
	"github.com/pivotal-golang/lager"
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
	DesiredLRPsStub        func(logger lager.Logger) ([]models.DesiredLRP, error)
	desiredLRPsMutex       sync.RWMutex
	desiredLRPsArgsForCall []struct {
		logger lager.Logger
	}
	desiredLRPsReturns struct {
		result1 []models.DesiredLRP
		result2 error
	}
	DesiredLRPsByDomainStub        func(logger lager.Logger, domain string) ([]models.DesiredLRP, error)
	desiredLRPsByDomainMutex       sync.RWMutex
	desiredLRPsByDomainArgsForCall []struct {
		logger lager.Logger
		domain string
	}
	desiredLRPsByDomainReturns struct {
		result1 []models.DesiredLRP
		result2 error
	}
	DesiredLRPByProcessGuidStub        func(logger lager.Logger, processGuid string) (models.DesiredLRP, error)
	desiredLRPByProcessGuidMutex       sync.RWMutex
	desiredLRPByProcessGuidArgsForCall []struct {
		logger      lager.Logger
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
	ActualLRPGroupsStub        func(logger lager.Logger) ([]models.ActualLRPGroup, error)
	actualLRPGroupsMutex       sync.RWMutex
	actualLRPGroupsArgsForCall []struct {
		logger lager.Logger
	}
	actualLRPGroupsReturns struct {
		result1 []models.ActualLRPGroup
		result2 error
	}
	ActualLRPGroupsByDomainStub        func(logger lager.Logger, domain string) ([]models.ActualLRPGroup, error)
	actualLRPGroupsByDomainMutex       sync.RWMutex
	actualLRPGroupsByDomainArgsForCall []struct {
		logger lager.Logger
		domain string
	}
	actualLRPGroupsByDomainReturns struct {
		result1 []models.ActualLRPGroup
		result2 error
	}
	ActualLRPGroupsByProcessGuidStub        func(logger lager.Logger, processGuid string) (models.ActualLRPGroupsByIndex, error)
	actualLRPGroupsByProcessGuidMutex       sync.RWMutex
	actualLRPGroupsByProcessGuidArgsForCall []struct {
		logger      lager.Logger
		processGuid string
	}
	actualLRPGroupsByProcessGuidReturns struct {
		result1 models.ActualLRPGroupsByIndex
		result2 error
	}
	ActualLRPGroupByProcessGuidAndIndexStub        func(logger lager.Logger, processGuid string, index int) (models.ActualLRPGroup, error)
	actualLRPGroupByProcessGuidAndIndexMutex       sync.RWMutex
	actualLRPGroupByProcessGuidAndIndexArgsForCall []struct {
		logger      lager.Logger
		processGuid string
		index       int
	}
	actualLRPGroupByProcessGuidAndIndexReturns struct {
		result1 models.ActualLRPGroup
		result2 error
	}
	RetireActualLRPsStub        func(lager.Logger, []models.ActualLRPKey)
	retireActualLRPsMutex       sync.RWMutex
	retireActualLRPsArgsForCall []struct {
		arg1 lager.Logger
		arg2 []models.ActualLRPKey
	}
	WatchForActualLRPChangesStub        func(logger lager.Logger, created func(models.ActualLRP, bool), changed func(models.ActualLRPChange, bool), deleted func(models.ActualLRP, bool)) (stop chan<- bool, errs <-chan error)
	watchForActualLRPChangesMutex       sync.RWMutex
	watchForActualLRPChangesArgsForCall []struct {
		logger  lager.Logger
		created func(models.ActualLRP, bool)
		changed func(models.ActualLRPChange, bool)
		deleted func(models.ActualLRP, bool)
	}
	watchForActualLRPChangesReturns struct {
		result1 chan<- bool
		result2 <-chan error
	}
	CellsStub        func() ([]models.CellPresence, error)
	cellsMutex       sync.RWMutex
	cellsArgsForCall []struct{}
	cellsReturns struct {
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
	domainsReturns struct {
		result1 []string
		result2 error
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

func (fake *FakeReceptorBBS) DesiredLRPs(logger lager.Logger) ([]models.DesiredLRP, error) {
	fake.desiredLRPsMutex.Lock()
	fake.desiredLRPsArgsForCall = append(fake.desiredLRPsArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.desiredLRPsMutex.Unlock()
	if fake.DesiredLRPsStub != nil {
		return fake.DesiredLRPsStub(logger)
	} else {
		return fake.desiredLRPsReturns.result1, fake.desiredLRPsReturns.result2
	}
}

func (fake *FakeReceptorBBS) DesiredLRPsCallCount() int {
	fake.desiredLRPsMutex.RLock()
	defer fake.desiredLRPsMutex.RUnlock()
	return len(fake.desiredLRPsArgsForCall)
}

func (fake *FakeReceptorBBS) DesiredLRPsArgsForCall(i int) lager.Logger {
	fake.desiredLRPsMutex.RLock()
	defer fake.desiredLRPsMutex.RUnlock()
	return fake.desiredLRPsArgsForCall[i].logger
}

func (fake *FakeReceptorBBS) DesiredLRPsReturns(result1 []models.DesiredLRP, result2 error) {
	fake.DesiredLRPsStub = nil
	fake.desiredLRPsReturns = struct {
		result1 []models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeReceptorBBS) DesiredLRPsByDomain(logger lager.Logger, domain string) ([]models.DesiredLRP, error) {
	fake.desiredLRPsByDomainMutex.Lock()
	fake.desiredLRPsByDomainArgsForCall = append(fake.desiredLRPsByDomainArgsForCall, struct {
		logger lager.Logger
		domain string
	}{logger, domain})
	fake.desiredLRPsByDomainMutex.Unlock()
	if fake.DesiredLRPsByDomainStub != nil {
		return fake.DesiredLRPsByDomainStub(logger, domain)
	} else {
		return fake.desiredLRPsByDomainReturns.result1, fake.desiredLRPsByDomainReturns.result2
	}
}

func (fake *FakeReceptorBBS) DesiredLRPsByDomainCallCount() int {
	fake.desiredLRPsByDomainMutex.RLock()
	defer fake.desiredLRPsByDomainMutex.RUnlock()
	return len(fake.desiredLRPsByDomainArgsForCall)
}

func (fake *FakeReceptorBBS) DesiredLRPsByDomainArgsForCall(i int) (lager.Logger, string) {
	fake.desiredLRPsByDomainMutex.RLock()
	defer fake.desiredLRPsByDomainMutex.RUnlock()
	return fake.desiredLRPsByDomainArgsForCall[i].logger, fake.desiredLRPsByDomainArgsForCall[i].domain
}

func (fake *FakeReceptorBBS) DesiredLRPsByDomainReturns(result1 []models.DesiredLRP, result2 error) {
	fake.DesiredLRPsByDomainStub = nil
	fake.desiredLRPsByDomainReturns = struct {
		result1 []models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeReceptorBBS) DesiredLRPByProcessGuid(logger lager.Logger, processGuid string) (models.DesiredLRP, error) {
	fake.desiredLRPByProcessGuidMutex.Lock()
	fake.desiredLRPByProcessGuidArgsForCall = append(fake.desiredLRPByProcessGuidArgsForCall, struct {
		logger      lager.Logger
		processGuid string
	}{logger, processGuid})
	fake.desiredLRPByProcessGuidMutex.Unlock()
	if fake.DesiredLRPByProcessGuidStub != nil {
		return fake.DesiredLRPByProcessGuidStub(logger, processGuid)
	} else {
		return fake.desiredLRPByProcessGuidReturns.result1, fake.desiredLRPByProcessGuidReturns.result2
	}
}

func (fake *FakeReceptorBBS) DesiredLRPByProcessGuidCallCount() int {
	fake.desiredLRPByProcessGuidMutex.RLock()
	defer fake.desiredLRPByProcessGuidMutex.RUnlock()
	return len(fake.desiredLRPByProcessGuidArgsForCall)
}

func (fake *FakeReceptorBBS) DesiredLRPByProcessGuidArgsForCall(i int) (lager.Logger, string) {
	fake.desiredLRPByProcessGuidMutex.RLock()
	defer fake.desiredLRPByProcessGuidMutex.RUnlock()
	return fake.desiredLRPByProcessGuidArgsForCall[i].logger, fake.desiredLRPByProcessGuidArgsForCall[i].processGuid
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

func (fake *FakeReceptorBBS) ActualLRPGroups(logger lager.Logger) ([]models.ActualLRPGroup, error) {
	fake.actualLRPGroupsMutex.Lock()
	fake.actualLRPGroupsArgsForCall = append(fake.actualLRPGroupsArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.actualLRPGroupsMutex.Unlock()
	if fake.ActualLRPGroupsStub != nil {
		return fake.ActualLRPGroupsStub(logger)
	} else {
		return fake.actualLRPGroupsReturns.result1, fake.actualLRPGroupsReturns.result2
	}
}

func (fake *FakeReceptorBBS) ActualLRPGroupsCallCount() int {
	fake.actualLRPGroupsMutex.RLock()
	defer fake.actualLRPGroupsMutex.RUnlock()
	return len(fake.actualLRPGroupsArgsForCall)
}

func (fake *FakeReceptorBBS) ActualLRPGroupsArgsForCall(i int) lager.Logger {
	fake.actualLRPGroupsMutex.RLock()
	defer fake.actualLRPGroupsMutex.RUnlock()
	return fake.actualLRPGroupsArgsForCall[i].logger
}

func (fake *FakeReceptorBBS) ActualLRPGroupsReturns(result1 []models.ActualLRPGroup, result2 error) {
	fake.ActualLRPGroupsStub = nil
	fake.actualLRPGroupsReturns = struct {
		result1 []models.ActualLRPGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeReceptorBBS) ActualLRPGroupsByDomain(logger lager.Logger, domain string) ([]models.ActualLRPGroup, error) {
	fake.actualLRPGroupsByDomainMutex.Lock()
	fake.actualLRPGroupsByDomainArgsForCall = append(fake.actualLRPGroupsByDomainArgsForCall, struct {
		logger lager.Logger
		domain string
	}{logger, domain})
	fake.actualLRPGroupsByDomainMutex.Unlock()
	if fake.ActualLRPGroupsByDomainStub != nil {
		return fake.ActualLRPGroupsByDomainStub(logger, domain)
	} else {
		return fake.actualLRPGroupsByDomainReturns.result1, fake.actualLRPGroupsByDomainReturns.result2
	}
}

func (fake *FakeReceptorBBS) ActualLRPGroupsByDomainCallCount() int {
	fake.actualLRPGroupsByDomainMutex.RLock()
	defer fake.actualLRPGroupsByDomainMutex.RUnlock()
	return len(fake.actualLRPGroupsByDomainArgsForCall)
}

func (fake *FakeReceptorBBS) ActualLRPGroupsByDomainArgsForCall(i int) (lager.Logger, string) {
	fake.actualLRPGroupsByDomainMutex.RLock()
	defer fake.actualLRPGroupsByDomainMutex.RUnlock()
	return fake.actualLRPGroupsByDomainArgsForCall[i].logger, fake.actualLRPGroupsByDomainArgsForCall[i].domain
}

func (fake *FakeReceptorBBS) ActualLRPGroupsByDomainReturns(result1 []models.ActualLRPGroup, result2 error) {
	fake.ActualLRPGroupsByDomainStub = nil
	fake.actualLRPGroupsByDomainReturns = struct {
		result1 []models.ActualLRPGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeReceptorBBS) ActualLRPGroupsByProcessGuid(logger lager.Logger, processGuid string) (models.ActualLRPGroupsByIndex, error) {
	fake.actualLRPGroupsByProcessGuidMutex.Lock()
	fake.actualLRPGroupsByProcessGuidArgsForCall = append(fake.actualLRPGroupsByProcessGuidArgsForCall, struct {
		logger      lager.Logger
		processGuid string
	}{logger, processGuid})
	fake.actualLRPGroupsByProcessGuidMutex.Unlock()
	if fake.ActualLRPGroupsByProcessGuidStub != nil {
		return fake.ActualLRPGroupsByProcessGuidStub(logger, processGuid)
	} else {
		return fake.actualLRPGroupsByProcessGuidReturns.result1, fake.actualLRPGroupsByProcessGuidReturns.result2
	}
}

func (fake *FakeReceptorBBS) ActualLRPGroupsByProcessGuidCallCount() int {
	fake.actualLRPGroupsByProcessGuidMutex.RLock()
	defer fake.actualLRPGroupsByProcessGuidMutex.RUnlock()
	return len(fake.actualLRPGroupsByProcessGuidArgsForCall)
}

func (fake *FakeReceptorBBS) ActualLRPGroupsByProcessGuidArgsForCall(i int) (lager.Logger, string) {
	fake.actualLRPGroupsByProcessGuidMutex.RLock()
	defer fake.actualLRPGroupsByProcessGuidMutex.RUnlock()
	return fake.actualLRPGroupsByProcessGuidArgsForCall[i].logger, fake.actualLRPGroupsByProcessGuidArgsForCall[i].processGuid
}

func (fake *FakeReceptorBBS) ActualLRPGroupsByProcessGuidReturns(result1 models.ActualLRPGroupsByIndex, result2 error) {
	fake.ActualLRPGroupsByProcessGuidStub = nil
	fake.actualLRPGroupsByProcessGuidReturns = struct {
		result1 models.ActualLRPGroupsByIndex
		result2 error
	}{result1, result2}
}

func (fake *FakeReceptorBBS) ActualLRPGroupByProcessGuidAndIndex(logger lager.Logger, processGuid string, index int) (models.ActualLRPGroup, error) {
	fake.actualLRPGroupByProcessGuidAndIndexMutex.Lock()
	fake.actualLRPGroupByProcessGuidAndIndexArgsForCall = append(fake.actualLRPGroupByProcessGuidAndIndexArgsForCall, struct {
		logger      lager.Logger
		processGuid string
		index       int
	}{logger, processGuid, index})
	fake.actualLRPGroupByProcessGuidAndIndexMutex.Unlock()
	if fake.ActualLRPGroupByProcessGuidAndIndexStub != nil {
		return fake.ActualLRPGroupByProcessGuidAndIndexStub(logger, processGuid, index)
	} else {
		return fake.actualLRPGroupByProcessGuidAndIndexReturns.result1, fake.actualLRPGroupByProcessGuidAndIndexReturns.result2
	}
}

func (fake *FakeReceptorBBS) ActualLRPGroupByProcessGuidAndIndexCallCount() int {
	fake.actualLRPGroupByProcessGuidAndIndexMutex.RLock()
	defer fake.actualLRPGroupByProcessGuidAndIndexMutex.RUnlock()
	return len(fake.actualLRPGroupByProcessGuidAndIndexArgsForCall)
}

func (fake *FakeReceptorBBS) ActualLRPGroupByProcessGuidAndIndexArgsForCall(i int) (lager.Logger, string, int) {
	fake.actualLRPGroupByProcessGuidAndIndexMutex.RLock()
	defer fake.actualLRPGroupByProcessGuidAndIndexMutex.RUnlock()
	return fake.actualLRPGroupByProcessGuidAndIndexArgsForCall[i].logger, fake.actualLRPGroupByProcessGuidAndIndexArgsForCall[i].processGuid, fake.actualLRPGroupByProcessGuidAndIndexArgsForCall[i].index
}

func (fake *FakeReceptorBBS) ActualLRPGroupByProcessGuidAndIndexReturns(result1 models.ActualLRPGroup, result2 error) {
	fake.ActualLRPGroupByProcessGuidAndIndexStub = nil
	fake.actualLRPGroupByProcessGuidAndIndexReturns = struct {
		result1 models.ActualLRPGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeReceptorBBS) RetireActualLRPs(arg1 lager.Logger, arg2 []models.ActualLRPKey) {
	fake.retireActualLRPsMutex.Lock()
	fake.retireActualLRPsArgsForCall = append(fake.retireActualLRPsArgsForCall, struct {
		arg1 lager.Logger
		arg2 []models.ActualLRPKey
	}{arg1, arg2})
	fake.retireActualLRPsMutex.Unlock()
	if fake.RetireActualLRPsStub != nil {
		fake.RetireActualLRPsStub(arg1, arg2)
	}
}

func (fake *FakeReceptorBBS) RetireActualLRPsCallCount() int {
	fake.retireActualLRPsMutex.RLock()
	defer fake.retireActualLRPsMutex.RUnlock()
	return len(fake.retireActualLRPsArgsForCall)
}

func (fake *FakeReceptorBBS) RetireActualLRPsArgsForCall(i int) (lager.Logger, []models.ActualLRPKey) {
	fake.retireActualLRPsMutex.RLock()
	defer fake.retireActualLRPsMutex.RUnlock()
	return fake.retireActualLRPsArgsForCall[i].arg1, fake.retireActualLRPsArgsForCall[i].arg2
}

func (fake *FakeReceptorBBS) WatchForActualLRPChanges(logger lager.Logger, created func(models.ActualLRP, bool), changed func(models.ActualLRPChange, bool), deleted func(models.ActualLRP, bool)) (stop chan<- bool, errs <-chan error) {
	fake.watchForActualLRPChangesMutex.Lock()
	fake.watchForActualLRPChangesArgsForCall = append(fake.watchForActualLRPChangesArgsForCall, struct {
		logger  lager.Logger
		created func(models.ActualLRP, bool)
		changed func(models.ActualLRPChange, bool)
		deleted func(models.ActualLRP, bool)
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

func (fake *FakeReceptorBBS) WatchForActualLRPChangesArgsForCall(i int) (lager.Logger, func(models.ActualLRP, bool), func(models.ActualLRPChange, bool), func(models.ActualLRP, bool)) {
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

var _ bbs.ReceptorBBS = new(FakeReceptorBBS)
