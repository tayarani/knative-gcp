package ktfe2e

import (
	"fmt"
	"testing"

	"github.com/google/knative-gcp/test/e2e/ktf/components"
	"github.com/google/knative-gcp/test/e2e/ktf/components/receiver"
	"github.com/google/knative-gcp/test/e2e/ktf/components/sender"

	"knative.dev/reconciler-test/pkg/framework"
)

type Config struct {
	framework.BaseConfig
	BrokerName string `desc:"The name of the broker"`
	Components components.Config
}

var myconfig = Config{}

func TestMain(m *testing.M) {
	fmt.Println(">>>>>>>>> 1")
	framework.
		NewSuite(m).
		Configure(&myconfig).
		Require(receiver.Component).
		Require(sender.Component).
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

			//time.Sleep(60 * time.Second)

			fmt.Println(">>>>>>>>> 5")

		})
}
