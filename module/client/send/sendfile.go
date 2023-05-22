package send

import (
	"github.com/signmem/tcpfiletransfer/g"
	"io"
	"net"
	"os"
	"strconv"
)


const BUFFERSIZE = 1024


func ReveFromHTTPToTcp(fileInfo g.BackupPeriod) {

	// client send file to server

	serverAddr := g.Config().TcpServer.Server
	serverPort := g.Config().TcpServer.Port
	tcpServer := serverAddr + ":" + serverPort

	connection, err := net.Dial("tcp", tcpServer)

	if err != nil {
		g.Logger.Errorf("net.Dial() err:%s", err)
	}

	defer connection.Close()

	defer connection.Close()
	filePath := g.Config().BackupPath + "/" +
		fileInfo.FSPath + "/" + fileInfo.FSName


	file, err := os.Open(filePath)

	if err != nil {
		g.Logger.Errorf("open file %s err:%s", filePath, err)
		return
	}

	fileDetail, err := file.Stat()
	if err != nil {
		g.Logger.Errorf("file %s stat err:%s", filePath, err)
		return
	}

	fileSize := g.FillString(strconv.FormatInt(fileDetail.Size(), 10),10)
	fileName := g.FillString(fileDetail.Name(),64)
	fileTempPath := g.FillString(fileInfo.FSPath,256)
	fileType := g.FillString(fileInfo.FSType,16)

	g.Logger.Debugf("filesize %, filename:%s, filepath:%s, type:%s",
		fileSize,fileName, fileTempPath, fileType)

	connection.Write([]byte(fileSize))
	connection.Write([]byte(fileName))
	connection.Write([]byte(fileTempPath))
	connection.Write([]byte(fileType))

	sendBuffer := make([]byte, BUFFERSIZE)

	g.Logger.Debugf("sending file:%s", filePath)

	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		connection.Write(sendBuffer)
	}
	g.Logger.Debugf("File %s has been sent, closing connection!", filePath)
	return


}


