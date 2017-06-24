package pull

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/types"
)

// All Fetch all remotes.
// --all
func All(g *types.Cmd) {
	g.AddOptions("--all")
}

// AllowUnrelatedHistories By default, git merge command refuses to merge histories that do not share a common ancestor.
// This option can be used to override this safety when merging histories of two projects that started their lives independently.
// As that is a very rare occasion, no configuration variable to enable this by default exists and will not be added.
// --allow-unrelated-histories
func AllowUnrelatedHistories(g *types.Cmd) {
	g.AddOptions("--allow-unrelated-histories")
}

// Append Append ref names and object names of fetched refs to the existing contents of .git/FETCH_HEAD. Without this option old data in .git/FETCH_HEAD will be overwritten.
// -a, --append
func Append(g *types.Cmd) {
	g.AddOptions("--append")
}

// Autostash Before starting rebase, stash local modifications away if needed, and apply the stash when done.
// --autostash
func Autostash(g *types.Cmd) {
	g.AddOptions("--autostash")
}

// Commit Perform the merge and commit the result. This option can be used to override --no-commit.
// --commit
func Commit(g *types.Cmd) {
	g.AddOptions("--commit")
}

// Deepen Similar to --depth, except it specifies the number of commits from the current shallow boundary instead of from the tip of each remote branch history.
// --deepen=<depth>
func Deepen(depth string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--deepen=%s", depth))
	}
}

// Depth Limit fetching to the specified number of commits from the tip of each remote branch history. If fetching to a shallow repository created by git clone with --depth=<depth> option (see git-clone(1)), deepen or shorten the history to the specified number of commits. Tags for the deepened commits are not fetched.
// --depth=<depth>
func Depth(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--depth=%s", value))
	}
}

// Edit Invoke an editor before committing successful mechanical merge to further edit the auto-generated merge message, so that the user can explain and justify the merge.
// --edit, -e
func Edit(g *types.Cmd) {
	g.AddOptions("--edit")
}

// Ff When the merge resolves as a fast-forward, only update the branch pointer, without creating a merge commit. This is the default behavior.
// --ff
func Ff(g *types.Cmd) {
	g.AddOptions("--ff")
}

// FfOnly Refuse to merge and exit with a non-zero status unless the current HEAD is already up-to-date or the merge can be resolved as a fast-forward.
// --ff-only
func FfOnly(g *types.Cmd) {
	g.AddOptions("--ff-only")
}

// Force When git fetch is used with <rbranch>:<lbranch> refspec, it refuses to update the local branch <lbranch> unless the remote branch <rbranch> it fetches is a descendant of <lbranch>. This option overrides that check.
// -f, --force
func Force(g *types.Cmd) {
	g.AddOptions("--force")
}

// Ipv4 Use IPv4 addresses only, ignoring IPv6 addresses.
// -4, --ipv4
func Ipv4(g *types.Cmd) {
	g.AddOptions("--ipv4")
}

// Ipv6 Use IPv6 addresses only, ignoring IPv4 addresses.
// -6, --ipv6
func Ipv6(g *types.Cmd) {
	g.AddOptions("--ipv6")
}

// Keep Keep downloaded pack.
// -k, --keep
func Keep(g *types.Cmd) {
	g.AddOptions("--keep")
}

// Log In addition to branch names, populate the log message with one-line descriptions from at most <n> actual commits that are being merged. See also git-fmt-merge-msg(1).
// --log[=<n>]
func Log(n string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(n) == 0 {
			g.AddOptions("--log")
		} else {
			g.AddOptions(fmt.Sprintf("--log=%s", n))
		}
	}
}

// NoAutostash --no-autostash is useful to override the rebase.autoStash configuration variable.
// --no-autostash
func NoAutostash(g *types.Cmd) {
	g.AddOptions("--no-autostash")
}

// NoCommit With --no-commit perform the merge but pretend the merge failed and do not autocommit, to give the user a chance to inspect and further tweak the merge result before committing.
// --no-commit
func NoCommit(g *types.Cmd) {
	g.AddOptions("--no-commit")
}

// NoEdit The --no-edit option can be used to accept the auto-generated message (this is generally discouraged).
// --no-edit
func NoEdit(g *types.Cmd) {
	g.AddOptions("--no-edit")
}

// NoFf Create a merge commit even when the merge resolves as a fast-forward. This is the default behaviour when merging an annotated (and possibly signed) tag.
// --no-ff
func NoFf(g *types.Cmd) {
	g.AddOptions("--no-ff")
}

// NoLog With --no-log do not list one-line descriptions from the actual commits being merged.
// --no-log
func NoLog(g *types.Cmd) {
	g.AddOptions("--no-log")
}

// NoRebase Override earlier --rebase.
// --no-rebase
func NoRebase(g *types.Cmd) {
	g.AddOptions("--no-rebase")
}

// NoRecurseSubmodules This option controls if new commits of all populated submodules should be fetched too.
// --no-recurse-submodules[=yes|on-demand|no]
func NoRecurseSubmodules(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(value) == 0 {
			g.AddOptions("--no-recurse-submodules")
		} else {
			g.AddOptions(fmt.Sprintf("--no-recurse-submodules=%s", value))
		}
	}
}

// NoSquash With --no-squash perform the merge and commit the result. This option can be used to override --squash.
// --no-squash
func NoSquash(g *types.Cmd) {
	g.AddOptions("--no-squash")
}

// NoStat With -n or --no-stat do not show a diffstat at the end of the merge.
// -n, --no-stat
func NoStat(g *types.Cmd) {
	g.AddOptions("--no-stat")
}

// NoTags By default, tags that point at objects that are downloaded from the remote repository are fetched and stored locally. This option disables this automatic tag following. The default behavior for a remote may be specified with the remote.<name>.tagOpt setting.
// --no-tags
func NoTags(g *types.Cmd) {
	g.AddOptions("--no-tags")
}

// NoVerifySignatures Verify that the tip commit of the side branch being merged is signed with a valid key, i.e. a key that has a valid uid: in the default trust model, this means the signing key has been signed by a trusted key.
// --no-verify-signatures
func NoVerifySignatures(g *types.Cmd) {
	g.AddOptions("--no-verify-signatures")
}

// Progress Progress status is reported on the standard error stream by default when it is attached to a terminal, unless -q is specified. This flag forces progress status even if the standard error stream is not directed to a terminal.
// --progress
func Progress(g *types.Cmd) {
	g.AddOptions("--progress")
}

// Quiet This is passed to both underlying git-fetch to squelch reporting of during transfer, and underlying git-merge to squelch output during merging.
// -q, --quiet
func Quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

// Rebase When true, rebase the current branch on top of the upstream branch after fetching. If there is a remote-tracking branch corresponding to the upstream branch and the upstream branch was rebased since last fetched, the rebase uses that information to avoid rebasing non-local changes.
// -r, --rebase[=false|true|preserve|interactive]
func Rebase(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(value) == 0 {
			g.AddOptions("--rebase")
		} else {
			g.AddOptions(fmt.Sprintf("--rebase=%s", value))
		}
	}
}

// RecurseSubmodules This option controls if new commits of all populated submodules should be fetched too.
// --recurse-submodules[=yes|on-demand|no]
func RecurseSubmodules(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(value) == 0 {
			g.AddOptions("--recurse-submodules")
		} else {
			g.AddOptions(fmt.Sprintf("--recurse-submodules=%s", value))
		}
	}
}

// ShallowExclude Deepen or shorten the history of a shallow repository to exclude commits reachable from a specified remote branch or tag. This option can be specified multiple times.
// --shallow-exclude=<revision>
func ShallowExclude(revision string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--shallow-exclude=%s", revision))
	}
}

// ShallowSince Deepen or shorten the history of a shallow repository to include all reachable commits after <date>.
// --shallow-since=<date>
func ShallowSince(date string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--shallow-since=%s", date))
	}
}

// Squash Produce the working tree and index state as if a real merge happened (except for the merge information), but do not actually make a commit, move the HEAD, or record $GIT_DIR/MERGE_HEAD (to cause the next git commit command to create a merge commit). This allows you to create a single commit on top of the current branch whose effect is the same as merging another branch (or more in case of an octopus).
// --squash
func Squash(g *types.Cmd) {
	g.AddOptions("--squash")
}

// Stat Show a diffstat at the end of the merge. The diffstat is also controlled by the configuration option merge.stat.
// --stat
func Stat(g *types.Cmd) {
	g.AddOptions("--stat")
}

// Strategy Use the given merge strategy; can be supplied more than once to specify them in the order they should be tried. If there is no -s option, a built-in list of strategies is used instead (git merge-recursive when merging a single head, git merge-octopus otherwise).
// -s <strategy>, --strategy=<strategy>
func Strategy(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--strategy=%s", value))
	}
}

// StrategyOption Pass merge strategy specific option through to the merge strategy.
// -X <option>, --strategy-option=<option>
func StrategyOption(option string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--strategy-option=%s", option))
	}
}

// Unshallow If the source repository is complete, convert a shallow repository to a complete one, removing all the limitations imposed by shallow repositories.
// If the source repository is shallow, fetch as much as possible so that the current repository has the same history as the source repository.
// --unshallow
func Unshallow(g *types.Cmd) {
	g.AddOptions("--unshallow")
}

// UpdateHeadOk By default git fetch refuses to update the head which corresponds to the current branch. This flag disables the check. This is purely for the internal use for git pull to communicate with git fetch, and unless you are implementing your own Porcelain you are not supposed to use it.
// -u, --update-head-ok
func UpdateHeadOk(g *types.Cmd) {
	g.AddOptions("--update-head-ok")
}

// UpdateShallow By default when fetching from a shallow repository, git fetch refuses refs that require updating .git/shallow. This option updates .git/shallow and accept such refs.
// --update-shallow
func UpdateShallow(g *types.Cmd) {
	g.AddOptions("--update-shallow")
}

// UploadPack When given, and the repository to fetch from is handled by git fetch-pack, --exec=<upload-pack> is passed to the command to specify non-default path for the command run on the other end.
// --upload-pack <upload-pack>
func UploadPack(value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions("--upload-pack")
		g.AddOptions(value)
	}
}

// Verbose Pass --verbose to git-fetch and git-merge.
// -v, --verbose
func Verbose(g *types.Cmd) {
	g.AddOptions("--verbose")
}

// VerifySignatures Verify that the tip commit of the side branch being merged is signed with a valid key, i.e. a key that has a valid uid: in the default trust model, this means the signing key has been signed by a trusted key.
// --verify-signatures
func VerifySignatures(g *types.Cmd) {
	g.AddOptions("--verify-signatures")
}

// Repository The "remote" repository that is the source of a fetch or pull operation. This parameter can be either a URL (see the section GIT URLS below) or the name of a remote (see the section REMOTES below).
func Repository(remote string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(remote)
	}
}

// Refspec Specifies which refs to fetch and which local refs to update. When no <refspec>s appear on the command line, the refs to fetch are read from remote.<repository>.fetch variables instead.
func Refspec(refs ...string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		for _, ref := range refs {
			g.AddOptions(ref)
		}
	}
}
