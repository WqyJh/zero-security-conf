package confz

import (
	"log"
	"os"
	"reflect"

	"github.com/WqyJh/confcrypt"
	"github.com/zeromicro/go-zero/core/conf"
)

type SecurityConf struct {
	Enable bool   `json:",default=true"`
	Env    string `json:",default=CONFIG_KEY"` // environment variable name stores the encryption key
}

func findSecurityConfInStruct(o interface{}) (_ SecurityConf, success bool) {
	v, ok := recursive(reflect.ValueOf(o))
	if ok {
		return v.Interface().(SecurityConf), true
	}
	return SecurityConf{}, false
}

func SecurityLoad(path string, v interface{}, opts ...conf.Option) error {
	if err := conf.Load(path, v, opts...); err != nil {
		return err
	}
	c, ok := findSecurityConfInStruct(v)
	if ok && c.Enable {
		key := os.Getenv(c.Env)
		decoded, err := confcrypt.Decode(v, key)
		if err != nil {
			return err
		}
		if reflect.TypeOf(v).Kind() == reflect.Ptr {
			reflect.ValueOf(v).Elem().Set(reflect.ValueOf(decoded).Elem())
			return nil
		}
		reflect.ValueOf(v).Set(reflect.ValueOf(decoded))
	}
	return nil
}

func SecurityMustLoad(path string, v interface{}, opts ...conf.Option) {
	if err := SecurityLoad(path, v, opts...); err != nil {
		log.Fatalf("error: config file %s, %s", path, err.Error())
	}
}
