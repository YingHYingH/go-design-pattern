package _0_composite

type IOrganization interface {
	Count() int
}
type Employee struct {
	Name string
}

func (e Employee) Count() int {
	return 1
}

type Department struct {
	Name             string
	SubOrganizations []IOrganization
}

func (d Department) Count() int {
	count := 0
	for _, sub := range d.SubOrganizations {
		count += sub.Count()
	}
	return count
}

func (d *Department) AddSub(org IOrganization) {
	d.SubOrganizations = append(d.SubOrganizations, org)
}

func NewOrganization() IOrganization {
	root := &Department{Name: "root"}
	for i := 0; i < 10; i++ {
		root.AddSub(&Employee{})
		root.AddSub(&Department{Name: "sub", SubOrganizations: []IOrganization{&Employee{}, &Department{}}})
	}
	return root
}
