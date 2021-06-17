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

func TestLoadReturnsError(t *testing.T) {
	responseType := Load()
	var rightError error
	if reflect.TypeOf(rightError) != reflect.TypeOf(responseType) {
		t.Errorf("The load function has to return an error, but returned %T instead", reflect.TypeOf(responseType))
	}
}

func TestBoolReturnTypeOnSetEnv(t *testing.T) {
	var key, value string
	var booleanType bool
	actual := setEnvValue(key, value)

	if reflect.TypeOf(booleanType) != reflect.TypeOf(actual) {
		t.Errorf("Expected to get %T, but got %T", reflect.TypeOf(booleanType), reflect.TypeOf(actual))
	}
}

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
