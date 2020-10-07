package sender

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"knative.dev/reconciler-test/pkg/config"
	"knative.dev/reconciler-test/pkg/framework"
	"knative.dev/reconciler-test/pkg/installer"
	"knative.dev/reconciler-test/pkg/manifest"
)

const packageName = "github.com/google/knative-gcp/test/test_images/sender"

// Deploy ...
func Deploy(rc framework.ResourceContext) corev1.ObjectReference {
	fmt.Println("senderComponent::Deploy")

	rc.Apply(manifest.FromString(podTemplate))
	rc.Apply(manifest.FromString(serviceTemplate))

	return corev1.ObjectReference{
		Namespace: rc.Namespace(),
		Name:      "sender",
	}
}

type senderComponent struct {
}

// Component blah blah
var Component = &senderComponent{}

var _ framework.Component = (*senderComponent)(nil)

func (s *senderComponent) Scope() framework.ComponentScope {
	return framework.ComponentScopeResource
}

func (s *senderComponent) Required(rc framework.ResourceContext, cfg config.Config) {
	fmt.Println("senderComponent::Required")
	installer.RegisterPackage(packageName)
}
