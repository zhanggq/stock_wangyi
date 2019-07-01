package monitor

import (
	"bytes"
	"os/exec"

	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego"
)

func BackupMysql() {
	var b bytes.Buffer
	log.Infof("Begin backup")
	do := "/usr/bin/mysql_backup.sh"
	cmd := exec.Command("sh", "-c", do)
	cmd.Stderr = &b
	res := cmd.Run()
	if res != nil {
		log.Errorf("mysql_backup err: %s", string(b.Bytes()[:]))
	}
}

func UpdateAndBackup() {
	beego.Info("Begin Update! ")
	InitializeChuangye()
	InitializeShenzhen()
	//BackupMysql()
	beego.Info("Update Finish! ")
}
