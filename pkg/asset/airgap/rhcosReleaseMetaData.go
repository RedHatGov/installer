package airgap

import (
	survey "gopkg.in/AlecAivazis/survey.v1"
	"fmt"
	"net/http"
	//"net/url"
	"os"
	"io"
	//"path"
	//"log"
	//"time"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/validate"
        //"github.com/openshift/installer/pkg/asset/releaseimage"
        //"github.com/openshift/installer/pkg/version"
	"runtime"
)

const (
	baseUrl = "https://mirror.openshift.com/pub/openshift-v4/"
)

type rhcosReleaseMetaData struct {
	distributionURL string
}

var _ asset.Asset = (*rhcosReleaseMetaData)(nil)

// Dependencies returns no dependencies.
func (a *rhcosReleaseMetaData) Dependencies() []asset.Asset {
	return nil
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	fmt.Println("Downloading: " + url)

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Failed to create: " + filepath)
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func (a *rhcosReleaseMetaData) pullVmware(airPackage *AirgapPackage) bool {

	fmt.Println("Downloading OVA")
	ova := "rhcos-" + airPackage.rhcos_ver + "-x86_64-vmware.x86_64.ova"
	DownloadFile(airPackage.dest + "/rhcos/" + ova, a.distributionURL + "rhcos-4.6.8-x86_64-vmware.x86_64.ova")

	return true
}

func (a *rhcosReleaseMetaData) pullAWS(airPackage *AirgapPackage) bool {
	return true
}

func (a *rhcosReleaseMetaData) pullAzure(airPackage *AirgapPackage) bool {
        return true
}

func (a *rhcosReleaseMetaData) pullGCP(airPackage *AirgapPackage) bool {
        return true
}

func (a *rhcosReleaseMetaData) pullOpenStack(airPackage *AirgapPackage) bool {
        return true
}

func (a *rhcosReleaseMetaData) pullBareMetal(airPackage *AirgapPackage) bool {
        return true
}


//func (a *rhcosReleaseMetaData) createAirgapPackage(rhcos_ver string, dest string, platform string) bool {
func (a *rhcosReleaseMetaData) createAirgapPackage(airPackage *AirgapPackage) bool {

	fmt.Println("Downloading RHCOS bits for: ", runtime.GOARCH)

	if runtime.GOARCH == "amd64" {
		a.distributionURL = baseUrl + "x86_64/dependencies/rhcos/4.6/" + airPackage.rhcos_ver + "/"
	} else {
		a.distributionURL= baseUrl + runtime.GOARCH + "/dependencies/rhcos/4.6/" + airPackage.rhcos_ver + "/"
	}

	a.pullVmware(airPackage)

	return true
}


// Generate queries for the cluster name from the user.
func (a *rhcosReleaseMetaData) Generate(parents asset.Parents) error {
	validator := survey.Required

	validator = survey.ComposeValidators(validator, func(ans interface{}) error {
		return validate.URI(ans.(string))
	})

        return nil

}

// Name returns the human-friendly name of the asset.
func (a *rhcosReleaseMetaData) Name() string {
	return "rhcosReleaseMetaData"
}

