/*
Copyright IBM Corp. 2017 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package policies

import (
	"fmt"
	"testing"

	"crypto_webtool/fabric/v2.2/msp"

	cb "github.com/hyperledger/fabric-protos-go/common"
	"crypto_webtool/fabric/v2.2/protoutil"
	"github.com/stretchr/testify/assert"
)

const TestPolicyName = "TestPolicyName"

type acceptPolicy struct{}

func (ap acceptPolicy) EvaluateSignedData(signedData []*protoutil.SignedData) error {
	return nil
}

func (ap acceptPolicy) EvaluateIdentities(identity []msp.Identity) error {
	return nil
}

func TestImplicitMarshalError(t *testing.T) {
	_, err := NewImplicitMetaPolicy([]byte("GARBAGE"), nil)
	assert.Error(t, err, "Should have errored unmarshaling garbage")
}

func makeManagers(count, passing int) map[string]*ManagerImpl {
	result := make(map[string]*ManagerImpl)
	remaining := passing
	for i := 0; i < count; i++ {
		policyMap := make(map[string]Policy)
		if remaining > 0 {
			policyMap[TestPolicyName] = acceptPolicy{}
		}
		remaining--

		result[fmt.Sprintf("%d", i)] = &ManagerImpl{
			Policies: policyMap,
		}
	}
	return result
}

// makePolicyTest creates an implicitMetaPolicy with a set of
func runPolicyTest(t *testing.T, rule cb.ImplicitMetaPolicy_Rule, managerCount int, passingCount int) error {
	imp, err := NewImplicitMetaPolicy(protoutil.MarshalOrPanic(&cb.ImplicitMetaPolicy{
		Rule:      rule,
		SubPolicy: TestPolicyName,
	}), makeManagers(managerCount, passingCount))
	if err != nil {
		panic(err)
	}

	errSD := imp.EvaluateSignedData(nil)

	imp, err = NewImplicitMetaPolicy(protoutil.MarshalOrPanic(&cb.ImplicitMetaPolicy{
		Rule:      rule,
		SubPolicy: TestPolicyName,
	}), makeManagers(managerCount, passingCount))
	if err != nil {
		panic(err)
	}

	errI := imp.EvaluateIdentities(nil)

	assert.False(t, ((errI == nil && errSD != nil) || (errSD == nil && errI != nil)))
	if errI != nil && errSD != nil {
		assert.Equal(t, errI.Error(), errSD.Error())
	}

	return errI
}

func TestImplicitMetaAny(t *testing.T) {
	assert.NoError(t, runPolicyTest(t, cb.ImplicitMetaPolicy_ANY, 1, 1))
	assert.NoError(t, runPolicyTest(t, cb.ImplicitMetaPolicy_ANY, 10, 1))
	assert.NoError(t, runPolicyTest(t, cb.ImplicitMetaPolicy_ANY, 10, 8))
	assert.NoError(t, runPolicyTest(t, cb.ImplicitMetaPolicy_ANY, 0, 0))

	err := runPolicyTest(t, cb.ImplicitMetaPolicy_ANY, 10, 0)
	assert.EqualError(t, err, "implicit policy evaluation failed - 0 sub-policies were satisfied, but this policy requires 1 of the 'TestPolicyName' sub-policies to be satisfied")
}

func TestImplicitMetaAll(t *testing.T) {
	assert.NoError(t, runPolicyTest(t, cb.ImplicitMetaPolicy_ALL, 1, 1))
	assert.NoError(t, runPolicyTest(t, cb.ImplicitMetaPolicy_ALL, 10, 10))
	assert.NoError(t, runPolicyTest(t, cb.ImplicitMetaPolicy_ALL, 0, 0))

	err := runPolicyTest(t, cb.ImplicitMetaPolicy_ALL, 10, 1)
	assert.EqualError(t, err, "implicit policy evaluation failed - 1 sub-policies were satisfied, but this policy requires 10 of the 'TestPolicyName' sub-policies to be satisfied")

	err = runPolicyTest(t, cb.ImplicitMetaPolicy_ALL, 10, 0)
	assert.EqualError(t, err, "implicit policy evaluation failed - 0 sub-policies were satisfied, but this policy requires 10 of the 'TestPolicyName' sub-policies to be satisfied")
}

func TestImplicitMetaMajority(t *testing.T) {
	assert.NoError(t, runPolicyTest(t, cb.ImplicitMetaPolicy_MAJORITY, 1, 1))
	assert.NoError(t, runPolicyTest(t, cb.ImplicitMetaPolicy_MAJORITY, 10, 6))
	assert.NoError(t, runPolicyTest(t, cb.ImplicitMetaPolicy_MAJORITY, 3, 2))
	assert.NoError(t, runPolicyTest(t, cb.ImplicitMetaPolicy_MAJORITY, 0, 0))

	err := runPolicyTest(t, cb.ImplicitMetaPolicy_MAJORITY, 10, 5)
	assert.EqualError(t, err, "implicit policy evaluation failed - 5 sub-policies were satisfied, but this policy requires 6 of the 'TestPolicyName' sub-policies to be satisfied")

	err = runPolicyTest(t, cb.ImplicitMetaPolicy_MAJORITY, 10, 0)
	assert.EqualError(t, err, "implicit policy evaluation failed - 0 sub-policies were satisfied, but this policy requires 6 of the 'TestPolicyName' sub-policies to be satisfied")
}
