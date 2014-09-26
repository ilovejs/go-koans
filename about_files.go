package go_koans

import "io/ioutil"
import "strings"

func aboutFiles() {
	filename := "about_files.go"

	contents, _ := ioutil.ReadFile(filename)

	lines := strings.Split(string(contents), "\n")

	//dlog(lines[5])
	assert(lines[0] == "package go_koans\r") // handling files is too trivial

	//WTF
	assert(lines[5] == "func aboutFiles() {\r")
}
