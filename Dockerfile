FROM golang:1.18.2-alpine as build

COPY build/app/go.mod src/go.mod
COPY build/app/cmd src/cmd/
COPY build/app/models src/models/
COPY build/app/restapi src/restapi/

RUN cd src/ && go mod tidy

RUN cd src && \
    export CGO_LDFLAGS="-static -w -s" && \
    go build -tags osusergo,netgo -o /application cmd/vmware-powercli-server/main.go; 

FROM mcr.microsoft.com/powershell:debian-11

ENV TERM linux
ENV TZ=Australia/Sydney

WORKDIR /root

RUN apt-get update && apt-get install curl wget git zip tar ca-certificates unzip jq -y && \
    apt-get install -y --no-install-recommends tzdata  

# Manually download powercli modules based on https://developer.vmware.com/powercli/installation-guide
RUN curl -s -o ./VMware-PowerCLI-13.0.0-20829139.zip -J -L https://vdc-repo.vmware.com/vmwb-repository/dcr-public/02830330-d306-4111-9360-be16afb1d284/c7b98bc2-fcce-44f0-8700-efed2b6275aa/VMware-PowerCLI-13.0.0-20829139.zip
RUN unzip -q -n VMware-PowerCLI-13.0.0-20829139.zip || echo "unzip hates backslashes"
RUN rm -f VMware-PowerCLI-13.0.0-20829139.zip && \
    mv VMware* /opt/microsoft/powershell/7/Modules/  

RUN pwsh -Command "\$ProgressPreference = \"SilentlyContinue\"; Set-PowerCLIConfiguration -Scope AllUsers -ParticipateInCEIP \$false -Confirm:\$false"
# RUN pwsh -Command "Set-PowerCLIConfiguration -Scope AllUsers -ParticipateInCEIP 0 -Confirm:0"

# Install extra SRM module
RUN curl -s -o master.zip -J -L https://github.com/benmeadowcroft/SRM-Cmdlets/archive/master.zip && \
    unzip -q -n master.zip && \
    rm master.zip && \
    mv SRM-Cmdlets-master /opt/microsoft/powershell/7/Modules/Meadowcroft.Srm

# Add in DLLs for SSH.NET
COPY resources /resources

# Configure a default profile option 
# See https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_profiles?view=powershell-7.3
RUN mkdir --parents /root/.config/powershell && \
    echo '$ProgressPreference = "SilentlyContinue"' > /root/.config/powershell/Microsoft.PowerShell_profile.ps1 

# DON'T CHANGE BELOW 
COPY --from=build /application /bin/application

EXPOSE 8080

CMD ["/bin/application", "--port=8080", "--host=0.0.0.0", "--write-timeout=0"]
