package manifest

import "testing"

func TestFailure(test *testing.T) {
	value, error := NewManifestColumnItem(0)
	if error == nil {
		test.Errorf("Expecting error!")
	}
	if value != nil {
		test.Errorf("Expecting no value!")
	}
}
func TestSuccess(test *testing.T) {
	value, error := NewManifestColumnItem(2803282938)
	if error != nil {
		test.Errorf("Expecting no error here!")
	}
	if value == nil {
		test.Errorf("Expecting some value here!")
	}
	if value.ID != -1491684358 {
		test.Errorf("Incorrectly converted!")
	}
}
