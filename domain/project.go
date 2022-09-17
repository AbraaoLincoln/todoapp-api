package domain

type Project struct {
	Id         string `json: id`
	Name       string `json: name`
	Color      string `json: color`
	CreateAt   string `json: createAt`
	ModifiedAt string `json: modifiedAt`
}

func (p *Project) IsEmpty() bool {
	return p.Id == "" && p.Name == "" && p.Color == ""
}
