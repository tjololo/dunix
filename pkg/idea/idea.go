package idea

//DefaultVersion default version of Idea to install
const DefaultVersion = "2018.1.2"

const uri = "https://download.jetbrains.com/idea/ideaIU-"

//GetDownloadURI generate idea download URL based on version
func GetDownloadURI(version string) string {
	if version == "" {
		return uri + DefaultVersion + fileExtension
	}
	return uri + version + fileExtension
}
