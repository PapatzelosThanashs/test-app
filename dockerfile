FROM jenkins/inbound-agent



# Download Go binary
RUN curl -LO https://go.dev/dl/go1.23.1.linux-amd64.tar.gz

# Extract the Go tarball to /usr/local
RUN tar -xzf go1.23.1.linux-amd64.tar.gz

ENV PATH="$PATH:/home/jenkins/go/bin"
