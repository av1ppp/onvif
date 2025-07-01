import please

from pathlib import Path


@please.task(name="git:push-tag")
def git_push_tag(ctx: please.TaskContext):
    new_tag = "v" + get_current_version()
    current_tag = get_current_git_tag()
    if new_tag != current_tag:
        please.shexec(f"git tag {new_tag}")

    please.shexec("git push -u origin main --tags")


def get_current_version() -> str:
    return Path("./.version").read_text(encoding="utf-8").strip()


def get_current_git_tag() -> str:
    return (
        please.shexec("git describe --tags --abbrev=0", capture_output=True)
        .stdout.decode("utf-8")
        .strip()
    )
