package minecraft

import (
	"fmt"
	"strconv"
	"time"
)

func SetupServerScript(jvmRam string) string {
	downloadUrl := "https://launcher.mojang.com/v1/objects/a412fd69db1f81db3f511c1463fd304675244077/server.jar"
	script := `
				mkdir /opt/minecraft
				echo "eula=true" > /opt/minecraft/eula.txt
				wget %s
				mv server.jar /opt/minecraft/
				apt-get update
				apt-get install screen default-jre sshpass -y
				screen -dm bash -c "cd /opt/minecraft; java -Xmx%sM -Xms%sM -jar server.jar nogui"
				`
	return fmt.Sprintf(script, downloadUrl, jvmRam, jvmRam)
}

func SaveScript(customerID string, game string) string {
	tarName := customerID + "-" + game + "-" + strconv.FormatInt(time.Now().Unix(), 10)
	pass := "a1b2c3d4"
	targetSsh := "marc@49.12.4.204:saves/"

	script := `
				pkill -SIGINT java
				sleep 10
				pkill java
				tar -czf %s.tar.gz -C /opt/ minecraft
				sshpass -p %s scp -oStrictHostKeyChecking=no %s.tar.gz %s
			`
	return fmt.Sprintf(script, tarName, pass, tarName, targetSsh)
}

func ListSavesByCustomerIDScript(customerID string) string {
	return "ls -A1 /home/marc/saves/" + customerID + "*.tar.gz"
}

func DeleteSaveScript(saveFileName string) string {
	return "rm /home/marc/saves/" + saveFileName
}

func SetupServerAndLoadSaveScript(saveFileName string, jvmRam string) string {
	pass := "a1b2c3d4"
	targetSsh := "marc@49.12.4.204:saves/"
	script := `
				apt-get update
				apt-get install screen default-jre sshpass -y
				sshpass -p %s scp -oStrictHostKeyChecking=no %s%s /opt
				cd /opt
				tar -xf %s
				screen -dm bash -c "cd /opt/minecraft; java -Xmx%sM -Xms%sM -jar server.jar nogui"
				`
	return fmt.Sprintf(script, pass, targetSsh, saveFileName, saveFileName, jvmRam, jvmRam)
}
