package main

import "testing"


type FactorioWithMocks struct {
	NewUserCalled bool
	DeletedUserCalled bool
	UnderTest Factorio
	t *testing.T
}

func createMock(t *testing.T) *FactorioWithMocks {
	this := &FactorioWithMocks{
		NewUserCalled: false,
		DeletedUserCalled: false,
		t: t,
	}

	this.UnderTest = CreateFactorio(func(s string) { this.NewUserCalled = true }, func(s string) { this.DeletedUserCalled = true })


	return this
}

func (this *FactorioWithMocks) reset() {
	this.DeletedUserCalled = false
	this.NewUserCalled = false
}

func (this *FactorioWithMocks) assertNewUserHasBeenCalled() {
	if ! this.NewUserCalled {
		this.t.Error("Expected the create user method to be called.")
	}
}

func (this *FactorioWithMocks) assertNewUserHasNotBeenCalled() {
	if this.NewUserCalled {
		this.t.Error("Expected the create user method to NOT be called.")
	}
}

func (this *FactorioWithMocks) assertDeleteUserHasBeenCalled() {
	if ! this.DeletedUserCalled {
		this.t.Error("Expected the delete user method to be called.")
	}
}

func (this *FactorioWithMocks) assertDeleteUserHasNotBeenCalled() {
	if this.DeletedUserCalled {
		this.t.Error("Expected the delete user method to NOT be called.")
	}
}

func TestNewUserShouldCallFunction(t *testing.T) {
	mock := createMock(t)

	mock.UnderTest.AddUser("fred")

	mock.assertNewUserHasBeenCalled()
}

func TestAddUserRunningTwice(t *testing.T) {
	mock := createMock(t)

	mock.UnderTest.AddUser("fred")
	mock.reset()
	mock.UnderTest.AddUser("fred")

	mock.assertNewUserHasNotBeenCalled()
}

func TestAddUserRunningTwiceWithDifferentUsers(t *testing.T) {
	mock := createMock(t)

	mock.UnderTest.AddUser("fred")
	mock.reset()
	mock.UnderTest.AddUser("john")

	mock.assertNewUserHasBeenCalled()
}


func TestDeleteUserShouldNotCall(t *testing.T) {
	mock := createMock(t)

	mock.UnderTest.DeleteUser("fred")

	mock.assertDeleteUserHasNotBeenCalled()
}

func TestDeleteUserCalledOnCreatedUsed(t *testing.T) {
	mock := createMock(t)

	mock.UnderTest.AddUser("fred")
	mock.UnderTest.DeleteUser("fred")

	mock.assertDeleteUserHasBeenCalled()
}

func TestDeleteUserCalledOnSecondCall(t *testing.T) {
	mock := createMock(t)

	mock.UnderTest.AddUser("fred")
	mock.UnderTest.DeleteUser("fred")
	mock.reset()
	mock.UnderTest.DeleteUser("fred")

	mock.assertDeleteUserHasNotBeenCalled()
}

func TestAddDeleteRunningTwiceWithDifferentUsers(t *testing.T) {
	mock := createMock(t)

	mock.UnderTest.AddUser("fred")
	mock.UnderTest.AddUser("john")

	mock.UnderTest.DeleteUser("fred")
	mock.reset()
	mock.UnderTest.DeleteUser("john")

	mock.assertDeleteUserHasBeenCalled()
}
