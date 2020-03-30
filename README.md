# Helm Mixin for Porter

<img src="https://porter.sh/images/mixins/helm.svg" align="right" width="150px"/>

This is a Helm mixin for [Porter](https://github.com/deislabs/porter). It executes the
appropriate helm command based on which action it is included within: `install`,
`upgrade`, or `delete`.

### Install or Upgrade

```shell
porter mixin install helm
```

### Mixin Configuration

helm client version

```yaml
- helm:
    version: v2.15.2
```

Repositories

```yaml
- helm:
    repositories:
      stable:
        url: "https://kubernetes-charts.storage.googleapis.com"
        cafile: "path/to/cafile"
        certfile: "path/to/certfile"
        keyfile: "path/to/keyfile"
        username: "username"
        password: "password"
```

### Mixin Syntax

Install

```yaml
install:
- helm:
    description: "Description of the command"
    name: RELEASE_NAME
    chart: STABLE_CHART_NAME
    version: CHART_VERSION
    namespace: NAMESPACE
    replace: BOOL
    devel: BOOL
    wait: BOOL # default true
    set:
      VAR1: VALUE1
      VAR2: VALUE2
```

Upgrade

```yaml
install:
- helm:
    description: "Description of the command"
    name: RELEASE_NAME
    chart: STABLE_CHART_NAME
    version: CHART_VERSION
    namespace: NAMESPACE
    resetValues: BOOL
    reuseValues: BOOL
    wait: BOOL # default true
    set:
      VAR1: VALUE1
      VAR2: VALUE2
```

Uninstall

```yaml
uninstall:
- helm:
    description: "Description of command"
    purge: BOOL
    releases:
      - RELEASE_NAME1
      - RELASE_NAME2
```

#### Outputs

The mixin supports saving secrets from Kuberentes as outputs.

```yaml
outputs:
    - name: NAME
      secret: SECRET_NAME
      key: SECRET_KEY
```

### Examples

Install

```yaml
install:
- helm:
    description: "Install MySQL"
    name: mydb
    chart: stable/mysql
    version: 0.10.2
    namespace: mydb
    replace: true
    set:
      mysqlDatabase: wordpress
      mysqlUser: wordpress
    outputs:
      - name: mysql-root-password
        secret: mydb-mysql
        key: mysql-root-password
      - name: mysql-password
        secret: mydb-mysql
        key: mysql-password
```

Upgrade

```yaml
upgrade:
- helm:
    description: "Upgrade MySQL"
    name: porter-ci-mysql
    chart: stable/mysql
    version: 0.10.2
    wait: true
    resetValues: true
    reuseValues: false
    set:
      mysqlDatabase: mydb
      mysqlUser: myuser
      livenessProbe.initialDelaySeconds: 30
      persistence.enabled: true
```

Uninstall

```yaml
uninstall:
- helm:
    description: "Uninstall MySQL"
    purge: true
    releases:
      - mydb
```