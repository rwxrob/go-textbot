package textbot

import (
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
