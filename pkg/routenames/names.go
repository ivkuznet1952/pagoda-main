package routenames

import (
	"fmt"
)

const (
	Home    = "home"
	About   = "about"
	Contact = "contact"
	//Shedule              = "contact"
	CreateGOrder         = "creategorder"
	Shedule              = "shedule"
	ContactSubmit        = "contact.submit"
	GOrderSubmit         = "gorder.submit"
	SheduleSubmit        = "shedule.submit"
	Login                = "login"
	LoginSubmit          = "login.submit"
	Register             = "register"
	RegisterSubmit       = "register.submit"
	ForgotPassword       = "forgot_password"
	ForgotPasswordSubmit = "forgot_password.submit"
	Logout               = "logout"
	VerifyEmail          = "verify_email"
	ResetPassword        = "reset_password"
	ResetPasswordSubmit  = "reset_password.submit"
	Search               = "search"
	Task                 = "task"
	TaskSubmit           = "task.submit"
	Cache                = "cache"
	CacheSubmit          = "cache.submit"
	Files                = "files"
	FilesSubmit          = "files.submit"
	AdminTasks           = "admin:tasks"
)

func AdminEntityList(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_list", entityTypeName)
}

func AdminEntityAdd(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_add", entityTypeName)
}

func AdminEntityEdit(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_edit", entityTypeName)
}

func AdminEntityDelete(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_delete", entityTypeName)
}

func AdminEntityAddSubmit(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_add.submit", entityTypeName)
}

func AdminEntityEditSubmit(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_edit.submit", entityTypeName)
}

func AdminEntityDeleteSubmit(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_delete.submit", entityTypeName)
}
