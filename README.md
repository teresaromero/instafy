# 🤳🏻 Instafy - backup your instagram posts

Instafy is a CLI to backup your instagram posts.

## Installation

All you need is install Go and run:

```
$ go get github.com/teresaromero/instafy
```

## Usage

For using the CLI and perform the backups you will need to set the `access_token`. You can retrieve this token by using the login command.

```
$ ig login [username]
```

### Help

```
ig help

Usage: ig [command] [options] [arguments]

Authorization Commands:
    login   Retrieve and save access_token.
    logout  Logout from the account. This deletes de token.
Backup Commands:
    backup Save into cws the content.
Basic Commands:
    help Show help.
    version Show version.
```

### Commands

#### backup

```
ig backup
```

Save the first 20 posts (image, caption, likes, comments) for the instagram account. This generates a package .zip that is downloaded into current working directory where the CLI is running.

| Flag      | Description | Default |
| ----------- | ----------- | ----------- |
| --all   | save all the post for the authorized account | false |
| --storage | select download destination | cwd |
| --scope | choose what data you want to save | images-caption-comments-likes |

#### delete

```
ig delete [post-id]
```

| Flag      | Description | Default |
| ----------- | ----------- | ----------- |
| --all   | deletes all the instagram data | false |

## Roadmap

- Oauth Instagram (work in progress)
- Backup first 20 posts into current working directory: image,caption
- Delete first 20 posts

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](/LICENSE)