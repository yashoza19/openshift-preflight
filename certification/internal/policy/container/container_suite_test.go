package container

import (
	"io"
	"testing"
	"errors"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
	"github.com/redhat-openshift-ecosystem/openshift-preflight/cli"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
)

func TestContainer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Container Suite")
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.TraceLevel)
}

type FakeLayer struct{}

func (fl FakeLayer) Digest() (v1.Hash, error) {
	return v1.Hash{}, nil
}

func (fl FakeLayer) DiffID() (v1.Hash, error) {
	return v1.Hash{}, nil
}

func (fl FakeLayer) Compressed() (io.ReadCloser, error) {
	return nil, nil
}

func (fl FakeLayer) Uncompressed() (io.ReadCloser, error) {
	return nil, nil
}

func (fl FakeLayer) Size() (int64, error) {
	return 0, nil
}

func (fl FakeLayer) MediaType() (types.MediaType, error) {
	return "mediatype", nil
}

type FakeSkopeoEngine struct {
	SkopeoReportStdout string
	SkopeoReportStderr string
	Tags               []string
	InspectReport      cli.SkopeoInspectReport
}

type SkopeoData struct {
	Repository string
	Tags       []string
}

func (fse FakeSkopeoEngine) ListTags(image string) (*cli.SkopeoListTagsReport, error) {
	skopeoReport := cli.SkopeoListTagsReport{
		Stdout: fse.SkopeoReportStdout,
		Stderr: fse.SkopeoReportStderr,
		Tags:   fse.Tags,
	}
	return &skopeoReport, nil
}

func (fse FakeSkopeoEngine) InspectImage(rawImage string, inspectOptions cli.SkopeoInspectOptions) (*cli.SkopeoInspectReport, error) {
	return &fse.InspectReport, nil
}

type BadSkopeoEngine struct{}

func (bse BadSkopeoEngine) ListTags(string) (*cli.SkopeoListTagsReport, error) {
	skopeoReport := cli.SkopeoListTagsReport{
		Stdout: "Bad Stdout",
		Stderr: "Bad stderr",
		Tags:   []string{""},
	}
	return &skopeoReport, errors.New("the Skopeo ListTags has failed")
}

func (bse BadSkopeoEngine) InspectImage(rawImage string, inspectOptions cli.SkopeoInspectOptions) (*cli.SkopeoInspectReport, error) {
	return &cli.SkopeoInspectReport{Stdout: "", Stderr: "some error output"}, errors.New("the skopeo image inspection has failed")
}