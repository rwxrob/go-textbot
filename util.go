package textbot

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path"
)

func HomeDotDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return path.Join(usr.HomeDir, "."+path.Base(os.Args[0]))
}

func Get(m map[string]interface{}, keys ...string) interface{} {

	if len(keys) == 1 {
		return m[keys[0]]
	}

	if len(keys) > 1 {
		i := m
		for n := 0; n < len(keys)-1; n++ {
			if newi, ok := i[keys[n]]; ok {
				i = newi.(map[string]interface{})
			} else {
				return nil
			}
		}
		return i[keys[len(keys)-1]]
	}

	if len(keys) < 1 {
		panic(MissingParams)
	}

	return nil
}

func Set(m map[string]interface{}, p ...interface{}) {

	if len(p) == 2 {
		m[p[0].(string)] = p[1]
		return
	}

	if len(p) > 2 {
		key := p[0].(string)
		if _, ok := m[key]; !ok {
			m[key] = map[string]interface{}{}
		}
		Set(m[key].(map[string]interface{}), p[1:]...)
	}

	if len(p) < 2 {
		panic(MissingParams)
	}
}

func SetDef(m map[string]interface{}, p ...interface{}) {
	keys := []string{}
	for _, k := range p[:len(p)-1] {
		keys = append(keys, k.(string))
	}
	cur := Get(m, keys...)
	if cur == nil {
		Set(m, p...)
	}
}

func JSON(m map[string]interface{}) string {
	byt, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return fmt.Sprintf("{\"ERROR\":\"%v\"}", err)
	}
	return string(byt)
}

func Print(m map[string]interface{}) {
	fmt.Print(JSON(m))
}
