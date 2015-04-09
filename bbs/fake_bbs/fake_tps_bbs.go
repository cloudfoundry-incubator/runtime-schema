// This file was generated by counterfeiter
package fake_bbs

import (
	"sync"
	"time"

	"github.com/cloudfoundry-incubator/runtime-schema/bbs"
	"github.com/tedsuo/ifrit"
)

type FakeTpsBBS struct {
	NewTpsWatcherLockStub        func(watcherID string, retryInterval time.Duration) ifrit.Runner
	newTpsWatcherLockMutex       sync.RWMutex
	newTpsWatcherLockArgsForCall []struct {
		watcherID     string
		retryInterval time.Duration
	}
	newTpsWatcherLockReturns struct {
		result1 ifrit.Runner
	}
}

func (fake *FakeTpsBBS) NewTpsWatcherLock(watcherID string, retryInterval time.Duration) ifrit.Runner {
	fake.newTpsWatcherLockMutex.Lock()
	fake.newTpsWatcherLockArgsForCall = append(fake.newTpsWatcherLockArgsForCall, struct {
		watcherID     string
		retryInterval time.Duration
	}{watcherID, retryInterval})
	fake.newTpsWatcherLockMutex.Unlock()
	if fake.NewTpsWatcherLockStub != nil {
		return fake.NewTpsWatcherLockStub(watcherID, retryInterval)
	} else {
		return fake.newTpsWatcherLockReturns.result1
	}
}

func (fake *FakeTpsBBS) NewTpsWatcherLockCallCount() int {
	fake.newTpsWatcherLockMutex.RLock()
	defer fake.newTpsWatcherLockMutex.RUnlock()
	return len(fake.newTpsWatcherLockArgsForCall)
}

func (fake *FakeTpsBBS) NewTpsWatcherLockArgsForCall(i int) (string, time.Duration) {
	fake.newTpsWatcherLockMutex.RLock()
	defer fake.newTpsWatcherLockMutex.RUnlock()
	return fake.newTpsWatcherLockArgsForCall[i].watcherID, fake.newTpsWatcherLockArgsForCall[i].retryInterval
}

func (fake *FakeTpsBBS) NewTpsWatcherLockReturns(result1 ifrit.Runner) {
	fake.NewTpsWatcherLockStub = nil
	fake.newTpsWatcherLockReturns = struct {
		result1 ifrit.Runner
	}{result1}
}

var _ bbs.TpsBBS = new(FakeTpsBBS)
