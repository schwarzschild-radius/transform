package spmdfy

import (
	"os"
	"os/exec"
	"log"
	"errors"
	"bytes"
	"hash/fnv"
	"io/ioutil"
)

func getSpmdfyPath() (string, error) {
	if path := os.Getenv("SPMDFYPATH"); len(path) > 0 {
		return path, nil
	}
	return "", errors.New("Please set SPMDFYPATH")
}

func runCMD(cmd []string) (string, string, bool) {
	log.SetPrefix("[spmdfy/run] ")
	handle := exec.Command(cmd[0], cmd[1:]...)
	log.Printf("running cmd %v", cmd)
	var stdout, stderr bytes.Buffer
	handle.Stderr = &stderr
	handle.Stdout = &stdout
	err := handle.Run()
	if err != nil {
		log.Fatal(err)
		return stdout.String(), stderr.String(), true
	}
	return stdout.String(), stderr.String(), true
}

func newTmpFile(filename string) (string, *os.File, error) {
	log.SetPrefix("[spmdfy/tmp] ")
	if err := os.MkdirAll("/tmp/spmdfy", 0755); err != nil {
		log.Panicln("unable to create tmpfile named: " + filename)
		return filename, nil, err
	}
	file, _ := os.Create("/tmp/spmdfy/" + filename + ".cu")
	log.Println(filename + " created in /tmp/spmdfy")
	return "/tmp/spmdfy/" + filename, file, nil
}

func getSrcHash(src []byte) (string) {
	h := fnv.New32a()
	h.Write(src)
	return string(h.Sum32())
}

func Spmdfy(src string) (string, error) {
	log.SetPrefix("[spmdfy] ")
	bsrc := []byte(src)
	filename := getSrcHash(bsrc)
	tmpfile, tmpfilehandle, _ := newTmpFile(filename)
	tmpfilehandle.Write(bsrc)
	spmdfypath, _ := getSpmdfyPath()
	stdout, _, _  := runCMD([]string{ spmdfypath + "/spmdfy", tmpfile + ".cu", "-o", tmpfile + ".ispc"})
	out, _ := ioutil.ReadFile(tmpfile + ".ispc")
	log.Println(stdout)
	return string(out) , errors.New("not sure what happened")
}


/// test function
func testSpmdfyCmd(){
	log.SetPrefix("[spmdfy/test] ")
	spmdfy_path , _ := getSpmdfyPath()
	stdout, _, _ := runCMD([]string{spmdfy_path + "/spmdfy", "--help"})
	log.Println(stdout)
}