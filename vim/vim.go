package vim

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/lepra-tsr/gdbt/config"

	"path/filepath"
)

type Vim struct {
}

func (u *Vim) OpenTemporaryFile(input string) (string, error) {
	tempFilePath, _ := filepath.Abs(filepath.Join(os.TempDir(), "/gdbt.md"))
	f, err := os.Create(tempFilePath)
	if err != nil {
		fmt.Println("failed to create temporary file: " + tempFilePath)
		return "", err
	}
	commentTemplate := "#! lines start with #! will be ignored.\n#!\n\n"

	inputBytes := []byte(commentTemplate + input)
	f.Write(inputBytes)
	f.Close()

	if err := u.open(tempFilePath); err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadFile(tempFilePath)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (u *Vim) open(absPath string) error {
	cmd := exec.Command("vim", "+", absPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println("cannot open temporary file with vim.")
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}

func (u *Vim) OpenDraftFile() error {
	u.open(config.DraftFilePath)
	return nil
}
