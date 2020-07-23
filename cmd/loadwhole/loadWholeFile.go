package loadwhole

import (
	"io/ioutil"
)

// LoadWholeFile loads the whole file applying function `apply` on it whole content
func LoadWholeFile(path string, apply func(a ...interface{}) (n int, err error)) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	content := string(b)
	apply(content)
}
