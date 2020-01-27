package visitor

import groups "github.com/RyzhAlexWork/gointern/Task3/pkg/groups"

// Visitor ...
type Visitor interface {
	VisitAdministrator(a groups.Administrator)
	VisitModerator(m groups.Moderator)
	VisitSponsor(s groups.Sponsor)
	VisitUser(u groups.User)
	GetLogins() []string
	GetNumberOfUsers() int
}

type visitor struct {
	numberOfUsers int
	userLogins    []string
}

// VisitAdministrator used when visiting group Administrator, add reputation.
func (v *visitor) VisitAdministrator(a groups.Administrator) {
	v.numberOfUsers++
	v.userLogins = append(v.userLogins, a.GetLogin())
	a.AddReputation(1000)
}

// VisitModerator used when visiting group Moderator, add reputation.
func (v *visitor) VisitModerator(m groups.Moderator) {
	v.numberOfUsers++
	v.userLogins = append(v.userLogins, m.GetLogin())
	m.AddReputation(500)
}

// VisitSponsor used when visiting group Sponsor, add reputation.
func (v *visitor) VisitSponsor(s groups.Sponsor) {
	v.numberOfUsers++
	v.userLogins = append(v.userLogins, s.GetLogin())
	s.AddReputation(100)
}

// VisitUser used when visiting group User, add reputation.
func (v *visitor) VisitUser(u groups.User) {
	v.numberOfUsers++
	v.userLogins = append(v.userLogins, u.GetLogin())
	u.AddReputation(50)
}

// GetLogins get all accounts login.
func (v *visitor) GetLogins() []string {
	return v.userLogins
}

// GetNumberOfUsers сounts the number of account visited.
func (v *visitor) GetNumberOfUsers() int {
	return v.numberOfUsers
}

// NewVisitor create user implementation for interface Visitor.
func NewVisitor() Visitor {
	return &visitor{
		numberOfUsers: 0,
		userLogins:    []string{},
	}
}
