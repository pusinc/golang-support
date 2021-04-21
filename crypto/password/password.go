package password

import "golang.org/x/crypto/bcrypt"

func Hash(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func HashAssign(password []byte, param *string) error {
	pwd, err := Hash(password)
	if err != nil {
		return err
	}
	*param = string(pwd)
	return nil
}

func Verify(hashedPassword, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	return err == nil
}
