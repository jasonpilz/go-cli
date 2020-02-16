package spec

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVersion(t *testing.T) {
	Convey("Given the need to assign a specification version", t, func() {
		Convey("When setting Version directly", func() {
			cases := []struct {
				v   Version
				s   string
				ver string
			}{
				// version unset
				{
					v:   Version{},
					ver: "0.0.0",
				},
				// version with patch
				{
					v:   Version{Major: 0, Minor: 0, Patch: 3},
					ver: "0.0.3",
				},
				// version with minor and patch
				{
					v:   Version{Major: 0, Minor: 1, Patch: 2},
					ver: "0.1.2",
				},
				// version with major, minor and patch
				{
					v:   Version{Major: 2, Minor: 33, Patch: 2},
					ver: "2.33.2",
				},
				// version with build
				{
					v:   Version{Major: 0, Minor: 1, Patch: 2, Build: "12345"},
					ver: "0.1.2",
				},
			}

			for _, c := range cases {
				So(c.v.String(), ShouldEqual, c.ver)

				if got, want := c.v.String(), c.ver; got != want {
					t.Errorf("version string for %#v = %q; want = %q", c.v, got, want)
				}
			}
		})
	})
}
