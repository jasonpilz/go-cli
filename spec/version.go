package spec

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/blang/semver"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	// TEMPLATE_TODO Repo is the github repository
	Repo = "go-cli"
	// TEMPLATE_TODO Owner is the github repository owner
	Owner = "jasonpilz"
)

var (
	// Github token required in environment if repo is private
	githubToken = os.Getenv("GITHUB_TOKEN")
	// Set at build time from goreleaser using ldflags
	Build, Major, Minor, Patch, Label string
	// Version is the application's current version
	AppVersion Version
)

// Version is the version info for the app specification
type Version struct {
	Major, Minor, Patch int
	Name, Build, Label  string
}

func init() {
	setAppVersion()
}

func setAppVersion() {
	if Build != "" {
		AppVersion.Build = Build
	}
	if Major != "" {
		i, _ := strconv.Atoi(Major)
		AppVersion.Major = i
	}
	if Minor != "" {
		i, _ := strconv.Atoi(Minor)
		AppVersion.Minor = i
	}
	if Patch != "" {
		i, _ := strconv.Atoi(Patch)
		AppVersion.Patch = i
	}
	if Label == "" {
		AppVersion.Label = "development"
	} else {
		AppVersion.Label = Label
	}
}

// String is the current version as formatted string.
func (v Version) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch))
	if v.Label != "" {
		buffer.WriteString("-" + v.Label)
	}

	return buffer.String()
}

// Complete is the complete version for app spec.
func (v Version) Complete(lv LatestVersioner) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s version %s", Repo, v.String()))

	if v.Build != "" {
		buffer.WriteString(fmt.Sprintf("\nGit commit hash: %s", v.Build))
	}

	if tagName, err := lv.LatestVersion(); err == nil {
		v0, err1 := semver.Make(tagName)
		v1, err2 := semver.Make(v.String())

		if len(v0.Build) == 0 {
			v0, err1 = semver.Make(tagName + "-release")
		}

		if err1 == nil && err2 == nil && v0.GT(v1) {
			update := fmt.Sprintf("\nUpdated version available: %s", tagName)
			buffer.WriteString(update)
		}
	}

	return buffer.String()
}

// LatestVersioner an interface for detecting the latest version.
type LatestVersioner interface {
	LatestVersion() (string, error)
}

// GithubLatestVersioner retrieves the latest version from Github.
type GithubLatestVersioner struct{}

var _ LatestVersioner = &GithubLatestVersioner{}

// LatestVersion retrieves the latest version from Github or returns an error.
func (glv *GithubLatestVersioner) LatestVersion() (string, error) {
	// Pass custom client using oauth to create GH client
	ctx := context.Background()
	tSrc := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	authClient := oauth2.NewClient(ctx, tSrc)
	client := github.NewClient(authClient)

	// Get latest release from Github API
	rel, _, err := client.Repositories.GetLatestRelease(ctx, Owner, Repo)
	if err != nil {
		return "", errors.New("Error getting latest release from GH")
	}

	return strings.TrimPrefix(*rel.TagName, "v"), nil
}
