package interfaces

type AuthRepository interface {
    CheckExistingUser(email, username string) (bool, error)
}