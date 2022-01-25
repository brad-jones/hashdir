FROM scratch
ENTRYPOINT ["/usr/bin/hashdir"]
COPY hashdir_linux_amd64 /usr/bin/hashdir
