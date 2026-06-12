# gitx

gitx (git extensions) is a CLI tool for adding custom commands to git.

Commands are just shell scripts, so they can be as simple or as complex as you want.

## Usage

Add an extension with `gitx add`

```shell
$ gitx add alexfu/git-pick-branch
```

> The argument must be a public GitHub repository in the form `owner/repo`

All custom commands in [alexfu/git-pick-branch](https://github.com/alexfu/git-pick-branch) are now available for use like so:

```shell
$ git pick-branch
```

List installed extensions with `gitx list`

```shell
$ gitx list
• alexfu/git-pick-branch/git-pick-branch
```

Remove extensions with `gitx remove`. You can remove an entire extension, a single command, or every extension from an owner.

```shell
$ gitx remove alexfu/git-pick-branch
$ gitx remove alexfu/git-pick-branch/git-pick-branch
$ gitx remove alexfu
```

## Install

<table>
  <thead>
    <tr><th>Platform</th><th width="600">Command</th></tr>
  </thead>
  <tbody>
    <tr>
      <td>macOS, Apple Silicon</td>
      <td width="600">

```shell
curl -sSL https://github.com/alexfu/gitx/releases/latest/download/gitx_darwin_arm64.tar.gz | sudo tar xz -C /usr/local/bin gitx
```

</td>
    </tr>
    <tr>
      <td>macOS, Intel</td>
      <td width="600">

```shell
curl -sSL https://github.com/alexfu/gitx/releases/latest/download/gitx_darwin_amd64.tar.gz | sudo tar xz -C /usr/local/bin gitx
```

</td>
    </tr>
    <tr>
      <td>Linux, x86_64</td>
      <td width="600">

```shell
curl -sSL https://github.com/alexfu/gitx/releases/latest/download/gitx_linux_amd64.tar.gz | sudo tar xz -C /usr/local/bin gitx
```

</td>
    </tr>
  </tbody>
</table>

Then add the following line to your shell's startup file (`~/.zshrc`, `~/.bashrc`, `~/.config/fish/config.fish`, etc.):

```shell
export PATH="$HOME/.gitx:$PATH"
```

Restart your shell or `source` the file for the change to take effect.

## Extensions

|Name|Description|
|----|-----------|
|[alexfu/git-pick-branch](https://github.com/alexfu/git-pick-branch)|Interactively pick a branch with `fzf`|
|[alexfu/git-pr-from-commit](https://github.com/alexfu/git-pr-from-commit)|Create pull requests from a commit|

## License

[MIT](LICENSE) © Alex Fu
