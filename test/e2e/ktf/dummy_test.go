package ktfe2e

import (
	"context"
	"fmt"
	"testing"

	//"github.com/google/knative-gcp/test/e2e/ktf/my_component"

	"knative.dev/reconciler-test/pkg/config"
	"knative.dev/reconciler-test/pkg/framework"
)

type Config struct {
	framework.BaseConfig
	BrokerName string `desc:"The name of the broker"`
}

var (
	//channelTestRunner eventingtestlib.ComponentsTestRunner
	//authConfig lib.AuthConfig
	projectID string
)
var cfg = Config{}

func TestMain(m *testing.M) {
	fmt.Println(">>>>>>>>> 1")
	framework.
		NewSuite(m).
		Configure(&cfg).
		Run()
	fmt.Println(">>>>>>>>> 2")
}

func TestCase(t *testing.T) {
	fmt.Println(">>>>>>>>> 3")
	framework.NewTest(t).
		Feature("Named-Broker").
		Stable().
		Must("Named-Broker").
		//Require(Component).
		Run(func(tc framework.TestContext) {
			fmt.Println(">>>>>>>>> 4")
			fmt.Println("broker name is: " + cfg.BrokerName)

			ctx := context.Background()
			fmt.Println(ctx)

			component := MyComponent{}
			fmt.Printf("%v", component)

			fmt.Println(">>>>>>>>> 5")

		})
}

// ===========
// MyComponent
type MyComponent struct {
}

// MyComponentConfig
type MyComponentConfig struct {
}

var _ framework.Component = (*MyComponent)(nil)

// Scope returns the component scope
func (s *MyComponent) Scope() framework.ComponentScope {
	return framework.ComponentScopeCluster
}

// Required
func (s *MyComponent) Required(rc framework.ResourceContext, cfg config.Config) {
	//ghcfg := config.GetConfig(cfg, "test/e2e/ktf").(MyComponentConfig)

	//rc.KoApply(ghcfg.Path, "test/e2e/ktf/config/")
}
