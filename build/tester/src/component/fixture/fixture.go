package fixture

import (
	"component/handler"
	"fmt"
	"os"
	"os/exec"
)

func Apply(fixturePath string) (err error) {
	cmd := exec.Command(
		"mysql",
		"--user=" + os.ExpandEnv("$MYSQL_ROOT_USER"),
		"--password=" + os.ExpandEnv("$MYSQL_ROOT_PASSWORD"),
		"--database=" + os.ExpandEnv("$MYSQL_DATABASE"),
		"--port=" + os.ExpandEnv("$MYSQL_PORT"),
		"--execute", "SOURCE " + fixturePath,
	)
	cmd.Env = os.Environ()

	out, err := cmd.CombinedOutput()
	if err != nil {
		handler.ErrorLog(err)
		return err
	}

	fmt.Println(string(out))

	return nil
}
