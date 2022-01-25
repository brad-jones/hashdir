package tests

import (
	"fmt"
	"path/filepath"
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
	assert.Equal(t, "2f598d7692dfca27bc7bc5a18d1c389b", testcli.Stdout())
}

func TestSameLogicalPathShouldHaveTheSameHash(t *testing.T) {
	p, err := filepath.Abs("./data/test_case_1")
	assert.Nil(t, err, err)
	testcli.Run(sutBin, "-alg", "md5", p)
	assert.Equal(t, true, testcli.Success(), testcli.Error())
	assert.Equal(t, "2f598d7692dfca27bc7bc5a18d1c389b", testcli.Stdout())
}

func TestSimpleSHA1(t *testing.T) {
	testcli.Run(sutBin, "-alg", "sha1", "./data/test_case_1")
	assert.Equal(t, true, testcli.Success(), testcli.Error())
	assert.Equal(t, "af1b109a5dbc8bb9f26d5e74c0d5459ba61b38d6", testcli.Stdout())
}

func TestSimpleSHA256(t *testing.T) {
	testcli.Run(sutBin, "-alg", "sha256", "./data/test_case_1")
	assert.Equal(t, true, testcli.Success(), testcli.Error())
	assert.Equal(t, "799c020fbc71cc5ec69f8169f194cdb120227015852fa085bc47c9f6ecb69508", testcli.Stdout())
}

func TestSimpleSHA512(t *testing.T) {
	testcli.Run(sutBin, "-alg", "sha512", "./data/test_case_1")
	assert.Equal(t, true, testcli.Success(), testcli.Error())
	assert.Equal(t, "d2a6616693c7da9fd17923462dfca974a591628091ba7d36e2dba22fa2bd042b41cbf4ab0021213e496506bc8c8fef9e679cbebcc0b845f337898f05e3ca61ed", testcli.Stdout())
}

func TestRecursiveMD5(t *testing.T) {
	testcli.Run(sutBin, "-alg", "md5", "./data/test_case_2")
	assert.Equal(t, true, testcli.Success(), testcli.Error())
	assert.Equal(t, "538362ba8956c37bacea786d92e7c264", testcli.Stdout())
}

func TestRecursiveSHA1(t *testing.T) {
	testcli.Run(sutBin, "-alg", "sha1", "./data/test_case_2")
	assert.Equal(t, true, testcli.Success(), testcli.Error())
	assert.Equal(t, "40819838dc428f97fe0d736f21a86a12613e8275", testcli.Stdout())
}

func TestRecursiveSHA256(t *testing.T) {
	testcli.Run(sutBin, "-alg", "sha256", "./data/test_case_2")
	assert.Equal(t, true, testcli.Success(), testcli.Error())
	assert.Equal(t, "9d8007ab4ddd268c2cdb428a2846bb221c90567ac2abc6dc741880750eef0bb7", testcli.Stdout())
}

func TestRecursiveSHA512(t *testing.T) {
	testcli.Run(sutBin, "-alg", "sha512", "./data/test_case_2")
	assert.Equal(t, true, testcli.Success(), testcli.Error())
	assert.Equal(t, "14438bfb48891eab6647f77358057e75a01b42d91bfd43e79925a9f59aa6a137f6633b559fe76fe25d63bf461da6b855c31d677addcb73c08c06f14aba39bb5b", testcli.Stdout())
}
