# gitx

gitx (git extensions) is a CLI tool for augmenting `git` through shell scripts.

## Install

<details>
  <summary>macOS, Apple Silicon</summary>

  ```shell
  curl -sSL https://github.com/alexfu/gitx/releases/latest/download/gitx_darwin_arm64.tar.gz | sudo tar xz -C /usr/local/bin gitx
  ```
</details>

<details>
  <summary>macOS, Intel</summary>

  ```shell
  curl -sSL https://github.com/alexfu/gitx/releases/latest/download/gitx_darwin_amd64.tar.gz | sudo tar xz -C /usr/local/bin gitx
  ```
</details>

<details>
  <summary>Linux, x86_64</summary>

  ```shell
  curl -sSL https://github.com/alexfu/gitx/releases/latest/download/gitx_linux_amd64.tar.gz | sudo tar xz -C /usr/local/bin gitx
  ```
</details>

Then add the following line to your shell's startup file (`~/.zshrc`, `~/.bashrc`, `~/.config/fish/config.fish`, etc.):

```shell
export PATH="$HOME/.gitx:$PATH"
```

Restart your shell or `source` the file for the change to take effect.

## Usage

Add extensions with `gitx add`. The argument must be a public GitHub repository in the form `owner/repo`.

```shell
$ gitx add alexfu/git-scripts
```

Then use any of the scripts the extension provides. Each script is named `git-<command>` and is invoked as `git <command>`. For example, `alexfu/git-scripts` provides a `git-open` script, which you call like so:

```shell
$ git open
```

List installed extensions with `gitx list`.

```shell
$ gitx list
• alexfu/git-scripts/git-fdf
• alexfu/git-scripts/git-open
• alexfu/git-scripts/git-peek
• alexfu/git-scripts/git-pr-from-commit
• alexfu/git-scripts/git-rename-branch
• alexfu/git-scripts/git-start
```

Remove extensions with `gitx remove`. You can remove an entire extension, a single script, or every extension from an owner.

```shell
$ gitx remove alexfu/git-scripts
$ gitx remove alexfu/git-scripts/git-open
$ gitx remove alexfu
```

## License

[MIT](LICENSE) © Alex Fu
