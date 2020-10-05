package components

import (
	"github.com/google/knative-gcp/test/e2e/ktf/components/receiver"
	"github.com/google/knative-gcp/test/e2e/ktf/components/sender"
)

// Config blah
type Config struct {
	Receiver receiver.Config
	Sender   sender.Config
}
