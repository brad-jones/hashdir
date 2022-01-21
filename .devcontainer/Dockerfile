# syntax=docker/dockerfile:1.3
# This Dockerfile requires the "buildkit" builder.
# see: https://github.com/moby/buildkit/blob/master/frontend/dockerfile/docs/syntax.md

# CONFIG
# ------------------------------------------------------------------------------
# Any of the values can easily be overridden via --build-arg FOO="bar" flags.
# see: https://docs.docker.com/engine/reference/builder/#arg
ARG DEBIAN_VERSION="11-slim"
ARG DOCKER_CLI_VERSION="5:20.10.8~3-0~debian-bullseye"
ARG DOCKER_COMPOSE_VERSION="1.29.2"
ARG DOCKER_COMPOSE_SHA256="f3f10cf3dbb8107e9ba2ea5f23c1d2159ff7321d16f0a23051d68d8e2547b323"
ARG GO_VERSION="1.17.2"
ARG GO_SHA256="f242a9db6a0ad1846de7b6d94d507915d14062660616a61ef7c808a76e4f1676"
ARG GOSU_VERSION="1.14"
ARG GOSU_SHA256="bd8be776e97ec2b911190a82d9ab3fa6c013ae6d3121eea3d0bfd5c82a0eaf8c"
ARG NODEJS_VERSION="16.10.0"
ARG NODEJS_SHA256="bca7f42ea3e61938cc28868614bb37908111b9ff190fe8022fa9954651b5665d"
ARG NPM_VERSION="7.24.1"
ARG TASK_VERSION="3.9.0"
ARG TASK_SHA256="cad72446d2b939ec611fea14c48f7ce28713c68cc902701fb4f1c2b12fe1fd1c"
ARG LEFTHOOK_VERSION="0.7.6"
ARG LEFTHOOK_SHA256="b762a35481539b4761c5703a3d3fb39df91da94cd49368d6d6c355af07a22d14"
ARG PNPM_VERSION="6.16.0"
ARG PNPM_SHA256="3b021b1d82a830d66588db568b33d94411642488b7b428115ea6af9932eea2fc"
ARG GOCI_LINT_VERSION="1.42.1"
ARG GOCI_SHA256="214b093c15863430c4b66dd39df677dab6e38fc873ded147e331740d50eea51f"
ARG JQ_VERSION="1.6"
ARG JQ_SHA256="af986793a515d500ab2d35f8d2aecd656e764504b789b66d7e1a0b727a124c44"
ARG SHFMT_VERSION="3.4.0"
ARG SHFMT_SHA256="5cd7a2b57a0592f919ca2e4249bd567ae3426801a28ae94d0b26f8f2c4ce17f9"
ARG SHELLCHECK_VERSION="0.7.2"
ARG SHELLCHECK_SHA256="70423609F27B504D6C0C47E340F33652AEA975E45F312324F2DBF91C95A3B188"
ARG DPRINT_VERSION="0.17.3"
ARG DPRINT_SHA256="dc8ee039ecdcadd2b5abe5c1b96e220ffc04aeb43a6b1461db5fc3940f18884d"
ARG GHCLI_VERSION="2.1.0"
ARG GHCLI_SHA256="eddadd0cf7ecf340614dd5914eaad1543d8491c966a3a3a41f8a03832dfac184"
ARG COMMITLINT_VERSION="13.2.0"
ARG SEMANTIC_RELEASE_VERSION="18.0.0"
ARG SEMANTIC_RELEASE_CHANGELOG_VERSION="6.0.0"
ARG SEMANTIC_RELEASE_COMMIT_ANALYZER_VERSION="9.0.1"
ARG SEMANTIC_RELEASE_EXEC_VERSION="6.0.1"
ARG SEMANTIC_RELEASE_GIT_VERSION="10.0.0"
ARG SEMANTIC_RELEASE_RELEASE_NOTES_GENERATOR_VERSION="10.0.2"
ARG CONVENTIONAL_COMMITS_CHANGELOG_VERSION="4.6.1"
ARG ORAS_VERSION="0.12.0"
ARG ORAS_SHA256="660a4ecd87414d1f29610b2ed4630482f1f0d104431576d37e59752c27de37ed"
ARG USERNAME="code"
ARG USER_UID="1000"
ARG USER_GID="${USER_UID}"
ARG SOURCE_SOCKET="/var/run/docker-host.sock"
ARG TARGET_SOCKET="/var/run/docker.sock"
ARG ENTRYPOINT_LOGS="/tmp/entrypoint-logs"

# UPSTREAM IMAGES
# ------------------------------------------------------------------------------
# This is a multi-stage Dockerfile which may consume many different base images.
# The idea here is to create aliases for all our "upstream" images that we depend
# on and then re-use the aliases, this ensures all the stages of a similar type
# do in fact use the same base image.
FROM debian:${DEBIAN_VERSION} as upstream-debian

# BUILDERS
# ------------------------------------------------------------------------------
# Each of the following stages build, download, install & configure some piece
# of software. And then isolate the changes into a "scratch" image, ignoring any
# changes like logs & tmp files.
#
# These changes are then copied into other stages as needed. This is essentially
# a modern version of the apt clean && rm -rf /var/log && rm ... RUN statements.
# One major advantage is we get the full benfits of docker layer caching without
# any of the downsides.
#
# For more details see the "bin/from-now" & "bin/to-now" scripts, which are bind
# mounted into the build when needed. Thus they don't contribute to any actual
# image themselves.

# Install root certs & configure apt to download packages over TLS.
#
# Also it would appear "apt-transport-https" is no longer needed
# see: https://askubuntu.com/a/1154737
FROM upstream-debian as ca-certificates-builder
ENV DEBIAN_FRONTEND="noninteractive"
RUN apt-get -y update
RUN --mount=type=bind,source=bin/from-now.bash,target=/bin/from-now from-now
RUN apt-get -y install --no-install-recommends ca-certificates
RUN sed -i 's/http:\/\//https:\/\//g' /etc/apt/sources.list
RUN --mount=type=bind,source=bin/to-now.bash,target=/bin/to-now to-now
FROM scratch as ca-certificates
COPY --from=ca-certificates-builder /tmp/changed-default /

# We will use this image as the base for all other builders
# This just saves uf from copying the ca-certificates into every other stage.
FROM upstream-debian as builder
ENV DEBIAN_FRONTEND="noninteractive"
COPY --from=ca-certificates / /
RUN apt-get -y update

# Install https://packages.debian.org/bullseye/xz-utils
FROM builder as xz-utils-builder
RUN --mount=type=bind,source=bin/from-now.bash,target=/bin/from-now from-now
RUN apt-get -y install --no-install-recommends xz-utils
RUN --mount=type=bind,source=bin/to-now.bash,target=/bin/to-now to-now
FROM scratch as xz-utils
COPY --from=xz-utils-builder /tmp/changed-default /

# Install https://linux.die.net/man/1/unzip
FROM builder as unzip-builder
RUN --mount=type=bind,source=bin/from-now.bash,target=/bin/from-now from-now
RUN apt-get -y install --no-install-recommends unzip
RUN --mount=type=bind,source=bin/to-now.bash,target=/bin/to-now to-now
FROM scratch as unzip
COPY --from=unzip-builder /tmp/changed-default /

# Install https://curl.se
FROM builder as curl-builder
RUN --mount=type=bind,source=bin/from-now.bash,target=/bin/from-now from-now
RUN apt-get -y install --no-install-recommends curl
RUN --mount=type=bind,source=bin/to-now.bash,target=/bin/to-now to-now
FROM scratch as curl
COPY --from=curl-builder /tmp/changed-default /

# Install https://www.openssh.com (just the client)
FROM builder as openssh-client-builder
RUN --mount=type=bind,source=bin/from-now.bash,target=/bin/from-now from-now
RUN apt-get -y install --no-install-recommends openssh-client
RUN --mount=type=bind,source=bin/to-now.bash,target=/bin/to-now to-now
FROM scratch as openssh-client
COPY --from=openssh-client-builder /tmp/changed-default /

# Install https://git-scm.com
FROM builder as git-builder
RUN --mount=type=bind,source=bin/from-now.bash,target=/bin/from-now from-now
RUN apt-get -y install --no-install-recommends git
RUN --mount=type=bind,source=bin/to-now.bash,target=/bin/to-now to-now
FROM scratch as git
COPY --from=git-builder /tmp/changed-default /

# Install https://gnupg.org
FROM builder as gnupg2-builder
RUN --mount=type=bind,source=bin/from-now.bash,target=/bin/from-now from-now
RUN apt-get -y install --no-install-recommends gnupg2
RUN --mount=type=bind,source=bin/to-now.bash,target=/bin/to-now to-now
FROM scratch as gnupg2
COPY --from=gnupg2-builder /tmp/changed-default /

# Install https://www.sudo.ws
FROM builder as sudo-builder
RUN --mount=type=bind,source=bin/from-now.bash,target=/bin/from-now from-now
RUN apt-get -y install --no-install-recommends sudo
RUN --mount=type=bind,source=bin/to-now.bash,target=/bin/to-now to-now
FROM scratch as sudo
COPY --from=sudo-builder /tmp/changed-default /

# Install https://github.com/tianon/gosu
# The entrypoint script makes use of this
FROM builder as gosu-builder
COPY --from=curl / /
ARG GOSU_VERSION
RUN curl -L "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-amd64" -o /usr/bin/gosu
ARG GOSU_SHA256
RUN echo "${GOSU_SHA256} /usr/bin/gosu" | sha256sum --check --status
RUN chmod +x /usr/bin/gosu
FROM scratch as gosu
COPY --from=gosu-builder /usr/bin/gosu /usr/bin/gosu

# Install https://linux.die.net/man/1/socat
# The entrypoint script makes use of this
FROM builder as socat-builder
RUN --mount=type=bind,source=bin/from-now.bash,target=/bin/from-now from-now
RUN apt-get -y install --no-install-recommends socat
RUN --mount=type=bind,source=bin/to-now.bash,target=/bin/to-now to-now
FROM scratch as socat
COPY --from=socat-builder /tmp/changed-default /

# Install https://man7.org/linux/man-pages/man1/procps.1.html
# The entrypoint script makes use of this
FROM builder as procps-builder
RUN --mount=type=bind,source=bin/from-now.bash,target=/bin/from-now from-now
RUN apt-get -y install --no-install-recommends procps
RUN --mount=type=bind,source=bin/to-now.bash,target=/bin/to-now to-now
FROM scratch as procps
COPY --from=procps-builder /tmp/changed-default /

# Install the docker cli tools - this does not install the docker engine
# see: https://docs.docker.com/engine/install/debian
FROM builder as docker-ce-cli-builder
COPY --from=curl / /
COPY --from=gnupg2 / /
RUN . /etc/os-release && \
  curl -fsSL "https://download.docker.com/linux/${ID}/gpg" | \
  gpg --dearmor -o "/usr/share/keyrings/docker-archive-keyring.gpg" && \
  echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/${ID} ${VERSION_CODENAME} stable" >/etc/apt/sources.list.d/docker.list
RUN apt-get -y update
RUN --mount=type=bind,source=bin/from-now.bash,target=/bin/from-now from-now
ARG DOCKER_CLI_VERSION
RUN apt-get -y install --no-install-recommends "docker-ce-cli=$DOCKER_CLI_VERSION"
RUN --mount=type=bind,source=bin/to-now.bash,target=/bin/to-now to-now
FROM scratch as docker-ce-cli
COPY --from=docker-ce-cli-builder /tmp/changed-default /

# Install https://github.com/docker/compose
FROM builder as docker-compose-builder
COPY --from=curl / /
ARG DOCKER_COMPOSE_VERSION
RUN curl -L "https://github.com/docker/compose/releases/download/$DOCKER_COMPOSE_VERSION/docker-compose-$(uname -s)-$(uname -m)" -o /usr/bin/docker-compose
ARG DOCKER_COMPOSE_SHA256
RUN echo "${DOCKER_COMPOSE_SHA256} /usr/bin/docker-compose" | sha256sum --check --status
RUN chmod +x /usr/bin/docker-compose
FROM scratch as docker-compose
COPY --from=docker-compose-builder /usr/bin/docker-compose /usr/bin/docker-compose

# Install https://golang.org
FROM builder as go-builder
COPY --from=curl / /
ARG GO_VERSION
RUN curl -L "https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz" -o /tmp/go.tar.gz
ARG GO_SHA256
RUN echo "${GO_SHA256} /tmp/go.tar.gz" | sha256sum --check --status
RUN --mount=type=bind,source=bin/from-now.bash,target=/bin/from-now from-now
RUN tar -C /usr/local -xzf /tmp/go.tar.gz
RUN ln -s /usr/local/go/bin/go /usr/bin/go
RUN ln -s /usr/local/go/bin/gofmt /usr/bin/gofmt
RUN --mount=type=bind,source=bin/to-now.bash,target=/bin/to-now to-now
FROM scratch as go
COPY --from=go-builder /tmp/changed-default /

# Install https://nodejs.org
FROM builder as nodejs-builder
COPY --from=curl / /
ARG NODEJS_VERSION
RUN curl -L "https://nodejs.org/dist/v${NODEJS_VERSION}/node-v${NODEJS_VERSION}-linux-x64.tar.gz" -o /tmp/nodejs.tar.gz
ARG NODEJS_SHA256
RUN echo "${NODEJS_SHA256} /tmp/nodejs.tar.gz" | sha256sum --check --status
RUN --mount=type=bind,source=bin/from-now.bash,target=/bin/from-now from-now
RUN tar -C /usr/local -xzf /tmp/nodejs.tar.gz
RUN mv /usr/local/node-v${NODEJS_VERSION}-linux-x64 /usr/local/nodejs
RUN ln -s /usr/local/nodejs/bin/node /usr/bin/node
RUN ln -s /usr/local/nodejs/bin/npm /usr/bin/npm
RUN ln -s /usr/local/nodejs/bin/npx /usr/bin/npx
ARG NPM_VERSION
RUN npm install -g npm@${NPM_VERSION}
RUN --mount=type=bind,source=bin/to-now.bash,target=/bin/to-now to-now
RUN rm -rf /tmp/changed-default/root
FROM scratch as nodejs
COPY --from=nodejs-builder /tmp/changed-default /

# Install https://pnpm.io
#
# We don't actually use this in the final dev container image
# We only use it to install commitlint & semantic-release as it
# reduces the final image size by some 100s of MBs.
FROM builder as pnpm-builder
COPY --from=curl / /
ARG PNPM_VERSION
RUN curl -L "https://github.com/pnpm/pnpm/releases/download/v${PNPM_VERSION}/pnpm-linuxstatic-x64" -o /usr/bin/pnpm
ARG PNPM_SHA256
RUN echo "${PNPM_SHA256} /usr/bin/pnpm" | sha256sum --check --status
RUN chmod +x /usr/bin/pnpm
FROM scratch as pnpm
COPY --from=pnpm-builder /usr/bin/pnpm /usr/bin/pnpm

# Install https://commitlint.js.org
FROM builder as commitlint-builder
COPY --from=nodejs / /
COPY --from=pnpm / /
RUN --mount=type=bind,source=bin/from-now.bash,target=/bin/from-now from-now
ARG COMMITLINT_VERSION
RUN --mount=type=cache,target=/root/.cache/pnpm \
  --mount=type=cache,target=/root/.local/state/pnpm \
  pnpm install --store-dir /.pnpm-store -g \
  "@commitlint/cli@${COMMITLINT_VERSION}" \
  "@commitlint/config-conventional@${COMMITLINT_VERSION}"
RUN echo '#!/usr/bin/env bash\nexec node "/usr/pnpm-global/5/node_modules/@commitlint/cli/cli.js" "$@"' >/usr/bin/commitlint
RUN chmod +x /usr/bin/commitlint
RUN --mount=type=bind,source=bin/to-now.bash,target=/bin/to-now to-now
FROM scratch as commitlint
COPY --from=commitlint-builder /tmp/changed-default /

# Install https://semantic-release.gitbook.io
FROM builder as semantic-release-builder
COPY --from=nodejs / /
COPY --from=pnpm / /
RUN --mount=type=bind,source=bin/from-now.bash,target=/bin/from-now from-now
ARG SEMANTIC_RELEASE_VERSION
ARG SEMANTIC_RELEASE_CHANGELOG_VERSION
ARG SEMANTIC_RELEASE_COMMIT_ANALYZER_VERSION
ARG SEMANTIC_RELEASE_EXEC_VERSION
ARG SEMANTIC_RELEASE_GIT_VERSION
ARG SEMANTIC_RELEASE_RELEASE_NOTES_GENERATOR_VERSION
RUN --mount=type=cache,target=/root/.cache/pnpm \
  --mount=type=cache,target=/root/.local/state/pnpm \
  pnpm install --store-dir /.pnpm-store -g \
  "semantic-release@${SEMANTIC_RELEASE_VERSION}" \
  "@semantic-release/changelog@${SEMANTIC_RELEASE_CHANGELOG_VERSION}" \
  "@semantic-release/commit-analyzer@${SEMANTIC_RELEASE_COMMIT_ANALYZER_VERSION}" \
  "@semantic-release/exec@${SEMANTIC_RELEASE_EXEC_VERSION}" \
  "@semantic-release/git@${SEMANTIC_RELEASE_GIT_VERSION}" \
  "@semantic-release/release-notes-generator@${SEMANTIC_RELEASE_RELEASE_NOTES_GENERATOR_VERSION}" \
  "conventional-changelog-conventionalcommits@${CONVENTIONAL_COMMITS_CHANGELOG_VERSION}"
RUN echo '#!/usr/bin/env bash\nexec node "/usr/pnpm-global/5/node_modules/semantic-release/bin/semantic-release.js" "$@"' >/usr/bin/semantic-release
RUN chmod +x /usr/bin/semantic-release
RUN --mount=type=bind,source=bin/to-now.bash,target=/bin/to-now to-now
FROM scratch as semantic-release
COPY --from=semantic-release-builder /tmp/changed-default /

# Install https://taskfile.dev
FROM builder as task-builder
COPY --from=curl / /
ARG TASK_VERSION
RUN curl -L "https://github.com/go-task/task/releases/download/v${TASK_VERSION}/task_linux_amd64.tar.gz" -o /tmp/task.tar.gz
ARG TASK_SHA256
RUN echo "${TASK_SHA256} /tmp/task.tar.gz" | sha256sum --check --status
RUN tar -C /tmp -xzf /tmp/task.tar.gz
RUN mv /tmp/task /usr/bin/task
RUN chmod +x /usr/bin/task
FROM scratch as task
COPY --from=task-builder /usr/bin/task /usr/bin/task

# Install https://github.com/evilmartians/lefthook
FROM builder as lefthook-builder
COPY --from=curl / /
ARG LEFTHOOK_VERSION
RUN curl -L "https://github.com/evilmartians/lefthook/releases/download/v${LEFTHOOK_VERSION}/lefthook_${LEFTHOOK_VERSION}_Linux_x86_64" -o /usr/bin/lefthook
ARG LEFTHOOK_SHA256
RUN echo "${LEFTHOOK_SHA256} /usr/bin/lefthook" | sha256sum --check --status
RUN chmod +x /usr/bin/lefthook
FROM scratch as lefthook
COPY --from=lefthook-builder /usr/bin/lefthook /usr/bin/lefthook

# Install https://golangci-lint.run
FROM builder as golangci-lint-builder
COPY --from=curl / /
ARG GOCI_LINT_VERSION
RUN curl -L "https://github.com/golangci/golangci-lint/releases/download/v${GOCI_LINT_VERSION}/golangci-lint-${GOCI_LINT_VERSION}-linux-amd64.tar.gz" -o /tmp/golangci-lint.tar.gz
ARG GOCI_SHA256
RUN echo "${GOCI_SHA256} /tmp/golangci-lint.tar.gz" | sha256sum --check --status
RUN tar -C /tmp -xzf /tmp/golangci-lint.tar.gz
RUN mv "/tmp/golangci-lint-${GOCI_LINT_VERSION}-linux-amd64/golangci-lint" "/usr/bin/golangci-lint"
RUN chmod +x /usr/bin/golangci-lint
FROM scratch as golangci-lint
COPY --from=golangci-lint-builder /usr/bin/golangci-lint /usr/bin/golangci-lint

# Install https://github.com/stedolan/jq
FROM builder as jq-builder
COPY --from=curl / /
ARG JQ_VERSION
RUN curl -L "https://github.com/stedolan/jq/releases/download/jq-${JQ_VERSION}/jq-linux64" -o /usr/bin/jq
ARG JQ_SHA256
RUN echo "${JQ_SHA256} /usr/bin/jq" | sha256sum --check --status
RUN chmod +x /usr/bin/jq
FROM scratch as jq
COPY --from=jq-builder /usr/bin/jq /usr/bin/jq

# Install shfmt from https://github.com/mvdan/sh
FROM builder as shfmt-builder
COPY --from=curl / /
ARG SHFMT_VERSION
RUN curl -L "https://github.com/mvdan/sh/releases/download/v${SHFMT_VERSION}/shfmt_v${SHFMT_VERSION}_linux_amd64" -o /usr/bin/shfmt
ARG SHFMT_SHA256
RUN echo "${SHFMT_SHA256} /usr/bin/shfmt" | sha256sum --check --status
RUN chmod +x /usr/bin/shfmt
FROM scratch as shfmt
COPY --from=shfmt-builder /usr/bin/shfmt /usr/bin/shfmt

# Install https://github.com/koalaman/shellcheck
FROM builder as shellcheck-builder
COPY --from=curl / /
COPY --from=xz-utils / /
ARG SHELLCHECK_VERSION
RUN curl -L "https://github.com/koalaman/shellcheck/releases/download/v${SHELLCHECK_VERSION}/shellcheck-v${SHELLCHECK_VERSION}.linux.x86_64.tar.xz" -o /tmp/shellcheck.tar.xz
ARG SHELLCHECK_SHA256
RUN echo "${SHELLCHECK_SHA256} /tmp/shellcheck.tar.xz" | sha256sum --check --status
RUN tar -C /tmp -xf /tmp/shellcheck.tar.xz
RUN mv "/tmp/shellcheck-v${SHELLCHECK_VERSION}/shellcheck" "/usr/bin/shellcheck"
RUN chmod +x /usr/bin/shellcheck
FROM scratch as shellcheck
COPY --from=shellcheck-builder /usr/bin/shellcheck /usr/bin/shellcheck

# Install https://github.com/dprint/dprint
FROM builder as dprint-builder
COPY --from=curl / /
COPY --from=unzip / /
ARG DPRINT_VERSION
RUN curl -L "https://github.com/dprint/dprint/releases/download/${DPRINT_VERSION}/dprint-x86_64-unknown-linux-gnu.zip" -o /tmp/dprint.zip
ARG DPRINT_SHA256
RUN echo "${DPRINT_SHA256} /tmp/dprint.zip" | sha256sum --check --status
RUN unzip /tmp/dprint.zip -d /usr/bin
RUN chmod +x /usr/bin/dprint
FROM scratch as dprint
COPY --from=dprint-builder /usr/bin/dprint /usr/bin/dprint

# Install https://github.com/cli/cli
FROM builder as ghcli-builder
COPY --from=curl / /
ARG GHCLI_VERSION
RUN curl -L "https://github.com/cli/cli/releases/download/v${GHCLI_VERSION}/gh_${GHCLI_VERSION}_linux_amd64.tar.gz" -o /tmp/gh.tar.gz
ARG GHCLI_SHA256
RUN echo "${GHCLI_SHA256} /tmp/gh.tar.gz" | sha256sum --check --status
RUN tar -C /tmp -xf /tmp/gh.tar.gz
RUN mv "/tmp/gh_${GHCLI_VERSION}_linux_amd64/bin/gh" /usr/bin/gh
RUN chmod +x /usr/bin/gh
FROM scratch as ghcli
COPY --from=ghcli-builder /usr/bin/gh /usr/bin/gh

# Install https://oras.land
FROM builder AS oras-builder
COPY --from=curl / /
ARG ORAS_VERSION
RUN curl -L "https://github.com/oras-project/oras/releases/download/v${ORAS_VERSION}/oras_${ORAS_VERSION}_linux_amd64.tar.gz" -o /tmp/oras.tar.gz
ARG ORAS_SHA256
RUN echo "${ORAS_SHA256} /tmp/oras.tar.gz" | sha256sum --check --status
RUN tar -C /tmp -xf /tmp/oras.tar.gz
RUN mv "/tmp/oras" /usr/bin/oras
RUN chmod +x /usr/bin/oras
FROM scratch as oras
COPY --from=oras-builder /usr/bin/oras /usr/bin/oras

# Create an unprivilaged user
#
# In conjuction with the "remoteUser" property this solves the bind mount
# permissions issue on Linux systems.
#
# see: https://code.visualstudio.com/remote/advancedcontainers/add-nonroot-user
FROM builder as user-builder
ARG USERNAME
ARG USER_UID
ARG USER_GID
RUN --mount=type=bind,source=bin/from-now.bash,target=/bin/from-now from-now
COPY --from=sudo / /
RUN groupadd --gid "$USER_GID" "$USERNAME"
RUN useradd --uid "$USER_UID" --gid "$USER_GID" -m "$USERNAME"
RUN echo "$USERNAME ALL=(root) NOPASSWD:ALL" >"/etc/sudoers.d/$USERNAME"
RUN chmod 0440 "/etc/sudoers.d/$USERNAME"
# also: https://code.visualstudio.com/remote/advancedcontainers/persist-bash-history
RUN mkdir /bashhistory
RUN touch /bashhistory/.bash_history
RUN chown -R "$USERNAME" /bashhistory
RUN echo "export PROMPT_COMMAND='history -a' && export HISTFILE=/bashhistory/.bash_history" >>"/home/$USERNAME/.bashrc"
RUN --mount=type=bind,source=bin/to-now.bash,target=/bin/to-now to-now
FROM scratch as user
COPY --from=user-builder /tmp/changed-default /

# FINAL IMAGE
# ------------------------------------------------------------------------------
FROM upstream-debian as vscode-devcontainer

# VsCode will respect the default shell as configured inside the container
# see: https://github.com/microsoft/vscode-remote-release/issues/350
ENV SHELL="/bin/bash"

# Commitlint needs to know where to find the globally installed modules
# see: https://github.com/conventional-changelog/commitlint/issues/613
ENV NODE_PATH="/usr/pnpm-global/5/node_modules"

# Disable CGO so we don't need gcc, etc & further bloat this image.
ENV CGO_ENABLED="0"

# Pull in all our results from our builders
COPY --from=ca-certificates / /
COPY --from=commitlint / /
COPY --from=curl / /
COPY --from=docker-ce-cli / /
COPY --from=docker-compose / /
COPY --from=dprint / /
COPY --from=ghcli / /
COPY --from=git / /
COPY --from=gnupg2 / /
COPY --from=go / /
COPY --from=golangci-lint / /
COPY --from=gosu / /
COPY --from=jq / /
COPY --from=lefthook / /
COPY --from=nodejs / /
COPY --from=openssh-client / /
COPY --from=procps / /
COPY --from=semantic-release / /
COPY --from=shellcheck / /
COPY --from=shfmt / /
COPY --from=socat / /
COPY --from=task / /
COPY --from=user / /
COPY --from=oras / /

# Execute all our tools just to make sure our from-now &
# to-now scripts all worked as expected.
RUN curl --version && \
  docker --version && \
  docker-compose --version && \
  git --version && \
  go version && \
  gosu --version && \
  gpg --version && \
  node --version && \
  npm --version && \
  npx --version && \
  task --version && \
  lefthook --help && \
  socat -h && \
  commitlint --version && \
  semantic-release --version && \
  golangci-lint --version && \
  jq --version && \
  shfmt --version && \
  shellcheck --version && \
  dprint --version && \
  gh --version && \
  ps --help && \
  ssh -V && \
  oras version

# Add our devcontainer hook scripts
COPY ./bin/hooks/. /usr/local/bin/.
RUN for file in /usr/local/bin/*.bash; do mv "$file" "${file%%.bash}"; done
RUN chmod +x /usr/local/bin/*

# Create our entrypoint that will execute each time the devcontainer is started.
#
# This is designed to work with: "overrideCommand": false
# If this property in the devcontainer.json config is set to true (the default)
# then the entrypoint will not be executed.
#
# NOTE: The entrypoint will execute with root permissions so it can do things
# like setup permissions to the mounted docker socket & then it will use gosu
# to switch to the unprivilaged user.
#
# see: https://code.visualstudio.com/docs/remote/containers-advanced#_adding-startup-commands-to-the-docker-image-instead
ARG USERNAME
ARG SOURCE_SOCKET
ARG TARGET_SOCKET
ARG ENTRYPOINT_LOGS
RUN --mount=type=bind,source=bin/install-entrypoint.bash,target=/bin/install-entrypoint install-entrypoint
ENTRYPOINT [ "/entrypoint.bash" ]
CMD [ "sleep", "infinity" ]
