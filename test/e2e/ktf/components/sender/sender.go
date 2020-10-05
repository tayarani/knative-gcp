package sender

import (
	"fmt"
	"io/ioutil"
	"log"

	"knative.dev/reconciler-test/pkg/config"
	"knative.dev/reconciler-test/pkg/framework"
	"knative.dev/reconciler-test/pkg/manifest"
)

type senderComponent struct {
}

// Component blah blah
var Component = &senderComponent{}

var _ framework.Component = (*senderComponent)(nil)

func (s *senderComponent) Scope() framework.ComponentScope {
	return framework.ComponentScopeCluster
}

func (s *senderComponent) Required(rc framework.ResourceContext, cfg config.Config) {
	fmt.Println("senderComponent::Required")

	fmt.Printf("SenderComponentConfig: %+v\n", config.GetConfig(cfg, "components/sender"))
	scfg := config.GetConfig(cfg, "components/sender").(Config)

	content, err := ioutil.ReadFile(scfg.YamlPath)
	if err != nil {
		log.Fatal(err)
	}

	yaml := manifest.FromString(string(content))
	rc.Apply(yaml, scfg)
}
