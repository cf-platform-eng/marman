// Code generated by counterfeiter. DO NOT EDIT.
package githubfakes

import (
	io "io"
	sync "sync"

	github "github.com/cf-platform-eng/marman/github"
	githuba "github.com/google/go-github/v25/github"
)

type FakeClient struct {
	DownloadReleaseAssetStub        func(string, string, int64) (io.ReadCloser, string, error)
	downloadReleaseAssetMutex       sync.RWMutex
	downloadReleaseAssetArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 int64
	}
	downloadReleaseAssetReturns struct {
		result1 io.ReadCloser
		result2 string
		result3 error
	}
	downloadReleaseAssetReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 string
		result3 error
	}
	ListReleasesStub        func(string, string, *githuba.ListOptions) ([]*githuba.RepositoryRelease, *githuba.Response, error)
	listReleasesMutex       sync.RWMutex
	listReleasesArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *githuba.ListOptions
	}
	listReleasesReturns struct {
		result1 []*githuba.RepositoryRelease
		result2 *githuba.Response
		result3 error
	}
	listReleasesReturnsOnCall map[int]struct {
		result1 []*githuba.RepositoryRelease
		result2 *githuba.Response
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) DownloadReleaseAsset(arg1 string, arg2 string, arg3 int64) (io.ReadCloser, string, error) {
	fake.downloadReleaseAssetMutex.Lock()
	ret, specificReturn := fake.downloadReleaseAssetReturnsOnCall[len(fake.downloadReleaseAssetArgsForCall)]
	fake.downloadReleaseAssetArgsForCall = append(fake.downloadReleaseAssetArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 int64
	}{arg1, arg2, arg3})
	fake.recordInvocation("DownloadReleaseAsset", []interface{}{arg1, arg2, arg3})
	fake.downloadReleaseAssetMutex.Unlock()
	if fake.DownloadReleaseAssetStub != nil {
		return fake.DownloadReleaseAssetStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.downloadReleaseAssetReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeClient) DownloadReleaseAssetCallCount() int {
	fake.downloadReleaseAssetMutex.RLock()
	defer fake.downloadReleaseAssetMutex.RUnlock()
	return len(fake.downloadReleaseAssetArgsForCall)
}

func (fake *FakeClient) DownloadReleaseAssetCalls(stub func(string, string, int64) (io.ReadCloser, string, error)) {
	fake.downloadReleaseAssetMutex.Lock()
	defer fake.downloadReleaseAssetMutex.Unlock()
	fake.DownloadReleaseAssetStub = stub
}

func (fake *FakeClient) DownloadReleaseAssetArgsForCall(i int) (string, string, int64) {
	fake.downloadReleaseAssetMutex.RLock()
	defer fake.downloadReleaseAssetMutex.RUnlock()
	argsForCall := fake.downloadReleaseAssetArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClient) DownloadReleaseAssetReturns(result1 io.ReadCloser, result2 string, result3 error) {
	fake.downloadReleaseAssetMutex.Lock()
	defer fake.downloadReleaseAssetMutex.Unlock()
	fake.DownloadReleaseAssetStub = nil
	fake.downloadReleaseAssetReturns = struct {
		result1 io.ReadCloser
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) DownloadReleaseAssetReturnsOnCall(i int, result1 io.ReadCloser, result2 string, result3 error) {
	fake.downloadReleaseAssetMutex.Lock()
	defer fake.downloadReleaseAssetMutex.Unlock()
	fake.DownloadReleaseAssetStub = nil
	if fake.downloadReleaseAssetReturnsOnCall == nil {
		fake.downloadReleaseAssetReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 string
			result3 error
		})
	}
	fake.downloadReleaseAssetReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) ListReleases(arg1 string, arg2 string, arg3 *githuba.ListOptions) ([]*githuba.RepositoryRelease, *githuba.Response, error) {
	fake.listReleasesMutex.Lock()
	ret, specificReturn := fake.listReleasesReturnsOnCall[len(fake.listReleasesArgsForCall)]
	fake.listReleasesArgsForCall = append(fake.listReleasesArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *githuba.ListOptions
	}{arg1, arg2, arg3})
	fake.recordInvocation("ListReleases", []interface{}{arg1, arg2, arg3})
	fake.listReleasesMutex.Unlock()
	if fake.ListReleasesStub != nil {
		return fake.ListReleasesStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.listReleasesReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeClient) ListReleasesCallCount() int {
	fake.listReleasesMutex.RLock()
	defer fake.listReleasesMutex.RUnlock()
	return len(fake.listReleasesArgsForCall)
}

func (fake *FakeClient) ListReleasesCalls(stub func(string, string, *githuba.ListOptions) ([]*githuba.RepositoryRelease, *githuba.Response, error)) {
	fake.listReleasesMutex.Lock()
	defer fake.listReleasesMutex.Unlock()
	fake.ListReleasesStub = stub
}

func (fake *FakeClient) ListReleasesArgsForCall(i int) (string, string, *githuba.ListOptions) {
	fake.listReleasesMutex.RLock()
	defer fake.listReleasesMutex.RUnlock()
	argsForCall := fake.listReleasesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClient) ListReleasesReturns(result1 []*githuba.RepositoryRelease, result2 *githuba.Response, result3 error) {
	fake.listReleasesMutex.Lock()
	defer fake.listReleasesMutex.Unlock()
	fake.ListReleasesStub = nil
	fake.listReleasesReturns = struct {
		result1 []*githuba.RepositoryRelease
		result2 *githuba.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) ListReleasesReturnsOnCall(i int, result1 []*githuba.RepositoryRelease, result2 *githuba.Response, result3 error) {
	fake.listReleasesMutex.Lock()
	defer fake.listReleasesMutex.Unlock()
	fake.ListReleasesStub = nil
	if fake.listReleasesReturnsOnCall == nil {
		fake.listReleasesReturnsOnCall = make(map[int]struct {
			result1 []*githuba.RepositoryRelease
			result2 *githuba.Response
			result3 error
		})
	}
	fake.listReleasesReturnsOnCall[i] = struct {
		result1 []*githuba.RepositoryRelease
		result2 *githuba.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downloadReleaseAssetMutex.RLock()
	defer fake.downloadReleaseAssetMutex.RUnlock()
	fake.listReleasesMutex.RLock()
	defer fake.listReleasesMutex.RUnlock()
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

var _ github.Client = new(FakeClient)
