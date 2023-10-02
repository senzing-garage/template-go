# -----------------------------------------------------------------------------
# Stages
# -----------------------------------------------------------------------------

ARG IMAGE_SENZINGAPI_RUNTIME=senzing/senzingapi-runtime:3.7.1
ARG IMAGE_GO_BUILDER=golang:1.21.0-bullseye
ARG IMAGE_FINAL=senzing/senzingapi-runtime:3.7.1

# -----------------------------------------------------------------------------
# Stage: senzingapi_runtime
# -----------------------------------------------------------------------------

FROM ${IMAGE_SENZINGAPI_RUNTIME} as senzingapi_runtime

# -----------------------------------------------------------------------------
# Stage: go_builder
# -----------------------------------------------------------------------------

FROM ${IMAGE_GO_BUILDER} as go_builder
ENV REFRESHED_AT=2023-10-02
LABEL Name="senzing/template-go-builder" \
      Maintainer="support@senzing.com" \
      Version="0.0.1"

# Copy local files from the Git repository.

COPY ./rootfs /
COPY . ${GOPATH}/src/template-go

# Copy files from prior stage.

COPY --from=senzingapi_runtime  "/opt/senzing/g2/lib/"   "/opt/senzing/g2/lib/"
COPY --from=senzingapi_runtime  "/opt/senzing/g2/sdk/c/" "/opt/senzing/g2/sdk/c/"

# Set path to Senzing libs.

ENV LD_LIBRARY_PATH=/opt/senzing/g2/lib/

# Build go program.

WORKDIR ${GOPATH}/src/template-go
RUN make build

# Copy binaries to /output.

RUN mkdir -p /output \
 && cp -R ${GOPATH}/src/template-go/target/*  /output/

# -----------------------------------------------------------------------------
# Stage: final
# -----------------------------------------------------------------------------

FROM ${IMAGE_FINAL} as final
ENV REFRESHED_AT=2023-08-01
LABEL Name="senzing/template-go" \
      Maintainer="support@senzing.com" \
      Version="0.0.1"

# Copy local files from the Git repository.

COPY ./rootfs /

# Copy files from prior stage.

COPY --from=go_builder "/output/linux-amd64/template-go" "/app/template-go"

# Runtime environment variables.

ENV LD_LIBRARY_PATH=/opt/senzing/g2/lib/

# Runtime execution.

WORKDIR /app
ENTRYPOINT ["/app/template-go"]
