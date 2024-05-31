package Users

import (
	"math/rand"

	"github.com/Petatookmykfc/prehnitelogs"
)

func init() {
	prehnitelogs.AddType("User", "User Action :")
	prehnitelogs.AddType("System", "Automated Action :")
	prehnitelogs.AddType("Admin", "Administrator Action :")
	prehnitelogs.AddType("User", "User")
}

type Users struct {
	name     string
	admin    bool
	password string
	logger   prehnitelogs.LogFunc
}

var AdminLog = prehnitelogs.GetCustomLogMethod("Admin")
var SystemLog = prehnitelogs.GetCustomLogMethod("System")
var UserLog = prehnitelogs.GetCustomLogMethod("User")

func (u *Users) HiddenAdminFunction() (bool, error) {

	if u.admin {
		AdminLog("Admin is doing something! That's fine...")
		return true, nil
	} else {
		err := UserLog("Attempted to access an admin only function! That's !fine...")
		HackerDetected(u)
		return false, err
	}
}
func HackerDetected(hacker *Users) {
	SystemLog("Attempted HACK By ? " + hacker.name)

	// TODO: Add Panic scream to notifie admins

	// BAN USER
	hacker.BanUser()
}

func (u *Users) ChangePassword(newPassword string) (bool, error) {
	u.password = newPassword
	u.logger("Password Changed")
	return true, nil
}

func (u *Users) BanUser() {
	banScript := RandStringBytes(32)
	SystemLog("Banning " + u.name + " new password " + banScript)
	u.password = banScript
	u.logger("Banned")
}

// Unauthorized access

// Generate Ban Script
func RandStringBytes(n int) string {
	SystemLog("Generating Ban code...")
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	SystemLog("Ban code generated!  " + string(b))
	return string(b)
}

func Imagine() {

	AdminUser := Users{name: "John", admin: true, password: "1234", logger: AdminLog}
	UnauthorizedUser := Users{name: "Dr.Evil", admin: false, password: "1234", logger: UserLog}

	AdminUser.HiddenAdminFunction()
	UnauthorizedUser.HiddenAdminFunction()

}

func ExampleEntry() {
	// Imaging this function working!
	Imagine()

	/*
		This is an example of how the system could be used.
		The aim of this script is to ban any user who triggers the function
		HiddenAdminFunction that isn't an admin. The logs produced by different functions will be stored with different prefixes
		making it clear what is triggering the log. For example: System will be logged with the prefix "Automated action", this is triggered
		when the system automaticaly bans the user who accessed the function. Anoher example of this could be automated systems such as backups,
		restarts, or anything that doesn't require human interaction inorder to execute.


		In this example the system will log along the lines of

		FILe -> Logs/HiddenAdminFunction.log {
			Admin accessed the hidden function ( thats fine)
			User accessed the hidden function (thats not ok)
		}
		FILE -> Logs/HackerDetected.log {
			I think i was hacked by $user
		}
		FILE -> Logs/BanUser.lgo {
			Systems banning $user
			System has banned $user
		}
		FILE -> RandStringBytes {
			Generating Ban code... ( setting as users password )
			Ban code generated! The code is $code
		}
	*/

}

// I know that users shouldn't even have that function attached to there struct.
