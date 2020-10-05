package ktfe2e

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
	"time"

	"knative.dev/reconciler-test/pkg/config"
	"knative.dev/reconciler-test/pkg/framework"
	"knative.dev/reconciler-test/pkg/manifest"
)

type Config struct {
	framework.BaseConfig
	BrokerName string `desc:"The name of the broker"`
	Myconfig   MyComponentConfig
}

var myconfig = Config{}

func TestMain(m *testing.M) {
	fmt.Println(">>>>>>>>> 1")
	framework.
		NewSuite(m).
		Configure(&myconfig).
		Require(Component).
		Run()
	fmt.Println(">>>>>>>>> 2")
}

func TestCase(t *testing.T) {
	fmt.Println(">>>>>>>>> 3")
	framework.NewTest(t).
		Feature("Named-Broker").
		Stable().
		//Must("Named-Broker").
		Run(func(tc framework.TestContext) {
			fmt.Println(">>>>>>>>> 4")
			fmt.Println("broker name is: " + myconfig.BrokerName)

			//ctx := context.Background()
			//fmt.Println(ctx)

			//component := MyComponent{}
			//fmt.Printf("%v", component)

			time.Sleep(60 * time.Second)

			fmt.Println(">>>>>>>>> 5")

		})
}

// ===========
// MyComponent
type MyComponent struct {
}

// MyComponentConfig
type MyComponentConfig struct {
	Jimmy string
}

var Component = &MyComponent{}
var _ framework.Component = (*MyComponent)(nil)

// Scope returns the component scope
func (s *MyComponent) Scope() framework.ComponentScope {
	fmt.Println("MyComponent::Scope")
	return framework.ComponentScopeCluster
}

// Required
func (s *MyComponent) Required(rc framework.ResourceContext, cfg config.Config) {
	fmt.Println("MyComponent::Required")

	fmt.Printf("MyComponentConfig: %+v\n", config.GetConfig(cfg, "myconfig"))
	mycfg := config.GetConfig(cfg, "myconfig").(MyComponentConfig)

	content, err := ioutil.ReadFile(mycfg.Jimmy)
	if err != nil {
		log.Fatal(err)
	}

	yaml := manifest.FromString(string(content))

	//rc.Logf("installing GitHubSource release ", ghcfg.Version)
	rc.Apply(yaml, mycfg)
}
