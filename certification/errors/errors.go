// Package errors defines all the single use as well as reusable errors within preflight
package errors

import "errors"

var (
	ErrNoChecksEnabled                 = errors.New("no checks have been enabled")
	ErrRequestedCheckNotFound          = errors.New("requested check not found")
	ErrRequestedFormatterNotFound      = errors.New("requested formatter is not known")
	ErrFormatterNameNotProvided        = errors.New("formatter name is required")
	ErrFormattingResults               = errors.New("error formatting results")
	ErrInsufficientPosArguments        = errors.New("not enough positional arguments")
	ErrGetRemoteContainerFailed        = errors.New("failed to pull remote container")
	ErrExtractingTarball               = errors.New("failed to extract tarball")
	ErrCreateTempDir                   = errors.New("failed to create temporary directory")
	ErrImageInspectFailed              = errors.New("failed to inspect image")
	ErrOperatorSdkScorecardFailed      = errors.New("failed to run operator-sdk scorecard")
	ErrOperatorSdkBundleValidateFailed = errors.New("failed to run operator-sdk bundle validate")
	ErrNoKubeconfig                    = errors.New("no environment variable KUBECONFIG could be found")
	ErrK8sAPICallFailed                = errors.New("unable to fetch the requested resource from k8s API server")
	ErrEmptyAnnotationFile             = errors.New("the annotations file was empty")
	ErrLicensesNotADir                 = errors.New("licenses is not a directory")
	ErrSupportCmdPromptFailed          = errors.New("prompt failed, please try re-running support command")
	ErrEmptyProjectID                  = errors.New("please enter a non empty project id")
	ErrRemovePFromProjectID            = errors.New("please remove leading character p from project id")
	ErrRemoveOSPIDFromProjectID        = errors.New("please remove leading character ospid- from project id")
	ErrRemoveSpecialCharFromProjectID  = errors.New("please remove all special characters from project id")
	ErrPullRequestURL                  = errors.New("please enter a valid url: including scheme, host, and path to pull request")
	ErrIndexImageUndefined             = errors.New("no environment variable PFLT_INDEXIMAGE could be found")
	ErrUnsupportedGoType               = errors.New("go type unsupported")
	ErrSaveFileFailed                  = errors.New("failed to save file to artifacts directory")
	ErrNon200StatusCode                = errors.New("error calling remote API")
	Err409StatusCode                   = errors.New("remote API returned conflict")
)
