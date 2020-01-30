package fixture

import (
	"component/file"
	"io"
	"os/exec"
)

func Apply(fixturePath string) (err error) {
	fixtureContent, err := file.GetContent(fixturePath)
	if err != nil {
		return err
	}
	cmd := exec.Command(
		"mysql",
		"--user", "$MYSQL_ROOT_USER",
		"--password", "$MYSQL_ROOT_PASSWORD",
		"--database", "$MYSQL_DATABASE",
		"--host", "$MYSQL_HOST",
		"--port", "$MYSQL_PORT",
		)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	err = cmd.Start()
	if err != nil {
		return err
	}
	_, err = io.WriteString(stdin, fixtureContent)
	if err != nil {
		return err
	}
	return nil
}
