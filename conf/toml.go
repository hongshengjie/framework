package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// Get Get
func Get(fileName string, dst interface{}) error {
	v, ok := client.fileNameValue[fileName]
	if !ok {
		return fmt.Errorf("not exist file %s", fileName)
	}
	return toml.Unmarshal([]byte(v), dst)
}
