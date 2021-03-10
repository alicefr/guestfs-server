FROM fedora:33 as LIBGUESTFS

RUN dnf update -y && dnf install --setopt=install_weak_deps=False -y  \
	libguestfs-tools-c  \
	&& dnf clean all \
	&& rm -rf /var/cache/dnf/

FROM LIBGUESTFS

COPY bin/server /usr/bin/virt-guestfs-tool

ENTRYPOINT ["/usr/bin/virt-guestfs-tool"]
