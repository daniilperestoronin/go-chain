package core

import (
  "fmt"
	"path/filepath"
	"runtime"
	"reflect"
	"testing"
)

func TestWallet(t *testing.T){
  private, public := newKeyPair()
  wallet := Wallet{private, public}
  twallet := NewWallet()
  equals(t, wallet, twallet)
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
