# pwdm
Simple application for encrypted password storage and quick recall.

The application uses the session key set by the `set-session` command to encrypt and decrypt values. On decryption it has no knowledge of what session key was used with which values, the only indicator is whether the decryption was successful or not. This means the user can use as many or as few session keys to manage different groups of data as they want as long as they can remember those keys. This is a core concept of this application and it will not change.

## Usage

First set a session key to use for encryption and decryption. Once the session key is set, all operations will use this key. If you want to change it, just run the `set-session` command again.

```
$ ./pwdm set-session example
```

To add a value to the store, use the `add` command:

```
$ ./pwdm add some-descriptive-key password-to-store
```

To retrieve a value, there's a few ways to use the `get` command.

By default it will print the value with no trailing newline characters for piping:

```
$ ./pwdm get some-descriptive-key
password-to-store%
```

To have the application append a newline, pass the `-n/--newline` flag:

```
$ ./pwdm get -n some-descriptive-key
password-to-store
```

In order to copy the retrieved value straight to the clipboard, run the command with the `--clipboard/-c` flag.

```
$ ./pwdm get -c some-descriptive-key
```
