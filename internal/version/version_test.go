package version_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xmlking/go-workspace/internal/version"
)

func TestGetBuildInfo(t *testing.T) {
	assert := assert.New(t)

	expected := ""
	actual := ""

	t.Log(version.GetBuildInfo())
	//t.Logf("build_info:%s", version.GetBuildInfo().PrettyString())
	//t.Log(version.GetSoftwareBOM())

	assert.Equal(actual, expected)
}

/* Skip ExampleXXX tests from CI
func ExampleGetBuildInfo() {
	fmt.Println(version.GetBuildInfo())
	fmt.Println(version.GetSoftwareBOM())
	//fmt.Printf("build_info:%s", version.GetBuildInfo().PrettyString())

	// Output:
	//{"tag":"","commit":"","branch":"main","state":"dirty","build_time":"","go_version":"go1.18rc1","compiler":"gc","platform":"darwin/amd64"}
	//  mod
}
*/
