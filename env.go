package env

import (
	"bufio"
	"io"
	"os"
	"path"
	"strings"
)

// Load takes in a list of string file names to be used
// as a collection of environmental variables to be used for
// application data kept outside of source
func Load(files ...string) error {

	// If no files provided, look for .env file
	// in the applications top level directory.
	// If none found, return nil
	if len(files) == 0 {
		file, err := os.Open(currentDir(".env"))

		if err != nil && os.IsNotExist(err) {
			return err
		}
		if err != nil {
			return err
		}

		setEnvFromFile(file)

		return nil
	}

	// If at least one file has been provided,
	// get all the env keys and assign then as env values
	// in the curent running process
	for _, file := range files {
		env, err := os.Open(file)
		if err != nil {
			return err
		}

		defer env.Close()
		setEnvFromFile(env)
	}

	return nil
}

func setEnvFromFile(file io.Reader) error {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		textLine := scanner.Text()
		kv := strings.Split(textLine, "=")
		setEnvValue(kv[0], kv[1])
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func setEnvValue(key, value string) bool {
	err := os.Setenv(key, value)
	return err == nil
}

func currentDir(file string) string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path.Join(dir, file)
}
