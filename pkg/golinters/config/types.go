package config

import "gopkg.in/yaml.v3"

type StringSet map[string]struct{}

func (s *StringSet) UnmarshalYAML(value *yaml.Node) error {
	var items []string
	if err := value.Decode(&items); err != nil {
		return err
	}
	*s = make(StringSet, len(items))
	for _, item := range items {
		(*s)[item] = struct{}{}
	}
	return nil
}

func (s *StringSet) MarshalYAML() (interface{}, error) {
	items := make([]string, 0, len(*s))
	for k := range *s {
		items = append(items, k)
	}
	return items, nil
}
