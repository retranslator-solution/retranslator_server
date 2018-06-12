package configs

import "github.com/spf13/viper"

type StorageBackend string

const (
	BadgerBackend = StorageBackend("badger")
)

type BadgerConf struct {
	Path string `yaml:"path"`
}
type StorageConf struct {
	Backend StorageBackend `yaml:"backend"`
	Badger BadgerConf `yaml:"badger"`
}

type Config struct {
	Storage StorageConf `yaml:"storage"`
}


func NewConfig(v *viper.Viper) Config{
	var conf Config

	conf.Storage.Backend = StorageBackend(v.GetString("storage.backend"))
	conf.Storage.Badger.Path = v.GetString("storage.badger.path")

	return conf
}
