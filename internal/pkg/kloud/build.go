package kloud

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	API_BUILD_STATIC              = "http://0.0.0.0:3000/api/core/build/%s/%s"
	API_BUILD_STATIC_CONTENT_TYPE = "application/javascript"
)

type BuildArg struct {
	Dockerfile string
}

type Build struct {
	BuildArg
	WorkDirectoryPath string
	gzipPathFile      string
}

func NewBuild() (b *Build, err error) {
	b = new(Build)
	b.Dockerfile = "./Dockerfile"
	if b.WorkDirectoryPath, err = os.Getwd(); err != nil {
		return nil, err
	}

	return b, nil
}

func (b *Build) Execute() error {
	// create archive of project
	if err := Gzip(b.WorkDirectoryPath, b.WorkDirectoryPath); err != nil {
		return err
	}

	// send archive
	b.gzipPathFile = gzipFilePath(b.WorkDirectoryPath)
	b.send("rubick", "test2")

	return nil
}

func (b *Build) Validate() error {
	// check workDirectory exist
	wd, err := os.Open(b.WorkDirectoryPath)
	if err != nil {
		return err
	}

	// check workDirectory is Directory
	if wdInfo, err := wd.Stat(); err != nil {
		return err
	} else if !wdInfo.IsDir() {
		return ErrWorkDirectory
	}

	// check dockerfile exist
	dockerfilePath := filepath.Join(b.WorkDirectoryPath, b.Dockerfile)
	if _, err := os.Stat(dockerfilePath); err != nil {
		return err
	}

	return nil
}

// TODO return ERROR
func (b *Build) send(co, project string) {

	// // we need to wait for everything to be done
	// wg := sync.WaitGroup{}
	// wg.Add(2)

	// get GZip file as input
	f, err := os.Open(b.gzipPathFile)
	if err != nil {
		Exit(err)
	}
	url := fmt.Sprintf(API_BUILD_STATIC, co, project)
	res, err := http.Post(url, API_BUILD_STATIC_CONTENT_TYPE, f)
	if err != nil {
		Exit(err)
	}

	log.Print(res)

	// // TeeReader gets the data from the file and also writes it to the PipeWriter
	// tr := io.TeeReader(f, pw)

	// go func() {
	// 	defer wg.Done()
	// 	defer pw.Close()

	// 	// get data from the TeeReader, which feeds the PipeReader through the PipeWriter
	// 	url := fmt.Sprintf(API_BUILD_STATIC, co, project)
	// 	res, err := http.Post(url, API_BUILD_STATIC_CONTENT_TYPE, tr)
	// 	if err != nil {
	// 		Exit(err)
	// 	}
	// 	ss <- res.Request.URL.RawFragment
	// }()

	// go func() {
	// 	defer wg.Done()

	// 	// read from the PipeReader to stdout
	// 	if _, err := io.Copy(os.Stdout, pr); err != nil {
	// 		Exit(err)
	// 	}

	// 	ss <- "stdout"
	// }()

	// wg.Wait()
	// fmt.Println(<-ss)
	// fmt.Println(<-ss)
}
