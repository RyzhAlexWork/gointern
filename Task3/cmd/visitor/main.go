package main

import (
	"fmt"

	"github.com/RyzhAlexWork/gointern/Task3/pkg/groups"
	"github.com/RyzhAlexWork/gointern/Task3/pkg/reputation"
	"github.com/RyzhAlexWork/gointern/Task3/pkg/visitor"
)

func main() {
	// Create visitor
	newVisitor := visitor.NewVisitor()

	// Create user
	reputationForUser := reputation.NewReputation(0)
	newUser := groups.NewUser("alex", reputationForUser)

	// Create sponsor
	reputationForSponsor := reputation.NewReputation(0)
	newSponsor := groups.NewSponsor("voidkol", reputationForSponsor)

	// Create moderator
	reputationForModerator := reputation.NewReputation(0)
	newModerator := groups.NewModerator("gooool", reputationForModerator)

	// Create administrator
	reputationForAdministrator := reputation.NewReputation(0)
	newAdministrator := groups.NewAdministrator("napore", reputationForAdministrator)

	accounts := []groups.Groups{newUser, newSponsor, newModerator, newAdministrator}

	// Use visitor
	for _, account := range accounts {
		account.Accept(newVisitor)
	}

	fmt.Printf("Reputation user after visitor: %d\n", newUser.GetReputation())
	fmt.Printf("Reputation sponsor after visitor: %d\n", newSponsor.GetReputation())
	fmt.Printf("Reputation moderator after visitor: %d\n", newModerator.GetReputation())
	fmt.Printf("Reputation administrator after visitor: %d\n", newAdministrator.GetReputation())
	fmt.Println("Visitor collected logins:", newVisitor.GetLogins())
	fmt.Println("Visitor counted logins:", newVisitor.GetNumberOfUsers())
}
