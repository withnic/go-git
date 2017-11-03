package goGitCmdWrapper

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const gitCmd = "git"

// Exec execs git command and return 1 (SUCSESS) or 0 (FAIL)
func Exec(args []string) int {
	cmd := exec.Command(gitCmd, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s Error:%s\n", gitCmd, err.Error())
		return 1
	}
	return 0
}

// Can returns able or unable.
func Can() (bool, error) {
	find := true
	_, err := exec.LookPath(gitCmd)
	if err != nil {
		find = false
	}
	return find, err
}

// Checkout changes current branch HEAD and returns error
func Checkout(name string) error {
	_, err := exec.Command(gitCmd, "checkout", name).Output()
	return err
}

// CurrentBranchName returns current branch name and error
func CurrentBranchName() (string, error) {
	res, err := exec.Command(gitCmd, "symbolic-ref", "--short", "HEAD").Output()
	return strings.Trim(string(res), "\n"), err
}
