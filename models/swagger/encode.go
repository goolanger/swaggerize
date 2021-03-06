package swagger

func (specs *Instance) Encode() map[string]interface{} {
	encoded := make(map[string]interface{}, 15)

	encoded["swagger"] = "2.0"
	if specs.info != nil {
		encoded["info"] = specs.info
	}
	if len(specs.schemes) > 0 {
		encoded["schemes"] = specs.schemes
	}
	if specs.host != nil {
		encoded["host"] = specs.host
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
			definitions[d.GetName()] = d.GetRep()
		}
		encoded["definitions"] = definitions
	}

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

	if len(specs.securities) > 0 {
		securities := make(map[string]interface{})
		for _, s := range specs.securities {
			securities[s.GetName()] = s.GetRep()
		}
		encoded["securityDefinitions"] = securities
	}

	if len(specs.secures) > 0 {
		var secures []interface{}
		for _, s := range specs.secures {
			secures= append(secures, s.GetRef().GetRep())
		}
		encoded["security"] = secures
	}

	if len(specs.tags) > 0 {
		encoded["tags"] = specs.tags
	}
	return encoded
}
