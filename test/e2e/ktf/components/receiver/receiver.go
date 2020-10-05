package receiver

import (
	"fmt"
	"io/ioutil"
	"log"

	"knative.dev/reconciler-test/pkg/config"
	"knative.dev/reconciler-test/pkg/framework"
	"knative.dev/reconciler-test/pkg/manifest"
)

// Receiver
type receiverComponent struct {
}

// Component blah blah
var Component = &receiverComponent{}

var _ framework.Component = (*receiverComponent)(nil)

func (s *receiverComponent) Scope() framework.ComponentScope {
	return framework.ComponentScopeCluster
}

func (s *receiverComponent) Required(rc framework.ResourceContext, cfg config.Config) {
	fmt.Println("receiverComponent::Required")

	fmt.Printf("ReceiverComponentConfig: %+v\n", config.GetConfig(cfg, "components/receiver"))
	rcfg := config.GetConfig(cfg, "components/receiver").(Config)

	content, err := ioutil.ReadFile(rcfg.YamlPath)
	if err != nil {
		log.Fatal(err)
	}

	yaml := manifest.FromString(string(content))
	rc.Apply(yaml, rcfg)
}
