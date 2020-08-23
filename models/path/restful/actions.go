package restful

type Actions struct {
	HasGet, HasFetch, HasPost, HasPut, HasDelete bool
}

func All() *Actions {
	return &Actions{
		HasGet:    true,
		HasFetch:  true,
		HasPost:   true,
		HasPut:    true,
		HasDelete: true,
	}
}

func Only() *Actions {
	return &Actions{}
}

func (a*Actions) Get() *Actions {
	a.HasGet = true
	return a
}
func (a*Actions) Fetch() *Actions {
	a.HasFetch = true
	return a
}
func (a*Actions) Post() *Actions {
	a.HasPost = true
	return a
}
func (a*Actions) Put() *Actions {
	a.HasPut = true
	return a
}
func (a*Actions) Delete() *Actions {
	a.HasDelete = true
	return a
}

func (a*Actions) DropGet() *Actions {
	a.HasGet = false
	return a
}
func (a*Actions) DropFetch() *Actions {
	a.HasFetch = false
	return a
}
func (a*Actions) DropPost() *Actions {
	a.HasPost = false
	return a
}
func (a*Actions) DropPut() *Actions {
	a.HasPut = false
	return a
}
func (a*Actions) DropDelete() *Actions {
	a.HasDelete = false
	return a
}

