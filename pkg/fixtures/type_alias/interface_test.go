package type_alias_test

import (
	"regexp"
	"testing"

	"github.com/chigopher/pathlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTypeAlias(t *testing.T) {
	for _, tt := range []struct {
		name          string
		filepath      string
		expectedRegex string
	}{
		{
			name:          "With alias resolved",
			filepath:      "./mock_InterfaceWithResolvedAlias_test.go",
			expectedRegex: `func \((_?[a-zA-Z]*)+ \*InterfaceWithResolvedAlias\) Foo\(\) int {`,
		},
		{
			name:          "With alias unresolved",
			filepath:      "./mock_InterfaceWithUnresolvedAlias_test.go",
			expectedRegex: `func \((_?[a-zA-Z]*)+ \*InterfaceWithUnresolvedAlias\) Foo\(\) type_alias.Type {`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			regex, err := regexp.Compile(tt.expectedRegex)
			require.NoError(t, err)
			path := pathlib.NewPath(tt.filepath)
			bytes, err := path.ReadFile()
			require.NoError(t, err)

			assert.True(t, regex.Match(bytes), "expected regex was not found in file")
		})
	}
}
