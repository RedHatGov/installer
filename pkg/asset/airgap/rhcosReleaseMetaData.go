package airgap

import (
	survey "gopkg.in/AlecAivazis/survey.v1"
	"fmt"
	"net/http"
	//"net/url"
	"os"
	"io"
	//"path"
	"log"
	//"time"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/validate"
        //"github.com/openshift/installer/pkg/asset/releaseimage"
        //"github.com/openshift/installer/pkg/version"
	"runtime"
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

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func pullVmware(baseUrl string) bool {

	DownloadFile("./rhcos-4.6.8-x86_64-vmware.x86_64.ova",baseUrl + "rhcos-4.6.8-x86_64-vmware.x86_64.ova")

	return true
}

func (a *rhcosReleaseMetaData) createAirgapPackage(dest string, platform string) bool {

	err := os.Chdir(dest)

	if err != nil {
		log.Println(err)
		return false
	}

	fmt.Println("Downloading RHCOS bits for: ", runtime.GOARCH)

	if runtime.GOARCH == "amd64" {
		a.distributionURL="https://mirror.openshift.com/pub/openshift-v4/x86_64/dependencies/rhcos/4.6/4.6.8/"
	} else {
		a.distributionURL="https://mirror.openshift.com/pub/openshift-v4/" + runtime.GOARCH + "/dependencies/rhcos/4.6/4.6.8/"
	}

/*
        versionString, err := version.Version()
        if err != nil {
                return false
        }
*/
	pullVmware(a.distributionURL)

/*
        fmt.Printf("%s %s\n", os.Args[0], versionString)
        if version.Commit != "" {
                fmt.Printf("built from commit %s\n", version.Commit)
        }
        if image, err := releaseimage.Default(); err == nil {
                fmt.Printf("release image %s\n", image)
        }
*/

/*
	res, err := http.Get(src)
	check(err)
	defer res.Body.Close()

	u, err := url.Parse(src)
	check(err)

	fileName := dest + path.Base(u.Path)
	fmt.Println(fileName)
	out, err := os.Create(path.Base(u.Path))
	defer out.Close()

	size := res.ContentLength
	bar := &Progbar{total: int(size)}
	written := make(chan int, 500)

	quit := make(chan bool)

	go func() {
		copied := 0
		c := 0
		tick := time.Tick(interval)

		for {
			select {
			case c = <-written:
				copied += c
			case <-tick:
				bar.PrintProg(copied)
			case <-quit:
				return		

			}
		}
	}()

	buf := make([]byte, 32*1024)
	for {
		rc, re := res.Body.Read(buf)
		if rc > 0 {
			wc, we := out.Write(buf[0:rc])
			check(we)

			if wc != rc {
				log.Fatal("Read and Write count mismatch")
			}

			if wc > 0 {
				written <- wc
			}
		}
		if re == io.EOF {
			break
		}
		check(re)
	}

	bar.PrintComplete()
	quit <- true
*/
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

