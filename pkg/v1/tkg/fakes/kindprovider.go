// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"sigs.k8s.io/kind/pkg/cluster"

	"github.com/vmware-tanzu-private/core/pkg/v1/tkg/kind"
)

type KindProvider struct {
	CreateStub        func(string, ...cluster.CreateOption) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 string
		arg2 []cluster.CreateOption
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(string, string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
		arg2 string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	KubeConfigStub        func(string, bool) (string, error)
	kubeConfigMutex       sync.RWMutex
	kubeConfigArgsForCall []struct {
		arg1 string
		arg2 bool
	}
	kubeConfigReturns struct {
		result1 string
		result2 error
	}
	kubeConfigReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *KindProvider) Create(arg1 string, arg2 ...cluster.CreateOption) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 string
		arg2 []cluster.CreateOption
	}{arg1, arg2})
	fake.recordInvocation("Create", []interface{}{arg1, arg2})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1
}

func (fake *KindProvider) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *KindProvider) CreateCalls(stub func(string, ...cluster.CreateOption) error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *KindProvider) CreateArgsForCall(i int) (string, []cluster.CreateOption) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *KindProvider) CreateReturns(result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *KindProvider) CreateReturnsOnCall(i int, result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *KindProvider) Delete(arg1 string, arg2 string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1
}

func (fake *KindProvider) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *KindProvider) DeleteCalls(stub func(string, string) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *KindProvider) DeleteArgsForCall(i int) (string, string) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *KindProvider) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *KindProvider) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *KindProvider) KubeConfig(arg1 string, arg2 bool) (string, error) {
	fake.kubeConfigMutex.Lock()
	ret, specificReturn := fake.kubeConfigReturnsOnCall[len(fake.kubeConfigArgsForCall)]
	fake.kubeConfigArgsForCall = append(fake.kubeConfigArgsForCall, struct {
		arg1 string
		arg2 bool
	}{arg1, arg2})
	fake.recordInvocation("KubeConfig", []interface{}{arg1, arg2})
	fake.kubeConfigMutex.Unlock()
	if fake.KubeConfigStub != nil {
		return fake.KubeConfigStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.kubeConfigReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *KindProvider) KubeConfigCallCount() int {
	fake.kubeConfigMutex.RLock()
	defer fake.kubeConfigMutex.RUnlock()
	return len(fake.kubeConfigArgsForCall)
}

func (fake *KindProvider) KubeConfigCalls(stub func(string, bool) (string, error)) {
	fake.kubeConfigMutex.Lock()
	defer fake.kubeConfigMutex.Unlock()
	fake.KubeConfigStub = stub
}

func (fake *KindProvider) KubeConfigArgsForCall(i int) (string, bool) {
	fake.kubeConfigMutex.RLock()
	defer fake.kubeConfigMutex.RUnlock()
	argsForCall := fake.kubeConfigArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *KindProvider) KubeConfigReturns(result1 string, result2 error) {
	fake.kubeConfigMutex.Lock()
	defer fake.kubeConfigMutex.Unlock()
	fake.KubeConfigStub = nil
	fake.kubeConfigReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *KindProvider) KubeConfigReturnsOnCall(i int, result1 string, result2 error) {
	fake.kubeConfigMutex.Lock()
	defer fake.kubeConfigMutex.Unlock()
	fake.KubeConfigStub = nil
	if fake.kubeConfigReturnsOnCall == nil {
		fake.kubeConfigReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.kubeConfigReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *KindProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.kubeConfigMutex.RLock()
	defer fake.kubeConfigMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *KindProvider) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ kind.KindClusterProvider = new(KindProvider)
