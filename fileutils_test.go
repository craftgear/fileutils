package fileutils

import (
	"os"
	"os/exec"
	"testing"
)

func TestMain(m *testing.M) {
	cmd := exec.Command("cp", "-r", "test", "_test")
	cmd.Run()
	m.Run()
	cmd = exec.Command("rm", "-rf", "test")
	cmd.Run()
	cmd = exec.Command("mv", "_test", "test")
	cmd.Run()
}

func TestCp(t *testing.T) {
	dest := "./test/fuga.txt"
	err := Cp(dest, "./test/hoge.txt", 0755)
	if err != nil {
		t.Error(err)
	}
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		t.Error(dest, "should exist", err)
	}
	//TODO
	//fuga.txtの中身がhoge.txtと同じこと

}

func TestCpDir(t *testing.T) {
	//TODO
}

func TestExist(t *testing.T) {
	//TODO
}
