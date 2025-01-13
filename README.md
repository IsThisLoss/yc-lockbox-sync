# yc-lockbox-sync

A tool to synchronize Yandex Cloud Lockbox secret to a file

## Usage

```bash
./yc-lockbox-sync --credentials ./authorized_keys.json --secret-id e6qclp5agq8e7oe2hpgg --secret-key .env --dst ./.env
```

Parameters:

`credentials` - path to file with credentials of service account

`secret-id` - id of the secret to sync in Yandex Cloud Lockbox

`secret-key` - key inside the secret to sync

`dst` - output file path
