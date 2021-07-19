// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"crypto_webtool/fabric/v2.2/msp"
	"crypto_webtool/fabric/v2.2/protoutil"
)

type Policy struct {
	EvaluateIdentitiesStub        func([]msp.Identity) error
	evaluateIdentitiesMutex       sync.RWMutex
	evaluateIdentitiesArgsForCall []struct {
		arg1 []msp.Identity
	}
	evaluateIdentitiesReturns struct {
		result1 error
	}
	evaluateIdentitiesReturnsOnCall map[int]struct {
		result1 error
	}
	EvaluateSignedDataStub        func([]*protoutil.SignedData) error
	evaluateSignedDataMutex       sync.RWMutex
	evaluateSignedDataArgsForCall []struct {
		arg1 []*protoutil.SignedData
	}
	evaluateSignedDataReturns struct {
		result1 error
	}
	evaluateSignedDataReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Policy) EvaluateIdentities(arg1 []msp.Identity) error {
	var arg1Copy []msp.Identity
	if arg1 != nil {
		arg1Copy = make([]msp.Identity, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.evaluateIdentitiesMutex.Lock()
	ret, specificReturn := fake.evaluateIdentitiesReturnsOnCall[len(fake.evaluateIdentitiesArgsForCall)]
	fake.evaluateIdentitiesArgsForCall = append(fake.evaluateIdentitiesArgsForCall, struct {
		arg1 []msp.Identity
	}{arg1Copy})
	fake.recordInvocation("EvaluateIdentities", []interface{}{arg1Copy})
	fake.evaluateIdentitiesMutex.Unlock()
	if fake.EvaluateIdentitiesStub != nil {
		return fake.EvaluateIdentitiesStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.evaluateIdentitiesReturns
	return fakeReturns.result1
}

func (fake *Policy) EvaluateIdentitiesCallCount() int {
	fake.evaluateIdentitiesMutex.RLock()
	defer fake.evaluateIdentitiesMutex.RUnlock()
	return len(fake.evaluateIdentitiesArgsForCall)
}

func (fake *Policy) EvaluateIdentitiesCalls(stub func([]msp.Identity) error) {
	fake.evaluateIdentitiesMutex.Lock()
	defer fake.evaluateIdentitiesMutex.Unlock()
	fake.EvaluateIdentitiesStub = stub
}

func (fake *Policy) EvaluateIdentitiesArgsForCall(i int) []msp.Identity {
	fake.evaluateIdentitiesMutex.RLock()
	defer fake.evaluateIdentitiesMutex.RUnlock()
	argsForCall := fake.evaluateIdentitiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Policy) EvaluateIdentitiesReturns(result1 error) {
	fake.evaluateIdentitiesMutex.Lock()
	defer fake.evaluateIdentitiesMutex.Unlock()
	fake.EvaluateIdentitiesStub = nil
	fake.evaluateIdentitiesReturns = struct {
		result1 error
	}{result1}
}

func (fake *Policy) EvaluateIdentitiesReturnsOnCall(i int, result1 error) {
	fake.evaluateIdentitiesMutex.Lock()
	defer fake.evaluateIdentitiesMutex.Unlock()
	fake.EvaluateIdentitiesStub = nil
	if fake.evaluateIdentitiesReturnsOnCall == nil {
		fake.evaluateIdentitiesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.evaluateIdentitiesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Policy) EvaluateSignedData(arg1 []*protoutil.SignedData) error {
	var arg1Copy []*protoutil.SignedData
	if arg1 != nil {
		arg1Copy = make([]*protoutil.SignedData, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.evaluateSignedDataMutex.Lock()
	ret, specificReturn := fake.evaluateSignedDataReturnsOnCall[len(fake.evaluateSignedDataArgsForCall)]
	fake.evaluateSignedDataArgsForCall = append(fake.evaluateSignedDataArgsForCall, struct {
		arg1 []*protoutil.SignedData
	}{arg1Copy})
	fake.recordInvocation("EvaluateSignedData", []interface{}{arg1Copy})
	fake.evaluateSignedDataMutex.Unlock()
	if fake.EvaluateSignedDataStub != nil {
		return fake.EvaluateSignedDataStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.evaluateSignedDataReturns
	return fakeReturns.result1
}

func (fake *Policy) EvaluateSignedDataCallCount() int {
	fake.evaluateSignedDataMutex.RLock()
	defer fake.evaluateSignedDataMutex.RUnlock()
	return len(fake.evaluateSignedDataArgsForCall)
}

func (fake *Policy) EvaluateSignedDataCalls(stub func([]*protoutil.SignedData) error) {
	fake.evaluateSignedDataMutex.Lock()
	defer fake.evaluateSignedDataMutex.Unlock()
	fake.EvaluateSignedDataStub = stub
}

func (fake *Policy) EvaluateSignedDataArgsForCall(i int) []*protoutil.SignedData {
	fake.evaluateSignedDataMutex.RLock()
	defer fake.evaluateSignedDataMutex.RUnlock()
	argsForCall := fake.evaluateSignedDataArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Policy) EvaluateSignedDataReturns(result1 error) {
	fake.evaluateSignedDataMutex.Lock()
	defer fake.evaluateSignedDataMutex.Unlock()
	fake.EvaluateSignedDataStub = nil
	fake.evaluateSignedDataReturns = struct {
		result1 error
	}{result1}
}

func (fake *Policy) EvaluateSignedDataReturnsOnCall(i int, result1 error) {
	fake.evaluateSignedDataMutex.Lock()
	defer fake.evaluateSignedDataMutex.Unlock()
	fake.EvaluateSignedDataStub = nil
	if fake.evaluateSignedDataReturnsOnCall == nil {
		fake.evaluateSignedDataReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.evaluateSignedDataReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Policy) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.evaluateIdentitiesMutex.RLock()
	defer fake.evaluateIdentitiesMutex.RUnlock()
	fake.evaluateSignedDataMutex.RLock()
	defer fake.evaluateSignedDataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Policy) recordInvocation(key string, args []interface{}) {
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
