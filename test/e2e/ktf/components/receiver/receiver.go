package receiver

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"knative.dev/reconciler-test/pkg/config"
	"knative.dev/reconciler-test/pkg/framework"
	"knative.dev/reconciler-test/pkg/installer"
	"knative.dev/reconciler-test/pkg/manifest"
)

const packageName = "github.com/google/knative-gcp/test/test_images/receiver"

// Deploy ...
func Deploy(rc framework.ResourceContext) corev1.ObjectReference {
	fmt.Println("receiverComponent::Deploy")

	rc.Apply(manifest.FromString(podTemplate))
	rc.Apply(manifest.FromString(serviceTemplate))

	return corev1.ObjectReference{
		Namespace: rc.Namespace(),
		Name:      "receiver",
	}
}

type receiverComponent struct {
}

// Component blah blah
var Component = &receiverComponent{}

var _ framework.Component = (*receiverComponent)(nil)

func (s *receiverComponent) Scope() framework.ComponentScope {
	return framework.ComponentScopeResource
}

func (s *receiverComponent) Required(rc framework.ResourceContext, cfg config.Config) {
	fmt.Println("receiverComponent::Required")
	installer.RegisterPackage(packageName)
}
