package runtime

import (
	"github.com/solo-io/gloo/test/testutils"
	"os"
)

// Context contains the set of properties that are defined at runtime by whoever is invoking tests.
// The intention of this is two-fold:
//  1. To provide a transparent definition for the set of runtime inputs that are accepted.
//  2. To ensure that tests are not aware of _how_ inputs are provided (command line, env variable), but
//     just are aware _that_ they exist
type Context struct {

	// ClusterName is the name of the cluster that will be used to execute the tests in
	ClusterName string

	ImageVariant string

	// RunSource identifies who/what triggered the test
	RunSource RunSource
}

func NewContext() Context {
	return Context{
		// ClusterName is derived from the environment variable
		ClusterName: os.Getenv(testutils.ClusterName),
	}
}
