package forgeio

import (
	"errors"
	"fmt"
	"strings"

	"github.com/onsi/gomega/types"
)

type DummyWriterMatcher struct {
	Pathes []string
}

func ContainFiles(pathes ...string) types.GomegaMatcher {
	return &DummyWriterMatcher{
		Pathes: pathes,
	}
}

func (matcher *DummyWriterMatcher) Match(actual interface{}) (success bool, err error) {
	dummyWriter, ok := actual.(DummyWriter)
	if !ok {
		return false, errors.New("ContainFiles expects DummyWriter")
	}

	for _, filepath := range matcher.Pathes {
		_, ok := dummyWriter.files[filepath]
		if !ok {
			return false, nil
		}
	}

	return true, nil
}

func (matcher *DummyWriterMatcher) FailureMessage(actual interface{}) (message string) {
	dummyWriter := actual.(DummyWriter)

	have := make([]string, 0, len(dummyWriter.files))

	for key := range dummyWriter.files {
		have = append(have, key)
	}

	pathes := strings.Join(matcher.Pathes, ",")
	expected := strings.Join(have, ",")

	return fmt.Sprintf("Expexted:\n[%s]\n to contain files:\n[%s]", expected, pathes)
}

func (matcher *DummyWriterMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	dummyWriter := actual.(DummyWriter)

	have := make([]string, 0, len(dummyWriter.files))

	for key := range dummyWriter.files {
		have = append(have, key)
	}

	pathes := strings.Join(matcher.Pathes, ",")
	expected := strings.Join(have, ",")

	return fmt.Sprintf("Expexted:\n[%s]\n to not contain files:\n[%s]", expected, pathes)
}
