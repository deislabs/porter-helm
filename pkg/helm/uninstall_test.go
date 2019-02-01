package helm

import (
	"bytes"
	"os"
	"testing"

	"github.com/deislabs/porter/pkg/test"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

type UninstallTest struct {
	expectedCommand string
	uninstallStep   UninstallStep
}

func TestMixin_Uninstall(t *testing.T) {
	releases := []string{
		"foo",
		"bar",
	}

	uninstallTests := []UninstallTest{
		UninstallTest{
			expectedCommand: `helm delete foo bar`,
			uninstallStep: UninstallStep{
				Arguments: UninstallArguments{
					Releases: releases,
				},
			},
		},
		UninstallTest{
			expectedCommand: `helm delete --purge foo bar`,
			uninstallStep: UninstallStep{
				Arguments: UninstallArguments{
					Purge:    true,
					Releases: releases,
				},
			},
		},
	}

	for _, uninstallTest := range uninstallTests {
		os.Setenv(test.ExpectedCommandEnv, uninstallTest.expectedCommand)
		defer os.Unsetenv(test.ExpectedCommandEnv)

		b, _ := yaml.Marshal(uninstallTest.uninstallStep)

		h := NewTestMixin(t)
		h.In = bytes.NewReader(b)

		err := h.Uninstall()

		require.NoError(t, err)
	}
}