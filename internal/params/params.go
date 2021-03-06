package params

import (
	libparams "github.com/qdm12/golibs/params"
)

// GetDataDir obtains the data directory from the environment
// variable DATADIR
func GetDataDir(e libparams.EnvParams, dir string) (string, error) {
	return e.GetEnv("DATADIR", libparams.Default(dir+"/data"))
}
