package idea

const defaultVersion = "2018.1.2"

const uri = "https://download.jetbrains.com/idea/ideaIU-"

//GetDownloadURI generate idea download URL based on version
func GetDownloadURI(version string) string {
	if version == "" {
		return uri + defaultVersion + fileExtension
	}
	return uri + version + fileExtension
}
