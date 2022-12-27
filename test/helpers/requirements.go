package helpers

import (
	"fmt"
	"os"
	"runtime"
	"strconv"

<<<<<<< HEAD
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/golang/protobuf/ptypes/wrappers"

=======
>>>>>>> master
	"github.com/hashicorp/go-multierror"
	"github.com/onsi/ginkgo"

	"k8s.io/apimachinery/pkg/util/sets"
)

const (
	// InvalidTestReqsEnvVar is used to define the behavior for running tests locally when the provided requirements
	// are not met. See ValidateRequirementsAndNotifyGinkgo for a detail of available behaviors
	InvalidTestReqsEnvVar = "INVALID_TEST_REQS"
)

// ValidateRequirementsAndNotifyGinkgo validates that the provided Requirements are met, and if they are not, uses
// the InvalidTestReqsEnvVar to determine how to proceed:
// Options are:
//	- `run`: Ignore any invalid requirements and execute the tests
//	- `skip`: Notify Ginkgo that the current spec was skipped
//	- `fail`: Notify Ginkgo that the current spec has failed [DEFAULT]
func ValidateRequirementsAndNotifyGinkgo(requirements ...Requirement) {
	err := ValidateRequirements(requirements)
	if err == nil {
		return
	}
	message := fmt.Sprintf("Test requirements not met: %v \n\n Consider using %s=skip to skip these tests", err, InvalidTestReqsEnvVar)
	switch os.Getenv(InvalidTestReqsEnvVar) {
	case "run":
		// ignore the error from validating requirements and let the tests proceed
		return

	case "skip":
		ginkgo.Skip(message)

	case "fail":
		fallthrough
	default:
		ginkgo.Fail(message)
	}
}

// ValidateRequirements returns an error if any of the Requirements are not met
func ValidateRequirements(requirements []Requirement) error {
	// default
	requiredConfiguration := &RequiredConfiguration{
		supportedOS:   sets.NewString(),
		supportedArch: sets.NewString(),
		reasons:       map[string]string{},
	}

	// apply requirements
	for _, requirement := range requirements {
		requirement(requiredConfiguration)
	}

	// perform validation
	return requiredConfiguration.Validate()
}

type RequiredConfiguration struct {
	supportedOS   sets.String
	supportedArch sets.String

	// Set of env variables which must be defined
	definedEnvVar []string

	// Set of env variables which must have a truthy value
	// Examples: "1", "t", "T", "true", "TRUE", "True"
	truthyEnvVar []string

	// User defined reasons for why particular environmental conditions are required
	reasons map[string]string
<<<<<<< HEAD

	// nil: no credentials required
	// false: credentials not defined
	// true: credentials defined
	awsCredentials *wrappers.BoolValue
=======
>>>>>>> master
}

// Validate returns an error is the RequiredConfiguration is not met
func (r RequiredConfiguration) Validate() error {
	var errs *multierror.Error

	errs = multierror.Append(
		errs,
		r.validateOS(),
		r.validateArch(),
		r.validateDefinedEnv(),
<<<<<<< HEAD
		r.validateTruthyEnv(),
		r.validateAwsCredentials())
=======
		r.validateTruthyEnv())
>>>>>>> master

	// If there are no errors, return
	if errs.ErrorOrNil() == nil {
		return nil
	}

	// If there are reasons defined, include them in the error message
	if len(r.reasons) > 0 {
		errs = multierror.Append(
			errs,
			fmt.Errorf("user defined reasons: %+v", r.reasons))
	}

	return errs.ErrorOrNil()
}

func (r RequiredConfiguration) validateOS() error {
	if r.supportedOS.Len() == 0 {
		// An empty set is considered to support all
		return nil
	}
	if r.supportedOS.Has(runtime.GOOS) {
		return nil
	}

	return fmt.Errorf("runtime os (%s), is not in supported set (%v)", runtime.GOOS, r.supportedOS.UnsortedList())
}

func (r RequiredConfiguration) validateArch() error {
	if r.supportedArch.Len() == 0 {
		// An empty set is considered to support all
		return nil
	}
	if r.supportedArch.Has(runtime.GOARCH) {
		return nil
	}

	return fmt.Errorf("runtime arch (%s), is not in supported set (%v)", runtime.GOARCH, r.supportedArch.UnsortedList())
}

func (r RequiredConfiguration) validateDefinedEnv() error {
	for _, env := range r.definedEnvVar {
		if _, found := os.LookupEnv(env); !found {
			return fmt.Errorf("env (%s) is not defined", env)
		}
	}
	return nil
}

func (r RequiredConfiguration) validateTruthyEnv() error {
	for _, env := range r.truthyEnvVar {
		envValue := os.Getenv(env)
		envBoolValue, _ := strconv.ParseBool(envValue)
		if !envBoolValue {
			return fmt.Errorf("env (%s) needs to be truthy, but is (%s)", env, envValue)
		}
	}
	return nil
}

<<<<<<< HEAD
func (r RequiredConfiguration) validateAwsCredentials() error {
	if r.awsCredentials == nil {
		// credentials not required
		return nil
	}

	if r.awsCredentials.GetValue() == true {
		// credentials required, but defined
		return nil
	}

	return fmt.Errorf("aws credentials needs to be set, but are not")
}

=======
>>>>>>> master
// Requirement represents a required property for tests.
type Requirement func(configuration *RequiredConfiguration)

func LinuxOnly(reason string) Requirement {
	return func(configuration *RequiredConfiguration) {
		configuration.supportedOS = sets.NewString("linux")
		configuration.reasons["linux"] = reason
	}
}

func DefinedEnv(env string) Requirement {
	return func(configuration *RequiredConfiguration) {
		configuration.definedEnvVar = append(configuration.definedEnvVar, env)
	}
}

func TruthyEnv(env string) Requirement {
	return func(configuration *RequiredConfiguration) {
		configuration.truthyEnvVar = append(configuration.truthyEnvVar, env)
	}
}

func Kubernetes(reason string) Requirement {
	return func(configuration *RequiredConfiguration) {
		configuration.reasons["kubernetes"] = reason
		TruthyEnv("RUN_KUBE_TESTS")(configuration)
	}
}

func Consul() Requirement {
	return func(configuration *RequiredConfiguration) {
		TruthyEnv("RUN_CONSUL_TESTS")(configuration)
	}
}

func Vault() Requirement {
	return func(configuration *RequiredConfiguration) {
		TruthyEnv("RUN_VAULT_TESTS")(configuration)
	}
}
<<<<<<< HEAD

func Aws() Requirement {
	return func(configuration *RequiredConfiguration) {
		_, err := credentials.NewSharedCredentials("", "").Get()
		configuration.awsCredentials = &wrappers.BoolValue{
			Value: err == nil,
		}

		DefinedEnv("AWS_ARN_ROLE_1")(configuration)
	}
}
=======
>>>>>>> master
