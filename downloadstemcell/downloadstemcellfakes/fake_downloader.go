// Code generated by counterfeiter. DO NOT EDIT.
package downloadstemcellfakes

import (
	sync "sync"

	downloadstemcell "github.com/cf-platform-eng/marman/downloadstemcell"
)

type FakeDownloader struct {
	DownloadFromPivnetStub        func(string, string, string, string, string) error
	downloadFromPivnetMutex       sync.RWMutex
	downloadFromPivnetArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}
	downloadFromPivnetReturns struct {
		result1 error
	}
	downloadFromPivnetReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDownloader) DownloadFromPivnet(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) error {
	fake.downloadFromPivnetMutex.Lock()
	ret, specificReturn := fake.downloadFromPivnetReturnsOnCall[len(fake.downloadFromPivnetArgsForCall)]
	fake.downloadFromPivnetArgsForCall = append(fake.downloadFromPivnetArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("DownloadFromPivnet", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.downloadFromPivnetMutex.Unlock()
	if fake.DownloadFromPivnetStub != nil {
		return fake.DownloadFromPivnetStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.downloadFromPivnetReturns
	return fakeReturns.result1
}

func (fake *FakeDownloader) DownloadFromPivnetCallCount() int {
	fake.downloadFromPivnetMutex.RLock()
	defer fake.downloadFromPivnetMutex.RUnlock()
	return len(fake.downloadFromPivnetArgsForCall)
}

func (fake *FakeDownloader) DownloadFromPivnetCalls(stub func(string, string, string, string, string) error) {
	fake.downloadFromPivnetMutex.Lock()
	defer fake.downloadFromPivnetMutex.Unlock()
	fake.DownloadFromPivnetStub = stub
}

func (fake *FakeDownloader) DownloadFromPivnetArgsForCall(i int) (string, string, string, string, string) {
	fake.downloadFromPivnetMutex.RLock()
	defer fake.downloadFromPivnetMutex.RUnlock()
	argsForCall := fake.downloadFromPivnetArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeDownloader) DownloadFromPivnetReturns(result1 error) {
	fake.downloadFromPivnetMutex.Lock()
	defer fake.downloadFromPivnetMutex.Unlock()
	fake.DownloadFromPivnetStub = nil
	fake.downloadFromPivnetReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDownloader) DownloadFromPivnetReturnsOnCall(i int, result1 error) {
	fake.downloadFromPivnetMutex.Lock()
	defer fake.downloadFromPivnetMutex.Unlock()
	fake.DownloadFromPivnetStub = nil
	if fake.downloadFromPivnetReturnsOnCall == nil {
		fake.downloadFromPivnetReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downloadFromPivnetReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDownloader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downloadFromPivnetMutex.RLock()
	defer fake.downloadFromPivnetMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDownloader) recordInvocation(key string, args []interface{}) {
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

var _ downloadstemcell.Downloader = new(FakeDownloader)
