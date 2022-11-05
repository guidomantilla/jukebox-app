package properties

var _ PropertySource = (*DefaultPropertySource)(nil)

type PropertySource interface {
	Get(property string) string
	AsMap() map[string]any
}

// DefaultPropertySource

type DefaultPropertySource struct {
	name       string
	properties Properties
}

func (source *DefaultPropertySource) Get(property string) string {
	return source.properties.Get(property)
}

func (source *DefaultPropertySource) AsMap() map[string]any {

	internalMap := make(map[string]any)
	internalMap["name"], internalMap["value"] = source.name, source.properties.AsMap()
	return internalMap
}

//

func NewDefaultPropertySource(name string, properties Properties) *DefaultPropertySource {
	return &DefaultPropertySource{
		name:       name,
		properties: properties,
	}
}
