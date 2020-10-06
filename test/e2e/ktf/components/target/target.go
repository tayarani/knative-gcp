package target

import (
	"fmt"
	"io/ioutil"
	"log"

	"knative.dev/reconciler-test/pkg/config"
	"knative.dev/reconciler-test/pkg/framework"
	"knative.dev/reconciler-test/pkg/manifest"
)

type targetComponent struct {
}

// Component blah blah
var Component = &targetComponent{}

var _ framework.Component = (*targetComponent)(nil)

func (t *targetComponent) Scope() framework.ComponentScope {
	return framework.ComponentScopeCluster
}

func (t *targetComponent) Required(rc framework.ResourceContext, cfg config.Config) {
	fmt.Println("targetComponent::Required")

	fmt.Printf("TargetComponentConfig: %+v\n", config.GetConfig(cfg, "components/target"))
	scfg := config.GetConfig(cfg, "components/target").(Config)

	content, err := ioutil.ReadFile(scfg.PodYaml)
	if err != nil {
		log.Fatal(err)
	}

	yaml := manifest.FromString(string(content))
	rc.Apply(yaml, scfg)

	content, err = ioutil.ReadFile(scfg.ServiceYaml)
	if err != nil {
		log.Fatal(err)
	}

	yaml = manifest.FromString(string(content))
	rc.Apply(yaml, scfg)
}
