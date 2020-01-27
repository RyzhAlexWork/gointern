package visitor

import (
	"testing"

	"github.com/RyzhAlexWork/gointern/Task3/pkg/groups"
	"github.com/RyzhAlexWork/gointern/Task3/pkg/reputation"
	"github.com/stretchr/testify/assert"
)

var (
	expectsReputation = []int{150, 400, 1000, 2500}
	expectsLogins     = []string{
		"alex",
		"voidkol",
		"gooool",
		"napore",
	}
	expectsNumberOfUsers = 4
)

const (
	visitorTest = "visitorTest"
)

func Test_Visitor(t *testing.T) {
	// Create visitor
	newVisitor := NewVisitor()

	// Create user
	reputationForUser := reputation.NewReputation(100)
	newUser := groups.NewUser("alex", reputationForUser)

	// Create sponsor
	reputationForSponsor := reputation.NewReputation(300)
	newSponsor := groups.NewSponsor("voidkol", reputationForSponsor)

	// Create moderator
	reputationForModerator := reputation.NewReputation(500)
	newModerator := groups.NewModerator("gooool", reputationForModerator)

	// Create administrator
	reputationForAdministrator := reputation.NewReputation(1500)
	newAdministrator := groups.NewAdministrator("napore", reputationForAdministrator)

	accounts := []groups.Groups{newUser, newSponsor, newModerator, newAdministrator}

	// Use visitor
	for _, account := range accounts {
		account.Accept(newVisitor)
	}

	t.Run(visitorTest, func(t *testing.T) {
		for i, reputation := range expectsReputation {
			assert.Equal(t, reputation, accounts[i].GetReputation())
		}
		for i, login := range expectsLogins {
			assert.Equal(t, login, newVisitor.GetLogins()[i])
		}
		assert.Equal(t, expectsNumberOfUsers, newVisitor.GetNumberOfUsers())
	})
}
