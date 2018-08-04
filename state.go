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
	Data Data `json:"data"`

	// Save is set to the last time a save successfully completed.
	Saved time.Time `json:"-"`

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
	jc.Data = Data{}
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

	go jc.autosave()

	return jc
}

func NewStateFromJSON(jsn []byte, args ...string) (*State, error) {
	jc := NewState(args...)
	err := json.Unmarshal(jsn, &jc)
	return jc, err
}

// NewStateFromFile uses the specific path to read a json file. However, the
// path does not automatically set the Dir and File, which must be
// provided as separate args as well if the default initial values are
// not wanted.
func NewStateFromFile(path string, args ...string) (*State, error) {
	byt, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return NewStateFromJSON(byt, args...)
}

// TODO: functions to add to or merge into the existing state
// func Import
// func ImportFile
// func ImportJSON

func (jc *State) Get(keys ...string) interface{} {
	jc.mu.Lock()
	defer jc.mu.Unlock()
	return jc.Data.Get(keys...)
}

func (jc *State) Set(p ...interface{}) {
	jc.mu.Lock()
	defer jc.mu.Unlock()
	jc.Data.Set(p...)
}

// Load initializes the State object with data freshly loaded from the
// current Path() throwing away the reference to any previous Data
// (which will be cleaned up with normal garbage collection).
func (jc *State) Load() error {
	newjc, err := NewStateFromFile(jc.Path())
	if err != nil {
		return err
	}
	jc.Data = newjc.Data
	return nil
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

// Save checks the last modified time of the file and refuses to
// overwrite it if the Saved time is older and returns
// a NewerCacheError.
func (jc *State) Save() error {

	//fmt.Println(jc)

	if !jc.unsaved {
		//println("nothing to save")
		return nil
	}

	jc.mu.Lock()
	defer jc.mu.Unlock()

	// do we even have a file?
	i, err := os.Stat(jc.Path())
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		} else {
			//println("no path found for ", jc.Path())
			os.MkdirAll(jc.Dir, 0700)
		}
	} else {
		if !jc.Saved.IsZero() && i.ModTime().After(jc.Saved) {
			return NewerCacheError
		}
	}

	prevs := jc.Saved
	prevu := jc.unsaved

	jc.Saved = time.Now()
	jc.unsaved = false

	err = ioutil.WriteFile(jc.Path(), []byte(jc.String()), 0600)
	if err != nil {
		jc.Saved = prevs
		jc.unsaved = prevu
		return err
	}

	return nil
}

func (jc *State) ForceSave() error {
	jc.mu.Lock()
	defer jc.mu.Unlock()
	jc.Saved = time.Now()
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
