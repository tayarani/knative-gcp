package ktfe2e

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/knative-gcp/test/e2e/ktf/components"
	"github.com/google/knative-gcp/test/e2e/ktf/components/target"

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
		//Require(receiver.Component).
		//Require(sender.Component).
		Require(target.Component).
		Run()
	fmt.Println(">>>>>>>>> 2")
}

func TestCase(t *testing.T) {
	fmt.Println(">>>>>>>>> 3")
	framework.NewTest(t).
		Feature("Named-Broker").
		Stable().
		Run(func(tc framework.TestContext) {
			fmt.Println(">>>>>>>>> 4")

			target.Deploy(tc)

			time.Sleep(60 * time.Second)

			fmt.Println(">>>>>>>>> 5")
		})
}
