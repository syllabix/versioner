package semver

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Parse will attempt to parse a string into a semver Version
// returning a zero value version and non nil error on failure
func Parse(s string) (Version, error) {
	r := bufio.NewReader(strings.NewReader(s))

	first, err := r.Peek(1)
	if err != nil {
		return Version{}, err
	}

	if bytes.EqualFold(first, []byte("v")) {
		r.ReadByte()
	}

	var v Version
	curnumtype := Major
	assign := func(num int) {
		switch curnumtype {
		case Major:
			v.major = num
		case Minor:
			v.minor = num
		case Patch:
			v.patch = num
		}
		curnumtype++
	}

	for {

		// if we are looking at the pre-release section
		// read til end of string and assign resulting value
		// unless there is a "hard" error (non io.EOF)
		if curnumtype == PreRelease {
			pre, err := r.ReadString('\n')
			if err != nil && err != io.EOF {
				return Version{}, err
			}
			v.prerelease = pre
			break
		}

		// peek the next two bytes to perform
		// a leading zero validation - if there is only
		// a single byte left - try to convert to an integer and assign
		b, err := r.Peek(2)
		if err != nil {
			if err == io.EOF && len(b) == 1 {
				vnum, err := strconv.Atoi(string(b[0]))
				if err != nil {
					return Version{}, err
				}
				r.ReadByte()
				assign(vnum)
				continue
			}
			return Version{}, err
		}

		// check out the peeked bytes -
		// if they are both zero, return an error
		// if first is a zero, and second is a '.' - then return 0
		// else we have a 0 and a non numeric char
		if b[0] == '0' {
			switch b[1] {
			case '0':
				return Version{}, fmt.Errorf("%s version number cannot have leading zeros", curnumtype)
			case '.':
				_, err := r.ReadBytes('.')
				if err != nil {
					return Version{}, err
				}
				assign(0)
				continue
			default:
				return Version{}, fmt.Errorf("%s version number contains invalid , non numeric [0-9] characters", curnumtype)
			}
		}

		// read the next set of bytes
		// validating they are numeric chars or a valid
		// delimiter ('.' or '-')
		var vbytes []byte
		for {
			b, err := r.ReadByte()
			if err != nil {
				if err == io.EOF {
					break
				}
				return Version{}, err
			}
			if b >= '0' && b <= '9' {
				vbytes = append(vbytes, b)
				continue
			}

			if b == '.' || b == '-' {
				break
			}

			return Version{}, fmt.Errorf("%s version number contains invalid , non numeric [0-9] characters", curnumtype)
		}

		// convert valid bytes into integer
		vnum, err := strconv.Atoi(string(vbytes))
		if err != nil {
			return Version{}, fmt.Errorf("%s version number contains invalid , non numeric [0-9] characters", curnumtype)
		}

		assign(vnum)
	}

	return v, nil
}
