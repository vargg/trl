package files

import (
	"fmt"
	"os"
	"testing"
	"time"
)

type FileFixture struct {
	fileName string
}

func (this *FileFixture) Write(data *[]byte) {
	os.WriteFile(this.fileName, *data, 0755)
}

func (this *FileFixture) Read() *[]byte {
	data, err := os.ReadFile(this.fileName)
	if err != nil {
		panic(err)
	}
	return &data
}

func (this *FileFixture) Remove() {
	os.Remove(this.fileName)
}

func TestRead(this *testing.T) {
	f := FileFixture{fileName: fmt.Sprintf("testRead-%s.txt", time.Now().String())}
	defer f.Remove()

	data := []byte("test-data")
	f.Write(&data)

	readData := Read(f.fileName)

	if string(*readData) != "test-data" {
		this.Errorf("unexpected file data: %s", string(*readData))
	}
}

func TestReadNotExists(this *testing.T) {
	f := FileFixture{fileName: fmt.Sprintf("testRead-%s.txt", time.Now().String())}
	defer f.Remove()

	data := []byte("test-data")
	f.Write(&data)

	readData := Read(f.fileName)

	if string(*readData) != "test-data" {
		this.Errorf("unexpected file data: %s", string(*readData))
	}
}
