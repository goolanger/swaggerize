package swagger

func (specs *Instance) Encode() map[string]interface{} {
	encoded := make(map[string]interface{}, 15)

	encoded["swagger"] = "2.0"
	if specs.info != nil {
		encoded["info"] = specs.info
	}
	if specs.basePath != nil {
		encoded["basePath"] = specs.basePath
	}
	if specs.externalDocs != nil {
		encoded["externalDocs"] = specs.externalDocs
	}
	if len(specs.definitions) > 0 {
		definitions := make(map[string]interface{}, len(specs.definitions))
		for _, d := range specs.definitions {
			definitions[d.GetName()] = d.GetRef()
		}
		encoded["definitions"] = definitions
	}
	if len(specs.paths) > 0 {
		paths := make(map[string]map[string]interface{})

		for i := len(specs.paths) - 1; i >= 0; i-- {
			route, rep := specs.paths[i].GetPath(), specs.paths[i].GetRep()
			if rep != nil {
				if paths[route] == nil {
					paths[route] = rep
				} else {
					for k, v := range rep {
						paths[route][k] = v
					}
				}
			}
		}
		encoded["paths"] = paths
	}
	if len(specs.tags) >0 {
		var tags []string
		for _, t := range specs.tags {
			tags = append(tags, t.GetName())
		}
		encoded["tags"] = tags
	}
	return encoded
}
