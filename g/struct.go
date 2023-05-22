package g

import "fmt"

type GlobalConfig struct {
	Debug			bool		`json:"debug"`
	LogFile			string		`json:"logfile"`
	LogMaxAge		int			`json:"logmaxage"`
	LogRotateAge	int			`json:"logrotateage"`
	BackupPath		string		`jsong:"backuppath"`
	TcpServer		*TcpServer	`json:"tcpserver"`
	Http 			*HTTP		`json:"http"`
}

type TcpServer struct {
	Server 			string 			`json:"server"`
	Port 			string			`json:"port"`
}

type HTTP struct {
	Address			string		`json:"address"`
	Port			string		`json:"port"`
}

type BackupPeriod struct {
	FSID                    int64           `json:"id"`
	FSPath                  string          `json:"fspath"`
	FSName                  string          `json:"fsname"`
	FSType                  string          `json:"fstype"`    // temp , perm
	FSStatus                int             `json:"status"`    // 0 create 9 done
	FSSize                  int64           `json:"size"`
	FSClient                string          `json:"client"`
	FSAgent                 string          `json:"agent"`
	FSCreateTime    string          `json:"create_time"`
	FSUpdateTime    string          `json:"update_time"`
	FSToken                 string          `json:"token"`
	FSStruct                string          `json:"struct"`         // file , directory
	FSRetry                 bool            `json:"retry"`
	GFSCluster              string          `json:"gfscluster"`
}

func (this *BackupPeriod) String() string {
	return fmt.Sprintf(
		"fsid: %d, fapath:%s, fsname:%s, fstype:%s, " +
			"client:%s, agent:%s, token:%s, retry:%v, gfscluster:%s ",
		this.FSID,
		this.FSPath,
		this.FSName,
		this.FSType,
		this.FSClient,
		this.FSAgent,
		this.FSToken,
		this.FSRetry,
		this.GFSCluster,
	)
}

