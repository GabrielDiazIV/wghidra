FROM ubuntu
RUN apt-get update
RUN apt-get install -y openjdk-17-jdk
RUN apt-get install -y unzip
RUN apt-get install -y curl
COPY container /container
RUN mkdir -p /container/input
RUN mkdir -p /container/output
WORKDIR /container
COPY /container/gjava /usr/local/bin
COPY /container/gpython /usr/local/bin
COPY /container/pyrun /usr/local/bin
RUN curl -L https://github.com/NationalSecurityAgency/ghidra/releases/download/Ghidra_10.2.3_build/ghidra_10.2.3_PUBLIC_20230208.zip > ghidra.zip
RUN unzip ghidra.zip
RUN rm ghidra.zip
RUN mv ghidra* ghidra
RUN apt-get install -y python3-pip
RUN apt-get install -y netcat
RUN apt-get install -y nano
RUN pip install ghidra_bridge
RUN python3 -m ghidra_bridge.install_server ./scripts
# uncomment this for testing
CMD ["tail", "-f", "/dev/null"]
