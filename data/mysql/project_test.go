// +build integration

package mysql

import (
	"testing"
)

func TestProjectGet(t *testing.T) {
	project := ProjectModel{DB: connect(t)}
	_, err := project.Get(2147483647)

	if err.Error() != "data: no matching record found" {
		t.Error("Expected to get no matching record")
	}
}
