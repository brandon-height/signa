package deployment

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestTriggerE2EPipeline(t *testing.T) {

	triggerE2EPipeline = func(path string) ([]byte, error) {
		b := []byte{123, 10, 32, 32, 34, 99, 111, 109, 112, 97, 114, 101, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 112, 114, 101, 118, 105, 111, 117, 115, 95, 115, 117, 99, 99, 101, 115, 115, 102, 117, 108, 95, 98, 117, 105, 108, 100, 34, 32, 58, 32, 123, 10, 32, 32, 32, 32, 34, 98, 117, 105, 108, 100, 95, 110, 117, 109, 34, 32, 58, 32, 51, 54, 55, 44, 10, 32, 32, 32, 32, 34, 115, 116, 97, 116, 117, 115, 34, 32, 58, 32, 34, 115, 117, 99, 99, 101, 115, 115, 34, 44, 10, 32, 32, 32, 32, 34, 98, 117, 105, 108, 100, 95, 116, 105, 109, 101, 95, 109, 105, 108, 108, 105, 115, 34, 32, 58, 32, 55, 48, 50, 57, 55, 10, 32, 32, 125, 44, 10, 32, 32, 34, 98, 117, 105, 108, 100, 95, 112, 97, 114, 97, 109, 101, 116, 101, 114, 115, 34, 32, 58, 32, 123, 10, 32, 32, 32, 32, 34, 67, 73, 82, 67, 76, 69, 95, 74, 79, 66, 34, 32, 58, 32, 34, 112, 114, 111, 100, 117, 99, 116, 105, 111, 110, 45, 98, 117, 105, 108, 100, 34, 10, 32, 32, 125, 44, 10, 32, 32, 34, 111, 115, 115, 34, 32, 58, 32, 102, 97, 108, 115, 101, 44, 10, 32, 32, 34, 99, 111, 109, 109, 105, 116, 116, 101, 114, 95, 100, 97, 116, 101, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 98, 111, 100, 121, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 117, 115, 97, 103, 101, 95, 113, 117, 101, 117, 101, 100, 95, 97, 116, 34, 32, 58, 32, 34, 50, 48, 49, 56, 45, 49, 48, 45, 50, 50, 84, 49, 52, 58, 51, 53, 58, 51, 51, 46, 52, 49, 57, 90, 34, 44, 10, 32, 32, 34, 102, 97, 105, 108, 95, 114, 101, 97, 115, 111, 110, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 114, 101, 116, 114, 121, 95, 111, 102, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 114, 101, 112, 111, 110, 97, 109, 101, 34, 32, 58, 32, 34, 112, 101, 120, 45, 101, 50, 101, 45, 116, 101, 115, 116, 105, 110, 103, 34, 44, 10, 32, 32, 34, 115, 115, 104, 95, 117, 115, 101, 114, 115, 34, 32, 58, 32, 91, 32, 93, 44, 10, 32, 32, 34, 98, 117, 105, 108, 100, 95, 117, 114, 108, 34, 32, 58, 32, 34, 104, 116, 116, 112, 115, 58, 47, 47, 99, 105, 114, 99, 108, 101, 99, 105, 46, 99, 111, 109, 47, 103, 104, 47, 115, 105, 103, 110, 97, 118, 105, 111, 47, 112, 101, 120, 45, 101, 50, 101, 45, 116, 101, 115, 116, 105, 110, 103, 47, 51, 54, 56, 34, 44, 10, 32, 32, 34, 112, 97, 114, 97, 108, 108, 101, 108, 34, 32, 58, 32, 49, 44, 10, 32, 32, 34, 102, 97, 105, 108, 101, 100, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 98, 114, 97, 110, 99, 104, 34, 32, 58, 32, 34, 109, 97, 115, 116, 101, 114, 34, 44, 10, 32, 32, 34, 117, 115, 101, 114, 110, 97, 109, 101, 34, 32, 58, 32, 34, 115, 105, 103, 110, 97, 118, 105, 111, 34, 44, 10, 32, 32, 34, 97, 117, 116, 104, 111, 114, 95, 100, 97, 116, 101, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 119, 104, 121, 34, 32, 58, 32, 34, 97, 112, 105, 34, 44, 10, 32, 32, 34, 117, 115, 101, 114, 34, 32, 58, 32, 123, 10, 32, 32, 32, 32, 34, 105, 115, 95, 117, 115, 101, 114, 34, 32, 58, 32, 116, 114, 117, 101, 44, 10, 32, 32, 32, 32, 34, 108, 111, 103, 105, 110, 34, 32, 58, 32, 34, 112, 97, 115, 99, 97, 108, 105, 115, 109, 34, 44, 10, 32, 32, 32, 32, 34, 97, 118, 97, 116, 97, 114, 95, 117, 114, 108, 34, 32, 58, 32, 34, 104, 116, 116, 112, 115, 58, 47, 47, 97, 118, 97, 116, 97, 114, 115, 49, 46, 103, 105, 116, 104, 117, 98, 117, 115, 101, 114, 99, 111, 110, 116, 101, 110, 116, 46, 99, 111, 109, 47, 117, 47, 51, 57, 48, 52, 55, 56, 51, 52, 63, 118, 61, 52, 34, 44, 10, 32, 32, 32, 32, 34, 110, 97, 109, 101, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 32, 32, 34, 118, 99, 115, 95, 116, 121, 112, 101, 34, 32, 58, 32, 34, 103, 105, 116, 104, 117, 98, 34, 44, 10, 32, 32, 32, 32, 34, 105, 100, 34, 32, 58, 32, 51, 57, 48, 52, 55, 56, 51, 52, 10, 32, 32, 125, 44, 10, 32, 32, 34, 118, 99, 115, 95, 114, 101, 118, 105, 115, 105, 111, 110, 34, 32, 58, 32, 34, 56, 100, 50, 49, 57, 56, 50, 48, 51, 57, 50, 54, 57, 97, 100, 52, 98, 97, 99, 98, 48, 50, 102, 52, 48, 101, 51, 49, 101, 50, 52, 57, 54, 98, 101, 49, 99, 98, 55, 48, 34, 44, 10, 32, 32, 34, 118, 99, 115, 95, 116, 97, 103, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 98, 117, 105, 108, 100, 95, 110, 117, 109, 34, 32, 58, 32, 51, 54, 56, 44, 10, 32, 32, 34, 105, 110, 102, 114, 97, 115, 116, 114, 117, 99, 116, 117, 114, 101, 95, 102, 97, 105, 108, 34, 32, 58, 32, 102, 97, 108, 115, 101, 44, 10, 32, 32, 34, 99, 111, 109, 109, 105, 116, 116, 101, 114, 95, 101, 109, 97, 105, 108, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 112, 114, 101, 118, 105, 111, 117, 115, 34, 32, 58, 32, 123, 10, 32, 32, 32, 32, 34, 98, 117, 105, 108, 100, 95, 110, 117, 109, 34, 32, 58, 32, 51, 54, 55, 44, 10, 32, 32, 32, 32, 34, 115, 116, 97, 116, 117, 115, 34, 32, 58, 32, 34, 115, 117, 99, 99, 101, 115, 115, 34, 44, 10, 32, 32, 32, 32, 34, 98, 117, 105, 108, 100, 95, 116, 105, 109, 101, 95, 109, 105, 108, 108, 105, 115, 34, 32, 58, 32, 55, 48, 50, 57, 55, 10, 32, 32, 125, 44, 10, 32, 32, 34, 115, 116, 97, 116, 117, 115, 34, 32, 58, 32, 34, 110, 111, 116, 95, 114, 117, 110, 110, 105, 110, 103, 34, 44, 10, 32, 32, 34, 99, 111, 109, 109, 105, 116, 116, 101, 114, 95, 110, 97, 109, 101, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 114, 101, 116, 114, 105, 101, 115, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 115, 117, 98, 106, 101, 99, 116, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 118, 99, 115, 95, 116, 121, 112, 101, 34, 32, 58, 32, 34, 103, 105, 116, 104, 117, 98, 34, 44, 10, 32, 32, 34, 116, 105, 109, 101, 100, 111, 117, 116, 34, 32, 58, 32, 102, 97, 108, 115, 101, 44, 10, 32, 32, 34, 100, 111, 110, 116, 95, 98, 117, 105, 108, 100, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 108, 105, 102, 101, 99, 121, 99, 108, 101, 34, 32, 58, 32, 34, 110, 111, 116, 95, 114, 117, 110, 110, 105, 110, 103, 34, 44, 10, 32, 32, 34, 110, 111, 95, 100, 101, 112, 101, 110, 100, 101, 110, 99, 121, 95, 99, 97, 99, 104, 101, 34, 32, 58, 32, 102, 97, 108, 115, 101, 44, 10, 32, 32, 34, 115, 116, 111, 112, 95, 116, 105, 109, 101, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 115, 115, 104, 95, 100, 105, 115, 97, 98, 108, 101, 100, 34, 32, 58, 32, 116, 114, 117, 101, 44, 10, 32, 32, 34, 98, 117, 105, 108, 100, 95, 116, 105, 109, 101, 95, 109, 105, 108, 108, 105, 115, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 112, 105, 99, 97, 114, 100, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 99, 105, 114, 99, 108, 101, 95, 121, 109, 108, 34, 32, 58, 32, 123, 10, 32, 32, 32, 32, 34, 115, 116, 114, 105, 110, 103, 34, 32, 58, 32, 34, 118, 101, 114, 115, 105, 111, 110, 58, 32, 50, 92, 110, 92, 110, 106, 111, 98, 115, 58, 92, 110, 32, 32, 98, 117, 105, 108, 100, 58, 92, 110, 32, 32, 32, 32, 100, 111, 99, 107, 101, 114, 58, 92, 110, 32, 32, 32, 32, 32, 32, 45, 32, 105, 109, 97, 103, 101, 58, 32, 99, 121, 112, 114, 101, 115, 115, 47, 98, 97, 115, 101, 58, 56, 92, 110, 32, 32, 32, 32, 32, 32, 32, 32, 101, 110, 118, 105, 114, 111, 110, 109, 101, 110, 116, 58, 92, 110, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 84, 69, 82, 77, 58, 32, 120, 116, 101, 114, 109, 92, 110, 32, 32, 32, 32, 112, 97, 114, 97, 108, 108, 101, 108, 105, 115, 109, 58, 32, 49, 92, 110, 32, 32, 32, 32, 115, 116, 101, 112, 115, 58, 92, 110, 32, 32, 32, 32, 32, 32, 45, 32, 99, 104, 101, 99, 107, 111, 117, 116, 92, 110, 32, 32, 32, 32, 32, 32, 45, 32, 114, 117, 110, 58, 32, 110, 112, 109, 32, 99, 105, 92, 110, 32, 32, 32, 32, 32, 32, 45, 32, 114, 117, 110, 58, 92, 110, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 110, 97, 109, 101, 58, 32, 82, 117, 110, 110, 105, 110, 103, 32, 69, 50, 69, 32, 116, 101, 115, 116, 115, 32, 111, 110, 32, 115, 116, 97, 103, 105, 110, 103, 92, 110, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 99, 111, 109, 109, 97, 110, 100, 58, 32, 121, 97, 114, 110, 32, 99, 121, 112, 114, 101, 115, 115, 58, 114, 117, 110, 115, 116, 97, 103, 105, 110, 103, 92, 110, 32, 32, 112, 114, 111, 100, 117, 99, 116, 105, 111, 110, 45, 98, 117, 105, 108, 100, 58, 92, 110, 32, 32, 32, 32, 100, 111, 99, 107, 101, 114, 58, 92, 110, 32, 32, 32, 32, 32, 32, 45, 32, 105, 109, 97, 103, 101, 58, 32, 99, 121, 112, 114, 101, 115, 115, 47, 98, 97, 115, 101, 58, 56, 92, 110, 32, 32, 32, 32, 32, 32, 32, 32, 101, 110, 118, 105, 114, 111, 110, 109, 101, 110, 116, 58, 92, 110, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 84, 69, 82, 77, 58, 32, 120, 116, 101, 114, 109, 92, 110, 32, 32, 32, 32, 112, 97, 114, 97, 108, 108, 101, 108, 105, 115, 109, 58, 32, 49, 92, 110, 32, 32, 32, 32, 115, 116, 101, 112, 115, 58, 92, 110, 32, 32, 32, 32, 32, 32, 45, 32, 99, 104, 101, 99, 107, 111, 117, 116, 92, 110, 32, 32, 32, 32, 32, 32, 45, 32, 114, 117, 110, 58, 32, 110, 112, 109, 32, 99, 105, 92, 110, 32, 32, 32, 32, 32, 32, 45, 32, 114, 117, 110, 58, 92, 110, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 110, 97, 109, 101, 58, 32, 82, 117, 110, 110, 105, 110, 103, 32, 69, 50, 69, 32, 116, 101, 115, 116, 115, 32, 111, 110, 32, 112, 114, 111, 100, 117, 99, 116, 105, 111, 110, 92, 110, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 99, 111, 109, 109, 97, 110, 100, 58, 32, 121, 97, 114, 110, 32, 99, 121, 112, 114, 101, 115, 115, 58, 114, 117, 110, 112, 114, 111, 100, 117, 99, 116, 105, 111, 110, 105, 110, 118, 92, 110, 92, 110, 119, 111, 114, 107, 102, 108, 111, 119, 115, 58, 92, 110, 32, 32, 118, 101, 114, 115, 105, 111, 110, 58, 32, 50, 92, 110, 32, 32, 98, 117, 105, 108, 100, 45, 111, 110, 45, 115, 116, 97, 103, 105, 110, 103, 58, 92, 110, 32, 32, 32, 32, 106, 111, 98, 115, 58, 92, 110, 32, 32, 32, 32, 45, 32, 98, 117, 105, 108, 100, 58, 92, 110, 32, 32, 32, 32, 32, 32, 32, 32, 102, 105, 108, 116, 101, 114, 115, 58, 92, 110, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 98, 114, 97, 110, 99, 104, 101, 115, 58, 92, 110, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 111, 110, 108, 121, 58, 32, 109, 97, 115, 116, 101, 114, 92, 110, 34, 10, 32, 32, 125, 44, 10, 32, 32, 34, 109, 101, 115, 115, 97, 103, 101, 115, 34, 32, 58, 32, 91, 32, 93, 44, 10, 32, 32, 34, 105, 115, 95, 102, 105, 114, 115, 116, 95, 103, 114, 101, 101, 110, 95, 98, 117, 105, 108, 100, 34, 32, 58, 32, 102, 97, 108, 115, 101, 44, 10, 32, 32, 34, 106, 111, 98, 95, 110, 97, 109, 101, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 115, 116, 97, 114, 116, 95, 116, 105, 109, 101, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 99, 97, 110, 99, 101, 108, 101, 114, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 112, 108, 97, 116, 102, 111, 114, 109, 34, 32, 58, 32, 34, 50, 46, 48, 34, 44, 10, 32, 32, 34, 111, 117, 116, 99, 111, 109, 101, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 118, 99, 115, 95, 117, 114, 108, 34, 32, 58, 32, 34, 104, 116, 116, 112, 115, 58, 47, 47, 103, 105, 116, 104, 117, 98, 46, 99, 111, 109, 47, 115, 105, 103, 110, 97, 118, 105, 111, 47, 112, 101, 120, 45, 101, 50, 101, 45, 116, 101, 115, 116, 105, 110, 103, 34, 44, 10, 32, 32, 34, 97, 117, 116, 104, 111, 114, 95, 110, 97, 109, 101, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 110, 111, 100, 101, 34, 32, 58, 32, 110, 117, 108, 108, 44, 10, 32, 32, 34, 99, 97, 110, 99, 101, 108, 101, 100, 34, 32, 58, 32, 102, 97, 108, 115, 101, 44, 10, 32, 32, 34, 97, 117, 116, 104, 111, 114, 95, 101, 109, 97, 105, 108, 34, 32, 58, 32, 110, 117, 108, 108, 10, 125}
		return b, nil
	}

	outputCircle, error := triggerE2EPipeline("curl -u $CIRCLECI_Token: -d build_parameters[CIRCLE_JOB]=production-build https://circleci.com/api/v1.1/project/github/signavio/pex-e2e-testing/tree/master")

	formattedOutput := []byte(fmt.Sprintf("%s\n", outputCircle))
	pipelineStarted := repoNameTestChecker(formattedOutput)

	if pipelineStarted != true {
		t.Fatal("E2E pipeline is not triggered")
	}
	if error != nil {
		t.Fatal("Error from CircleCI")
	}

}

func repoNameTestChecker(output []byte) bool {
	var f interface{}

	json.Unmarshal(output, &f)
	m := f.(map[string]interface{})

	x := false

	for key, value := range m {

		if strings.Compare(key, "reponame") == 0 {

			if strings.Compare(value.(string), "pex-e2e-testing") == 0 {
				x = true
			}
		}

	}
	return x
}