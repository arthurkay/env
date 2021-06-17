package env

import (
	"os"
	"reflect"
	"testing"
)

func init() {
	createEnv("TEST=hello_gopher")
}

func createEnv(kv string) error {
	file, err := os.Create(".env")
	if err != nil {
		return err
	}
	_, er := file.WriteString(kv)
	if er != nil {
		return er
	}

	return nil
}

func deleteEnvFile() {
	os.Remove(".env")
}

// TestMultiEnvFiles tests the different number of environmental
// files passed to the package
func TestMultiEnvFiles(t *testing.T) {
	files := []struct {
		Name   string
		Output error
		Input  []string
	}{
		{
			"No file provided",
			nil,
			[]string{},
		},
		{
			"One file provided",
			nil,
			[]string{".env"},
		},
		{
			"Two files provided",
			nil,
			[]string{".env", ".env"},
		},
		{
			"Three files provided",
			nil,
			[]string{".env", ".env", ".env"},
		},
		{
			"Four files provided",
			nil,
			[]string{".env", ".env", ".env", ".env"},
		},
	}

	for _, data := range files {
		t.Run(data.Name, func(t *testing.T) {
			if err := Load(data.Input...); err != nil {
				t.Errorf("Expected nil, but got %v", err)
			}
		})
	}
}

// TestLoadReturnError tests the Load function return type
func TestLoadReturnsError(t *testing.T) {
	responseType := Load()
	var rightError error
	if reflect.TypeOf(rightError) != reflect.TypeOf(responseType) {
		t.Errorf("The load function has to return an error, but returned %T instead", reflect.TypeOf(responseType))
	}
}

// TestBoolReturnOnSetEnv checks the return type of the
// setEnvValue function, making sure it returns a boolean
func TestBoolReturnTypeOnSetEnv(t *testing.T) {
	var key, value string
	var booleanType bool
	actual := setEnvValue(key, value)

	if reflect.TypeOf(booleanType) != reflect.TypeOf(actual) {
		t.Errorf("Expected to get %T, but got %T", reflect.TypeOf(booleanType), reflect.TypeOf(actual))
	}
}

// TestStringCapture tests the ability to get the set value from the package
func TestStringCapture(t *testing.T) {
	err := createEnv("TEST=Arthur")
	if err != nil {
		t.Errorf("Expected a .env file, but found none")
	}
	Load(".env")

	if os.Getenv("TEST") != "Arthur" {
		t.Errorf("Expected Arthur, but got %s", os.Getenv("TEST"))
	}
}

// TestNullString tests handling of null string value
func TestNullString(t *testing.T) {
	err := createEnv("TEST=null")
	if err != nil {
		t.Errorf("expected nil, got %T", err)
	}

	Load()
	if os.Getenv("TEST") != "" {
		t.Errorf("Expected nil, but got %s", os.Getenv("TEST"))
	}

	deleteEnvFile()
}

// TestHandleEmptyLines tests how the package handles empty lines in the
// .TestHandleEmptyLinesenv files
func TestHandleEmptyLines(t *testing.T) {
	err := createEnv("                  ")
	if err != nil {
		t.Errorf("Expected nil, but got %v", err)
	}

	er := Load()
	if er != nil {
		t.Errorf("Expected nil, but got %v", er)
	}
}
