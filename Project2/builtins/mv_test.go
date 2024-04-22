package builtins

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestRenamefile(t *testing.T) {
	sourceFile, err := ioutil.TempFile("", "source")
	if err != nil {
		t.Fatal("creating temp file failed:", err)
	}
	defer os.Remove(sourceFile.Name())

	targetPath := sourceFile.Name() + "_renamed"
	err = Renamefile(sourceFile.Name(), targetPath)
	if err != nil {
		t.Errorf("Renamefile failed: %v", err)
	}

	_, err = os.Stat(targetPath)
	if os.IsNotExist(err) {
		t.Errorf("File was not renamed")
	}
	defer os.Remove(targetPath)
}


