package tag

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/syllabix/versioner/internal/git"
)

// Git Integration Test for GetLatest
func TestGetLatest(t *testing.T) {
	// retrieve current branch
	cur, err := git.CurrentBranch()
	if err != nil {
		t.Fatalf("Integration test failed to run: %s", err.Error())
	}

	defer cleanupBranch(t, cur)

	// create and checkout test branch
	cmd := exec.Command("git", "checkout", "-b", branchname)
	if err := cmd.Run(); err != nil {
		t.Fatalf("integration test with Git failed: %v", err)
	}

	// create a dummy file to commit
	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("integration test with Git failed: %v", err)
	}

	defer func() {
		// remove dummy file
		file.Close()
		err := os.Remove(file.Name())
		if err != nil {
			t.Logf("Integration test file %s failed be removed, please remove manually - sorry :(", file.Name())
		}
	}()

	file.WriteString("this file is for git integration testing\n")
	addcommit(t, "feat: added file for integration test")

	file.WriteString("file modification\n")
	addcommit(t, "feat: modified test file")

	file.WriteString("another file modification\n")
	addcommit(t, "chore: quick modification for testing purposes")

	// TODO: determine best way to assert this state - mock of some sort?
	// assert an error occurs when no tags exist
	// _, err = GetLatest()
	// if err == nil {
	// 	t.Errorf("GetLatest() should have returned an error as history has no tags")
	// 	t.FailNow()
	// }

	file.WriteString("awesome new line in the file!\n")
	addcommit(t, "feat: important update")

	tagone := "v0.0.1"
	addTag(t, tagone)
	defer cleanupTag(t, tagone)
	testGetLatest(t, tagone)

	file.WriteString("hard work, really hard work\n")
	addcommit(t, "feat: commit hard work")

	file.WriteString("boring work, really boring work\n")
	addcommit(t, "feat: commit boring work")

	file.WriteString("fun work, really fun work\n")
	addcommit(t, "feat: commit fun work")

	file.WriteString("cool work, really cool work\n")
	addcommit(t, "feat: commit cool work")

	tagtwo := "v0.0.2"
	addTag(t, tagtwo)
	defer cleanupTag(t, tagtwo)
	testGetLatest(t, tagtwo)
}

////////////////////////
// TestGetLatest Helpers
////////////////////////
const (
	branchname = "test/integration/get_latest_func"
	filename   = "getlatest_test.txt"
)

func addcommit(t *testing.T, message string) {
	cmd := exec.Command("git", "add", filename)
	if err := cmd.Run(); err != nil {
		t.Fatalf("integration test with Git failed: %v", err)
	}

	cmd = exec.Command("git", "commit", "-m", message)
	if err := cmd.Run(); err != nil {
		t.Fatalf("integration test with Git failed: %v", err)
	}
}

func addTag(t *testing.T, tag string) {
	tagCmd := exec.Command("git", "tag", "-a", tag, "-m", fmt.Sprintf("release %s", tag))
	if err := tagCmd.Run(); err != nil {
		t.Fatalf("Integration test failed: %v", err)
	}
}

func cleanupTag(t *testing.T, tag string) {
	cmd := exec.Command("git", "tag", "-d", tag)
	if err := cmd.Run(); err != nil {
		t.Fatalf("Integration test failed to clean up - please remove tag: %s", string(tag))
	}
}

func testGetLatest(t *testing.T, tag string) {
	version, err := GetLatest()
	if err != nil {
		t.Errorf("GetLatest() should have found a valid tag version and returned a nil error")
	}

	if version != tag {
		t.Errorf("GetLatest() expected %s, received %s", tag, version)
	}
}

func cleanupBranch(t *testing.T, branch string) {
	// return to branch user was in before test ran
	cmd := exec.Command("git", "checkout", branch)
	if err := cmd.Run(); err != nil {
		t.Logf("Integration test branch %s failed be removed, please remove manually - sorry :(", branchname)
	}

	// delete test branch
	cleanup := exec.Command("git", "branch", "-D", branchname)
	err := cleanup.Run()
	if err != nil {
		t.Logf("Integration test branch %s failed be removed, please remove manually - sorry :(", branchname)
	}
}
