package monitor

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMonitor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Monitor Suite")
}
