/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User     : ICI
   Name     : Ibrahim ÇOBANİ
   Date     : 11.08.2022
   Time     : 16:15
*/

package libs

import (
	"os"
	"path/filepath"
)

func GetAppDataPath(fileName string) (string, error) {
	fullExecPATH, err := os.Executable()
	if err != nil {
		return "", err
	}
	dir, _ := filepath.Split(fullExecPATH)
	return filepath.Join(dir, fileName), nil
}
