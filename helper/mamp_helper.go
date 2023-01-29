package helper

import (
	"fmt"
	"io/ioutil"
	"mamp-vhosts-manager/model"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type MampHelper struct {

}

func (mh *MampHelper) InitializeVhosts(){
	var path string = mh.GetVirtualHostsDirectoryPath()

	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		// Directory doesn't exists, let's create

		err := os.Mkdir(path, 0755)
		if err != nil {
			// Error while creating the folder
		}
	}

	//TODO: Check for permissions on directory

	mh.AddApacheIncludesVhostsDirectory()
}

// Function to get file host path
func (mh *MampHelper) GetFileHostPath() string{
	return "/etc/hosts";
}

// Function to get mamp root path
func (mh *MampHelper) GetMampRootPath() string{
	return "/Applications/MAMP"
}

func (mh *MampHelper) GetStopApacheMampPath() string{
	return mh.GetMampRootPath() + "/bin/stopApache.sh"
}

func (mh *MampHelper) GetStartApacheMampPath() string{
	return mh.GetMampRootPath() + "/bin/startApache.sh"
}

// Function to guess the probably base location
// of the document root for new vhosts
// fetching a random vhost and getting the document root
// without the part after the vhost name
func (mh *MampHelper) GuessDocumentRoot() string{
	vhosts := mh.GetVirtualHosts()
	if(len(vhosts) > 0){
		vhost := vhosts[0]
		name := vhost.Name
		documentRoot := vhost.DocumentRoot

		finalStringPieces := []string{}
		documentRootPieces := strings.Split(documentRoot, "/")
		for _, piece := range documentRootPieces{
			if piece == name {
				break
			}
			finalStringPieces = append(finalStringPieces, piece)
		}

		finalString := strings.Join(finalStringPieces, "/")
		return finalString
	}
	return ""
}

// Function to restart apache launching
// the stop and start scripts
func (mh *MampHelper) RestartApache() {
	// Stop apache by calling sh file
	stopApacheCommand := mh.GetStopApacheMampPath()
	_, err := os.Stat(stopApacheCommand)
	if err != nil && os.IsNotExist(err) {
		// File doesn't exists
		fmt.Println("File doesn't exists")
	}
	// File exists, let's execute it
	exec.Command(stopApacheCommand).Run()

	// Start apache by calling sh file
	startApacheCommand := mh.GetStartApacheMampPath()
	_, err = os.Stat(startApacheCommand)
	if err != nil && os.IsNotExist(err) {
		// File doesn't exists
		fmt.Println("File doesn't exists")
	}
	// File exists, let's execute it
	exec.Command(startApacheCommand).Run()
}

func (mh *MampHelper) GetVirtualHostsDirectoryPath() string{
	// Get the path to the vhosts directory relative to the root path
	return mh.GetMampRootPath() + "/conf/apache/vhosts"
}

// Get mamp main apache config file relative to root path
func (mh *MampHelper) GetMampApacheConfigFilePath() string{
	return mh.GetMampRootPath() + "/conf/apache/httpd.conf"
}

// Check main file contains inclusion for vhosts directory
// if not adds it with a placeholder tag that let's
// it know if it was added by this program
func (mh *MampHelper) AddApacheIncludesVhostsDirectory() {
	// get apache file
	file, err := os.Open(mh.GetMampApacheConfigFilePath())
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)

	// Check if the file contains the vhosts directory
	// using this placeholder tag: #mamp-vhosts-manager-vhosts-directory
	if !strings.Contains(str, "#mamp-vhosts-manager-vhosts-directory") {
		// The file doesn't contain the tag, let's add it
		// Get the path to the vhosts directory relative to the root path
		vhostsDirectoryPath := mh.GetVirtualHostsDirectoryPath()

		// create the inclusion string
		inclusionString := fmt.Sprintf("Include %s/*", vhostsDirectoryPath)

		// create the placeholder tag
		placeholderTag := "#mamp-vhosts-manager-vhosts-directory"

		// create the final string

		finalString := fmt.Sprintf("%s\n%s\n", placeholderTag, inclusionString )

		// append final string to the file

		f, err := os.OpenFile(mh.GetMampApacheConfigFilePath(), os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {

		}

		defer f.Close()

		if _, err = f.WriteString(finalString); err != nil {

		}
	}
}

func (mh *MampHelper) GetVirtualHosts() []*model.Vhost {
	var path string = mh.GetVirtualHostsDirectoryPath()

	vhosts := make([]*model.Vhost, 0, 10)
	ext := ".conf"
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ext {

			file, err := os.Open(path)
			if err != nil {
				fmt.Print(err)
			}
			defer file.Close()

			b, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Print(err)
			}
			str := string(b)

			// The name should be the filename without the extension
			vhostName := filepath.Base(path)
			vhostName = strings.TrimSuffix(vhostName, filepath.Ext(vhostName))

			lines := strings.Split(str, "\n")
			vhost := &model.Vhost{
				Name: vhostName,
			}
			for _, line := range lines {
				if strings.Contains(line, "DocumentRoot") {
					vhost.DocumentRoot = strings.TrimSpace(line)
					vhost.DocumentRoot = strings.TrimPrefix(vhost.DocumentRoot, "DocumentRoot")
					vhost.DocumentRoot = strings.TrimSpace(vhost.DocumentRoot)
					// Remove quotes from vhost.DocumentRoot
					vhost.DocumentRoot = strings.TrimPrefix(vhost.DocumentRoot, "\"")
					vhost.DocumentRoot = strings.TrimSuffix(vhost.DocumentRoot, "\"")
				}
				if strings.Contains(line, "ServerAdmin") {
					vhost.ServerAdmin = strings.TrimSpace(strings.TrimPrefix(line, "ServerAdmin"))
				}
				if strings.Contains(line, "ServerName") {
					vhost.ServerName = strings.TrimSpace(line)
					vhost.ServerName = strings.TrimPrefix(vhost.ServerName, "ServerName")
					vhost.ServerName = strings.TrimSpace(vhost.ServerName)
				}
			}

			vhosts = append(vhosts, vhost)
		}
		return nil
	})
	return vhosts
}

// Function to get virtual host from name
func (mh *MampHelper) GetVirtualHost(name string) *model.Vhost {
	for _, vhost := range mh.GetVirtualHosts(){
		if(vhost.Name == name){
			return vhost
		}
	}
	return nil
}

func (mh *MampHelper) CreateVirtualHost(name string, serverName string, serverAdmin string, documentRoot string) *model.Vhost {
	var path string = mh.GetVirtualHostsDirectoryPath()

	// Delete the virtual host if it already exists
	if mh.GetVirtualHost(name) != nil {
		mh.DeleteVirtualHost(name)
	}

	file, err := os.Create(path + "/" + name + ".conf")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	// If serveradmin is empty set it to
	// webmaster@dummy-host2.example.com

	if serverAdmin == "" {
		serverAdmin = "webmaster@dummy-host2.example.com"
	}

	file.WriteString("<VirtualHost *:80>\n")
	file.WriteString("	ServerAdmin " + serverAdmin + "\n")
	file.WriteString("	DocumentRoot " + documentRoot + "\n")
	file.WriteString("	ServerName " + serverName + "\n")
	file.WriteString("	#ErrorLog \"logs/" + name + "-error_log\"\n")
	file.WriteString("	#CustomLog \"logs/" + name + "-access_log\" common\n")
	file.WriteString("</VirtualHost>\n")

	vhost := mh.GetVirtualHost(name)
	mh.AddVirtualHostToHostsFile(vhost)

	return vhost
}

// Function to remove virtual host record
// from the file host
func (mh *MampHelper) RemoveVirtualHostFromHostsFile(vhost *model.Vhost) {

	fh := FileHelper{}

	// Get the path to the hosts file
	hostsFilePath := mh.GetFileHostPath()

	// Create the tag to remove from the hosts file
	hostsFileTag := fmt.Sprintf("#mamp-vhosts-manager-vhost-%s", vhost.Name)

	// Check if the tag is already present
	// in the hosts file
	file, _ := os.ReadFile(hostsFilePath)

	// Convert file to string
	stringFile := string(file)

	if strings.Contains(stringFile, hostsFileTag) {

		// Use the file helper to find line containg tag
		lineNumber := fh.GetLineNumberThatContainsString(stringFile, hostsFileTag)

		// print line number
		fmt.Println(lineNumber)

		if lineNumber > 0 {
			// Remove the line
			stringFile = fh.RemoveLineFromFileContent(stringFile, lineNumber)
			stringFile = fh.RemoveLineFromFileContent(stringFile, lineNumber)

			// Write the file
			err := ioutil.WriteFile(hostsFilePath, []byte(stringFile), 0644)
			if err != nil {
				fmt.Print(err)
			}
		}
	}
}

// Function to add the virtual host record
// to the file host with a tag that let
// recognize if already added
// and avoid duplicates
// #mamp-vhosts-manager-vhost-{name}
func (mh *MampHelper) AddVirtualHostToHostsFile(vhost *model.Vhost) {
	// Get the path to the hosts file
	hostsFilePath := mh.GetFileHostPath()

	// Create the string to add to the hosts file
	hostsFileString := fmt.Sprintf("127.0.0.1\t%s\n", vhost.ServerName)

	// Create the tag to add to the hosts file
	hostsFileTag := fmt.Sprintf("#mamp-vhosts-manager-vhost-%s", vhost.Name)

	// Check if the tag is already present
	// in the hosts file
	file, _ := os.ReadFile(hostsFilePath)

	// Convert file to string
	stringFile := string(file)

	if strings.Contains(stringFile, hostsFileTag) {
		mh.RemoveVirtualHostFromHostsFile(vhost)
	}

	stringFile = stringFile + "\n"

	// Add the tag
	stringFile = stringFile + hostsFileTag + "\n"

	// Add the string
	stringFile = stringFile + hostsFileString + "\n"

	// Write the file
	err := os.WriteFile(hostsFilePath, []byte(stringFile), 0644)
	if err != nil {
		fmt.Print(err)
	}
}

func (mh *MampHelper) DeleteVirtualHost(name string) {

	vhost := mh.GetVirtualHost(name)
	mh.RemoveVirtualHostFromHostsFile(vhost)

	// Delete the virtual host if it already exists
	// search virtualhost path
	var path string = mh.GetVirtualHostsDirectoryPath()
	var vhostPath string = path + "/" + name + ".conf"
	// Delete it
	err := os.Remove(vhostPath)
	if err != nil {
		fmt.Print(err)
	}


}