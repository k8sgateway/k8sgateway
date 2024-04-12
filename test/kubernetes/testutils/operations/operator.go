package operations

import (
	"context"
	"fmt"
	"io"
	"slices"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

// Operator is responsible for executing Operation against a Kubernetes Cluster
// This is meant to mirror the behavior of an Operator of the Gloo Gateway Product
// Although we operate against a Kubernetes Cluster, an Operator is intentionally
// unaware of Kubernetes behavior, and instead is more of a scheduler of Operation.
// This allows us to test its functionality, and also more easily inject behaviors
type Operator struct {
	progressWriter       io.Writer
	assertionInterceptor func(func()) error
}

// NewOperator returns an Operator
func NewOperator() *Operator {
	return &Operator{
		progressWriter: io.Discard,
		assertionInterceptor: func(f func()) error {
			// do nothing, assertions will bubble up and lead to a panic
			return nil
		},
	}
}

// NewGinkgoOperator returns an Operator used for the Ginkgo test framework
func NewGinkgoOperator() *Operator {
	return NewOperator().
		WithProgressWriter(ginkgo.GinkgoWriter).
		WithAssertionInterceptor(gomega.InterceptGomegaFailure)
}

// WithProgressWriter sets the io.Writer used by the Operator
func (o *Operator) WithProgressWriter(writer io.Writer) *Operator {
	o.progressWriter = writer
	return o
}

// WithAssertionInterceptor sets the function that will be used to intercept ScenarioAssertion failures
func (o *Operator) WithAssertionInterceptor(assertionInterceptor func(func()) error) *Operator {
	o.assertionInterceptor = assertionInterceptor
	return o
}

// ExecuteOperations executes a set of Operation
// NOTE: The Operator doesn't attempt to undo any of these Operation so if you are modifying
// resources on the Cluster, it is your responsibility to perform Operation to undo those changes
// If you would like to rely on this functionality, please see ExecuteReversibleOperations
func (o *Operator) ExecuteOperations(ctx context.Context, operations ...Operation) error {
	return o.executeSafe(func() error {
		return o.executeOperations(ctx, operations...)
	})
}

// ExecuteReversibleOperations executes a set of ReversibleOperation
// In order, the ReversibleOperation.Do will be executed, and then on success or failure
// the ReversibleOperation.Undo will also be executed
// This way, developers do not need to worry about resources being cleaned up appropriately in tests
func (o *Operator) ExecuteReversibleOperations(ctx context.Context, operations ...ReversibleOperation) error {
	return o.executeSafe(func() error {
		return o.executeReversibleOperations(ctx, operations...)
	})
}

func (o *Operator) executeSafe(fnMayPanic func() error) error {
	// Intercept failures, so that we can return an error to the test code,
	// and it can decide what to do with it
	var executionErr error
	interceptedErr := o.assertionInterceptor(func() {
		executionErr = fnMayPanic()
	})
	if interceptedErr != nil {
		return interceptedErr
	}
	return executionErr
}

func (o *Operator) executeOperations(ctx context.Context, operations ...Operation) error {
	for _, op := range operations {
		if err := o.executeOperation(ctx, op); err != nil {
			return err
		}
	}
	return nil
}

func (o *Operator) executeReversibleOperations(ctx context.Context, operations ...ReversibleOperation) (err error) {
	var undoOperations []Operation

	defer func() {
		// We need to perform the undo operations in reverse order
		// This way, if we execute: do-A -> do-B -> do-C
		// We should undo it by executing: undo-C -> undo-B -> undo-A
		slices.Reverse(undoOperations)
		undoErr := o.executeOperations(ctx, undoOperations...)
		if undoErr != nil {
			err = undoErr
		}
	}()

	for _, op := range operations {
		undoOperations = append(undoOperations, op.Undo)

		doErr := o.executeOperation(ctx, op.Do)
		if doErr != nil {
			return doErr
		}

	}
	return nil
}

func (o *Operator) executeOperation(ctx context.Context, operation Operation) error {
	o.writeProgress(operation, "starting operation")

	op := operation.Execute()
	o.writeProgress(operation, "executing operation")
	if err := op(ctx); err != nil {
		return err
	}

	o.writeProgress(operation, "asserting operation")
	assertion := operation.ExecutionAssertion()
	assertion(ctx)
	o.writeProgress(operation, "completing operation")
	return nil
}

func (o *Operator) writeProgress(operation Operation, progress string) {
	_, _ = o.progressWriter.Write([]byte(fmt.Sprintf("%s: %s\n", operation.Name(), progress)))
}
