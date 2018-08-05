package textbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"
	"time"
)

type State struct {
	mu      sync.Mutex
	unsaved bool

	// Dir defaults to a full path to a home directory named after
	// the running executable with a dot in front ("~/.myproggy").
	Dir string `json:"dir,omitempty"`

	// File name defaults to "cache.json", which will be placed into the
	// Dir directory.
	File string `json:"file,omitempty"`

	// Data is the actual key/value data.
	Data map[string]interface{} `json:"data"`

	// Every indicates the duration interval for automatic Saves (if any).
	Every string `json:"every,omitempty"`
}

func (jc *State) Path() string {
	return path.Join(jc.Dir, jc.File)
}

// NewState accepts up to three variadic arguments and returns a new State
// pointer. First argument is the Dir path string. Second is File name.
// Third is a parsable time.Duration string for the interval of Every
// automatic Save.
func NewState(args ...string) *State {
	jc := new(State)
	jc.Data = map[string]interface{}{}
	jc.Every = "0s" // default

	if len(args) > 0 {
		jc.Dir = args[0]
	} else {
		jc.Dir = HomeDotDir() // default
	}

	if len(args) > 1 {
		jc.File = args[1]
	} else {
		jc.File = "cache.json" // default
	}

	if len(args) > 2 {
		_, err := time.ParseDuration(args[2])
		if err != nil {
			panic(err)
		}
		jc.Every = args[2]
	}

	byt, err := ioutil.ReadFile(jc.Path())
	if err != nil {
		if !os.IsNotExist(err) {
			return nil
		}
		jc.Save()
	} else {
		fmt.Println(string(byt))
		err := json.Unmarshal(byt, &jc)
		fmt.Println(jc)
		if err != nil {
			panic(err)
		}
	}

	go jc.autosave()

	return jc
}

// TODO: functions to add to or merge into the existing state
// func Import
// func ImportFile
// func ImportJSON

func (jc *State) Get(keys ...string) interface{} {
	jc.mu.Lock()
	defer jc.mu.Unlock()
	return Get(jc.Data, keys...)
}

func (jc *State) Set(p ...interface{}) {
	jc.mu.Lock()
	defer jc.mu.Unlock()
	Set(jc.Data, p...)
	jc.unsaved = true
}

// autosave is always started for every new cache but does nothing
// unless the Every duration is set to something other than zero, which
// it checks for every second.
func (jc *State) autosave() {
	for {
		dur, err := time.ParseDuration(jc.Every)
		if err != nil {
			panic(err)
		}
		if dur > 0 {
			time.Sleep(dur)
			err := jc.Save()
			if err != nil {
				log.Println(err)
				jc.Every = ""
			}
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

// Save saves the file if it needs saving. No attempt to check if
// something else has modified the file is made.
func (jc *State) Save() error {

	if !jc.unsaved {
		return nil
	}

	jc.mu.Lock()
	defer jc.mu.Unlock()

	_, err := os.Stat(jc.Path())
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		} else {
			os.MkdirAll(jc.Dir, 0700)
		}
	}

	prevu := jc.unsaved
	jc.unsaved = false

	err = ioutil.WriteFile(jc.Path(), []byte(jc.String()), 0600)
	if err != nil {
		jc.unsaved = prevu
		return err
	}

	return nil
}

func (jc *State) ForceSave() error {
	jc.mu.Lock()
	defer jc.mu.Unlock()
	return ioutil.WriteFile(jc.Path(), []byte(jc.String()), 0600)
}

func (jc *State) String() string {
	return JSONString(jc)
}

func (jc *State) Pretty() string {
	byt, err := json.MarshalIndent(jc, "", "  ")
	if err != nil {
		fmt.Sprintf("{\"ERROR\":\"%v\"}", err)
	}
	return string(byt)
}

func (jc *State) Print() {
	fmt.Print(jc.Pretty())
}

//TODO delegate add Marshal and Unmarshal JSON to data
