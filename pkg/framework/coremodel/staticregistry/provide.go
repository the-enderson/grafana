package staticregistry

import (
	"github.com/grafana/grafana/pkg/framework/coremodel"
	"github.com/grafana/thema"
)

// ProvideExplicitRegistry provides access to individual coremodels via explicit method calls.
//
// Prefer this to the generic ProvideRegistry type when your code works with known,
// specific coremodels(s), rather than generically across all of them. This allows
// standard Go static analysis tools to determine which code is depending on
// particular coremodels.
//
// This will use the default Grafana thema.Library, defined in pkg/cuectx, which
// will avoid duplicate parsing of Thema CUE schemas. If you need control over the
// thema.Library in use, use ProvideExplicitRegistryWithLib instead.
func ProvideExplicitRegistry() (ExplicitRegistry, error) {
	return provideExplicitRegistry(nil)
}

// ProvideExplicitRegistryWithLib is the same as ProvideExplicitRegistry, but allows control
// over the thema.Library used to initialize the underlying
//
// Prefer ProvideExplicitRegistry unless you absolutely need this control.
func ProvideExplicitRegistryWithLib(lib thema.Library) (ExplicitRegistry, error) {
	return provideExplicitRegistry(&lib)
}

// ProvideRegistry provides a simple static Registry for coremodels.
//
// Prefer this to the static ProvideExplicitRegistry when your code needs to work with
// all coremodels generically, rather than known, specific coremodel(s).
func ProvideRegistry() (*coremodel.Registry, error) {
	return provideRegistry()
}
