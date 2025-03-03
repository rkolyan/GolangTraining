package domains

type Member struct {
	//TODO
}

type MemberService interface {
	GetMembers(projectID string) ([]Member, error)
	UpdateMember(member *Member) error
	DeleteMember(member *Member) error
}

type MemberRepository interface {
	Select(projectID string) ([]Member, error)
	Update(member *Member) error
	Delete(member *Member) error
}
