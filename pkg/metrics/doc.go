/*
Package metrics provides a metrics wrapper.

Metrics must be initialized at startup as early as possible:
  import "github.com/urjitbhatia/goyaad/pkg/metrics"
  ...
  metrics.InitMetrics(metricsCollectorAddr)
*/
package metrics
