package golang

const DefaultVersion = "1.10.3"

const uri = "https://dl.google.com/go/go"

func GetDownloadURI(version string) string {
	if version == "" {
		return uri + DefaultVersion + uriPostfix
	}
	return uri + version + uriPostfix
}