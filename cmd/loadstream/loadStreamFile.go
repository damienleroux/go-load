package loadstream

import (
	"bufio"
	"io"
	"os"
	"sync"
)

func processChunk(chunk []byte, apply *func(a ...interface{}) (n int, err error), stringPool *sync.Pool) {
	logs := stringPool.Get().(string)
	logs = string(chunk)
	(*apply)(logs)
}

func readFile(f *os.File, apply *func(a ...interface{}) (n int, err error)) (err error) {
	r := bufio.NewReader(f)

	stringPool := sync.Pool{New: func() interface{} {
		lines := ""
		return lines
	}}

	var wg sync.WaitGroup

	for {
		b, err := r.ReadBytes('\n')

		if err != nil {
			if err == io.EOF {
				if len(b) != 0 {
					wg.Add(1)
					go func() {
						processChunk(b, apply, &stringPool)
						wg.Done()
					}()
				}
				wg.Wait()
				return nil
			}
			return err
		}
		wg.Add(1)
		go func() {
			processChunk(b, apply, &stringPool)
			wg.Done()
		}()
	}
}

// LoadStreamFile loads the file by chunks applying function `apply` on every chunks of streamed content from file
func LoadStreamFile(path string, apply func(a ...interface{}) (n int, err error)) {
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	err = readFile(f, &apply)
	if err != nil {
		panic(err)
	}

}
