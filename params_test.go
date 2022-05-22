package errorist_test

import (
	"errors"
	"fmt"
	"github.com/canmor/errorist"
	"github.com/stretchr/testify/assert"
	"testing"
)

const NoPermission = "no permission"
const FolderNameConflict = "folder name conflict"

func TestWrapParams(t *testing.T) {
	var err = errors.New(NoPermission)

	params := []interface{}{"param1", "param2"}
	errWithParams := errorist.WrapParams(err, params...)

	assert.ErrorIs(t, errWithParams, err)
	assert.Equal(t, params, errorist.UnwrapParams(errWithParams))
	assert.Equal(t, fmt.Sprintf(NoPermission+", params %v", params), errWithParams.Error())
}

func ExampleWrapParams_errorsIs() {
	var ErrNoPermission = errors.New(NoPermission)
	var ErrFolderNameConflict = errors.New(FolderNameConflict)

	err := errorist.WrapParams(ErrNoPermission, "param 1", "param 2")

	if errors.Is(err, ErrNoPermission) {
		fmt.Printf("no permission, params: %+v\n", errorist.UnwrapParams(err))
	} else if errors.Is(err, ErrFolderNameConflict) {
		fmt.Printf("folder name conflict: %s\n", err)
	} else if err != nil {
		fmt.Println("internal error")
	}
}
