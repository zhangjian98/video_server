package dbops

import "log"

func AddUserCrredential(loginName string, pwd string) error {
	stmtIns, err := dbConnent.Prepare("INSTER INTO users (loginName, pwd) VALUES (?, ?)")
	if err != nil {
		return err
	}
	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return nil

}
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConnent.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()
	return pwd, nil
}
func DeleteUser(lognName, pwd string) error {
	stmtDelete, err := dbConnent.Prepare("DELETE FROM users WHERE login_name = ? and pwd = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	stmtDelete.Exec(lognName, pwd)
	stmtDelete.Close()
	return nil
}
