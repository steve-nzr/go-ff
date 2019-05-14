package defines

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

// List of all defines
var List = load()

func load() (list map[string]int64) {
	list = make(map[string]int64)
	files, err := ioutil.ReadDir("resources/")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !strings.HasPrefix(file.Name(), "define") && !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		b, err := ioutil.ReadFile(fmt.Sprintf("resources/%s", file.Name()))
		if err != nil {
			panic(err)
		}

		fileDefines := make(map[string]int64)
		json.Unmarshal(b, &fileDefines)
		for k, v := range fileDefines {
			if _, ok := list[k]; ok {
				panic(fmt.Errorf("define %v already used", k))
			}
			list[k] = v
		}

	}

	return list
}

// Get a value by key & return it with the ok value (found or not)
func Get(key string) (v int64, ok bool) {
	v, ok = List[key]
	return
}

// MustGet returns a value from a key or panic
func MustGet(key string) int64 {
	v, ok := List[key]
	if !ok {
		panic(fmt.Errorf("defines key %s not found", key))
	}
	return v
}
