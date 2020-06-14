package managers

import (
	"fmt"
	"testing"
)

func TestSix910Manager_hashPassword(t *testing.T) {
	var pw = "tester123"
	var sm Six910Manager
	suc, hpw := sm.hashPassword(pw)
	fmt.Println("hashed pw: ", hpw)
	valid := sm.validatePassword(pw, hpw)
	if !suc || hpw == "" || !valid {
		t.Fail()
	}
}
