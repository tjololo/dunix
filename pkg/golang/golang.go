package golang

//DefaultVersion default version of Golang to install
const DefaultVersion = "1.10.3"

const uri = "https://dl.google.com/go/go"

//GetDownloadURI Generate download URI for golang based in version
func GetDownloadURI(version string) string {
	if version == "" {
		return uri + DefaultVersion + uriPostfix
	}
	return uri + version + uriPostfix
}
