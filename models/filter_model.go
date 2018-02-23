package models

var Filters []*Filter

func init() {
	filter := new(Filter)
	filter.Name = "State"
	Filters = append(Filters, filter)

	//filter = new(Filter)
	//filter.Name = "City"
	//Filters = append(Filters, filter)
}

type Filter struct {
	Name 			string 				`json:"name"`
	Options 		[]string 			`json:"options"`
	Subfilters		[]*Filter			`json:"subfilters"`
}

func (f *Filter) GetOptions(properties []*Property) {
	var options []string
	options = append(options, "All")

	for _, property := range properties {
		found := false
		for _, option := range options {
			if option == property.Address.State {
				found = true
			}
		}

		if found == false {
			options = append(options, property.Address.State)
		}
	}

	f.Options = options
}
