# ch08/ex02 (FTP Server)

## Supported FTP Commands

- :heavy_check_mark: supported
- :warning: partially supported
- :no_entry: not supported

| status             | command | note                               |
| ------------------ | ------- | ---------------------------------- |
| :warning:          | USER    | not implemented authentication     |
| :warning:          | PASS    | not implemented authentication     |
| :no_entry:         | ACCT    |                                    |
| :no_entry:         | CWD     |                                    |
| :no_entry:         | CDUP    |                                    |
| :no_entry:         | SMNT    |                                    |
| :heavy_check_mark: | QUIT    |                                    |
| :no_entry:         | REIN    |                                    |
| :heavy_check_mark: | PORT    |                                    |
| :no_entry:         | PASV    |                                    |
| :no_entry:         | TYPE    |                                    |
| :no_entry:         | STRU    |                                    |
| :no_entry:         | MODE    |                                    |
| :warning:          | RETR    | not implemented ASCII, EBCDIC mode |
| :warning:          | STOR    | not implemented ASCII, EBCDIC mode |
| :no_entry:         | STOU    |                                    |
| :no_entry:         | APPE    |                                    |
| :no_entry:         | ALLO    |                                    |
| :no_entry:         | REST    |                                    |
| :no_entry:         | RNFR    |                                    |
| :no_entry:         | RNTO    |                                    |
| :no_entry:         | ABOR    |                                    |
| :no_entry:         | DELE    |                                    |
| :no_entry:         | RMD     |                                    |
| :no_entry:         | MKD     |                                    |
| :no_entry:         | PWD     |                                    |
| :no_entry:         | LIST    |                                    |
| :no_entry:         | NLST    |                                    |
| :no_entry:         | SITE    |                                    |
| :heavy_check_mark: | SYST    |                                    |
| :no_entry:         | STAT    |                                    |
| :no_entry:         | HELP    |                                    |
| :heavy_check_mark: | NOOP    |                                    |
| :heavy_check_mark: | EPRT    |                                    |
| :no_entry:         | EPSV    |                                    |

## References

- [rfc959](https://datatracker.ietf.org/doc/html/rfc959)
- [rfc2428](https://datatracker.ietf.org/doc/html/rfc2428)
