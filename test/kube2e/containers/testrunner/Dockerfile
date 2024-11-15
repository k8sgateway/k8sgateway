# TODO: build/push @ quay.io/solo-io/testrunner:v1.7.0-beta18
# docker buildx build --platform linux/amd64,linux/arm64 -t us-docker.pkg.dev/developers-369321/rav-test/testrunner:v1.7.0-beta17
FROM ubuntu:22.04

RUN apt update && apt install -y curl
COPY --from=lachlanevenson/k8s-kubectl:v1.10.3 /usr/local/bin/kubectl /usr/local/bin/kubectl

# Python
RUN apt-get install -y python2; apt clean
RUN ln -s /bin/python2 /bin/python

COPY root.crt /

CMD ["/bin/sh", "-c", "echo 'STARTING SLEEP! Access me.' && /bin/sleep 36000"]