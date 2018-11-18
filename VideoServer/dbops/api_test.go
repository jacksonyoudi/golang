package dbops

import (
	"fmt"
	"golang/VideoServer/dbops/models"
	"testing"
)

/*
init: dbconn, init sql, trucate tables -> run test-> clear test;
 */

func clearTables()  {
	models.ClearTables()
}


func TestMain(m *testing.M)  {
	m.Run()
}

func TestUserWorkFlow(t *testing.T)  {
	//testAddUser(t)
	testGetUser(t)
}

func testAddUser(t *testing.T)  {
	err := models.CreateUser("youdi", "123")
	if err != nil {
		t.Errorf("Error of createUser, %v", err)
	}
}

func testGetUser(t *testing.T)  {
	user, err := models.GetUser("youdi")
	if err != nil {
		t.Errorf("Error of createUser, %v", err)
	} else {
		fmt.Printf("User: %v", user)
	}
}