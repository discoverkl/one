package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func InstallAndExit() {
	install := (len(os.Args) == 2 && os.Args[1] == "install")
	if install {
		err := installToSysPath()
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
}

func installToSysPath() error {
	sysPath := os.Getenv("PATH")
	sp := strings.Split(sysPath, ":")

	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	exePath, _ = readlink(exePath)
	name := filepath.Base(exePath)

	// check if already installed
	for _, path := range sp {
		target := filepath.Join(path, name)
		target, _ = readlink(target)
		if target == exePath {
			fmt.Printf("already installed in directory: %s\n", path)
			return nil
		}
	}

	for _, path := range sp {
		target := filepath.Join(path, name)
		// err = os.Rename(exePath, target)
		err = move(exePath, target)
		if err == nil {
			fmt.Printf("OK. %s is moved to directory: %s\n", name, path)
			return nil
		} else {
			log.Printf("try move to '%s' failed: %v", path, err)
		}
	}
	return fmt.Errorf("not installed")
}

func readlink(path string) (string, error) {
	hold, err := filepath.EvalSymlinks(path)
	if err != nil {
		return "", err
	}
	return filepath.Abs(hold)
}

// move and chmod +x
func move(path string, target string) error {
	fin, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fin.Close()

	if _, err = os.Stat(target); err == nil {
		err = os.Remove(target)
		if err != nil {
			return fmt.Errorf("remove target failed: %w", err)
		}
	}

	fout, err := os.Create(target)
	if err != nil {
		return err
	}
	defer fout.Close()

	_, err = io.Copy(fout, fin)
	if err != nil {
		return err
	}

	err = os.Chmod(target, 0755)
	if err != nil {
		return err
	}

	err = os.Remove(path)
	if err != nil {
		log.Printf("WARN remove '%s' failed: %v", path, err)
	}
	return nil
}
