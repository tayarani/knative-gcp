package target

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"knative.dev/pkg/test/helpers"
	"knative.dev/reconciler-test/pkg/config"
	"knative.dev/reconciler-test/pkg/framework"
	"knative.dev/reconciler-test/pkg/installer"
	"knative.dev/reconciler-test/pkg/manifest"
)

const packageName = "github.com/google/knative-gcp/test/test_images/target"

// Deploy ...
func Deploy(rc framework.ResourceContext) corev1.ObjectReference {
	fmt.Println("targetComponent::Deploy")

	image := rc.ImageName(packageName)
	name := helpers.AppendRandomString("target")

	data := struct {
		Name  string
		Image string
	}{
		Name:  name,
		Image: image,
	}

	rc.Apply(manifest.FromString(podTemplate), data)
	rc.Apply(manifest.FromString(serviceTemplate), data)

	return corev1.ObjectReference{
		Namespace: rc.Namespace(),
		Name:      name,
	}
}

type targetComponent struct {
}

// Component blah blah
var Component = &targetComponent{}

var _ framework.Component = (*targetComponent)(nil)

func (t *targetComponent) Scope() framework.ComponentScope {
	return framework.ComponentScopeResource
}

func (t *targetComponent) Required(rc framework.ResourceContext, cfg config.Config) {
	fmt.Println("targetComponent::Required")
	installer.RegisterPackage(packageName)
}
