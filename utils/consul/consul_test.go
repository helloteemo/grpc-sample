package consul

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"reflect"
	"sync"
	"testing"
	"time"
	"zlutils/logger"
	"zlutils/request"
	zt "zlutils/time"
)

type Tmp struct {
	D *zt.Duration `json:"d" validate:"required"`
}

func TestWatchJson(t *testing.T) {
	Init(":8500", "tmp")
	var tmp Tmp
	ValiStruct().WatchJson("d", &tmp, func() {
		//panic(1)
		fmt.Println("change to", tmp)
	})
	fmt.Println(tmp) //{0s}
	select {}
}

func TestGetJson(t *testing.T) {
	Init(":8500", "test/service/counter")
	var tmp Tmp
	fmt.Println(TryGetJson("tmp", &tmp))
	GetJson("tmp", &tmp)
	fmt.Println(tmp)
}

func TestGetYaml(t *testing.T) {
	Init(":8500", "test/service/counter")
	var ty struct {
		I int               `yaml:"i" validate:"required"`
		S string            `yaml:"s" validate:"required"`
		M map[string]string `yaml:"m" validate:"required"`
		F float64           `yaml:"f" validate:"required"`
	}
	ValiStruct().GetYaml("t.yaml", &ty)
	fmt.Println(ty)
}

func TestGetJsonValiStruct(t *testing.T) {
	Init(":8500", "test/service/counter")
	var yoyo request.Config
	ValiStruct().GetJson("yoyo", &yoyo)
	var b int
	ValiVar("min=2").GetJson("b", &b)
}

func TestWatchJsonHandler(t *testing.T) {
	Init(":8500", "test/service/counter")
	ValiStruct().WatchJsonVarious("tmp", func(tmp Tmp) {
		fmt.Println(tmp)
	})
	select {}
}

func TestGetJsonHandler(t *testing.T) {
	Init(":8500", "test/service/counter")
	GetJson("log_watch", func(log logger.Config) {
		fmt.Println(log.Level)
	})
}

func TestWithPrefix(t *testing.T) {
	Init(":8500", "test/service/counter")
	lo := WithPrefix("test/service/example")
	lo.ValiVar("len=15").GetJson("eee", func(re string) {
		fmt.Println(re)
	})
	ValiStruct().GetJson("redis", func(redis struct {
		Url      string        `json:"url" validate:"url"`
		Duration time.Duration `json:"duration"`
	}) {
		fmt.Println(redis)
	})
}

func TestWatchJsonVariousVar(t *testing.T) {
	Init(":8500", "tmp")
	var i *int
	WatchJsonVarious("i", &i)
	for {
		fmt.Println(*i)
		time.Sleep(time.Second)
	}
}
func TestWatchJsonVariousFunc(t *testing.T) {
	Init(":8500", "tmp")
	WatchJsonVarious("i", func(i *int) {
		fmt.Println(*i)
	})
	select {}
}

func TestWatchWithLocker(t *testing.T) {
	Init(":8500", "tmp")
	mu := &sync.Mutex{}
	var i int
	WithLocker(mu).WatchJsonVarious("i", &i)
	go func() {
		for {
			mu.Lock()
			fmt.Println("访问i的这几秒中, consul的修改不会生效")
			time.Sleep(time.Second * 5)
			mu.Unlock()
			fmt.Println("观察i此时才发生变化")
			time.Sleep(time.Second * 5)
		}
	}()
	for {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func TestWatchYamlVarious(t *testing.T) {
	Init(":8500", "tmp")
	var v struct {
		B bool `yaml:"b" json:"b"`
	}
	WatchYamlVarious("v.yaml", &v)
	for {
		time.Sleep(time.Second * 2)
		fmt.Println(v)
	}
}

func un(bs []byte, unmarshal Unmarshal) {
	var v *struct { // 即使是var v struct, 下面的结果也一样
		B bool `yaml:"b" json:"b"`
	}
	t := reflect.TypeOf(v)

	fmt.Println("v", unmarshal(bs, &v), v)

	v2 := reflect.New(t).Interface()
	fmt.Println("v2", unmarshal(bs, &v2), v2, reflect.TypeOf(v2))
	// yaml会将类型变成了map[string]interface {}
	// json不会改变类型, 仍是struct

	v3 := reflect.New(t).Interface()
	fmt.Println("v3", unmarshal(bs, v3), v3, reflect.TypeOf(v3))
	// yaml/json都没将类型改变, 仍是struct
}

func TestYaml(_ *testing.T) {
	un([]byte(`b: true`), yaml.Unmarshal)
}

func TestJson(_ *testing.T) {
	un([]byte(`{"b":true}`), json.Unmarshal)
}
