package fixture

import (
	"component/handler"
	"fmt"
	"os"
	"os/exec"
)

func Apply(fixturePath string, hostname string) (err error) {
	cmd := exec.Command(
		"mysql",
		"--host=" + os.ExpandEnv("$MYSQL_MASTER_HOSTNAME"),
		"--user=" + os.ExpandEnv("$MYSQL_ROOT_USER"),
		"--password=" + os.ExpandEnv("$MYSQL_ROOT_PASSWORD"),
		"--database=" + os.ExpandEnv("$MYSQL_DATABASE"),
		"--port=" + os.ExpandEnv("$MYSQL_PORT"),
		"--execute=" + "SOURCE " + fixturePath + "",
	)
	cmd.Env = os.Environ()

	fmt.Println(
		"mysql",
		"--host=" + os.ExpandEnv("$MYSQL_MASTER_HOSTNAME"),
		"--user=" + os.ExpandEnv("$MYSQL_ROOT_USER"),
		"--password=" + os.ExpandEnv("$MYSQL_ROOT_PASSWORD"),
		"--database=" + os.ExpandEnv("$MYSQL_DATABASE"),
		"--port=" + os.ExpandEnv("$MYSQL_PORT"),
		"--execute=" + "SOURCE " + fixturePath + "",
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		handler.ErrorLog(err)
		return err
	}

	fmt.Println(string(out))

	return nil
}
