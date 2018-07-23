package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"time"
)

func fmtHeader(t time.Time) string {
	return fmt.Sprintf("%s\n\n", t.Format(time.UnixDate))
}

func loadFile(dir string) error {
	// Start up a temp file
	tmpFile, err := ioutil.TempFile("", "did")
	if err != nil {
		return err
	}

	tNow := time.Now()
	// Write Record Header
	header := fmtHeader(tNow)
	tmpFile.WriteString(header)
	tmpFile.Close()

	// Open vim for additionally documenting stuff
	cmd := exec.Command("vim", "+2", tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}

	// Read contents of just written vim file
	contents, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		log.Printf("Error while editing. Error: %v\n", err)
	} else {
		if header == string(contents) {
			log.Fatal("Aborting do to empty commit message")
		}
		log.Printf("Successfully edited.")
	}

	// Write record
	didFile := filepath.Join(dir, ".did.txt")
	fmt.Println("Writing to ", didFile)
	f, err := os.OpenFile(didFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	record := fmt.Sprintf("\n<<RECORD\n%s\n", string(contents))
	_, err = f.WriteString(record)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func main() {
	// get home directory
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	configDir := filepath.Join(usr.HomeDir, ".did")
	err = os.MkdirAll(configDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	loadFile(filepath.Join(usr.HomeDir, ".did"))
}
