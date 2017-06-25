# enved

Simple command to expand environment variable in file.

## Example

```
$ env
WEB_DOCKER_IMAGE=hogehoge

$ cat docker-compose.yml.template
version: '2'
services:
  web:
    image: $WEB_DOCKER_IMAGE
    container_name: app
    env_file:
...

$ enved docker-compose.yml.template
version: '2'
services:
  web:
    image: hogehoge
    container_name: app
    env_file:
...
```
