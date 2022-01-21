package tests

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/rendon/testcli"
	"github.com/stretchr/testify/assert"
)

var sutBin = fmt.Sprintf("../bin/hashdir_%s_%s%s",
	runtime.GOOS, runtime.GOARCH, exeExt(),
)

func exeExt() string {
	if runtime.GOOS == "windows" {
		return ".exe"
	}
	return ""
}

func TestSimpleMD5(t *testing.T) {
	testcli.Run(sutBin, "-alg", "md5", "./data/test_case_1")
	assert.Equal(t, true, testcli.Success(), testcli.Error())
	assert.Equal(t, "9f9fe4e25d87330a51903f16f5cb6b7f", testcli.Stdout())
}

func TestSimpleSHA1(t *testing.T) {
	testcli.Run(sutBin, "-alg", "sha1", "./data/test_case_1")
	assert.Equal(t, true, testcli.Success(), testcli.Error())
	assert.Equal(t, "6830ecc3d3eb29277f314dcae45f37c38fe03c84", testcli.Stdout())
}

func TestSimpleSHA256(t *testing.T) {
	testcli.Run(sutBin, "-alg", "sha256", "./data/test_case_1")
	assert.Equal(t, true, testcli.Success(), testcli.Error())
	assert.Equal(t, "733e450b4a1beaea7a6f454aa5963428d6d0044f5503b0d6eee6ca4091725325", testcli.Stdout())
}

func TestSimpleSHA512(t *testing.T) {
	testcli.Run(sutBin, "-alg", "sha512", "./data/test_case_1")
	assert.Equal(t, true, testcli.Success(), testcli.Error())
	assert.Equal(t, "d9b04f14025173be560ada5c798686e016f6acab8971e0dea690401d31dead3dccd29e4b6f5fac62f5b9c105c4a2ba02d3afe2f4dd1d5134ef8c3a927ea0b7b9", testcli.Stdout())
}

func TestRecursiveMD5(t *testing.T) {
	testcli.Run(sutBin, "-alg", "md5", "./data/test_case_2")
	assert.Equal(t, true, testcli.Success(), testcli.Error())
	assert.Equal(t, "1838ca8cd9e7b80f5f136342c76ef2a0", testcli.Stdout())
}

func TestRecursiveSHA1(t *testing.T) {
	testcli.Run(sutBin, "-alg", "sha1", "./data/test_case_2")
	assert.Equal(t, true, testcli.Success(), testcli.Error())
	assert.Equal(t, "7760376449090ebd5007a3450451d6ea05d16284", testcli.Stdout())
}

func TestRecursiveSHA256(t *testing.T) {
	testcli.Run(sutBin, "-alg", "sha256", "./data/test_case_2")
	assert.Equal(t, true, testcli.Success(), testcli.Error())
	assert.Equal(t, "6b6a3bea390cb56e4b025999d4a0bb07ff7d9e121b525ae54840917edd866c17", testcli.Stdout())
}

func TestRecursiveSHA512(t *testing.T) {
	testcli.Run(sutBin, "-alg", "sha512", "./data/test_case_2")
	assert.Equal(t, true, testcli.Success(), testcli.Error())
	assert.Equal(t, "70c27be1b1e98cb1fac3acf0cd92c3b48026683370cdd9d50f8494790fc57dddd31d3cf670953bfca8bafe62ec79c173d16b8090d77838f0bf641980bdf1ec70", testcli.Stdout())
}
