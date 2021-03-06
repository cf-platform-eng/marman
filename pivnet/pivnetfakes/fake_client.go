// Code generated by counterfeiter. DO NOT EDIT.
package pivnetfakes

import (
	sync "sync"

	semver "github.com/Masterminds/semver"
	pivnet "github.com/cf-platform-eng/marman/pivnet"
	pivneta "github.com/pivotal-cf/go-pivnet"
)

type FakeClient struct {
	AcceptEULAStub        func(string, int) error
	acceptEULAMutex       sync.RWMutex
	acceptEULAArgsForCall []struct {
		arg1 string
		arg2 int
	}
	acceptEULAReturns struct {
		result1 error
	}
	acceptEULAReturnsOnCall map[int]struct {
		result1 error
	}

	DownloadFileStub        func(string, int, *pivneta.ProductFile) error
	downloadFileMutex       sync.RWMutex
	downloadFileArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 *pivneta.ProductFile
	}
	downloadFileReturns struct {
		result1 error
	}
	downloadFileReturnsOnCall map[int]struct {
		result1 error
	}
	FindReleaseByVersionConstraintStub        func(string, *semver.Constraints) (*pivneta.Release, error)
	findReleaseByVersionConstraintMutex       sync.RWMutex
	findReleaseByVersionConstraintArgsForCall []struct {
		arg1 string
		arg2 *semver.Constraints
	}
	findReleaseByVersionConstraintReturns struct {
		result1 *pivneta.Release
		result2 error
	}
	findReleaseByVersionConstraintReturnsOnCall map[int]struct {
		result1 *pivneta.Release
		result2 error
	}

	ListFilesForReleaseStub        func(string, int) ([]pivneta.ProductFile, error)
	listFilesForReleaseMutex       sync.RWMutex
	listFilesForReleaseArgsForCall []struct {
		arg1 string
		arg2 int
	}
	listFilesForReleaseReturns struct {
		result1 []pivneta.ProductFile
		result2 error
	}
	listFilesForReleaseReturnsOnCall map[int]struct {
		result1 []pivneta.ProductFile
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) AcceptEULA(arg1 string, arg2 int) error {
	fake.acceptEULAMutex.Lock()
	ret, specificReturn := fake.acceptEULAReturnsOnCall[len(fake.acceptEULAArgsForCall)]
	fake.acceptEULAArgsForCall = append(fake.acceptEULAArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("AcceptEULA", []interface{}{arg1, arg2})
	fake.acceptEULAMutex.Unlock()
	if fake.AcceptEULAStub != nil {
		return fake.AcceptEULAStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.acceptEULAReturns
	return fakeReturns.result1
}

func (fake *FakeClient) AcceptEULACallCount() int {
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	return len(fake.acceptEULAArgsForCall)
}

func (fake *FakeClient) AcceptEULACalls(stub func(string, int) error) {
	fake.acceptEULAMutex.Lock()
	defer fake.acceptEULAMutex.Unlock()
	fake.AcceptEULAStub = stub
}

func (fake *FakeClient) AcceptEULAArgsForCall(i int) (string, int) {
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	argsForCall := fake.acceptEULAArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) AcceptEULAReturns(result1 error) {
	fake.acceptEULAMutex.Lock()
	defer fake.acceptEULAMutex.Unlock()
	fake.AcceptEULAStub = nil
	fake.acceptEULAReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) AcceptEULAReturnsOnCall(i int, result1 error) {
	fake.acceptEULAMutex.Lock()
	defer fake.acceptEULAMutex.Unlock()
	fake.AcceptEULAStub = nil
	if fake.acceptEULAReturnsOnCall == nil {
		fake.acceptEULAReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.acceptEULAReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DownloadFile(arg1 string, arg2 int, arg3 *pivneta.ProductFile) error {
	fake.downloadFileMutex.Lock()
	ret, specificReturn := fake.downloadFileReturnsOnCall[len(fake.downloadFileArgsForCall)]
	fake.downloadFileArgsForCall = append(fake.downloadFileArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 *pivneta.ProductFile
	}{arg1, arg2, arg3})
	fake.recordInvocation("DownloadFile", []interface{}{arg1, arg2, arg3})
	fake.downloadFileMutex.Unlock()
	if fake.DownloadFileStub != nil {
		return fake.DownloadFileStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.downloadFileReturns
	return fakeReturns.result1
}

func (fake *FakeClient) DownloadFileCallCount() int {
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	return len(fake.downloadFileArgsForCall)
}

func (fake *FakeClient) DownloadFileCalls(stub func(string, int, *pivneta.ProductFile) error) {
	fake.downloadFileMutex.Lock()
	defer fake.downloadFileMutex.Unlock()
	fake.DownloadFileStub = stub
}

func (fake *FakeClient) DownloadFileArgsForCall(i int) (string, int, *pivneta.ProductFile) {
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	argsForCall := fake.downloadFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClient) DownloadFileReturns(result1 error) {
	fake.downloadFileMutex.Lock()
	defer fake.downloadFileMutex.Unlock()
	fake.DownloadFileStub = nil
	fake.downloadFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DownloadFileReturnsOnCall(i int, result1 error) {
	fake.downloadFileMutex.Lock()
	defer fake.downloadFileMutex.Unlock()
	fake.DownloadFileStub = nil
	if fake.downloadFileReturnsOnCall == nil {
		fake.downloadFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downloadFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) FindReleaseByVersionConstraint(arg1 string, arg2 *semver.Constraints) (*pivneta.Release, error) {
	fake.findReleaseByVersionConstraintMutex.Lock()
	ret, specificReturn := fake.findReleaseByVersionConstraintReturnsOnCall[len(fake.findReleaseByVersionConstraintArgsForCall)]
	fake.findReleaseByVersionConstraintArgsForCall = append(fake.findReleaseByVersionConstraintArgsForCall, struct {
		arg1 string
		arg2 *semver.Constraints
	}{arg1, arg2})
	fake.recordInvocation("FindReleaseByVersionConstraint", []interface{}{arg1, arg2})
	fake.findReleaseByVersionConstraintMutex.Unlock()
	if fake.FindReleaseByVersionConstraintStub != nil {
		return fake.FindReleaseByVersionConstraintStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findReleaseByVersionConstraintReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) FindReleaseByVersionConstraintCallCount() int {
	fake.findReleaseByVersionConstraintMutex.RLock()
	defer fake.findReleaseByVersionConstraintMutex.RUnlock()
	return len(fake.findReleaseByVersionConstraintArgsForCall)
}

func (fake *FakeClient) FindReleaseByVersionConstraintCalls(stub func(string, *semver.Constraints) (*pivneta.Release, error)) {
	fake.findReleaseByVersionConstraintMutex.Lock()
	defer fake.findReleaseByVersionConstraintMutex.Unlock()
	fake.FindReleaseByVersionConstraintStub = stub
}

func (fake *FakeClient) FindReleaseByVersionConstraintArgsForCall(i int) (string, *semver.Constraints) {
	fake.findReleaseByVersionConstraintMutex.RLock()
	defer fake.findReleaseByVersionConstraintMutex.RUnlock()
	argsForCall := fake.findReleaseByVersionConstraintArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) FindReleaseByVersionConstraintReturns(result1 *pivneta.Release, result2 error) {
	fake.findReleaseByVersionConstraintMutex.Lock()
	defer fake.findReleaseByVersionConstraintMutex.Unlock()
	fake.FindReleaseByVersionConstraintStub = nil
	fake.findReleaseByVersionConstraintReturns = struct {
		result1 *pivneta.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FindReleaseByVersionConstraintReturnsOnCall(i int, result1 *pivneta.Release, result2 error) {
	fake.findReleaseByVersionConstraintMutex.Lock()
	defer fake.findReleaseByVersionConstraintMutex.Unlock()
	fake.FindReleaseByVersionConstraintStub = nil
	if fake.findReleaseByVersionConstraintReturnsOnCall == nil {
		fake.findReleaseByVersionConstraintReturnsOnCall = make(map[int]struct {
			result1 *pivneta.Release
			result2 error
		})
	}
	fake.findReleaseByVersionConstraintReturnsOnCall[i] = struct {
		result1 *pivneta.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListFilesForRelease(arg1 string, arg2 int) ([]pivneta.ProductFile, error) {
	fake.listFilesForReleaseMutex.Lock()
	ret, specificReturn := fake.listFilesForReleaseReturnsOnCall[len(fake.listFilesForReleaseArgsForCall)]
	fake.listFilesForReleaseArgsForCall = append(fake.listFilesForReleaseArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("ListFilesForRelease", []interface{}{arg1, arg2})
	fake.listFilesForReleaseMutex.Unlock()
	if fake.ListFilesForReleaseStub != nil {
		return fake.ListFilesForReleaseStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listFilesForReleaseReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) ListFilesForReleaseCallCount() int {
	fake.listFilesForReleaseMutex.RLock()
	defer fake.listFilesForReleaseMutex.RUnlock()
	return len(fake.listFilesForReleaseArgsForCall)
}

func (fake *FakeClient) ListFilesForReleaseCalls(stub func(string, int) ([]pivneta.ProductFile, error)) {
	fake.listFilesForReleaseMutex.Lock()
	defer fake.listFilesForReleaseMutex.Unlock()
	fake.ListFilesForReleaseStub = stub
}

func (fake *FakeClient) ListFilesForReleaseArgsForCall(i int) (string, int) {
	fake.listFilesForReleaseMutex.RLock()
	defer fake.listFilesForReleaseMutex.RUnlock()
	argsForCall := fake.listFilesForReleaseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) ListFilesForReleaseReturns(result1 []pivneta.ProductFile, result2 error) {
	fake.listFilesForReleaseMutex.Lock()
	defer fake.listFilesForReleaseMutex.Unlock()
	fake.ListFilesForReleaseStub = nil
	fake.listFilesForReleaseReturns = struct {
		result1 []pivneta.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListFilesForReleaseReturnsOnCall(i int, result1 []pivneta.ProductFile, result2 error) {
	fake.listFilesForReleaseMutex.Lock()
	defer fake.listFilesForReleaseMutex.Unlock()
	fake.ListFilesForReleaseStub = nil
	if fake.listFilesForReleaseReturnsOnCall == nil {
		fake.listFilesForReleaseReturnsOnCall = make(map[int]struct {
			result1 []pivneta.ProductFile
			result2 error
		})
	}
	fake.listFilesForReleaseReturnsOnCall[i] = struct {
		result1 []pivneta.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()

	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	fake.findReleaseByVersionConstraintMutex.RLock()
	defer fake.findReleaseByVersionConstraintMutex.RUnlock()

	fake.listFilesForReleaseMutex.RLock()
	defer fake.listFilesForReleaseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ pivnet.Client = new(FakeClient)
