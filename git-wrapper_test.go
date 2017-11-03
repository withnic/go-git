package goGitCmdWrapper

import "testing"

func TestCan(t *testing.T) {
	c, err := Can()
	if !c {
		t.Errorf("Can not use this library. Need :%s", gitCmd)
	}
	if err != nil {
		t.Error(err)
	}
}

func TestExec(t *testing.T) {
	cmd := []string{"diff"}
	res := Exec(cmd)
	if res != 0 {
		t.Errorf("Error: %v", gitCmd)
	}
}

// DANGER TEST!!!
func TestCurrentBranchName(t *testing.T) {
	current, err := CurrentBranchName()
	if err != nil {
		t.Errorf("error: %v", err.Error())
	}
	name := "master"
	if err := Checkout(name); err != nil {
		t.Errorf("error: %v", err.Error())
	}

	if err := Checkout(current); err != nil {
		t.Errorf("error: %v", err.Error())
	}

}
